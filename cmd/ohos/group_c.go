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

// ===================== 群组创建与成员加入 API =====================

// OpenIM_CreateGroup 创建群组（异步）
//
//export OpenIM_CreateGroup
func OpenIM_CreateGroup(cb *C.base_callback_t, operationID, groupReqInfo *C.char) {
	// 将 C 回调转换为 Go 回调，并将 C 字符串转换为 Go 字符串后调用 SDK
	open_im_sdk.CreateGroup(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupReqInfo))
}

// OpenIM_JoinGroup 加入群组（异步）
//
//export OpenIM_JoinGroup
func OpenIM_JoinGroup(cb *C.base_callback_t, operationID, groupID, reqMsg *C.char, joinSource C.int32_t, ex *C.char) {
	// joinSource 为入群来源（搜索/邀请等），ex 为扩展字段
	open_im_sdk.JoinGroup(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupID), C.GoString(reqMsg), int32(joinSource), C.GoString(ex))
}

// OpenIM_QuitGroup 退出群组（异步）
//
//export OpenIM_QuitGroup
func OpenIM_QuitGroup(cb *C.base_callback_t, operationID, groupID *C.char) {
	open_im_sdk.QuitGroup(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupID))
}

// OpenIM_DismissGroup 解散群组（异步）
//
//export OpenIM_DismissGroup
func OpenIM_DismissGroup(cb *C.base_callback_t, operationID, groupID *C.char) {
	open_im_sdk.DismissGroup(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupID))
}

// ===================== 群组信息与成员管理 API =====================

// OpenIM_ChangeGroupMute 修改群组禁言状态（异步）
//
//export OpenIM_ChangeGroupMute
func OpenIM_ChangeGroupMute(cb *C.base_callback_t, operationID, groupID *C.char, isMute C.int) {
	// isMute: 1=禁言, 0=取消禁言，需转换为 Go bool 类型
	open_im_sdk.ChangeGroupMute(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupID), isMute != 0)
}

// OpenIM_ChangeGroupMemberMute 修改群成员禁言时长（异步）
//
//export OpenIM_ChangeGroupMemberMute
func OpenIM_ChangeGroupMemberMute(cb *C.base_callback_t, operationID, groupID, userID *C.char, mutedSeconds C.int) {
	// mutedSeconds 为禁言秒数，0 表示取消禁言
	open_im_sdk.ChangeGroupMemberMute(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupID), C.GoString(userID), int(mutedSeconds))
}

// OpenIM_TransferGroupOwner 转让群主（异步）
//
//export OpenIM_TransferGroupOwner
func OpenIM_TransferGroupOwner(cb *C.base_callback_t, operationID, groupID, newOwnerUserID *C.char) {
	open_im_sdk.TransferGroupOwner(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupID), C.GoString(newOwnerUserID))
}

// OpenIM_KickGroupMember 踢出群成员（异步）
//
//export OpenIM_KickGroupMember
func OpenIM_KickGroupMember(cb *C.base_callback_t, operationID, groupID, reason, userIDList *C.char) {
	// userIDList 为 JSON 格式的用户 ID 列表
	open_im_sdk.KickGroupMember(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupID), C.GoString(reason), C.GoString(userIDList))
}

// OpenIM_SetGroupInfo 设置群组信息（异步）
//
//export OpenIM_SetGroupInfo
func OpenIM_SetGroupInfo(cb *C.base_callback_t, operationID, groupInfo *C.char) {
	// groupInfo 为 JSON 格式的群组信息
	open_im_sdk.SetGroupInfo(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupInfo))
}

// OpenIM_SetGroupMemberInfo 设置群成员信息（异步）
//
//export OpenIM_SetGroupMemberInfo
func OpenIM_SetGroupMemberInfo(cb *C.base_callback_t, operationID, groupMemberInfo *C.char) {
	// groupMemberInfo 为 JSON 格式的群成员信息
	open_im_sdk.SetGroupMemberInfo(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupMemberInfo))
}

// ===================== 群组列表与信息查询 API =====================

// OpenIM_GetJoinedGroupList 获取已加入的群组列表（异步）
//
//export OpenIM_GetJoinedGroupList
func OpenIM_GetJoinedGroupList(cb *C.base_callback_t, operationID *C.char) {
	open_im_sdk.GetJoinedGroupList(newBaseCallback(cb), C.GoString(operationID))
}

