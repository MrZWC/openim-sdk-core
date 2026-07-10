# 鸿蒙 (HarmonyOS) 开发流程指南

## 概述

本文档说明如何在鸿蒙 Native 项目中集成 openim-sdk-core 编译产物（`libopenim_sdk.so` + `libopenim_sdk.h`），并通过 NAPI 桥接供 ArkTS 层调用。

### 编译产物

| 文件 | 说明 |
|------|------|
| `libopenim_sdk.so` | ARM64 动态共享库，包含完整 IM SDK 功能 |
| `libopenim_sdk.h` | C 头文件，定义所有导出函数和回调结构体 |

> 编译方法请参考 [鸿蒙编译指南](./ohos-build-setup-cn.md)

### SDK 架构

```
ArkTS 层 (UI/业务逻辑)
    ↕ napi 桥接
C++ Native 层 (openim_sdk.h)
    ↕ C ABI
Go SDK (libopenim_sdk.so)
    ↕ 网络通信
OpenIM 服务端
```

---

## 一、项目集成

### 1.1 目录结构

在鸿蒙 DevEco Studio 项目中，将编译产物放置如下：

```
entry/src/main/
├── cpp/
│   ├── CMakeLists.txt
│   ├── types/
│   │   └── libentry/
│   │       ├── Index.d.ts          # NAPI 接口声明
│   │       └── oh-package.json5
│   ├── napi_init.cpp               # NAPI 初始化
│   ├── openim_wrapper.cpp          # SDK 封装
│   └── openim_listeners.cpp        # 监听器实现
├── libs/
│   └── arm64-v8a/
│       └── libopenim_sdk.so        # SDK 动态库
└── resources/
```

### 1.2 CMakeLists.txt 配置

```cmake
cmake_minimum_required(VERSION 3.5.0)
project(openim_entry)

set(NATIVERENDER_ROOT_PATH ${OHOS_SDK_DIR}/native)

# 指定 C++ 标准
set(CMAKE_CXX_STANDARD 17)

# 包含 SDK 头文件目录
# 方式1: 头文件放在 cpp 目录下
include_directories(${CMAKE_CURRENT_SOURCE_DIR})

# 方式2: 头文件放在单独的 include 目录
# include_directories(${CMAKE_CURRENT_SOURCE_DIR}/../../../include)

# 链接 SDK 动态库
add_library(openim_sdk SHARED IMPORTED)
set_target_properties(openim_sdk PROPERTIES
    IMPORTED_LOCATION ${CMAKE_CURRENT_SOURCE_DIR}/../../../libs/${OHOS_ARCH}/libopenim_sdk.so
)

# Native 模块
add_library(entry SHARED
    napi_init.cpp
    openim_wrapper.cpp
    openim_listeners.cpp
)

# 链接依赖
target_link_libraries(entry PUBLIC
    openim_sdk                    # OpenIM SDK
    ${NATIVERENDER_ROOT_PATH}/build/cmake/ohos_napi-cpp-lib.a  # NAPI
)

# ABI 过滤: 仅支持 arm64-v8a
# 在 build-profile.json5 中配置:
# "abiFilters": ["arm64-v8a"]
```

### 1.3 build-profile.json5 配置

在模块级 `build-profile.json5` 中添加 ABI 过滤：

```json5
{
  "module": {
    "name": "entry",
    "type": "entry",
    "deviceTypes": ["phone", "tablet"],
    "abiFilters": ["arm64-v8a"]
  }
}
```

---

## 二、NAPI 桥接层

### 2.1 数据类型转换

ArkTS 与 C 之间需要通过 NAPI 进行数据转换：

```cpp
#include <napi/native_api.h>

// ArkTS string -> C string
std::string napiToString(napi_env env, napi_value value) {
    size_t len = 0;
    napi_get_value_string_utf8(env, value, nullptr, 0, &len);
    std::string str(len, '\0');
    napi_get_value_string_utf8(env, value, str.data(), len + 1, &len);
    return str;
}

// C string -> ArkTS string
napi_value stringToNapi(napi_env env, const char* str) {
    napi_value result;
    napi_create_string_utf8(env, str ? str : "", NAPI_AUTO_LENGTH, &result);
    return result;
}

// C int -> ArkTS number
napi_value intToNapi(napi_env env, int value) {
    napi_value result;
    napi_create_int32(env, value, &result);
    return result;
}
```

