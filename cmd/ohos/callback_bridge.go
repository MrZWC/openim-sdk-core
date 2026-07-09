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
	"unsafe"

	"github.com/openimsdk/openim-sdk-core/v3/open_im_sdk_callback"
)

// ===================== Go 回调桥接实现 =====================
// 以下结构体将 C 函数指针包装为 Go 接口实现，使 Go SDK 能通过 C 回调通知鸿蒙侧。

// cString 是辅助函数，将 Go string 转换为 C 字符串并返回释放函数
func cString(s string) (*C.char, func()) {
	if s == "" {
		return nil, func() {}
	}
	cs := C.CString(s)
	return cs, func() { C.free(unsafe.Pointer(cs)) }
}

// ---- Base 回调桥接 ----

// baseCallback 实现 open_im_sdk_callback.Base 接口
type baseCallback struct {
	cb *C.base_callback_t
}

// newBaseCallback 创建 Base 回调桥接对象
func newBaseCallback(cb *C.base_callback_t) *baseCallback {
	return &baseCallback{cb: cb}
}

func (b *baseCallback) OnSuccess(data string) {
	if b.cb == nil {
		return
	}
	cData, free := cString(data)
	defer free()
	C.base_cb_on_success(b.cb, cData)
}

func (b *baseCallback) OnError(errCode int32, errMsg string) {
	if b.cb == nil {
		return
	}
	cMsg, free := cString(errMsg)
	defer free()
	C.base_cb_on_error(b.cb, C.int32_t(errCode), cMsg)
}

// ---- SendMsgCallBack 回调桥接 ----

// sendMsgCallback 实现 open_im_sdk_callback.SendMsgCallBack 接口
type sendMsgCallback struct {
	cb *C.send_msg_callback_t
}

// newSendMsgCallback 创建 SendMsgCallBack 回调桥接对象
func newSendMsgCallback(cb *C.send_msg_callback_t) *sendMsgCallback {
	return &sendMsgCallback{cb: cb}
}

func (s *sendMsgCallback) OnSuccess(data string) {
	if s.cb == nil {
		return
	}
	cData, free := cString(data)
	defer free()
	C.base_cb_on_success((*C.base_callback_t)(unsafe.Pointer(s.cb)), cData)
}

func (s *sendMsgCallback) OnError(errCode int32, errMsg string) {
	if s.cb == nil {
		return
	}
	cMsg, free := cString(errMsg)
	defer free()
	C.base_cb_on_error((*C.base_callback_t)(unsafe.Pointer(s.cb)), C.int32_t(errCode), cMsg)
}

func (s *sendMsgCallback) OnProgress(progress int) {
	if s.cb == nil {
		return
	}
	C.send_msg_cb_on_progress(s.cb, C.int(progress))
}

// ---- OnConnListener 回调桥接 ----

// connListener 实现 open_im_sdk_callback.OnConnListener 接口
type connListener struct {
	cb *C.conn_listener_t
}

// newConnListener 创建 OnConnListener 回调桥接对象
func newConnListener(cb *C.conn_listener_t) *connListener {
	return &connListener{cb: cb}
}

func (l *connListener) OnConnecting() {
	if l.cb == nil {
		return
	}
	C.conn_cb_on_connecting(l.cb)
}

func (l *connListener) OnConnectSuccess() {
	if l.cb == nil {
		return
	}
	C.conn_cb_on_connect_success(l.cb)
}

func (l *connListener) OnConnectFailed(errCode int32, errMsg string) {
	if l.cb == nil {
		return
	}
	cMsg, free := cString(errMsg)
	defer free()
	C.conn_cb_on_connect_failed(l.cb, C.int32_t(errCode), cMsg)
}

func (l *connListener) OnKickedOffline() {
	if l.cb == nil {
		return
	}
	C.conn_cb_on_kicked_offline(l.cb)
}

func (l *connListener) OnUserTokenExpired() {
	if l.cb == nil {
		return
	}
	C.conn_cb_on_user_token_expired(l.cb)
}

func (l *connListener) OnUserTokenInvalid(errMsg string) {
	if l.cb == nil {
		return
	}
	cMsg, free := cString(errMsg)
	defer free()
	C.conn_cb_on_user_token_invalid(l.cb, cMsg)
}

// ---- OnConversationListener 回调桥接 ----

// conversationListener 实现 open_im_sdk_callback.OnConversationListener 接口
type conversationListener struct {
	cb *C.conversation_listener_t
}

func newConversationListener(cb *C.conversation_listener_t) *conversationListener {
	return &conversationListener{cb: cb}
}

