# 将 openim-sdk-core 编译为鸿蒙(HarmonyOS)可运行的 C 库

## 概述

将 Go 语言编写的 openim-sdk-core 项目编译为鸿蒙系统可调用的 `.so` 动态链接库 + `.h` 头文件，导出全部约 100+ 个 SDK API，供鸿蒙应用通过 NAPI/CMake 链接调用。

## 当前状态分析

### 项目架构

* **语言**：Go 1.25.0，模块名 `github.com/openimsdk/openim-sdk-core/v3`

* **已有构建**：Android（gomobile bind → AAR）、iOS（gomobile bind → xcframework）、WASM

* **cgo 依赖**：通过 `gorm.io/driver/sqlite` → `github.com/mattn/go-sqlite3` 间接依赖 cgo（SQLite 的 C 实现）

* **无现有 cgo 导出代码**：项目中没有 `import "C"`、`//export` 注解

### API 架构

SDK 的导出函数分布在 `open_im_sdk/` 包的多个文件中，有三种调用模式：

1. **异步回调型**：`func XxxApi(callback Base, operationID string, ...args)` — 通过 `call()` 异步执行，结果通过 `Base` 接口的 `OnSuccess(data string)` / `OnError(errCode int32, errMsg string)` 回调返回
2. **同步返回型**：`func CreateXxxMessage(operationID string, ...args) string` — 通过 `syncCall()` 同步执行，直接返回 JSON 字符串
3. **Listener 设置型**：`func SetXxxListener(listener XxxListener)` — 设置事件监听器

### 回调接口

`open_im_sdk_callback/callback_client.go` 定义了 11 个回调接口：

* `Base`：OnSuccess / OnError

* `SendMsgCallBack`：继承 Base + OnProgress

* `OnConnListener`：6 个连接事件方法

* `OnGroupListener`：11 个群组事件方法

* `OnFriendshipListener`：9 个好友事件方法

* `OnConversationListener`：8 个会话事件方法

* `OnAdvancedMsgListener`：6 个消息事件方法

* `OnUserListener`：2 个用户事件方法

* `OnCustomBusinessListener`：1 个方法

* `OnMessageKvInfoListener`：1 个方法

* `OnSignalingListener`：10 个信令事件方法

* `UploadFileCallback`：9 个文件上传方法

* `UploadLogProgress`：1 个方法

### 技术路线（基于调研确认）

**使用 ohos\_golang\_go（OpenHarmony 官方 Go fork）+ 鸿蒙 NDK 工具链**

* OpenHarmony 社区维护了 Go 的 fork 版本：`https://gitcode.com/openharmony-sig/ohos_golang_go`

* 支持 `GOOS=openharmony GOARCH=arm64` + `-buildmode=c-shared`

* 配合鸿蒙 Command Line Tools 的 LLVM/Clang 工具链编译 cgo 部分（sqlite3）

* 编译环境要求：Ubuntu 18.04+（glibc 2.27+）

**备选方案**：若 `GOOS=openharmony` 的 c-shared 模式不完善，可使用 `GOOS=linux GOARCH=arm64` + 鸿蒙 NDK 工具链（部分用户已验证可行）。

## 实施计划

### 步骤 1：创建 cgo 导出层 — 回调桥接模块

**新建文件**：`cmd/ohos/callback_bridge.go`

这是最核心的模块，负责将 C 函数指针包装为 Go 接口实现，使 Go SDK 能通过 C 回调通知鸿蒙侧。

#### 1.1 C 类型定义（在 `import "C"` 注释块中）

定义所有回调的 C 函数指针类型和结构体：