### 2.2 异步回调传递到 ArkTS

SDK 的异步操作通过 C 函数指针回调，需要将结果传递到 ArkTS 层。使用 `napi_threadsafe_function` 实现跨线程调用：

```cpp
#include <napi/native_api.h>
#include <thread>

// 全局 NAPI 环境和回调引用
static napi_threadsafe_function g_ts_fn = nullptr;

// 初始化 threadsafe function（在 NAPI 注册时调用）
void initThreadSafeFunction(napi_env env, napi_value callback) {
    napi_value work_name;
    napi_create_string_utf8(env, "OpenIMCallback", NAPI_AUTO_LENGTH, &work_name);

    napi_create_threadsafe_function(
        env, callback, nullptr, work_name,
        0, 1, nullptr, nullptr, nullptr,
        [](napi_env env, napi_value js_cb, void* context, void* data) {
            // 在 ArkTS 主线程执行
            auto* cbData = static_cast<CallbackData*>(data);
            napi_value result;
            napi_create_string_utf8(env, cbData->json.c_str(), NAPI_AUTO_LENGTH, &result);
            napi_call_function(env, nullptr, js_cb, 1, &result, nullptr);
            delete cbData;
        },
        &g_ts_fn
    );
}

// 回调数据结构
struct CallbackData {
    std::string json;
    int errCode = 0;
    std::string errMsg;
};
```

---

## 三、完整开发示例

### 3.1 引入头文件

```cpp
#include "libopenim_sdk.h"
#include <napi/native_api.h>
#include <string>
#include <hilog/log.h>

#define LOG_TAG "OpenIM"
#define LOGI(...) OH_LOG_Print(LOG_APP, LOG_INFO, 0xFF00, LOG_TAG, __VA_ARGS__)
```

### 3.2 初始化 SDK

```cpp
// ===== 连接监听器回调 =====
static void on_connecting() {
    LOGI("OpenIM: connecting");
}

static void on_connect_success() {
    LOGI("OpenIM: connect success");
}

static void on_connect_failed(int32_t err_code, const char* err_msg) {
    LOGI("OpenIM: connect failed: %{public}d %{public}s", err_code, err_msg);
}

static void on_kicked_offline() {
    LOGI("OpenIM: kicked offline");
}

static void on_user_token_expired() {
    LOGI("OpenIM: token expired");
}

static void on_user_token_invalid(const char* err_msg) {
    LOGI("OpenIM: token invalid: %{public}s", err_msg);
}

// 连接监听器结构体（需保持生命周期）
static conn_listener_t g_conn_listener = {
    .on_connecting = on_connecting,
    .on_connect_success = on_connect_success,
    .on_connect_failed = on_connect_failed,
    .on_kicked_offline = on_kicked_offline,
    .on_user_token_expired = on_user_token_expired,
    .on_user_token_invalid = on_user_token_invalid,
};

// ===== NAPI: 初始化 SDK =====
// ArkTS 调用: openim.initSDK(config)
static napi_value InitSDK(napi_env env, napi_callback_info info) {
    size_t argc = 1;
    napi_value args[1];
    napi_get_cb_info(env, info, &argc, args, nullptr, nullptr);

    std::string config = napiToString(env, args[0]);

    // config JSON 示例:
    // {
    //   "platformID": 1,
    //   "apiAddr": "http://your-api:10002",
    //   "wsAddr": "ws://your-ws:10001",
    //   "dataDir": "/data/app/el2/100/base/com.example.app/files/",
    //   "logLevel": 5,
    //   "isLogStandardOutput": false,
    //   "logFilePath": "/data/app/el2/100/base/com.example.app/files/logs/",
    //   "systemType": "ohos"
    // }

    int result = OpenIM_InitSDK(&g_conn_listener, "init_sdk", config.c_str());
    return intToNapi(env, result);
}
```

### 3.3 登录

