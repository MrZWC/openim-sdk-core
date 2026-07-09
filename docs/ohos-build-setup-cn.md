# 编译 openim-sdk-core 为鸿蒙 (HarmonyOS) C 共享库

## 概述

本文档介绍如何将 openim-sdk-core 编译为鸿蒙系统可调用的 `.so` 动态链接库 + `.h` 头文件。编译产物可通过 NAPI/CMake 在鸿蒙 Native 项目中集成调用。

### 编译原理

- 使用标准 Go（1.20+）配合 `GOOS=linux GOARCH=arm64` 交叉编译
- 通过 `go build -buildmode=c-shared` 生成 `.so` + `.h`
- 使用鸿蒙 NDK 的 LLVM/Clang 工具链交叉编译 cgo 依赖（SQLite3）
- Go runtime 的 Linux ARM64 系统调用与鸿蒙系统兼容

> **关于 GOOS=openharmony**：OpenHarmony 社区维护了 Go fork（ohos_golang_go）支持 `GOOS=openharmony`，但其 `-buildmode=c-shared` 模式尚未完善，会报错 `buildmode=c-shared not supported on openharmony/arm64`。因此推荐使用 `GOOS=linux GOARCH=arm64` 方案。

### 与 Android/iOS 的区别

| 平台 | 工具 | 产物 | 回调机制 |
|------|------|------|----------|
| Android | gomobile bind | .aar | 自动生成 Java 接口 |
| iOS | gomobile bind | .xcframework | 自动生成 Swift 协议 |
| 鸿蒙 | go build -buildmode=c-shared | .so + .h | C 函数指针桥接 |

---

## 环境准备

### 系统要求

- **Windows**: Windows 10/11 64 位
- **Linux**: Ubuntu 18.04+ (glibc 2.27+)

### 第一步：安装 Go 语言环境

安装标准 Go 1.20 或更高版本（无需 ohos_golang_go）。

#### Windows