// OpenIM_GetJoinedGroupListPage 分页获取已加入的群组列表（异步）
//
//export OpenIM_GetJoinedGroupListPage
func OpenIM_GetJoinedGroupListPage(cb *C.base_callback_t, operationID *C.char, offset, count C.int32_t) {
	open_im_sdk.GetJoinedGroupListPage(newBaseCallback(cb), C.GoString(operationID), int32(offset), int32(count))
}

// OpenIM_GetSpecifiedGroupsInfo 获取指定群组信息（异步）
//
//export OpenIM_GetSpecifiedGroupsInfo
func OpenIM_GetSpecifiedGroupsInfo(cb *C.base_callback_t, operationID, groupIDList *C.char) {
	// groupIDList 为 JSON 格式的群 ID 列表
	open_im_sdk.GetSpecifiedGroupsInfo(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupIDList))
}

// OpenIM_SearchGroups 搜索群组（异步）
//
//export OpenIM_SearchGroups
func OpenIM_SearchGroups(cb *C.base_callback_t, operationID, searchParam *C.char) {
	// searchParam 为 JSON 格式的搜索参数
	open_im_sdk.SearchGroups(newBaseCallback(cb), C.GoString(operationID), C.GoString(searchParam))
}

// ===================== 群成员查询 API =====================

// OpenIM_GetGroupMemberOwnerAndAdmin 获取群主和管理员列表（异步）
//
//export OpenIM_GetGroupMemberOwnerAndAdmin
func OpenIM_GetGroupMemberOwnerAndAdmin(cb *C.base_callback_t, operationID, groupID *C.char) {
	open_im_sdk.GetGroupMemberOwnerAndAdmin(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupID))
}

// OpenIM_GetGroupMemberListByJoinTimeFilter 按入群时间过滤获取群成员列表（异步）
//
//export OpenIM_GetGroupMemberListByJoinTimeFilter
func OpenIM_GetGroupMemberListByJoinTimeFilter(cb *C.base_callback_t, operationID, groupID *C.char, offset, count C.int32_t, joinTimeBegin, joinTimeEnd C.int64_t, filterUserIDList *C.char) {
	// joinTimeBegin/joinTimeEnd 为入群时间范围（毫秒），filterUserIDList 为需过滤的用户 ID 列表
	open_im_sdk.GetGroupMemberListByJoinTimeFilter(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupID), int32(offset), int32(count), int64(joinTimeBegin), int64(joinTimeEnd), C.GoString(filterUserIDList))
}

// OpenIM_GetSpecifiedGroupMembersInfo 获取指定群成员信息（异步）
//
//export OpenIM_GetSpecifiedGroupMembersInfo
func OpenIM_GetSpecifiedGroupMembersInfo(cb *C.base_callback_t, operationID, groupID, userIDList *C.char) {
	// userIDList 为 JSON 格式的用户 ID 列表
	open_im_sdk.GetSpecifiedGroupMembersInfo(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupID), C.GoString(userIDList))
}

// OpenIM_GetGroupMemberList 获取群成员列表（异步）
//
//export OpenIM_GetGroupMemberList
func OpenIM_GetGroupMemberList(cb *C.base_callback_t, operationID, groupID *C.char, filter, offset, count C.int32_t) {
	// filter 为成员过滤类型（0=全部, 1=群主, 2=管理员, 3=普通成员等）
	open_im_sdk.GetGroupMemberList(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupID), int32(filter), int32(offset), int32(count))
}

// OpenIM_SearchGroupMembers 搜索群成员（异步）
//
//export OpenIM_SearchGroupMembers
func OpenIM_SearchGroupMembers(cb *C.base_callback_t, operationID, searchParam *C.char) {
	// searchParam 为 JSON 格式的搜索参数
	open_im_sdk.SearchGroupMembers(newBaseCallback(cb), C.GoString(operationID), C.GoString(searchParam))
}

// OpenIM_IsJoinGroup 判断是否已加入群组（异步）
//
//export OpenIM_IsJoinGroup
func OpenIM_IsJoinGroup(cb *C.base_callback_t, operationID, groupID *C.char) {
	open_im_sdk.IsJoinGroup(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupID))
}