```cpp
// ===== 基础回调（用于登录等异步操作）=====
// 回调中收到的字符串仅在回调期间有效，如需保留需拷贝
static std::string g_login_result;

static void on_login_success(const char* data) {
    LOGI("OpenIM: login success: %{public}s", data ? data : "");
    if (data) g_login_result = data;
    // 通知 ArkTS 层登录成功
}

static void on_login_error(int32_t err_code, const char* err_msg) {
    LOGI("OpenIM: login error: %{public}d %{public}s", err_code, err_msg);
    // 通知 ArkTS 层登录失败
}

static base_callback_t g_login_callback = {
    .on_success = on_login_success,
    .on_error = on_login_error,
};

// ===== NAPI: 登录 =====
// ArkTS 调用: openim.login(userID, token)
static napi_value Login(napi_env env, napi_callback_info info) {
    size_t argc = 2;
    napi_value args[2];
    napi_get_cb_info(env, info, &argc, args, nullptr, nullptr);

    std::string userID = napiToString(env, args[0]);
    std::string token = napiToString(env, args[1]);

    OpenIM_Login(&g_login_callback, "login_op", userID.c_str(), token.c_str());
    return nullptr;
}

// ===== NAPI: 获取登录状态 =====
// ArkTS 调用: openim.getLoginStatus()
// 返回: 1=已登出, 2=登录中, 3=已登录
static napi_value GetLoginStatus(napi_env env, napi_callback_info info) {
    int status = OpenIM_GetLoginStatus("get_status");
    return intToNapi(env, status);
}

// ===== NAPI: 登出 =====
static napi_value Logout(napi_env env, napi_callback_info info) {
    OpenIM_Logout(&g_login_callback, "logout_op");
    return nullptr;
}
```

### 3.4 设置消息监听器

```cpp
// ===== 消息监听器回调 =====
static void on_recv_new_message(const char* message) {
    // message 是 JSON 格式的消息内容
    LOGI("OpenIM: recv new message: %{public}s", message ? message : "");
    // 传递到 ArkTS 层处理
}

static void on_recv_c2c_read_receipt(const char* msg_receipt_list) {
    LOGI("OpenIM: recv c2c read receipt");
}

static void on_new_recv_message_revoked(const char* message_revoked) {
    LOGI("OpenIM: message revoked");
}

static void on_recv_offline_new_message(const char* message) {
    LOGI("OpenIM: recv offline message");
}

static void on_msg_deleted(const char* message) {
    LOGI("OpenIM: message deleted");
}

static void on_recv_online_only_message(const char* message) {
    LOGI("OpenIM: recv online only message");
}

// 消息监听器结构体
static advanced_msg_listener_t g_msg_listener = {
    .on_recv_new_message = on_recv_new_message,
    .on_recv_c2c_read_receipt = on_recv_c2c_read_receipt,
    .on_new_recv_message_revoked = on_new_recv_message_revoked,
    .on_recv_offline_new_message = on_recv_offline_new_message,
    .on_msg_deleted = on_msg_deleted,
    .on_recv_online_only_message = on_recv_online_only_message,
};

// ===== NAPI: 设置消息监听器 =====
// ArkTS 调用: openim.setAdvancedMsgListener()
static napi_value SetAdvancedMsgListener(napi_env env, napi_callback_info info) {
    OpenIM_SetAdvancedMsgListener(&g_msg_listener);
    return nullptr;
}
```

### 3.5 设置会话监听器

```cpp
// ===== 会话监听器回调 =====
static void on_sync_server_start(int reinstalled) {
    LOGI("OpenIM: sync server start, reinstalled=%{public}d", reinstalled);
}

static void on_sync_server_finish(int reinstalled) {
    LOGI("OpenIM: sync server finish");
}

static void on_sync_server_progress(int progress) {
    LOGI("OpenIM: sync progress: %{public}d%%", progress);
}

static void on_sync_server_failed(int reinstalled) {
    LOGI("OpenIM: sync server failed");
}

static void on_new_conversation(const char* conversation_list) {
    LOGI("OpenIM: new conversation: %{public}s", conversation_list);
}

static void on_conversation_changed(const char* conversation_list) {
    LOGI("OpenIM: conversation changed");
}

static void on_total_unread_count_changed(int32_t total_unread_count) {
    LOGI("OpenIM: total unread: %{public}d", total_unread_count);
}

static void on_conversation_user_input_status_changed(const char* change) {
    LOGI("OpenIM: input status changed");
}

static conversation_listener_t g_conv_listener = {
    .on_sync_server_start = on_sync_server_start,
    .on_sync_server_finish = on_sync_server_finish,
    .on_sync_server_progress = on_sync_server_progress,
    .on_sync_server_failed = on_sync_server_failed,
    .on_new_conversation = on_new_conversation,
    .on_conversation_changed = on_conversation_changed,
    .on_total_unread_count_changed = on_total_unread_count_changed,
    .on_conversation_user_input_status_changed = on_conversation_user_input_status_changed,
};

static napi_value SetConversationListener(napi_env env, napi_callback_info info) {
    OpenIM_SetConversationListener(&g_conv_listener);
    return nullptr;
}
```