1. 访问 [Go 官方下载页面](https://go.dev/dl/) 下载 Windows 安装包（如 `go1.25.0.windows-amd64.msi`）
2. 运行安装程序，默认安装到 `C:\Program Files\Go`
3. 验证安装：
   ```powershell
   go version
   # 输出: go version go1.25.0 windows/amd64
   ```

#### Linux

```bash
wget https://go.dev/dl/go1.25.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.25.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
```

### 第二步：安装鸿蒙 NDK 工具链

#### 方式 A：通过 OpenHarmony Command Line Tools 安装（推荐）

1. 访问 [华为开发者下载页面](https://developer.huawei.com/consumer/cn/download/) 下载 Command Line Tools
   - **Windows**: 下载 Windows 版并解压
   - **Linux**: 下载 Linux 版并解压
2. 解压后，NDK 工具链位于：
   ```
   <解压目录>/sdk/default/openharmony/native/llvm/bin/
   ```
3. 验证 NDK 工具链存在：
   ```powershell
   # Windows - 检查以下文件是否存在
   Test-Path "C:\Users\<用户名>\AppData\Local\OpenHarmony\command-line-tools\sdk\default\openharmony\native\llvm\bin\clang.exe"
   # 应输出 True
   ```

   > **注意**：Windows 版 NDK 中编译器文件名为 `clang.exe` 和 `clang++.exe`，而非 `aarch64-unknown-linux-ohos-clang.exe`。目标平台通过 `--target=aarch64-linux-ohos` 参数指定。

#### 方式 B：通过 DevEco Studio 安装

1. 下载并安装 [DevEco Studio](https://developer.huawei.com/consumer/cn/download/deveco-studio)
2. DevEco Studio 内置 HarmonyOS SDK，NDK 位于：
   - **Windows**: `C:\Program Files\DevEco Studio\sdk\default\openharmony\native`
   - 或用户自定义安装路径下的 `sdk\default\openharmony\native`

#### NDK 目录结构

```
native/
├── llvm/
│   └── bin/
│       ├── clang.exe              # C 编译器（Windows）
│       ├── clang++.exe            # C++ 编译器（Windows）
│       ├── llvm-ar.exe            # 静态库归档工具
│       ├── aarch64-unknown-linux-ohos-clang     # Linux 下的包装脚本
│       └── ...
└── sysroot/                       # 鸿蒙系统根目录（C 标准库头文件等）
```

### 第三步：验证环境

```powershell
# Windows PowerShell
Write-Host "Go version:" (go version)
Write-Host "NDK clang exists:" (Test-Path "C:\Users\<用户名>\AppData\Local\OpenHarmony\command-line-tools\sdk\default\openharmony\native\llvm\bin\clang.exe")
```

```bash
# Linux
echo "Go version: $(go version)"
ls /path/to/command-line-tools/sdk/default/openharmony/native/llvm/bin/clang
```

---

## 编译

### Windows 编译

编译脚本会自动检测 NDK 路径（支持 OpenHarmony Command Line Tools 和 DevEco Studio 的常见安装路径）。

```powershell
# 方式1: 自动检测 NDK 路径（推荐）
powershell -ExecutionPolicy Bypass -File scripts\build-ohos.ps1

# 方式2: 手动指定 NDK 路径
powershell -ExecutionPolicy Bypass -File scripts\build-ohos.ps1 -OHOS_SDK_PATH "C:\Users\<用户名>\AppData\Local\OpenHarmony\command-line-tools\sdk\default\openharmony\native"

# 方式3: 通过环境变量指定
$env:OHOS_SDK_PATH = "C:\path\to\native"
powershell -ExecutionPolicy Bypass -File scripts\build-ohos.ps1

# 方式4: 通过 Makefile 运行（需安装 GNU Make for Windows）
make ohos
```

### Linux 编译

```bash
# 方式1: 直接运行 bash 脚本
OHOS_SDK_PATH=/path/to/openharmony/native bash scripts/build-ohos.sh

# 方式2: 通过 Makefile 运行
make ohos OHOS_SDK_PATH=/path/to/openharmony/native
```

### 编译参数说明

| 参数 | 说明 | 默认值 |
|------|------|--------|
| `-OHOS_SDK_PATH` | 鸿蒙 NDK native 目录路径 | 自动检测 |
| `-USE_OHOS` | 使用 `GOOS=openharmony` 方案（需 ohos_golang_go，c-shared 可能不支持） | 不启用，默认使用 `GOOS=linux` |

> **默认方案**：脚本默认使用 `GOOS=linux GOARCH=arm64`，这是经过社区验证可行的方案，使用标准 Go 即可编译。

### 编译产物

编译成功后，产物位于：

```
build/ohos/arm64-v8a/
├── libopenim_sdk.so    # 动态共享库
└── libopenim_sdk.h     # C 头文件
```

---

## 鸿蒙项目集成

### 1. 复制文件

将编译产物复制到鸿蒙 Native 项目中：

```
项目根目录/
├── libs/
│   └── arm64-v8a/
│       └── libopenim_sdk.so
├── include/
│   └── libopenim_sdk.h
└── src/main/cpp/       # Native 源码目录
```

### 2. 配置 CMakeLists.txt

在鸿蒙 Native 项目的 `CMakeLists.txt` 中添加：

```cmake
# 链接 openim-sdk-core 动态库
target_link_libraries(entry PRIVATE
    ${NATIVERENDER_ROOT_PATH}/../../../libs/arm64-v8a/libopenim_sdk.so
)

# 添加头文件搜索路径
target_include_directories(entry PRIVATE
    ${NATIVERENDER_ROOT_PATH}/../../../include
)
```

### 3. 调用 SDK API

在 C/C++ 代码中引入头文件并调用 API：

```cpp
#include "libopenim_sdk.h"

// 定义连接监听器回调函数
static void on_connecting() { /* 连接中 */ }
static void on_connect_success() { /* 连接成功 */ }
static void on_connect_failed(int32_t err_code, const char* err_msg) { /* 连接失败 */ }
static void on_kicked_offline() { /* 被踢下线 */ }
static void on_user_token_expired() { /* Token 过期 */ }
static void on_user_token_invalid(const char* err_msg) { /* Token 无效 */ }

// 定义基础回调函数
static void on_success(const char* data) { /* 操作成功，data 为 JSON 结果 */ }
static void on_error(int32_t err_code, const char* err_msg) { /* 操作失败 */ }

// 初始化 SDK
void init_sdk() {
    conn_listener_t conn_listener = {
        .on_connecting = on_connecting,
        .on_connect_success = on_connect_success,
        .on_connect_failed = on_connect_failed,
        .on_kicked_offline = on_kicked_offline,
        .on_user_token_expired = on_user_token_expired,
        .on_user_token_invalid = on_user_token_invalid,
    };

    const char* config = "{\"platformID\":1,\"apiAddr\":\"http://your-api\",\"wsAddr\":\"ws://your-ws\",\"dataDir\":\"/data/\"}";
    int result = OpenIM_InitSDK(&conn_listener, "init_op", config);
    // result: 1=成功, 0=失败
}

// 登录
void login() {
    base_callback_t callback = {
        .on_success = on_success,
        .on_error = on_error,
    };
    OpenIM_Login(&callback, "login_op", "user_id", "token");
}
```

### 4. 内存管理

SDK 返回的 C 字符串（`*char`）由 Go 侧在 C 堆上分配，使用完毕后必须调用 `OpenIM_FreeString` 释放：

```cpp
const char* version = OpenIM_GetSdkVersion();
printf("SDK version: %s\n", version);
OpenIM_FreeString(version);  // 释放内存
```

---

## 常见问题

### Q1: 报错 `-buildmode=c-shared not supported on openharmony/arm64`

**原因**：`GOOS=openharmony` 的 c-shared 模式尚未被 Go 或 ohos_golang_go 完善支持。

**解决**：使用默认的 `GOOS=linux GOARCH=arm64` 方案（脚本默认行为，无需额外参数）：

```powershell
powershell -ExecutionPolicy Bypass -File scripts\build-ohos.ps1
```

如果脚本误用了 `GOOS=openharmony`（旧版本脚本），去掉 `-USE_OHOS` 参数即可。

### Q2: 报错找不到 `aarch64-unknown-linux-ohos-clang.exe`

**原因**：Windows 版鸿蒙 NDK 中编译器文件名为 `clang.exe`，而非 `aarch64-unknown-linux-ohos-clang.exe`（后者是 Linux 版的包装脚本）。

**解决**：确保使用最新版编译脚本（`build-ohos.ps1`），已改为使用 `clang.exe` + `--target` 参数。

### Q3: 编译报错找不到 `sqlite3.h` 或 C 编译器错误

确认鸿蒙 NDK 的 sysroot 路径正确。检查 `OHOS_SDK_PATH` 是否指向 `native` 目录（包含 `sysroot` 子目录）：

```powershell
# 正确的路径（指向 native 目录）
-OHOS_SDK_PATH "C:\Users\xxx\AppData\Local\OpenHarmony\command-line-tools\sdk\default\openharmony\native"

# 错误的路径（指向 llvm\bin 目录）
-OHOS_SDK_PATH "C:\Users\xxx\AppData\Local\OpenHarmony\command-line-tools\sdk\default\openharmony\native\llvm\bin"
```

### Q4: PowerShell 脚本报语法错误（意外标记 `}`）

**原因**：PowerShell 5.1 在中文 Windows 上默认用 GBK 编码读取无 BOM 的 UTF-8 文件，中文注释被错误解析。

**解决**：确保使用最新版编译脚本（`build-ohos.ps1`），已改为纯 ASCII 英文注释。

### Q5: 编译成功但鸿蒙应用加载 .so 报错

确认所有依赖的 `.so` 都已鸿蒙化。Go 编译的 `.so` 可能依赖 `libc.so` 等系统库，鸿蒙系统应已内置。

### Q6: 回调函数没有被调用

检查 C 侧回调结构体中的函数指针是否正确设置（非 NULL）。Go 侧会在调用前检查 NULL。

---

## 导出的 API 列表

SDK 共导出约 120 个 C 函数，按模块分为：

| 模块 | 函数数量 | 说明 |
|------|----------|------|
| 初始化/登录 | 9 | InitSDK、Login、Logout、GetLoginStatus 等 |
| Listener 设置 | 7 | SetConversationListener、SetGroupListener 等 |
| 会话管理 | ~15 | GetAllConversationList、SetConversation 等 |
| 消息创建 | ~20 | CreateTextMessage、CreateImageMessage 等 |
| 消息操作 | ~15 | SendMessage、RevokeMessage、DeleteMessage 等 |
| 群组管理 | 29 | CreateGroup、JoinGroup、GetGroupMemberList 等 |
| 好友关系 | 16 | GetFriendList、AddFriend、AddBlack 等 |
| 用户管理 | 4 | GetUsersInfo、SetSelfInfo 等 |
| 第三方功能 | 5 | UploadFile、UploadLogs、SetAppBadge 等 |
| 在线状态 | 4 | SubscribeUsersStatus、GetUserStatus 等 |
| 内存管理 | 1 | OpenIM_FreeString |

完整的函数签名请参考编译生成的 `libopenim_sdk.h` 头文件。
