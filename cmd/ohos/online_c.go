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

// OpenIM_SubscribeUsersStatus 订阅用户在线状态（异步）
//
//export OpenIM_SubscribeUsersStatus
func OpenIM_SubscribeUsersStatus(cb *C.base_callback_t, operationID, userIDs *C.char) {
	open_im_sdk.SubscribeUsersStatus(newBaseCallback(cb), C.GoString(operationID), C.GoString(userIDs))
}

// OpenIM_UnsubscribeUsersStatus 取消订阅用户在线状态（异步）
//
//export OpenIM_UnsubscribeUsersStatus
func OpenIM_UnsubscribeUsersStatus(cb *C.base_callback_t, operationID, userIDs *C.char) {
	open_im_sdk.UnsubscribeUsersStatus(newBaseCallback(cb), C.GoString(operationID), C.GoString(userIDs))
}

// OpenIM_GetSubscribeUsersStatus 获取已订阅用户在线状态（异步）
//
//export OpenIM_GetSubscribeUsersStatus
func OpenIM_GetSubscribeUsersStatus(cb *C.base_callback_t, operationID *C.char) {
	open_im_sdk.GetSubscribeUsersStatus(newBaseCallback(cb), C.GoString(operationID))
}

// OpenIM_GetUserStatus 获取用户在线状态（异步）
//
//export OpenIM_GetUserStatus
func OpenIM_GetUserStatus(cb *C.base_callback_t, operationID, userIDs *C.char) {
	open_im_sdk.GetUserStatus(newBaseCallback(cb), C.GoString(operationID), C.GoString(userIDs))
}