### 3.6 发送消息

```cpp
// ===== 发送消息回调（含进度）=====
static void on_send_success(const char* data) {
    LOGI("OpenIM: send success: %{public}s", data);
}

static void on_send_error(int32_t err_code, const char* err_msg) {
    LOGI("OpenIM: send error: %{public}d", err_code);
}

static void on_send_progress(int progress) {
    LOGI("OpenIM: send progress: %{public}d%%", progress);
}

static send_msg_callback_t g_send_callback = {
    .on_success = on_send_success,
    .on_error = on_send_error,
    .on_progress = on_send_progress,
};

// ===== NAPI: 发送文本消息 =====
// ArkTS 调用: openim.sendMessage(recvID, text)
static napi_value SendMessage(napi_env env, napi_callback_info info) {
    size_t argc = 3;
    napi_value args[3];
    napi_get_cb_info(env, info, &argc, args, nullptr, nullptr);

    std::string recvID = napiToString(env, args[0]);
    std::string groupID = napiToString(env, args[1]);
    std::string text = napiToString(env, args[2]);

    // 1. 创建文本消息（同步返回 JSON）
    const char* msgJson = OpenIM_CreateTextMessage("create_msg", text.c_str());
    if (!msgJson) {
        return intToNapi(env, 0);
    }

    // 2. 发送消息（异步）
    std::string msgStr(msgJson);
    OpenIM_FreeString(const_cast<char*>(msgJson));  // 释放创建消息返回的内存

    OpenIM_SendMessage(&g_send_callback, "send_msg",
        msgStr.c_str(), recvID.c_str(), groupID.c_str(), "", 0);
    return intToNapi(env, 1);
}
```

### 3.7 获取会话列表

```cpp
// ===== NAPI: 获取所有会话列表 =====
// ArkTS 调用: openim.getAllConversationList()
static base_callback_t g_query_callback;

static void on_query_success(const char* data) {
    LOGI("OpenIM: query success, data length: %{public}zu", data ? strlen(data) : 0);
    // data 是 JSON 数组，传递到 ArkTS 层
}

static void on_query_error(int32_t err_code, const char* err_msg) {
    LOGI("OpenIM: query error: %{public}d", err_code);
}

static napi_value GetAllConversationList(napi_env env, napi_callback_info info) {
    g_query_callback.on_success = on_query_success;
    g_query_callback.on_error = on_query_error;
    OpenIM_GetAllConversationList(&g_query_callback, "get_conv_list");
    return nullptr;
}
```

### 3.8 NAPI 模块注册

