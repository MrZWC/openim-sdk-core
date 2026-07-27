#ifndef OPENIM_CALLBACK_TYPES_H
#define OPENIM_CALLBACK_TYPES_H

#include <stdint.h>
#include <stdlib.h>
#include <stdbool.h>

// ===================== C callback function pointer types =====================

// Base callback
typedef void (*on_success_t)(const char* data);
typedef void (*on_error_t)(int32_t err_code, const char* err_msg);

typedef struct {
    on_success_t on_success;
    on_error_t   on_error;
} base_callback_t;

// SendMsg callback (Base + progress)
typedef void (*on_progress_t)(int progress);

typedef struct {
    on_success_t  on_success;
    on_error_t    on_error;
    on_progress_t on_progress;
} send_msg_callback_t;

// OnConnListener
typedef void (*on_connecting_t)();
typedef void (*on_connect_success_t)();
typedef void (*on_connect_failed_t)(int32_t err_code, const char* err_msg);
typedef void (*on_kicked_offline_t)();
typedef void (*on_user_token_expired_t)();
typedef void (*on_user_token_invalid_t)(const char* err_msg);

typedef struct {
    on_connecting_t         on_connecting;
    on_connect_success_t    on_connect_success;
    on_connect_failed_t     on_connect_failed;
    on_kicked_offline_t     on_kicked_offline;
    on_user_token_expired_t on_user_token_expired;
    on_user_token_invalid_t on_user_token_invalid;
} conn_listener_t;

// OnConversationListener
typedef void (*on_sync_server_start_t)(bool reinstalled);
typedef void (*on_sync_server_finish_t)(bool reinstalled);
typedef void (*on_sync_server_progress_t)(int progress);
typedef void (*on_sync_server_failed_t)(bool reinstalled);
typedef void (*on_new_conversation_t)(const char* conversation_list);
typedef void (*on_conversation_changed_t)(const char* conversation_list);
typedef void (*on_total_unread_count_changed_t)(int32_t total_unread_count);
typedef void (*on_conversation_user_input_status_changed_t)(const char* change);

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

// OnAdvancedMsgListener
typedef void (*on_recv_new_message_t)(const char* message);
typedef void (*on_recv_c2c_read_receipt_t)(const char* msg_receipt_list);
typedef void (*on_new_recv_message_revoked_t)(const char* message_revoked);
typedef void (*on_recv_offline_new_message_t)(const char* message);
typedef void (*on_msg_deleted_t)(const char* message);
typedef void (*on_recv_online_only_message_t)(const char* message);

typedef struct {
    on_recv_new_message_t          on_recv_new_message;
    on_recv_c2c_read_receipt_t     on_recv_c2c_read_receipt;
    on_new_recv_message_revoked_t  on_new_recv_message_revoked;
    on_recv_offline_new_message_t  on_recv_offline_new_message;
    on_msg_deleted_t               on_msg_deleted;
    on_recv_online_only_message_t  on_recv_online_only_message;
} advanced_msg_listener_t;

// OnFriendshipListener
typedef void (*on_friend_application_added_t)(const char* friend_application);
typedef void (*on_friend_application_deleted_t)(const char* friend_application);
typedef void (*on_friend_application_accepted_t)(const char* friend_application);
typedef void (*on_friend_application_rejected_t)(const char* friend_application);
typedef void (*on_friend_added_t)(const char* friend_info);
typedef void (*on_friend_deleted_t)(const char* friend_info);
typedef void (*on_friend_info_changed_t)(const char* friend_info);
typedef void (*on_black_added_t)(const char* black_info);
typedef void (*on_black_deleted_t)(const char* black_info);

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

// OnGroupListener
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

// OnUserListener
typedef void (*on_self_info_updated_t)(const char* user_info);
typedef void (*on_user_status_changed_t)(const char* user_online_status);

typedef struct {
    on_self_info_updated_t     on_self_info_updated;
    on_user_status_changed_t   on_user_status_changed;
} user_listener_t;

// OnCustomBusinessListener
typedef void (*on_recv_custom_business_message_t)(const char* business_message);

typedef struct {
    on_recv_custom_business_message_t on_recv_custom_business_message;
} custom_business_listener_t;

// OnMessageKvInfoListener
typedef void (*on_message_kv_info_changed_t)(const char* message_changed_list);

typedef struct {
    on_message_kv_info_changed_t on_message_kv_info_changed;
} message_kv_info_listener_t;

// UploadFileCallback
typedef void (*upload_file_open_t)(int64_t size);
typedef void (*upload_file_part_size_t)(int64_t part_size, int num);
typedef void (*upload_file_hash_part_progress_t)(int index, int64_t size, const char* part_hash);
typedef void (*upload_file_hash_part_complete_t)(const char* parts_hash, const char* file_hash);
typedef void (*upload_file_upload_id_t)(const char* upload_id);
typedef void (*upload_file_upload_part_complete_t)(int index, int64_t part_size, const char* part_hash);
typedef void (*upload_file_upload_complete_t)(int64_t file_size, int64_t stream_size, int64_t storage_size);
typedef void (*upload_file_complete_t)(int64_t size, const char* url, int typ);

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

// UploadLogProgress
typedef void (*upload_log_progress_cb_t)(int64_t current, int64_t size);

typedef struct {
    upload_log_progress_cb_t on_progress;
} upload_log_progress_t;

// ===================== C wrapper functions (safe function pointer calls) =====================

// Base callback wrappers
static void base_cb_on_success(base_callback_t* cb, const char* data) {
    if (cb && cb->on_success) cb->on_success(data);
}
static void base_cb_on_error(base_callback_t* cb, int32_t err_code, const char* err_msg) {
    if (cb && cb->on_error) cb->on_error(err_code, err_msg);
}

// SendMsg callback wrappers
static void send_msg_cb_on_progress(send_msg_callback_t* cb, int progress) {
    if (cb && cb->on_progress) cb->on_progress(progress);
}

// OnConnListener wrappers
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

// OnConversationListener wrappers
static void conv_cb_on_sync_server_start(conversation_listener_t* cb, bool reinstalled) {
    if (cb && cb->on_sync_server_start) cb->on_sync_server_start(reinstalled);
}
static void conv_cb_on_sync_server_finish(conversation_listener_t* cb, bool reinstalled) {
    if (cb && cb->on_sync_server_finish) cb->on_sync_server_finish(reinstalled);
}
static void conv_cb_on_sync_server_progress(conversation_listener_t* cb, int progress) {
    if (cb && cb->on_sync_server_progress) cb->on_sync_server_progress(progress);
}
static void conv_cb_on_sync_server_failed(conversation_listener_t* cb, bool reinstalled) {
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

// OnAdvancedMsgListener wrappers
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

// OnFriendshipListener wrappers
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

// OnGroupListener wrappers
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

// OnUserListener wrappers
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

// UploadFileCallback wrappers
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

#endif // OPENIM_CALLBACK_TYPES_H
