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

// ===================== 好友信息查询 API =====================

// OpenIM_GetSpecifiedFriendsInfo 获取指定好友信息（异步）
//
//export OpenIM_GetSpecifiedFriendsInfo
func OpenIM_GetSpecifiedFriendsInfo(cb *C.base_callback_t, operationID, userIDList *C.char, filterBlack C.int) {
	// userIDList 为 JSON 格式的用户 ID 列表，filterBlack 为是否过滤黑名单用户（1=过滤, 0=不过滤）
	open_im_sdk.GetSpecifiedFriendsInfo(newBaseCallback(cb), C.GoString(operationID), C.GoString(userIDList), filterBlack != 0)
}

// OpenIM_GetFriendList 获取好友列表（异步）
//
//export OpenIM_GetFriendList
func OpenIM_GetFriendList(cb *C.base_callback_t, operationID *C.char, filterBlack C.int) {
	// filterBlack 为是否过滤黑名单用户（1=过滤, 0=不过滤）
	open_im_sdk.GetFriendList(newBaseCallback(cb), C.GoString(operationID), filterBlack != 0)
}

// OpenIM_GetFriendListPage 分页获取好友列表（异步）
//
//export OpenIM_GetFriendListPage
func OpenIM_GetFriendListPage(cb *C.base_callback_t, operationID *C.char, offset, count C.int32_t, filterBlack C.int) {
	// offset 为分页偏移量，count 为每页数量，filterBlack 为是否过滤黑名单用户
	open_im_sdk.GetFriendListPage(newBaseCallback(cb), C.GoString(operationID), int32(offset), int32(count), filterBlack != 0)
}

// OpenIM_SearchFriends 搜索好友（异步）
//
//export OpenIM_SearchFriends
func OpenIM_SearchFriends(cb *C.base_callback_t, operationID, searchParam *C.char) {
	// searchParam 为 JSON 格式的搜索参数
	open_im_sdk.SearchFriends(newBaseCallback(cb), C.GoString(operationID), C.GoString(searchParam))
}

// OpenIM_CheckFriend 检查好友关系（异步）
//
//export OpenIM_CheckFriend
func OpenIM_CheckFriend(cb *C.base_callback_t, operationID, userIDList *C.char) {
	// userIDList 为 JSON 格式的用户 ID 列表
	open_im_sdk.CheckFriend(newBaseCallback(cb), C.GoString(operationID), C.GoString(userIDList))
}

// ===================== 好友操作 API =====================

// OpenIM_AddFriend 添加好友（异步）
//
//export OpenIM_AddFriend
func OpenIM_AddFriend(cb *C.base_callback_t, operationID, userIDReqMsg *C.char) {
	// userIDReqMsg 为 JSON 格式的添加好友请求（含目标用户 ID 和验证消息）
	open_im_sdk.AddFriend(newBaseCallback(cb), C.GoString(operationID), C.GoString(userIDReqMsg))
}

// OpenIM_UpdateFriends 更新好友信息（异步）
//
//export OpenIM_UpdateFriends
func OpenIM_UpdateFriends(cb *C.base_callback_t, operationID, req *C.char) {
	// req 为 JSON 格式的更新好友请求参数
	open_im_sdk.UpdateFriends(newBaseCallback(cb), C.GoString(operationID), C.GoString(req))
}

// OpenIM_DeleteFriend 删除好友（异步）
//
//export OpenIM_DeleteFriend
func OpenIM_DeleteFriend(cb *C.base_callback_t, operationID, friendUserID *C.char) {
	// friendUserID 为待删除好友的用户 ID
	open_im_sdk.DeleteFriend(newBaseCallback(cb), C.GoString(operationID), C.GoString(friendUserID))
}

// ===================== 好友申请管理 API =====================

// OpenIM_GetFriendApplicationListAsRecipient 获取作为接收者的好友申请列表（异步）
//
//export OpenIM_GetFriendApplicationListAsRecipient
func OpenIM_GetFriendApplicationListAsRecipient(cb *C.base_callback_t, operationID, req *C.char) {
	// req 为 JSON 格式的请求参数
	open_im_sdk.GetFriendApplicationListAsRecipient(newBaseCallback(cb), C.GoString(operationID), C.GoString(req))
}

// OpenIM_GetFriendApplicationListAsApplicant 获取作为申请者的好友申请列表（异步）
//
//export OpenIM_GetFriendApplicationListAsApplicant
func OpenIM_GetFriendApplicationListAsApplicant(cb *C.base_callback_t, operationID, req *C.char) {
	// req 为 JSON 格式的请求参数
	open_im_sdk.GetFriendApplicationListAsApplicant(newBaseCallback(cb), C.GoString(operationID), C.GoString(req))
}

// OpenIM_AcceptFriendApplication 接受好友申请（异步）
//
//export OpenIM_AcceptFriendApplication
func OpenIM_AcceptFriendApplication(cb *C.base_callback_t, operationID, userIDHandleMsg *C.char) {
	// userIDHandleMsg 为 JSON 格式的处理参数（含申请者用户 ID 和处理留言）
	open_im_sdk.AcceptFriendApplication(newBaseCallback(cb), C.GoString(operationID), C.GoString(userIDHandleMsg))
}

// OpenIM_RefuseFriendApplication 拒绝好友申请（异步）
//
//export OpenIM_RefuseFriendApplication
func OpenIM_RefuseFriendApplication(cb *C.base_callback_t, operationID, userIDHandleMsg *C.char) {
	// userIDHandleMsg 为 JSON 格式的处理参数（含申请者用户 ID 和拒绝留言）
	open_im_sdk.RefuseFriendApplication(newBaseCallback(cb), C.GoString(operationID), C.GoString(userIDHandleMsg))
}

// OpenIM_GetFriendApplicationUnhandledCount 获取好友申请未处理数量（异步）
//
//export OpenIM_GetFriendApplicationUnhandledCount
func OpenIM_GetFriendApplicationUnhandledCount(cb *C.base_callback_t, operationID, req *C.char) {
	// req 为 JSON 格式的请求参数
	open_im_sdk.GetFriendApplicationUnhandledCount(newBaseCallback(cb), C.GoString(operationID), C.GoString(req))
}

// ===================== 黑名单管理 API =====================

// OpenIM_AddBlack 拉黑用户（异步）
//
//export OpenIM_AddBlack
func OpenIM_AddBlack(cb *C.base_callback_t, operationID, blackUserID, ex *C.char) {
	// blackUserID 为被拉黑的用户 ID，ex 为扩展字段
	open_im_sdk.AddBlack(newBaseCallback(cb), C.GoString(operationID), C.GoString(blackUserID), C.GoString(ex))
}

// OpenIM_GetBlackList 获取黑名单列表（异步）
//
//export OpenIM_GetBlackList
func OpenIM_GetBlackList(cb *C.base_callback_t, operationID *C.char) {
	open_im_sdk.GetBlackList(newBaseCallback(cb), C.GoString(operationID))
}

// OpenIM_RemoveBlack 移除黑名单（异步）
//
//export OpenIM_RemoveBlack
func OpenIM_RemoveBlack(cb *C.base_callback_t, operationID, removeUserID *C.char) {
	// removeUserID 为需移出黑名单的用户 ID
	open_im_sdk.RemoveBlack(newBaseCallback(cb), C.GoString(operationID), C.GoString(removeUserID))
}