```cpp
// napi_init.cpp
#include "libopenim_sdk.h"

// 声明所有 NAPI 函数
extern napi_value InitSDK(napi_env, napi_callback_info);
extern napi_value Login(napi_env, napi_callback_info);
extern napi_value Logout(napi_env, napi_callback_info);
extern napi_value GetLoginStatus(napi_env, napi_callback_info);
extern napi_value SetAdvancedMsgListener(napi_env, napi_callback_info);
extern napi_value SetConversationListener(napi_env, napi_callback_info);
extern napi_value SendMessage(napi_env, napi_callback_info);
extern napi_value GetAllConversationList(napi_env, napi_callback_info);

// 模块注册
static napi_value Init(napi_env env, napi_value exports) {
    napi_property_descriptor desc[] = {
        {"initSDK", nullptr, InitSDK, nullptr, nullptr, nullptr, napi_default, nullptr},
        {"login", nullptr, Login, nullptr, nullptr, nullptr, napi_default, nullptr},
        {"logout", nullptr, Logout, nullptr, nullptr, nullptr, napi_default, nullptr},
        {"getLoginStatus", nullptr, GetLoginStatus, nullptr, nullptr, nullptr, napi_default, nullptr},
        {"setAdvancedMsgListener", nullptr, SetAdvancedMsgListener, nullptr, nullptr, nullptr, napi_default, nullptr},
        {"setConversationListener", nullptr, SetConversationListener, nullptr, nullptr, nullptr, napi_default, nullptr},
        {"sendMessage", nullptr, SendMessage, nullptr, nullptr, nullptr, napi_default, nullptr},
        {"getAllConversationList", nullptr, GetAllConversationList, nullptr, nullptr, nullptr, napi_default, nullptr},
    };
    napi_define_properties(env, exports, sizeof(desc) / sizeof(desc[0]), desc);
    return exports;
}

static napi_module openimModule = {
    .nm_version = 1,
    .nm_flags = 0,
    .nm_filename = nullptr,
    .nm_register_func = Init,
    .nm_modname = "entry",
    .nm_priv = nullptr,
    .reserved = {0},
};

extern "C" __attribute__((constructor)) void RegisterModule() {
    napi_module_register(&openimModule);
}
```

### 3.9 ArkTS 层调用

```typescript
// Index.ets
import openim from 'libentry.so';

@Entry
@Component
struct Index {
  build() {
    Column() {
      Button('初始化SDK')
        .onClick(() => {
          const config = JSON.stringify({
            platformID: 1,
            apiAddr: 'http://your-api:10002',
            wsAddr: 'ws://your-ws:10001',
            dataDir: getContext(this).filesDir + '/',
            logLevel: 5,
            isLogStandardOutput: false,
            logFilePath: getContext(this).filesDir + '/logs/',
            systemType: 'ohos'
          });
          const result = openim.initSDK(config);
          console.info('InitSDK result: ' + result);
        })

      Button('登录')
        .onClick(() => {
          openim.login('user123', 'tokenxxx');
        })

      Button('设置监听器')
        .onClick(() => {
          openim.setAdvancedMsgListener();
          openim.setConversationListener();
        })

      Button('发送消息')
        .onClick(() => {
          openim.sendMessage('recvUserID', '', 'Hello HarmonyOS!');
        })

      Button('获取会话列表')
        .onClick(() => {
          openim.getAllConversationList();
        })
    }
  }
}
```

---

## 四、内存管理

### 4.1 SDK 返回的字符串

以下函数返回的 `char*` 由 Go 在 C 堆上分配，**必须调用 `OpenIM_FreeString` 释放**：

| 函数 | 说明 |
|------|------|
| `OpenIM_GetSdkVersion()` | SDK 版本号 |
| `OpenIM_GetLoginUserID()` | 当前登录用户 ID |
| `OpenIM_GetAtAllTag()` | @全部 标识 |
| `OpenIM_CreateTextMessage()` | 创建文本消息 |
| `OpenIM_CreateImageMessage()` | 创建图片消息 |
| 所有 `Create*Message()` 函数 | 创建各类消息 |
| `OpenIM_GetConversationIDBySessionType()` | 会话 ID |

```cpp
const char* version = OpenIM_GetSdkVersion();
printf("Version: %s\n", version);
OpenIM_FreeString(const_cast<char*>(version));  // 必须释放
```

### 4.2 回调中的字符串生命周期

回调函数中收到的 `const char*` 字符串**仅在回调执行期间有效**。Go 侧在回调返回后会释放内存。如需保留，必须拷贝：

```cpp
static void on_recv_new_message(const char* message) {
    // 错误: 直接保存指针，回调返回后悬空
    // g_last_message = message;

    // 正确: 拷贝字符串
    std::string msgCopy(message);
    // 后续使用 msgCopy
}
```

---

## 五、API 参考速查表

### 5.1 初始化/登录

