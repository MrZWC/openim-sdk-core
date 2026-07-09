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
#include <stdint.h>
*/
import "C"

import (
	"github.com/openimsdk/openim-sdk-core/v3/open_im_sdk"
)

// OpenIM_GetUsersInfo 获取用户信息（异步）
//
//export OpenIM_GetUsersInfo
func OpenIM_GetUsersInfo(cb *C.base_callback_t, operationID, userIDs *C.char) {
	open_im_sdk.GetUsersInfo(newBaseCallback(cb), C.GoString(operationID), C.GoString(userIDs))
}

// OpenIM_SetSelfInfo 设置自身用户信息（异步）
//
//export OpenIM_SetSelfInfo
func OpenIM_SetSelfInfo(cb *C.base_callback_t, operationID, userInfo *C.char) {
	open_im_sdk.SetSelfInfo(newBaseCallback(cb), C.GoString(operationID), C.GoString(userInfo))
}

// OpenIM_GetSelfUserInfo 获取自身用户信息（异步）
//
//export OpenIM_GetSelfUserInfo
func OpenIM_GetSelfUserInfo(cb *C.base_callback_t, operationID *C.char) {
	open_im_sdk.GetSelfUserInfo(newBaseCallback(cb), C.GoString(operationID))
}

// OpenIM_GetUserClientConfig 获取用户客户端配置（异步）
//
//export OpenIM_GetUserClientConfig
func OpenIM_GetUserClientConfig(cb *C.base_callback_t, operationID *C.char) {
	open_im_sdk.GetUserClientConfig(newBaseCallback(cb), C.GoString(operationID))
}