func (l *conversationListener) OnSyncServerStart(reinstalled bool) {
	if l.cb == nil {
		return
	}
	r := 0
	if reinstalled {
		r = 1
	}
	C.conv_cb_on_sync_server_start(l.cb, C.int(r))
}

func (l *conversationListener) OnSyncServerFinish(reinstalled bool) {
	if l.cb == nil {
		return
	}
	r := 0
	if reinstalled {
		r = 1
	}
	C.conv_cb_on_sync_server_finish(l.cb, C.int(r))
}

func (l *conversationListener) OnSyncServerProgress(progress int) {
	if l.cb == nil {
		return
	}
	C.conv_cb_on_sync_server_progress(l.cb, C.int(progress))
}

func (l *conversationListener) OnSyncServerFailed(reinstalled bool) {
	if l.cb == nil {
		return
	}
	r := 0
	if reinstalled {
		r = 1
	}
	C.conv_cb_on_sync_server_failed(l.cb, C.int(r))
}

func (l *conversationListener) OnNewConversation(conversationList string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(conversationList)
	defer free()
	C.conv_cb_on_new_conversation(l.cb, cData)
}

func (l *conversationListener) OnConversationChanged(conversationList string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(conversationList)
	defer free()
	C.conv_cb_on_conversation_changed(l.cb, cData)
}

func (l *conversationListener) OnTotalUnreadMessageCountChanged(totalUnreadCount int32) {
	if l.cb == nil {
		return
	}
	C.conv_cb_on_total_unread_count_changed(l.cb, C.int32_t(totalUnreadCount))
}

func (l *conversationListener) OnConversationUserInputStatusChanged(change string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(change)
	defer free()
	C.conv_cb_on_conversation_user_input_status_changed(l.cb, cData)
}

// ---- OnAdvancedMsgListener 回调桥接 ----

// advancedMsgListener 实现 open_im_sdk_callback.OnAdvancedMsgListener 接口
type advancedMsgListener struct {
	cb *C.advanced_msg_listener_t
}

func newAdvancedMsgListener(cb *C.advanced_msg_listener_t) *advancedMsgListener {
	return &advancedMsgListener{cb: cb}
}

func (l *advancedMsgListener) OnRecvNewMessage(message string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(message)
	defer free()
	C.msg_cb_on_recv_new_message(l.cb, cData)
}

func (l *advancedMsgListener) OnRecvC2CReadReceipt(msgReceiptList string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(msgReceiptList)
	defer free()
	C.msg_cb_on_recv_c2c_read_receipt(l.cb, cData)
}

func (l *advancedMsgListener) OnNewRecvMessageRevoked(messageRevoked string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(messageRevoked)
	defer free()
	C.msg_cb_on_new_recv_message_revoked(l.cb, cData)
}

func (l *advancedMsgListener) OnRecvOfflineNewMessage(message string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(message)
	defer free()
	C.msg_cb_on_recv_offline_new_message(l.cb, cData)
}

func (l *advancedMsgListener) OnMsgDeleted(message string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(message)
	defer free()
	C.msg_cb_on_msg_deleted(l.cb, cData)
}

func (l *advancedMsgListener) OnRecvOnlineOnlyMessage(message string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(message)
	defer free()
	C.msg_cb_on_recv_online_only_message(l.cb, cData)
}

// ---- OnFriendshipListener 回调桥接 ----

// friendshipListener 实现 open_im_sdk_callback.OnFriendshipListener 接口
type friendshipListener struct {
	cb *C.friendship_listener_t
}

func newFriendshipListener(cb *C.friendship_listener_t) *friendshipListener {
	return &friendshipListener{cb: cb}
}

func (l *friendshipListener) OnFriendApplicationAdded(friendApplication string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(friendApplication)
	defer free()
	C.friend_cb_on_application_added(l.cb, cData)
}

func (l *friendshipListener) OnFriendApplicationDeleted(friendApplication string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(friendApplication)
	defer free()
	C.friend_cb_on_application_deleted(l.cb, cData)
}

func (l *friendshipListener) OnFriendApplicationAccepted(friendApplication string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(friendApplication)
	defer free()
	C.friend_cb_on_application_accepted(l.cb, cData)
}

func (l *friendshipListener) OnFriendApplicationRejected(friendApplication string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(friendApplication)
	defer free()
	C.friend_cb_on_application_rejected(l.cb, cData)
}

func (l *friendshipListener) OnFriendAdded(friendInfo string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(friendInfo)
	defer free()
	C.friend_cb_on_friend_added(l.cb, cData)
}

