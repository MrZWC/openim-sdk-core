# OHOS NAPI 桥接层布尔类型同步改造指南

## 1. 背景与目标

openim-sdk-core 的 OHOS 导出层已完成"布尔语义统一改造":将所有用 `int`/`C.int` 表达布尔语义的对外 C 接口,统一改为 C99 的 `_Bool`/`bool`,与 Android 端的布尔语义对齐。

本次改造是**破坏性 ABI 变更**(`int` 4 字节 → `_Bool` 1 字节)。OHOS 侧 NAPI 桥接代码(ArkTS/JS ↔ native C)必须同步更新,否则会出现参数截断、读取错位、回调值错误等问题。

本文档列出 NAPI 侧需要同步修改的全部内容。

---

## 2. ABI 变更影响

| 维度 | 改造前 | 改造后 | 字节数 |
|------|--------|--------|--------|
| C 参数/返回类型 | `int` | `_Bool` (`bool`) | 4 → 1 |
| NAPI 取值 API | `napi_get_value_int32` | `napi_get_value_bool` | — |
| NAPI 建值 API | `napi_create_int32` | `napi_create_bool` | — |

> 关键风险:`int` 与 `_Bool` 大小不同。若 NAPI 侧仍按 `int32_t` 读写,而 native 侧按 `_Bool` 读写,会导致栈/寄存器参数解析错位,行为未定义。务必全量同步。

---

## 3. 改动分类总览

NAPI 侧改动分三类:

| 类别 | 方向 | 涉及 NAPI 操作 |
|------|------|----------------|
| A. 导出函数入参 | ArkTS → native | `napi_get_value_int32` → `napi_get_value_bool` |
| B. 导出函数返回值 | native → ArkTS | `napi_create_int32` → `napi_create_bool` |
| C. 回调函数签名 | native → ArkTS | 回调实现签名 `int` → `bool`;回调内 `napi_create_int32` → `napi_create_bool` |

---

## 4. 类别 A:导出函数入参(ArkTS → native)

### 4.1 涉及接口清单

以下 8 个导出函数的布尔参数,类型由 `int` 变为 `_Bool`:

| 导出函数 | 参数 | 改前 C 签名 | 改后 C 签名 |
|----------|------|-------------|-------------|
| `OpenIM_SendMessage` | `isOnlineOnly` | `int` | `_Bool` |
| `OpenIM_SendMessageNotOss` | `isOnlineOnly` | `int` | `_Bool` |
| `OpenIM_ChangeInputStates` | `focus` | `int` | `_Bool` |
| `OpenIM_ChangeGroupMute` | `isMute` | `int` | `_Bool` |
| `OpenIM_SetAppBackgroundStatus` | `isBackground` | `int` | `_Bool` |
| `OpenIM_GetSpecifiedFriendsInfo` | `filterBlack` | `int` | `_Bool` |
| `OpenIM_GetFriendList` | `filterBlack` | `int` | `_Bool` |
| `OpenIM_GetFriendListPage` | `filterBlack` | `int` | `_Bool` |

> 注:`OpenIM_GetFriendListPage` 的 `offset`、`count` 仍为 `int32_t`,**不要改**。

### 4.2 NAPI 代码改法

每个涉及上述布尔参数的 NAPI 包装函数,都需要把"从 JS 值提取参数"的步骤从 `int32` 改为 `bool`。

**改前示例**(以 `OpenIM_ChangeGroupMute` 为例):

```cpp
// 改前:按 int32 提取
int32_t isMuteInt = 0;
napi_status status = napi_get_value_int32(env, argv[3], &isMuteInt);
if (status != napi_ok) { /* 错误处理 */ }

// 调用 native(签名: void OpenIM_ChangeGroupMute(..., int isMute))
OpenIM_ChangeGroupMute(cb, operationID, groupID, isMuteInt);
```

**改后示例**:

```cpp
// 改后:按 bool 提取
bool isMute = false;
napi_status status = napi_get_value_bool(env, argv[3], &isMute);
if (status != napi_ok) { /* 错误处理 */ }

// 调用 native(签名已变为: void OpenIM_ChangeGroupMute(..., _Bool isMute))
OpenIM_ChangeGroupMute(cb, operationID, groupID, (_Bool)isMute);
```