```c
// Base 回调
typedef void (*on_success_t)(const char* data);
typedef void (*on_error_t)(int32_t err_code, const char* err_msg);

typedef struct {
    on_success_t on_success;
    on_error_t   on_error;
} base_callback_t;

// SendMsg 回调（继承 Base + 进度）
typedef void (*on_progress_t)(int progress);
typedef struct {
    on_success_t  on_success;
    on_error_t    on_error;
    on_progress_t on_progress;
} send_msg_callback_t;

// OnConnListener
typedef void (*on_connecting_t)();
typedef void (*on_connect_success_t)();
typedef void (*on_connect_failed_t)(int32_t err_code, const char* err_msg);
typedef void (*on_kicked_offline_t)();
typedef void (*on_user_token_expired_t)();
typedef void (*on_user_token_invalid_t)(const char* err_msg);
typedef struct { ... } conn_listener_t;

// OnConversationListener
typedef void (*on_sync_server_start_t)(int reinstalled);
typedef void (*on_sync_server_finish_t)(int reinstalled);
// ... 其他 listener 结构体
```

同时定义 C 侧的 wrapper 调用函数（用于安全调用函数指针）：

```c
static void call_on_success(on_success_t cb, const char* data) {
    if (cb) cb(data);
}
// ... 其他 wrapper
```

#### 1.2 Go 侧接口实现

为每个回调接口创建 Go 结构体，持有 C 函数指针，实现对应接口：

```go
// baseCallback 实现 open_im_sdk_callback.Base
type baseCallback struct {
    cb *C.base_callback_t
}

func (b *baseCallback) OnSuccess(data string) {
    cData := C.CString(data)
    defer C.free(unsafe.Pointer(cData))
    C.call_on_success(b.cb.on_success, cData)
}

func (b *baseCallback) OnError(errCode int32, errMsg string) {
    cMsg := C.CString(errMsg)
    defer C.free(unsafe.Pointer(cMsg))
    C.call_on_error(b.cb.on_error, C.int32_t(errCode), cMsg)
}

// connListener 实现 open_im_sdk_callback.OnConnListener
type connListener struct {
    cb *C.conn_listener_t
}
// ... 实现所有方法
```

对所有 11+ 个回调接口都创建对应的 Go 桥接实现。

### 步骤 2：创建 cgo 导出层 — API 导出模块

按模块分文件，每个文件用 `//export` 导出对应模块的 C 函数。所有文件在 `cmd/ohos/` 目录下，属于 `main` 包。

#### 2.1 `cmd/ohos/main.go` — 入口文件

```go
package main

import "C"

// main 函数是 c-shared 模式必需的
func main() {}
```

#### 2.2 `cmd/ohos/init_login_c.go` — 初始化/登录（约 10 个函数）

导出：InitSDK、UnInitSDK、Login、Logout、GetLoginStatus、GetLoginUserID、GetSdkVersion、SetAppBackgroundStatus、NetworkStatusChanged

```go
//export OpenIM_InitSDK
func OpenIM_InitSDK(connListener *C.conn_listener_t, operationID, config *C.char) C.int {
    listener := newConnListener(connListener)
    result := open_im_sdk.InitSDK(listener, C.GoString(operationID), C.GoString(config))
    if result { return 1 }
    return 0
}

//export OpenIM_Login
func OpenIM_Login(cb *C.base_callback_t, operationID, userID, token *C.char) {
    callback := newBaseCallback(cb)
    open_im_sdk.Login(callback, C.GoString(operationID), C.GoString(userID), C.GoString(token))
}

//export OpenIM_GetLoginStatus
func OpenIM_GetLoginStatus(operationID *C.char) C.int {
    return C.int(open_im_sdk.GetLoginStatus(C.GoString(operationID)))
}

//export OpenIM_GetSdkVersion
func OpenIM_GetSdkVersion() *C.char {
    return C.CString(open_im_sdk.GetSdkVersion())
}
// ... 其他
```

#### 2.3 `cmd/ohos/listener_c.go` — Listener 设置（约 7 个函数）

导出：SetGroupListener、SetConversationListener、SetAdvancedMsgListener、SetUserListener、SetFriendListener、SetCustomBusinessListener、SetMessageKvInfoListener

```go
//export OpenIM_SetConversationListener
func OpenIM_SetConversationListener(listener *C.conversation_listener_t) {
    open_im_sdk.SetConversationListener(newConversationListener(listener))
}
```