| 函数 | 参数 | 返回 | 说明 |
|------|------|------|------|
| `OpenIM_InitSDK` | conn_listener_t*, operationID, config | int(1=成功) | 初始化 |
| `OpenIM_Login` | base_callback_t*, operationID, userID, token | void(异步) | 登录 |
| `OpenIM_Logout` | base_callback_t*, operationID | void(异步) | 登出 |
| `OpenIM_GetLoginStatus` | operationID | int(1-3) | 登录状态 |
| `OpenIM_GetLoginUserID` | - | char*(需释放) | 当前用户 |
| `OpenIM_GetSdkVersion` | - | char*(需释放) | SDK 版本 |
| `OpenIM_UnInitSDK` | operationID | void | 反初始化 |
| `OpenIM_SetAppBackgroundStatus` | base_callback_t*, operationID, isBackground | void(异步) | 前后台状态 |
| `OpenIM_NetworkStatusChanged` | base_callback_t*, operationID | void(异步) | 网络变化 |

### 5.2 监听器设置

| 函数 | 参数 |
|------|------|
| `OpenIM_SetConversationListener` | conversation_listener_t* |
| `OpenIM_SetAdvancedMsgListener` | advanced_msg_listener_t* |
| `OpenIM_SetGroupListener` | group_listener_t* |
| `OpenIM_SetFriendListener` | friendship_listener_t* |
| `OpenIM_SetUserListener` | user_listener_t* |
| `OpenIM_SetCustomBusinessListener` | custom_business_listener_t* |
| `OpenIM_SetMessageKvInfoListener` | message_kv_info_listener_t* |

### 5.3 消息创建（同步返回 char*，需释放）

| 函数 | 参数 |
|------|------|
| `OpenIM_CreateTextMessage` | operationID, text |
| `OpenIM_CreateTextAtMessage` | operationID, text, atUserList, atUsersInfo, message |
| `OpenIM_CreateImageMessage` | operationID, imagePath |
| `OpenIM_CreateSoundMessage` | operationID, soundPath, duration |
| `OpenIM_CreateVideoMessage` | operationID, videoPath, videoType, duration, snapshotPath |
| `OpenIM_CreateFileMessage` | operationID, filePath, fileName |
| `OpenIM_CreateCustomMessage` | operationID, data, extension, description |
| `OpenIM_CreateLocationMessage` | operationID, description, longitude, latitude |
| `OpenIM_CreateQuoteMessage` | operationID, text, message |
| `OpenIM_CreateCardMessage` | operationID, cardInfo |
| `OpenIM_CreateMergerMessage` | operationID, messageList, title, summaryList |
| `OpenIM_CreateFaceMessage` | operationID, index, data |
| `OpenIM_CreateForwardMessage` | operationID, message |

### 5.4 消息操作（异步，使用 base_callback_t 或 send_msg_callback_t）

| 函数 | 参数 | 回调类型 |
|------|------|----------|
| `OpenIM_SendMessage` | send_msg_callback_t*, operationID, message, recvID, groupID, offlinePushInfo, isOnlineOnly | send_msg_callback_t |
| `OpenIM_RevokeMessage` | base_callback_t*, operationID, conversationID, clientMsgID | base_callback_t |
| `OpenIM_DeleteMessage` | base_callback_t*, operationID, conversationID, clientMsgID | base_callback_t |
| `OpenIM_MarkConversationMessageAsRead` | base_callback_t*, operationID, conversationID | base_callback_t |
| `OpenIM_GetAdvancedHistoryMessageList` | base_callback_t*, operationID, getMessageOptions | base_callback_t |
| `OpenIM_SearchLocalMessages` | base_callback_t*, operationID, searchParam | base_callback_t |

### 5.5 会话管理（异步，使用 base_callback_t）

| 函数 | 参数 |
|------|------|
| `OpenIM_GetAllConversationList` | base_callback_t*, operationID |
| `OpenIM_GetConversationListSplit` | base_callback_t*, operationID, offset, count |
| `OpenIM_GetOneConversation` | base_callback_t*, operationID, sessionType, sourceID |
| `OpenIM_SetConversationDraft` | base_callback_t*, operationID, conversationID, draftText |
| `OpenIM_GetTotalUnreadMsgCount` | base_callback_t*, operationID |
| `OpenIM_HideConversation` | base_callback_t*, operationID, conversationID |

### 5.6 群组/好友/用户

群组（29 个）、好友（16 个）、用户（4 个）、第三方（5 个）、在线状态（4 个）的完整函数列表请参考编译生成的 `libopenim_sdk.h` 头文件。

