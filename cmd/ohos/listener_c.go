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

// OpenIM_SetGroupListener 设置群组事件监听器
// 参数:
//   - listener: 群组监听器回调结构体指针
//
//export OpenIM_SetGroupListener
func OpenIM_SetGroupListener(listener *C.group_listener_t) {
	open_im_sdk.SetGroupListener(newGroupListener(listener))
}

// OpenIM_SetConversationListener 设置会话事件监听器
// 参数:
//   - listener: 会话监听器回调结构体指针
//
//export OpenIM_SetConversationListener
func OpenIM_SetConversationListener(listener *C.conversation_listener_t) {
	open_im_sdk.SetConversationListener(newConversationListener(listener))
}

// OpenIM_SetAdvancedMsgListener 设置高级消息事件监听器
// 参数:
//   - listener: 消息监听器回调结构体指针
//
//export OpenIM_SetAdvancedMsgListener
func OpenIM_SetAdvancedMsgListener(listener *C.advanced_msg_listener_t) {
	open_im_sdk.SetAdvancedMsgListener(newAdvancedMsgListener(listener))
}

// OpenIM_SetUserListener 设置用户事件监听器
// 参数:
//   - listener: 用户监听器回调结构体指针
//
//export OpenIM_SetUserListener
func OpenIM_SetUserListener(listener *C.user_listener_t) {
	open_im_sdk.SetUserListener(newUserListener(listener))
}

// OpenIM_SetFriendListener 设置好友事件监听器
// 参数:
//   - listener: 好友监听器回调结构体指针
//
//export OpenIM_SetFriendListener
func OpenIM_SetFriendListener(listener *C.friendship_listener_t) {
	open_im_sdk.SetFriendListener(newFriendshipListener(listener))
}

// OpenIM_SetCustomBusinessListener 设置自定义业务事件监听器
// 参数:
//   - listener: 自定义业务监听器回调结构体指针
//
//export OpenIM_SetCustomBusinessListener
func OpenIM_SetCustomBusinessListener(listener *C.custom_business_listener_t) {
	open_im_sdk.SetCustomBusinessListener(newCustomBusinessListener(listener))
}

// OpenIM_SetMessageKvInfoListener 设置消息 KV 信息变更监听器
// 参数:
//   - listener: 消息 KV 监听器回调结构体指针
//
//export OpenIM_SetMessageKvInfoListener
func OpenIM_SetMessageKvInfoListener(listener *C.message_kv_info_listener_t) {
	open_im_sdk.SetMessageKvInfoListener(newMessageKvInfoListener(listener))
}