// OpenIM_GetUsersInGroup 查询用户是否在群组中（异步）
//
//export OpenIM_GetUsersInGroup
func OpenIM_GetUsersInGroup(cb *C.base_callback_t, operationID, groupID, userIDList *C.char) {
	// userIDList 为 JSON 格式的用户 ID 列表
	open_im_sdk.GetUsersInGroup(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupID), C.GoString(userIDList))
}

// ===================== 群申请与邀请管理 API =====================

// OpenIM_GetGroupApplicationListAsRecipient 获取作为接收者的群申请列表（异步）
//
//export OpenIM_GetGroupApplicationListAsRecipient
func OpenIM_GetGroupApplicationListAsRecipient(cb *C.base_callback_t, operationID, req *C.char) {
	// req 为 JSON 格式的请求参数
	open_im_sdk.GetGroupApplicationListAsRecipient(newBaseCallback(cb), C.GoString(operationID), C.GoString(req))
}

// OpenIM_GetGroupApplicationListAsApplicant 获取作为申请者的群申请列表（异步）
//
//export OpenIM_GetGroupApplicationListAsApplicant
func OpenIM_GetGroupApplicationListAsApplicant(cb *C.base_callback_t, operationID, req *C.char) {
	// req 为 JSON 格式的请求参数
	open_im_sdk.GetGroupApplicationListAsApplicant(newBaseCallback(cb), C.GoString(operationID), C.GoString(req))
}

// OpenIM_InviteUserToGroup 邀请用户加入群组（异步）
//
//export OpenIM_InviteUserToGroup
func OpenIM_InviteUserToGroup(cb *C.base_callback_t, operationID, groupID, reason, userIDList *C.char) {
	// userIDList 为 JSON 格式的被邀请用户 ID 列表
	open_im_sdk.InviteUserToGroup(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupID), C.GoString(reason), C.GoString(userIDList))
}

// OpenIM_AcceptGroupApplication 接受群申请（异步）
//
//export OpenIM_AcceptGroupApplication
func OpenIM_AcceptGroupApplication(cb *C.base_callback_t, operationID, groupID, fromUserID, handleMsg *C.char) {
	// fromUserID 为申请者用户 ID，handleMsg 为处理留言
	open_im_sdk.AcceptGroupApplication(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupID), C.GoString(fromUserID), C.GoString(handleMsg))
}

// OpenIM_RefuseGroupApplication 拒绝群申请（异步）
//
//export OpenIM_RefuseGroupApplication
func OpenIM_RefuseGroupApplication(cb *C.base_callback_t, operationID, groupID, fromUserID, handleMsg *C.char) {
	// fromUserID 为申请者用户 ID，handleMsg 为拒绝留言
	open_im_sdk.RefuseGroupApplication(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupID), C.GoString(fromUserID), C.GoString(handleMsg))
}

// OpenIM_GetGroupApplicationUnhandledCount 获取群申请未处理数量（异步）
//
//export OpenIM_GetGroupApplicationUnhandledCount
func OpenIM_GetGroupApplicationUnhandledCount(cb *C.base_callback_t, operationID, req *C.char) {
	// req 为 JSON 格式的请求参数
	open_im_sdk.GetGroupApplicationUnhandledCount(newBaseCallback(cb), C.GoString(operationID), C.GoString(req))
}

// ===================== 群组同步检查 API =====================

// OpenIM_CheckLocalGroupFullSync 检查本地群组全量同步（异步）
//
//export OpenIM_CheckLocalGroupFullSync
func OpenIM_CheckLocalGroupFullSync(cb *C.base_callback_t, operationID *C.char) {
	open_im_sdk.CheckLocalGroupFullSync(newBaseCallback(cb), C.GoString(operationID))
}

// OpenIM_CheckGroupMemberFullSync 检查群成员全量同步（异步）
//
//export OpenIM_CheckGroupMemberFullSync
func OpenIM_CheckGroupMemberFullSync(cb *C.base_callback_t, operationID, groupID *C.char) {
	open_im_sdk.CheckGroupMemberFullSync(newBaseCallback(cb), C.GoString(operationID), C.GoString(groupID))
}
