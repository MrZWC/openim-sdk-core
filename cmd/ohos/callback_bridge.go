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
#include <stdlib.h>

// ===================== C 回调函数指针类型定义 =====================

// Base 回调函数指针
typedef void (*on_success_t)(const char* data);
typedef void (*on_error_t)(int32_t err_code, const char* err_msg);

// Base 回调结构体
typedef struct {
    on_success_t on_success;
    on_error_t   on_error;
} base_callback_t;

// SendMsg 回调函数指针（继承 Base + 进度）
typedef void (*on_progress_t)(int progress);

// SendMsg 回调结构体
typedef struct {
    on_success_t  on_success;
    on_error_t    on_error;
    on_progress_t on_progress;
} send_msg_callback_t;

// OnConnListener 回调函数指针
typedef void (*on_connecting_t)();
typedef void (*on_connect_success_t)();
typedef void (*on_connect_failed_t)(int32_t err_code, const char* err_msg);
typedef void (*on_kicked_offline_t)();
typedef void (*on_user_token_expired_t)();
typedef void (*on_user_token_invalid_t)(const char* err_msg);

// OnConnListener 回调结构体
typedef struct {
    on_connecting_t         on_connecting;
    on_connect_success_t    on_connect_success;
    on_connect_failed_t     on_connect_failed;
    on_kicked_offline_t     on_kicked_offline;
    on_user_token_expired_t on_user_token_expired;
    on_user_token_invalid_t on_user_token_invalid;
} conn_listener_t;

// OnConversationListener 回调函数指针
typedef void (*on_sync_server_start_t)(int reinstalled);
typedef void (*on_sync_server_finish_t)(int reinstalled);
typedef void (*on_sync_server_progress_t)(int progress);
typedef void (*on_sync_server_failed_t)(int reinstalled);
typedef void (*on_new_conversation_t)(const char* conversation_list);
typedef void (*on_conversation_changed_t)(const char* conversation_list);
typedef void (*on_total_unread_count_changed_t)(int32_t total_unread_count);
typedef void (*on_conversation_user_input_status_changed_t)(const char* change);

// OnConversationListener 回调结构体
typedef struct {
    on_sync_server_start_t                      on_sync_server_start;
    on_sync_server_finish_t                     on_sync_server_finish;
    on_sync_server_progress_t                   on_sync_server_progress;
    on_sync_server_failed_t                     on_sync_server_failed;
    on_new_conversation_t                       on_new_conversation;
    on_conversation_changed_t                   on_conversation_changed;
    on_total_unread_count_changed_t             on_total_unread_count_changed;
    on_conversation_user_input_status_changed_t on_conversation_user_input_status_changed;
} conversation_listener_t;

// OnAdvancedMsgListener 回调函数指针
typedef void (*on_recv_new_message_t)(const char* message);
typedef void (*on_recv_c2c_read_receipt_t)(const char* msg_receipt_list);
typedef void (*on_new_recv_message_revoked_t)(const char* message_revoked);
typedef void (*on_recv_offline_new_message_t)(const char* message);
typedef void (*on_msg_deleted_t)(const char* message);
typedef void (*on_recv_online_only_message_t)(const char* message);

// OnAdvancedMsgListener 回调结构体
typedef struct {
    on_recv_new_message_t          on_recv_new_message;
    on_recv_c2c_read_receipt_t     on_recv_c2c_read_receipt;
    on_new_recv_message_revoked_t  on_new_recv_message_revoked;
    on_recv_offline_new_message_t  on_recv_offline_new_message;
    on_msg_deleted_t               on_msg_deleted;
    on_recv_online_only_message_t  on_recv_online_only_message;
} advanced_msg_listener_t;

