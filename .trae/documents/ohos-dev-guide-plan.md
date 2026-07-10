# 新建鸿蒙开发流程文档计划

## 概述

新建一份独立的鸿蒙开发流程文档，涵盖从编译产物集成到完整业务开发的全流程，作为 `ohos-build-setup-cn.md`（编译指南）的姊妹篇。

## 修改文件

**新建**：`docs/ohos-dev-guide-cn.md`

文档结构：

1. **概述**：SDK 架构、编译产物说明
2. **项目集成**：目录结构、CMakeLists.txt 配置、ABI 设置
3. **NAPI 桥接层**：C++ <-> ArkTS 数据转换、异步回调传递到 ArkTS 的机制（napi_threadsafe_function）
4. **完整开发示例**：
   - 初始化 SDK（conn_listener_t + config JSON）
   - 登录/登出
   - 设置监听器（消息/会话/群组/好友）
   - 发送消息（send_msg_callback_t 含进度）
   - 获取会话列表和消息历史
   - 内存管理（OpenIM_FreeString）
5. **API 参考速查表**：按模块列出所有导出函数签名
6. **线程安全**：Go 回调线程、NAPI 主线程、线程间通信
7. **常见问题**

## 验证

检查文档中所有 C 函数签名与 `callback_types.h` 和 `//export` 一致，config JSON 字段与 `sdk_struct.go` 一致。