#### 2.4 `cmd/ohos/conversation_msg_c.go` — 会话/消息（约 45 个函数）

导出所有会话管理和消息操作函数，包括：

* 会话列表：GetAllConversationList、GetConversationListSplit、GetOneConversation 等

* 消息创建：CreateTextMessage、CreateImageMessage 等（约 20 个 Create\* 函数）

* 消息操作：SendMessage、RevokeMessage、DeleteMessage、MarkConversationMessageAsRead 等

* 消息历史：GetAdvancedHistoryMessageList、SearchLocalMessages 等

#### 2.5 `cmd/ohos/group_c.go` — 群组（约 25 个函数）

导出：CreateGroup、JoinGroup、QuitGroup、DismissGroup、GetJoinedGroupList 等

#### 2.6 `cmd/ohos/relation_c.go` — 好友关系（约 16 个函数）

导出：GetFriendList、AddFriend、DeleteFriend、AcceptFriendApplication 等

#### 2.7 `cmd/ohos/user_c.go` — 用户（约 4 个函数）

导出：GetUsersInfo、SetSelfInfo、GetSelfUserInfo、GetUserClientConfig

#### 2.8 `cmd/ohos/third_c.go` — 第三方功能（约 5 个函数）

导出：UpdateFcmToken、SetAppBadge、UploadLogs、Logs、UploadFile
注意：UploadFile 和 UploadLogs 有特殊回调类型（UploadFileCallback、UploadLogProgress），需要在 callback\_bridge.go 中添加对应桥接。

#### 2.9 `cmd/ohos/online_c.go` — 在线状态（约 4 个函数）

导出：SubscribeUsersStatus、UnsubscribeUsersStatus、GetSubscribeUsersStatus、GetUserStatus

### 步骤 3：创建构建脚本

#### 3.1 `scripts/build-ohos.sh` — 鸿蒙编译脚本

```bash
#!/bin/bash
# 鸿蒙 HarmonyOS C 库编译脚本
# 前置条件：
#   1. 已安装 ohos_golang_go（GOOS=openharmony 支持）
#   2. 已下载鸿蒙 Command Line Tools

# 配置鸿蒙 NDK 工具链路径（需用户根据实际路径修改）
OHOS_SDK_PATH="${OHOS_SDK_PATH:-/path/to/command-line-tools/sdk/default/openharmony/native}"

export GOOS=openharmony
export GOARCH=arm64
export CGO_ENABLED=1
export CC="${OHOS_SDK_PATH}/llvm/bin/aarch64-unknown-linux-ohos-clang"
export CXX="${OHOS_SDK_PATH}/llvm/bin/aarch64-unknown-linux-ohos-clang++"
export AR="${OHOS_SDK_PATH}/llvm/bin/llvm-ar"

# 编译为 c-shared 动态库
go build -buildmode=c-shared -trimpath -ldflags "-s -w" \
    -o build/ohos/arm64-v8a/libopenim_sdk.so \
    ./cmd/ohos/
```

#### 3.2 Makefile 新增目标

在根目录 `Makefile` 中新增：

```makefile
## ohos: Build the HarmonyOS shared library
.PHONY: ohos
ohos:
	@bash scripts/build-ohos.sh
```

### 步骤 4：字符串内存管理

C 侧字符串（`*C.char`）的内存管理是关键点：

* **入参字符串**：Go 侧用 `C.GoString()` 转换后，C 侧内存由调用方管理

* **返回字符串**：Go 侧用 `C.CString()` 分配，C 侧使用后需调用 `OpenIM_FreeString()` 释放

* 在 `main.go` 中导出内存释放函数：

```go
//export OpenIM_FreeString
func OpenIM_FreeString(s *C.char) {
    C.free(unsafe.Pointer(s))
}
```

## 假设与决策