func (l *friendshipListener) OnFriendDeleted(friendInfo string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(friendInfo)
	defer free()
	C.friend_cb_on_friend_deleted(l.cb, cData)
}

func (l *friendshipListener) OnFriendInfoChanged(friendInfo string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(friendInfo)
	defer free()
	C.friend_cb_on_friend_info_changed(l.cb, cData)
}

func (l *friendshipListener) OnBlackAdded(blackInfo string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(blackInfo)
	defer free()
	C.friend_cb_on_black_added(l.cb, cData)
}

func (l *friendshipListener) OnBlackDeleted(blackInfo string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(blackInfo)
	defer free()
	C.friend_cb_on_black_deleted(l.cb, cData)
}

// ---- OnGroupListener 回调桥接 ----

// groupListener 实现 open_im_sdk_callback.OnGroupListener 接口
type groupListener struct {
	cb *C.group_listener_t
}

func newGroupListener(cb *C.group_listener_t) *groupListener {
	return &groupListener{cb: cb}
}

func (l *groupListener) OnJoinedGroupAdded(groupInfo string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(groupInfo)
	defer free()
	C.group_cb_on_joined_group_added(l.cb, cData)
}

func (l *groupListener) OnJoinedGroupDeleted(groupInfo string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(groupInfo)
	defer free()
	C.group_cb_on_joined_group_deleted(l.cb, cData)
}

func (l *groupListener) OnGroupMemberAdded(groupMemberInfo string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(groupMemberInfo)
	defer free()
	C.group_cb_on_group_member_added(l.cb, cData)
}

func (l *groupListener) OnGroupMemberDeleted(groupMemberInfo string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(groupMemberInfo)
	defer free()
	C.group_cb_on_group_member_deleted(l.cb, cData)
}

func (l *groupListener) OnGroupApplicationAdded(groupApplication string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(groupApplication)
	defer free()
	C.group_cb_on_group_application_added(l.cb, cData)
}

func (l *groupListener) OnGroupApplicationDeleted(groupApplication string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(groupApplication)
	defer free()
	C.group_cb_on_group_application_deleted(l.cb, cData)
}

func (l *groupListener) OnGroupInfoChanged(groupInfo string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(groupInfo)
	defer free()
	C.group_cb_on_group_info_changed(l.cb, cData)
}

func (l *groupListener) OnGroupDismissed(groupInfo string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(groupInfo)
	defer free()
	C.group_cb_on_group_dismissed(l.cb, cData)
}

func (l *groupListener) OnGroupMemberInfoChanged(groupMemberInfo string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(groupMemberInfo)
	defer free()
	C.group_cb_on_group_member_info_changed(l.cb, cData)
}

func (l *groupListener) OnGroupApplicationAccepted(groupApplication string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(groupApplication)
	defer free()
	C.group_cb_on_group_application_accepted(l.cb, cData)
}

func (l *groupListener) OnGroupApplicationRejected(groupApplication string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(groupApplication)
	defer free()
	C.group_cb_on_group_application_rejected(l.cb, cData)
}

// ---- OnUserListener 回调桥接 ----

// userListener 实现 open_im_sdk_callback.OnUserListener 接口
type userListener struct {
	cb *C.user_listener_t
}

func newUserListener(cb *C.user_listener_t) *userListener {
	return &userListener{cb: cb}
}

func (l *userListener) OnSelfInfoUpdated(userInfo string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(userInfo)
	defer free()
	C.user_cb_on_self_info_updated(l.cb, cData)
}

func (l *userListener) OnUserStatusChanged(userOnlineStatus string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(userOnlineStatus)
	defer free()
	C.user_cb_on_user_status_changed(l.cb, cData)
}

// ---- OnCustomBusinessListener 回调桥接 ----

// customBusinessListener 实现 open_im_sdk_callback.OnCustomBusinessListener 接口
type customBusinessListener struct {
	cb *C.custom_business_listener_t
}

func newCustomBusinessListener(cb *C.custom_business_listener_t) *customBusinessListener {
	return &customBusinessListener{cb: cb}
}

func (l *customBusinessListener) OnRecvCustomBusinessMessage(businessMessage string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(businessMessage)
	defer free()
	C.business_cb_on_recv(l.cb, cData)
}

// ---- OnMessageKvInfoListener 回调桥接 ----

// messageKvInfoListener 实现 open_im_sdk_callback.OnMessageKvInfoListener 接口
type messageKvInfoListener struct {
	cb *C.message_kv_info_listener_t
}

func newMessageKvInfoListener(cb *C.message_kv_info_listener_t) *messageKvInfoListener {
	return &messageKvInfoListener{cb: cb}
}

