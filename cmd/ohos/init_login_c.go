// Copyright © 2023 OpenIM SDK. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

/*
#include "callback_types.h"
*/
import "C"

import (
	"github.com/openimsdk/openim-sdk-core/v3/open_im_sdk"
)

// OpenIM_GetSdkVersion 获取 SDK 版本号
// 返回: C 字符串（需调用 OpenIM_FreeString 释放）
//
//export OpenIM_GetSdkVersion
func OpenIM_GetSdkVersion() *C.char {
	return C.CString(open_im_sdk.GetSdkVersion())
}

// OpenIM_InitSDK 初始化 SDK
// 参数:
//   - connListener: 连接监听器回调结构体指针
//   - operationID: 操作ID，用于链路追踪
//   - config: JSON 格式的配置字符串
//
// 返回: 1 表示成功，0 表示失败
//
//export OpenIM_InitSDK
func OpenIM_InitSDK(connListener *C.conn_listener_t, operationID, config *C.char) C.int {
	// 将 C 回调结构体包装为 Go 接口实现
	listener := newConnListener(connListener)
	// 调用 SDK 初始化
	result := open_im_sdk.InitSDK(listener, C.GoString(operationID), C.GoString(config))
	if result {
		return 1
	}
	return 0
}

// OpenIM_UnInitSDK 反初始化 SDK，释放资源
// 参数:
//   - operationID: 操作ID
//
//export OpenIM_UnInitSDK
func OpenIM_UnInitSDK(operationID *C.char) {
	open_im_sdk.UnInitSDK(C.GoString(operationID))
}

// OpenIM_Login 登录（异步，结果通过 callback 返回）
// 参数:
//   - cb: 基础回调结构体指针
//   - operationID: 操作ID
//   - userID: 用户ID
//   - token: 登录令牌
//
//export OpenIM_Login
func OpenIM_Login(cb *C.base_callback_t, operationID, userID, token *C.char) {
	callback := newBaseCallback(cb)
	open_im_sdk.Login(callback, C.GoString(operationID), C.GoString(userID), C.GoString(token))
}

// OpenIM_Logout 登出（异步，结果通过 callback 返回）
//
//export OpenIM_Logout
func OpenIM_Logout(cb *C.base_callback_t, operationID *C.char) {
	callback := newBaseCallback(cb)
	open_im_sdk.Logout(callback, C.GoString(operationID))
}

// OpenIM_GetLoginStatus 获取登录状态（同步）
// 返回: 1=已登出, 2=登录中, 3=已登录
//
//export OpenIM_GetLoginStatus
func OpenIM_GetLoginStatus(operationID *C.char) C.int {
	return C.int(open_im_sdk.GetLoginStatus(C.GoString(operationID)))
}

// OpenIM_GetLoginUserID 获取当前登录用户ID（同步）
// 返回: C 字符串（需调用 OpenIM_FreeString 释放）
//
//export OpenIM_GetLoginUserID
func OpenIM_GetLoginUserID() *C.char {
	return C.CString(open_im_sdk.GetLoginUserID())
}

// OpenIM_SetAppBackgroundStatus 设置 App 前后台状态（异步）
// 参数:
//   - isBackground: 1 表示后台，0 表示前台
//
//export OpenIM_SetAppBackgroundStatus
func OpenIM_SetAppBackgroundStatus(cb *C.base_callback_t, operationID *C.char, isBackground C.int) {
	callback := newBaseCallback(cb)
	open_im_sdk.SetAppBackgroundStatus(callback, C.GoString(operationID), isBackground != 0)
}

// OpenIM_NetworkStatusChanged 通知网络状态变化（异步）
//
//export OpenIM_NetworkStatusChanged
func OpenIM_NetworkStatusChanged(cb *C.base_callback_t, operationID *C.char) {
	callback := newBaseCallback(cb)
	open_im_sdk.NetworkStatusChanged(callback, C.GoString(operationID))
}