**注意事项**:
- ArkTS 侧调用时,直接传 `boolean` 即可(NAPI 会自动用 `napi_get_value_bool` 解析),无需再传 `0/1`。
- 若原先 ArkTS 侧传的是数字 `0`/`1`,建议同步改为 `false`/`true`;若暂时保留数字,`napi_get_value_bool` 对非零数字会返回 `true`,语义兼容,但建议统一为布尔字面量。

---

## 5. 类别 B:导出函数返回值(native → ArkTS)

### 5.1 涉及接口

| 导出函数 | 返回值 | 改前 | 改后 |
|----------|--------|------|------|
| `OpenIM_InitSDK` | 初始化结果 | `int` (1=成功, 0=失败) | `_Bool` (true=成功, false=失败) |

> 注:`OpenIM_GetLoginStatus` 返回值仍是 `int`(1=已登出, 2=登录中, 3=已登录),**不要改**。

### 5.2 NAPI 代码改法

`OpenIM_InitSDK` 是同步导出函数,其返回值需转成 JS 值返回给 ArkTS。

**改前示例**:

```cpp
// 改前:native 返回 int
int ret = OpenIM_InitSDK(connListener, operationID, config);

napi_value jsResult;
napi_create_int32(env, ret, &jsResult);  // 把 0/1 作为数字返回
return jsResult;
```

**改后示例**:

```cpp
// 改后:native 返回 _Bool
_Bool ret = OpenIM_InitSDK(connListener, operationID, config);

napi_value jsResult;
napi_create_bool(env, (bool)ret, &jsResult);  // 作为 boolean 返回
return jsResult;
```

**注意事项**:
- ArkTS 侧接收 `OpenIM_InitSDK` 返回值的变量,类型应从 `number` 改为 `boolean`。
- 若 ArkTS 侧原先用 `if (ret === 1)` 判断,需改为 `if (ret === true)` 或直接 `if (ret)`。

---

## 6. 类别 C:回调函数签名(native → ArkTS)

### 6.1 涉及回调

以下 3 个回调函数指针的类型签名,`reinstalled` 参数由 `int` 变为 `bool`:

| 回调函数指针类型 | 参数 | 改前 C 签名 | 改后 C 签名 |
|------------------|------|-------------|-------------|
| `on_sync_server_start_t` | `reinstalled` | `int` | `bool` |
| `on_sync_server_finish_t` | `reinstalled` | `int` | `bool` |
| `on_sync_server_failed_t` | `reinstalled` | `int` | `bool` |

> 注:`on_sync_server_progress_t(int progress)` **未改**,仍为 `int`(进度值)。

### 6.2 OHOS 侧回调实现改法

OHOS 侧需要实现这些回调函数(注册到 `conversation_listener_t` 结构体),其函数签名必须与新的 C 类型一致。

**改前示例**:

```cpp
// 改前:回调实现按 int 接收
static void on_sync_server_start_cb(int reinstalled) {
    // 把 reinstalled 传回 ArkTS 层
    napi_value jsArg;
    napi_create_int32(env, reinstalled, &jsArg);  // 作为数字传回
    // ... 调用 ArkTS 回调
}

// 注册
listener->on_sync_server_start = on_sync_server_start_cb;
```

**改后示例**:

```cpp
// 改后:回调实现按 bool 接收
static void on_sync_server_start_cb(bool reinstalled) {
    // 把 reinstalled 传回 ArkTS 层
    napi_value jsArg;
    napi_create_bool(env, reinstalled, &jsArg);  // 作为 boolean 传回
    // ... 调用 ArkTS 回调
}

// 注册(注册方式不变)
listener->on_sync_server_start = on_sync_server_start_cb;
```

三个回调(`on_sync_server_start`、`on_sync_server_finish`、`on_sync_server_failed`)都按同样方式修改。