func (l *messageKvInfoListener) OnMessageKvInfoChanged(messageChangedList string) {
	if l.cb == nil {
		return
	}
	cData, free := cString(messageChangedList)
	defer free()
	C.msgkv_cb_on_changed(l.cb, cData)
}

// ---- UploadFileCallback 回调桥接 ----

// uploadFileCallback 实现 open_im_sdk_callback.UploadFileCallback 接口
type uploadFileCallback struct {
	cb *C.upload_file_callback_t
}

func newUploadFileCallback(cb *C.upload_file_callback_t) *uploadFileCallback {
	return &uploadFileCallback{cb: cb}
}

func (l *uploadFileCallback) Open(size int64) {
	if l.cb == nil {
		return
	}
	C.upload_file_cb_open(l.cb, C.int64_t(size))
}

func (l *uploadFileCallback) PartSize(partSize int64, num int) {
	if l.cb == nil {
		return
	}
	C.upload_file_cb_part_size(l.cb, C.int64_t(partSize), C.int(num))
}

func (l *uploadFileCallback) HashPartProgress(index int, size int64, partHash string) {
	if l.cb == nil {
		return
	}
	cHash, free := cString(partHash)
	defer free()
	C.upload_file_cb_hash_part_progress(l.cb, C.int(index), C.int64_t(size), cHash)
}

func (l *uploadFileCallback) HashPartComplete(partsHash string, fileHash string) {
	if l.cb == nil {
		return
	}
	cParts, freeParts := cString(partsHash)
	defer freeParts()
	cFile, freeFile := cString(fileHash)
	defer freeFile()
	C.upload_file_cb_hash_part_complete(l.cb, cParts, cFile)
}

func (l *uploadFileCallback) UploadID(uploadID string) {
	if l.cb == nil {
		return
	}
	cID, free := cString(uploadID)
	defer free()
	C.upload_file_cb_upload_id(l.cb, cID)
}

func (l *uploadFileCallback) UploadPartComplete(index int, partSize int64, partHash string) {
	if l.cb == nil {
		return
	}
	cHash, free := cString(partHash)
	defer free()
	C.upload_file_cb_upload_part_complete(l.cb, C.int(index), C.int64_t(partSize), cHash)
}

func (l *uploadFileCallback) UploadComplete(fileSize int64, streamSize int64, storageSize int64) {
	if l.cb == nil {
		return
	}
	C.upload_file_cb_upload_complete(l.cb, C.int64_t(fileSize), C.int64_t(streamSize), C.int64_t(storageSize))
}

func (l *uploadFileCallback) Complete(size int64, url string, typ int) {
	if l.cb == nil {
		return
	}
	cURL, free := cString(url)
	defer free()
	C.upload_file_cb_complete(l.cb, C.int64_t(size), cURL, C.int(typ))
}

// ---- UploadLogProgress 回调桥接 ----

// uploadLogProgress 实现 open_im_sdk_callback.UploadLogProgress 接口
type uploadLogProgress struct {
	cb *C.upload_log_progress_t
}

func newUploadLogProgress(cb *C.upload_log_progress_t) *uploadLogProgress {
	return &uploadLogProgress{cb: cb}
}

func (l *uploadLogProgress) OnProgress(current int64, size int64) {
	if l.cb == nil {
		return
	}
	C.upload_log_cb_on_progress(l.cb, C.int64_t(current), C.int64_t(size))
}

// 确保所有桥接类型实现了对应的接口（编译时检查）
var _ open_im_sdk_callback.Base = (*baseCallback)(nil)
var _ open_im_sdk_callback.SendMsgCallBack = (*sendMsgCallback)(nil)
var _ open_im_sdk_callback.OnConnListener = (*connListener)(nil)
var _ open_im_sdk_callback.OnConversationListener = (*conversationListener)(nil)
var _ open_im_sdk_callback.OnAdvancedMsgListener = (*advancedMsgListener)(nil)
var _ open_im_sdk_callback.OnFriendshipListener = (*friendshipListener)(nil)
var _ open_im_sdk_callback.OnGroupListener = (*groupListener)(nil)
var _ open_im_sdk_callback.OnUserListener = (*userListener)(nil)
var _ open_im_sdk_callback.OnCustomBusinessListener = (*customBusinessListener)(nil)
var _ open_im_sdk_callback.OnMessageKvInfoListener = (*messageKvInfoListener)(nil)
var _ open_im_sdk_callback.UploadFileCallback = (*uploadFileCallback)(nil)
var _ open_im_sdk_callback.UploadLogProgress = (*uploadLogProgress)(nil)