// OnFriendshipListener 回调函数指针
typedef void (*on_friend_application_added_t)(const char* friend_application);
typedef void (*on_friend_application_deleted_t)(const char* friend_application);
typedef void (*on_friend_application_accepted_t)(const char* friend_application);
typedef void (*on_friend_application_rejected_t)(const char* friend_application);
typedef void (*on_friend_added_t)(const char* friend_info);
typedef void (*on_friend_deleted_t)(const char* friend_info);
typedef void (*on_friend_info_changed_t)(const char* friend_info);
typedef void (*on_black_added_t)(const char* black_info);
typedef void (*on_black_deleted_t)(const char* black_info);

// OnFriendshipListener 回调结构体
typedef struct {
    on_friend_application_added_t    on_friend_application_added;
    on_friend_application_deleted_t  on_friend_application_deleted;
    on_friend_application_accepted_t on_friend_application_accepted;
    on_friend_application_rejected_t on_friend_application_rejected;
    on_friend_added_t                on_friend_added;
    on_friend_deleted_t              on_friend_deleted;
    on_friend_info_changed_t         on_friend_info_changed;
    on_black_added_t                 on_black_added;
    on_black_deleted_t               on_black_deleted;
} friendship_listener_t;

// OnGroupListener 回调函数指针
typedef void (*on_joined_group_added_t)(const char* group_info);
typedef void (*on_joined_group_deleted_t)(const char* group_info);
typedef void (*on_group_member_added_t)(const char* group_member_info);
typedef void (*on_group_member_deleted_t)(const char* group_member_info);
typedef void (*on_group_application_added_t)(const char* group_application);
typedef void (*on_group_application_deleted_t)(const char* group_application);
typedef void (*on_group_info_changed_t)(const char* group_info);
typedef void (*on_group_dismissed_t)(const char* group_info);
typedef void (*on_group_member_info_changed_t)(const char* group_member_info);
typedef void (*on_group_application_accepted_t)(const char* group_application);
typedef void (*on_group_application_rejected_t)(const char* group_application);

// OnGroupListener 回调结构体
typedef struct {
    on_joined_group_added_t           on_joined_group_added;
    on_joined_group_deleted_t         on_joined_group_deleted;
    on_group_member_added_t           on_group_member_added;
    on_group_member_deleted_t         on_group_member_deleted;
    on_group_application_added_t      on_group_application_added;
    on_group_application_deleted_t    on_group_application_deleted;
    on_group_info_changed_t           on_group_info_changed;
    on_group_dismissed_t              on_group_dismissed;
    on_group_member_info_changed_t    on_group_member_info_changed;
    on_group_application_accepted_t   on_group_application_accepted;
    on_group_application_rejected_t   on_group_application_rejected;
} group_listener_t;

// OnUserListener 回调函数指针
typedef void (*on_self_info_updated_t)(const char* user_info);
typedef void (*on_user_status_changed_t)(const char* user_online_status);

// OnUserListener 回调结构体
typedef struct {
    on_self_info_updated_t     on_self_info_updated;
    on_user_status_changed_t   on_user_status_changed;
} user_listener_t;

// OnCustomBusinessListener 回调结构体
typedef void (*on_recv_custom_business_message_t)(const char* business_message);

typedef struct {
    on_recv_custom_business_message_t on_recv_custom_business_message;
} custom_business_listener_t;

// OnMessageKvInfoListener 回调结构体
typedef void (*on_message_kv_info_changed_t)(const char* message_changed_list);

typedef struct {
    on_message_kv_info_changed_t on_message_kv_info_changed;
} message_kv_info_listener_t;

// UploadFileCallback 回调函数指针
typedef void (*upload_file_open_t)(int64_t size);
typedef void (*upload_file_part_size_t)(int64_t part_size, int num);
typedef void (*upload_file_hash_part_progress_t)(int index, int64_t size, const char* part_hash);
typedef void (*upload_file_hash_part_complete_t)(const char* parts_hash, const char* file_hash);
typedef void (*upload_file_upload_id_t)(const char* upload_id);
typedef void (*upload_file_upload_part_complete_t)(int index, int64_t part_size, const char* part_hash);
typedef void (*upload_file_upload_complete_t)(int64_t file_size, int64_t stream_size, int64_t storage_size);
typedef void (*upload_file_complete_t)(int64_t size, const char* url, int typ);