---

## 六、线程安全

### 6.1 回调线程

SDK 的回调（监听器、异步操作结果）在 **Go 运行时的 goroutine 线程**中触发，**不是** NAPI 主线程。直接在回调中操作 ArkTS 对象会导致崩溃。

### 6.2 跨线程通信

使用 `napi_threadsafe_function` 将回调从 Go 线程安全地传递到 ArkTS 主线程：

```cpp
// 1. 初始化时创建 threadsafe function
static napi_threadsafe_function g_ts_fn = nullptr;

void createTsFunction(napi_env env, napi_value callback) {
    napi_value name;
    napi_create_string_utf8(env, "OpenIMEvent", NAPI_AUTO_LENGTH, &name);
    napi_create_threadsafe_function(env, callback, nullptr, name,
        0, 1, nullptr, nullptr, nullptr,
        tsCallJs, &g_ts_fn);
}

// 2. tsCallJs: 在主线程执行的函数
static void tsCallJs(napi_env env, napi_value js_cb, void* context, void* data) {
    auto* eventData = static_cast<std::pair<std::string, std::string>*>(data);
    // eventData->first = 事件类型, eventData->second = JSON 数据
    napi_value type, payload;
    napi_create_string_utf8(env, eventData->first.c_str(), NAPI_AUTO_LENGTH, &type);
    napi_create_string_utf8(env, eventData->second.c_str(), NAPI_AUTO_LENGTH, &payload);
    napi_value args[2] = {type, payload};
    napi_call_function(env, nullptr, js_cb, 2, args, nullptr);
    delete eventData;
}

// 3. 在 C 回调中通过 threadsafe function 发送事件
static void on_recv_new_message(const char* message) {
    if (g_ts_fn) {
        auto* data = new std::pair<std::string, std::string>(
            "onRecvNewMessage", message ? message : "");
        napi_acquire_threadsafe_function(g_ts_fn);
        napi_call_threadsafe_function(g_ts_fn, data, napi_tsfn_nonblocking);
        napi_release_threadsafe_function(g_ts_fn, napi_tsfn_release);
    }
}
```

### 6.3 监听器生命周期

监听器结构体（如 `conn_listener_t`、`advanced_msg_listener_t`）的内存需在 SDK 使用期间保持有效。Go 侧保存的是 C 结构体指针，如果结构体内存被释放或移动，会导致崩溃。

**推荐做法**：使用 `static` 全局变量存储监听器结构体。

---

## 七、常见问题

### Q1: 加载 .so 报错 `dlopen failed`

确认 `libopenim_sdk.so` 放置在 `entry/src/main/libs/arm64-v8a/` 目录下，且 `build-profile.json5` 中 `abiFilters` 包含 `arm64-v8a`。

### Q2: 链接报错 `undefined reference to OpenIM_xxx`

确认 CMakeLists.txt 中正确链接了 `openim_sdk` 库，且 `IMPORTED_LOCATION` 路径正确。

### Q3: 回调函数不触发

1. 确认监听器结构体的函数指针字段已正确赋值（非 NULL）
2. 确认监听器结构体内存有效（使用 static 全局变量）
3. 确认在 `InitSDK` 和 `Login` 成功后再调用 `SetXxxListener`

### Q4: 回调中操作 UI 崩溃

回调在 Go goroutine 线程中执行，不能直接操作 ArkTS UI。必须使用 `napi_threadsafe_function` 切换到主线程。

### Q5: `OpenIM_CreateTextMessage` 返回 NULL

确认 SDK 已成功初始化（`OpenIM_InitSDK` 返回 1）且用户已登录（`OpenIM_GetLoginStatus` 返回 3）。

### Q6: config JSON 格式错误

`apiAddr` 必须包含 `http`，`wsAddr` 必须包含 `ws`，`platformID` 不能为 0。`dataDir` 需为应用可写目录（如 `getContext().filesDir`）。

### Q7: 内存泄漏

所有 `OpenIM_Create*Message`、`OpenIM_GetSdkVersion`、`OpenIM_GetLoginUserID` 等返回 `char*` 的函数，返回值必须调用 `OpenIM_FreeString` 释放。