1. **编译环境**：Ubuntu 18.04+ (glibc 2.27+)，因为 ohos\_golang\_go 需要在 Linux 上编译
2. **Go 工具链**：使用 ohos\_golang\_go fork（支持 `GOOS=openharmony`），而非标准 Go
3. **目标架构**：仅 arm64-v8a（鸿蒙设备主流架构），暂不支持 armeabi-v7a
4. **c-shared 模式**：使用 `-buildmode=c-shared` 生成 `.so` + `.h`
5. **备选方案**：若 `GOOS=openharmony` 的 c-shared 不支持，回退到 `GOOS=linux GOARCH=arm64` + 鸿蒙 NDK
6. **回调线程安全**：Go 回调可能在任意 goroutine 触发，cgo 调用 C 函数指针是安全的，但 C 侧回调函数需自行处理线程同步
7. **SQLite 依赖**：sqlite3 的 C 代码通过鸿蒙 NDK 的 clang 交叉编译，无需额外预编译
8. **不修改现有源码**：所有变更都在新增的 `cmd/ohos/` 目录和构建脚本中，不修改 `open_im_sdk/` 和 `open_im_sdk_callback/` 中的现有代码

## 产出文件清单

| 文件                               | 说明                         |
| -------------------------------- | -------------------------- |
| `cmd/ohos/main.go`               | c-shared 入口 + 内存释放函数       |
| `cmd/ohos/callback_bridge.go`    | C 函数指针 → Go 接口桥接（全部回调接口）   |
| `cmd/ohos/init_login_c.go`       | 初始化/登录 API 导出（\~10 函数）     |
| `cmd/ohos/listener_c.go`         | Listener 设置 API 导出（\~7 函数） |
| `cmd/ohos/conversation_msg_c.go` | 会话/消息 API 导出（\~45 函数）      |
| `cmd/ohos/group_c.go`            | 群组 API 导出（\~25 函数）         |
| `cmd/ohos/relation_c.go`         | 好友关系 API 导出（\~16 函数）       |
| `cmd/ohos/user_c.go`             | 用户 API 导出（\~4 函数）          |
| `cmd/ohos/third_c.go`            | 第三方功能 API 导出（\~5 函数）       |
| `cmd/ohos/online_c.go`           | 在线状态 API 导出（\~4 函数）        |
| `scripts/build-ohos.sh`          | 鸿蒙编译脚本                     |
| `Makefile`（修改）                   | 新增 `ohos` 构建目标             |

## 验证步骤

1. **编译验证**：在配置好 ohos\_golang\_go + 鸿蒙 NDK 的 Ubuntu 环境中执行 `make ohos`，确认生成 `build/ohos/arm64-v8a/libopenim_sdk.so` 和 `libopenim_sdk.h`
2. **头文件检查**：检查生成的 `.h` 文件包含所有 `OpenIM_*` 导出函数声明
3. **符号检查**：使用 `nm -D libopenim_sdk.so` 确认所有导出符号存在
4. **集成测试**：创建鸿蒙 Native 项目，在 CMakeLists.txt 中链接 .so，调用 `OpenIM_InitSDK` → `OpenIM_Login` 验证基本流程
5. **回调验证**：验证 Login 回调能正确通过 C 函数指针传递到鸿蒙侧

## 风险与缓解

| 风险                              | 缓解方案                                                |
| ------------------------------- | --------------------------------------------------- |
| `GOOS=openharmony` c-shared 不支持 | 回退到 `GOOS=linux GOARCH=arm64` + 鸿蒙 NDK，社区已验证可行      |
| sqlite3 交叉编译失败                  | 确保鸿蒙 NDK 的 sysroot 包含必要的 C 头文件；或预编译 sqlite3 静态库     |
| Go 运行时在鸿蒙上系统调用不兼容               | 使用 `GOOS=linux` 方案时，Go runtime 的 Linux 系统调用大部分与鸿蒙兼容 |
| 回调线程安全问题                        | 在 C 侧回调函数中避免阻塞操作，必要时使用消息队列传递到主线程                    |