// UploadFileCallback 回调结构体
typedef struct {
    upload_file_open_t                     open;
    upload_file_part_size_t                part_size;
    upload_file_hash_part_progress_t       hash_part_progress;
    upload_file_hash_part_complete_t       hash_part_complete;
    upload_file_upload_id_t                upload_id;
    upload_file_upload_part_complete_t     upload_part_complete;
    upload_file_upload_complete_t          upload_complete;
    upload_file_complete_t                 complete;
} upload_file_callback_t;

// UploadLogProgress 回调结构体
typedef void (*upload_log_progress_cb_t)(int64_t current, int64_t size);

typedef struct {
    upload_log_progress_cb_t on_progress;
} upload_log_progress_t;

// ===================== C wrapper 函数（安全调用函数指针） =====================

// Base 回调 wrapper
static void base_cb_on_success(base_callback_t* cb, const char* data) {
    if (cb && cb->on_success) cb->on_success(data);
}
static void base_cb_on_error(base_callback_t* cb, int32_t err_code, const char* err_msg) {
    if (cb && cb->on_error) cb->on_error(err_code, err_msg);
}

// SendMsg 回调 wrapper
static void send_msg_cb_on_progress(send_msg_callback_t* cb, int progress) {
    if (cb && cb->on_progress) cb->on_progress(progress);
}

// OnConnListener wrapper
static void conn_cb_on_connecting(conn_listener_t* cb) {
    if (cb && cb->on_connecting) cb->on_connecting();
}
static void conn_cb_on_connect_success(conn_listener_t* cb) {
    if (cb && cb->on_connect_success) cb->on_connect_success();
}
static void conn_cb_on_connect_failed(conn_listener_t* cb, int32_t err_code, const char* err_msg) {
    if (cb && cb->on_connect_failed) cb->on_connect_failed(err_code, err_msg);
}
static void conn_cb_on_kicked_offline(conn_listener_t* cb) {
    if (cb && cb->on_kicked_offline) cb->on_kicked_offline();
}
static void conn_cb_on_user_token_expired(conn_listener_t* cb) {
    if (cb && cb->on_user_token_expired) cb->on_user_token_expired();
}
static void conn_cb_on_user_token_invalid(conn_listener_t* cb, const char* err_msg) {
    if (cb && cb->on_user_token_invalid) cb->on_user_token_invalid(err_msg);
}

// OnConversationListener wrapper
static void conv_cb_on_sync_server_start(conversation_listener_t* cb, int reinstalled) {
    if (cb && cb->on_sync_server_start) cb->on_sync_server_start(reinstalled);
}
static void conv_cb_on_sync_server_finish(conversation_listener_t* cb, int reinstalled) {
    if (cb && cb->on_sync_server_finish) cb->on_sync_server_finish(reinstalled);
}
static void conv_cb_on_sync_server_progress(conversation_listener_t* cb, int progress) {
    if (cb && cb->on_sync_server_progress) cb->on_sync_server_progress(progress);
}
static void conv_cb_on_sync_server_failed(conversation_listener_t* cb, int reinstalled) {
    if (cb && cb->on_sync_server_failed) cb->on_sync_server_failed(reinstalled);
}
static void conv_cb_on_new_conversation(conversation_listener_t* cb, const char* conversation_list) {
    if (cb && cb->on_new_conversation) cb->on_new_conversation(conversation_list);
}
static void conv_cb_on_conversation_changed(conversation_listener_t* cb, const char* conversation_list) {
    if (cb && cb->on_conversation_changed) cb->on_conversation_changed(conversation_list);
}
static void conv_cb_on_total_unread_count_changed(conversation_listener_t* cb, int32_t total_unread_count) {
    if (cb && cb->on_total_unread_count_changed) cb->on_total_unread_count_changed(total_unread_count);
}
static void conv_cb_on_conversation_user_input_status_changed(conversation_listener_t* cb, const char* change) {
    if (cb && cb->on_conversation_user_input_status_changed) cb->on_conversation_user_input_status_changed(change);
}