**注意事项**:
- 回调函数签名必须从 `int` 改为 `bool`,否则注册时函数指针类型不匹配(C 编译器会报警告/错误,或运行时栈布局错位)。
- ArkTS 侧 `OnSyncServerStart` 等回调方法接收的参数类型,应从 `number` 改为 `boolean`。

---

## 7. 明确不需要改动的接口(避免误改)

以下接口虽为 `int` 类型,但语义为**真整数**(非布尔),NAPI 侧**保持原样**:

| 接口 | 参数/返回 | 语义 |
|------|-----------|------|
| `on_sync_server_progress_t` | `progress` | 进度值(0-100) |
| `OpenIM_GetLoginStatus` 返回值 | `int` | 登录状态码(1/2/3) |
| `OpenIM_ChangeGroupMemberMute` | `mutedSeconds` | 禁言秒数 |
| `OpenIM_GetConversationListSplit` | `offset`, `count` | 分页参数 |
| `OpenIM_GetConversationIDBySessionType` | `sessionType` | 会话类型枚举 |
| `OpenIM_CreateFaceMessage` | `index` | 表情索引 |
| `OpenIM_UploadLogs` / `OpenIM_Logs` | `line`, `logLevel` | 行号/日志级别 |
| `base_callback_t` / `conn_listener_t` 的 `err_code` | `int32_t` | 错误码 |
| `on_total_unread_count_changed_t` | `totalUnread_count` | 未读总数 |
| `upload_file_callback_t` 各 `index`/`num`/`typ` | `int` | 索引/数量/类型 |

---

## 8. 改造检查清单

完成 NAPI 侧修改后,按以下清单逐项确认:

- [ ] **入参提取**:8 个导出函数的布尔参数,均已从 `napi_get_value_int32` 改为 `napi_get_value_bool`
- [ ] **返回值建值**:`OpenIM_InitSDK` 返回值已从 `napi_create_int32` 改为 `napi_create_bool`
- [ ] **回调签名**:3 个 `on_sync_server_*` 回调实现函数签名已从 `int` 改为 `bool`
- [ ] **回调内建值**:3 个回调内传回 ArkTS 的 `reinstalled`,已从 `napi_create_int32` 改为 `napi_create_bool`
- [ ] **未误改真整数**:`progress`、登录状态码、`mutedSeconds`、分页参数等仍走 `int32` 通道
- [ ] **ArkTS 侧类型同步**:对应 ArkTS 接口的参数/返回值类型,已从 `number` 改为 `boolean`
- [ ] **头文件更新**:已使用最新生成的 `libopenim_sdk.h`(通过 `make ohos` 重新构建获取)
- [ ] **联调验证**:`OnSyncServerStart(true/false)`、`OpenIM_ChangeGroupMute(true/false)`、`OpenIM_InitSDK()` 返回值在 ArkTS 侧接收正确

---

## 9. 附:Go/C 侧改动来源(供对照)

本次 NAPI 同步改造对应 Go/C 侧的以下文件改动(详见仓库提交):

- `cmd/ohos/callback_types.h`:新增 `#include <stdbool.h>`;3 个回调 typedef + 3 个 C 包装函数 `int` → `bool`
- `cmd/ohos/callback_bridge.go`:`OnSyncServerStart/Finish/Failed` 改用 `C._Bool(reinstalled)`
- `cmd/ohos/conversation_msg_c.go`:`SendMessage`/`SendMessageNotOss`/`ChangeInputStates` 布尔参数 `C.int` → `C._Bool`
- `cmd/ohos/group_c.go`:`ChangeGroupMute` 布尔参数 `C.int` → `C._Bool`
- `cmd/ohos/init_login_c.go`:`SetAppBackgroundStatus` 布尔参数 `C.int` → `C._Bool`;`OpenIM_InitSDK` 返回 `C.int` → `C._Bool`
- `cmd/ohos/relation_c.go`:3 个 `filterBlack` 参数 `C.int` → `C._Bool`

重新生成头文件命令:`make ohos`(Windows: `make ohos` 调用 `scripts/build-ohos.ps1`;Linux: 调用 `scripts/build-ohos.sh`)。