// OnAdvancedMsgListener wrapper
static void msg_cb_on_recv_new_message(advanced_msg_listener_t* cb, const char* message) {
    if (cb && cb->on_recv_new_message) cb->on_recv_new_message(message);
}
static void msg_cb_on_recv_c2c_read_receipt(advanced_msg_listener_t* cb, const char* msg_receipt_list) {
    if (cb && cb->on_recv_c2c_read_receipt) cb->on_recv_c2c_read_receipt(msg_receipt_list);
}
static void msg_cb_on_new_recv_message_revoked(advanced_msg_listener_t* cb, const char* message_revoked) {
    if (cb && cb->on_new_recv_message_revoked) cb->on_new_recv_message_revoked(message_revoked);
}
static void msg_cb_on_recv_offline_new_message(advanced_msg_listener_t* cb, const char* message) {
    if (cb && cb->on_recv_offline_new_message) cb->on_recv_offline_new_message(message);
}
static void msg_cb_on_msg_deleted(advanced_msg_listener_t* cb, const char* message) {
    if (cb && cb->on_msg_deleted) cb->on_msg_deleted(message);
}
static void msg_cb_on_recv_online_only_message(advanced_msg_listener_t* cb, const char* message) {
    if (cb && cb->on_recv_online_only_message) cb->on_recv_online_only_message(message);
}

// OnFriendshipListener wrapper
static void friend_cb_on_application_added(friendship_listener_t* cb, const char* friend_application) {
    if (cb && cb->on_friend_application_added) cb->on_friend_application_added(friend_application);
}
static void friend_cb_on_application_deleted(friendship_listener_t* cb, const char* friend_application) {
    if (cb && cb->on_friend_application_deleted) cb->on_friend_application_deleted(friend_application);
}
static void friend_cb_on_application_accepted(friendship_listener_t* cb, const char* friend_application) {
    if (cb && cb->on_friend_application_accepted) cb->on_friend_application_accepted(friend_application);
}
static void friend_cb_on_application_rejected(friendship_listener_t* cb, const char* friend_application) {
    if (cb && cb->on_friend_application_rejected) cb->on_friend_application_rejected(friend_application);
}
static void friend_cb_on_friend_added(friendship_listener_t* cb, const char* friend_info) {
    if (cb && cb->on_friend_added) cb->on_friend_added(friend_info);
}
static void friend_cb_on_friend_deleted(friendship_listener_t* cb, const char* friend_info) {
    if (cb && cb->on_friend_deleted) cb->on_friend_deleted(friend_info);
}
static void friend_cb_on_friend_info_changed(friendship_listener_t* cb, const char* friend_info) {
    if (cb && cb->on_friend_info_changed) cb->on_friend_info_changed(friend_info);
}
static void friend_cb_on_black_added(friendship_listener_t* cb, const char* black_info) {
    if (cb && cb->on_black_added) cb->on_black_added(black_info);
}
static void friend_cb_on_black_deleted(friendship_listener_t* cb, const char* black_info) {
    if (cb && cb->on_black_deleted) cb->on_black_deleted(black_info);
}

// OnGroupListener wrapper
static void group_cb_on_joined_group_added(group_listener_t* cb, const char* group_info) {
    if (cb && cb->on_joined_group_added) cb->on_joined_group_added(group_info);
}
static void group_cb_on_joined_group_deleted(group_listener_t* cb, const char* group_info) {
    if (cb && cb->on_joined_group_deleted) cb->on_joined_group_deleted(group_info);
}
static void group_cb_on_group_member_added(group_listener_t* cb, const char* group_member_info) {
    if (cb && cb->on_group_member_added) cb->on_group_member_added(group_member_info);
}
static void group_cb_on_group_member_deleted(group_listener_t* cb, const char* group_member_info) {
    if (cb && cb->on_group_member_deleted) cb->on_group_member_deleted(group_member_info);
}
static void group_cb_on_group_application_added(group_listener_t* cb, const char* group_application) {
    if (cb && cb->on_group_application_added) cb->on_group_application_added(group_application);
}
static void group_cb_on_group_application_deleted(group_listener_t* cb, const char* group_application) {
    if (cb && cb->on_group_application_deleted) cb->on_group_application_deleted(group_application);
}
static void group_cb_on_group_info_changed(group_listener_t* cb, const char* group_info) {
    if (cb && cb->on_group_info_changed) cb->on_group_info_changed(group_info);
}
static void group_cb_on_group_dismissed(group_listener_t* cb, const char* group_info) {
    if (cb && cb->on_group_dismissed) cb->on_group_dismissed(group_info);
}
static void group_cb_on_group_member_info_changed(group_listener_t* cb, const char* group_member_info) {
    if (cb && cb->on_group_member_info_changed) cb->on_group_member_info_changed(group_member_info);
}
static void group_cb_on_group_application_accepted(group_listener_t* cb, const char* group_application) {
    if (cb && cb->on_group_application_accepted) cb->on_group_application_accepted(group_application);
}
static void group_cb_on_group_application_rejected(group_listener_t* cb, const char* group_application) {
    if (cb && cb->on_group_application_rejected) cb->on_group_application_rejected(group_application);
}

// OnUserListener wrapper
static void user_cb_on_self_info_updated(user_listener_t* cb, const char* user_info) {
    if (cb && cb->on_self_info_updated) cb->on_self_info_updated(user_info);
}
static void user_cb_on_user_status_changed(user_listener_t* cb, const char* user_online_status) {
    if (cb && cb->on_user_status_changed) cb->on_user_status_changed(user_online_status);
}

// OnCustomBusinessListener wrapper
static void business_cb_on_recv(custom_business_listener_t* cb, const char* business_message) {
    if (cb && cb->on_recv_custom_business_message) cb->on_recv_custom_business_message(business_message);
}

// OnMessageKvInfoListener wrapper
static void msgkv_cb_on_changed(message_kv_info_listener_t* cb, const char* message_changed_list) {
    if (cb && cb->on_message_kv_info_changed) cb->on_message_kv_info_changed(message_changed_list);
}

// UploadFileCallback wrapper
static void upload_file_cb_open(upload_file_callback_t* cb, int64_t size) {
    if (cb && cb->open) cb->open(size);
}
static void upload_file_cb_part_size(upload_file_callback_t* cb, int64_t part_size, int num) {
    if (cb && cb->part_size) cb->part_size(part_size, num);
}
static void upload_file_cb_hash_part_progress(upload_file_callback_t* cb, int index, int64_t size, const char* part_hash) {
    if (cb && cb->hash_part_progress) cb->hash_part_progress(index, size, part_hash);
}
static void upload_file_cb_hash_part_complete(upload_file_callback_t* cb, const char* parts_hash, const char* file_hash) {
    if (cb && cb->hash_part_complete) cb->hash_part_complete(parts_hash, file_hash);
}
static void upload_file_cb_upload_id(upload_file_callback_t* cb, const char* upload_id) {
    if (cb && cb->upload_id) cb->upload_id(upload_id);
}
static void upload_file_cb_upload_part_complete(upload_file_callback_t* cb, int index, int64_t part_size, const char* part_hash) {
    if (cb && cb->upload_part_complete) cb->upload_part_complete(index, part_size, part_hash);
}
static void upload_file_cb_upload_complete(upload_file_callback_t* cb, int64_t file_size, int64_t stream_size, int64_t storage_size) {
    if (cb && cb->upload_complete) cb->upload_complete(file_size, stream_size, storage_size);
}
static void upload_file_cb_complete(upload_file_callback_t* cb, int64_t size, const char* url, int typ) {
    if (cb && cb->complete) cb->complete(size, url, typ);
}

// UploadLogProgress wrapper
static void upload_log_cb_on_progress(upload_log_progress_t* cb, int64_t current, int64_t size) {
    if (cb && cb->on_progress) cb->on_progress(current, size);
}
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
