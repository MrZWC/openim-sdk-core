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

// ===================== 会话管理 API =====================

// OpenIM_GetAllConversationList 获取所有会话列表（异步）
//
//export OpenIM_GetAllConversationList
func OpenIM_GetAllConversationList(cb *C.base_callback_t, operationID *C.char) {
	open_im_sdk.GetAllConversationList(newBaseCallback(cb), C.GoString(operationID))
}

// OpenIM_GetConversationListSplit 分页获取会话列表（异步）
//
//export OpenIM_GetConversationListSplit
func OpenIM_GetConversationListSplit(cb *C.base_callback_t, operationID *C.char, offset, count C.int) {
	open_im_sdk.GetConversationListSplit(newBaseCallback(cb), C.GoString(operationID), int(offset), int(count))
}

// OpenIM_GetOneConversation 获取指定会话（异步）
//
//export OpenIM_GetOneConversation
func OpenIM_GetOneConversation(cb *C.base_callback_t, operationID *C.char, sessionType C.int32_t, sourceID *C.char) {
	open_im_sdk.GetOneConversation(newBaseCallback(cb), C.GoString(operationID), int32(sessionType), C.GoString(sourceID))
}

// OpenIM_GetMultipleConversation 批量获取会话（异步）
//
//export OpenIM_GetMultipleConversation
func OpenIM_GetMultipleConversation(cb *C.base_callback_t, operationID *C.char, conversationIDList *C.char) {
	open_im_sdk.GetMultipleConversation(newBaseCallback(cb), C.GoString(operationID), C.GoString(conversationIDList))
}

// OpenIM_SetConversation 设置会话属性（异步）
//
//export OpenIM_SetConversation
func OpenIM_SetConversation(cb *C.base_callback_t, operationID, conversationID, req *C.char) {
	open_im_sdk.SetConversation(newBaseCallback(cb), C.GoString(operationID), C.GoString(conversationID), C.GoString(req))
}

// OpenIM_HideConversation 隐藏会话（异步）
//
//export OpenIM_HideConversation
func OpenIM_HideConversation(cb *C.base_callback_t, operationID, conversationID *C.char) {
	open_im_sdk.HideConversation(newBaseCallback(cb), C.GoString(operationID), C.GoString(conversationID))
}

// OpenIM_SetConversationDraft 设置会话草稿（异步）
//
//export OpenIM_SetConversationDraft
func OpenIM_SetConversationDraft(cb *C.base_callback_t, operationID, conversationID, draftText *C.char) {
	open_im_sdk.SetConversationDraft(newBaseCallback(cb), C.GoString(operationID), C.GoString(conversationID), C.GoString(draftText))
}

// OpenIM_GetTotalUnreadMsgCount 获取总未读消息数（异步）
//
//export OpenIM_GetTotalUnreadMsgCount
func OpenIM_GetTotalUnreadMsgCount(cb *C.base_callback_t, operationID *C.char) {
	open_im_sdk.GetTotalUnreadMsgCount(newBaseCallback(cb), C.GoString(operationID))
}

// OpenIM_HideAllConversations 隐藏所有会话（异步）
//
//export OpenIM_HideAllConversations
func OpenIM_HideAllConversations(cb *C.base_callback_t, operationID *C.char) {
	open_im_sdk.HideAllConversations(newBaseCallback(cb), C.GoString(operationID))
}

// OpenIM_ClearConversationAndDeleteAllMsg 清空会话并删除所有消息（异步）
//
//export OpenIM_ClearConversationAndDeleteAllMsg
func OpenIM_ClearConversationAndDeleteAllMsg(cb *C.base_callback_t, operationID, conversationID *C.char) {
	open_im_sdk.ClearConversationAndDeleteAllMsg(newBaseCallback(cb), C.GoString(operationID), C.GoString(conversationID))
}

// OpenIM_DeleteConversationAndDeleteAllMsg 删除会话及所有消息（异步）
//
//export OpenIM_DeleteConversationAndDeleteAllMsg
func OpenIM_DeleteConversationAndDeleteAllMsg(cb *C.base_callback_t, operationID, conversationID *C.char) {
	open_im_sdk.DeleteConversationAndDeleteAllMsg(newBaseCallback(cb), C.GoString(operationID), C.GoString(conversationID))
}

// ===================== 消息创建 API（同步返回 JSON 字符串） =====================

// OpenIM_GetAtAllTag 获取 @全部 标识（同步）
//
//export OpenIM_GetAtAllTag
func OpenIM_GetAtAllTag(operationID *C.char) *C.char {
	return C.CString(open_im_sdk.GetAtAllTag(C.GoString(operationID)))
}

// OpenIM_CreateTextMessage 创建文本消息（同步）
//
//export OpenIM_CreateTextMessage
func OpenIM_CreateTextMessage(operationID, text *C.char) *C.char {
	return C.CString(open_im_sdk.CreateTextMessage(C.GoString(operationID), C.GoString(text)))
}

// OpenIM_CreateAdvancedTextMessage 创建高级文本消息（同步）
//
//export OpenIM_CreateAdvancedTextMessage
func OpenIM_CreateAdvancedTextMessage(operationID, text, messageEntityList *C.char) *C.char {
	return C.CString(open_im_sdk.CreateAdvancedTextMessage(C.GoString(operationID), C.GoString(text), C.GoString(messageEntityList)))
}

// OpenIM_CreateTextAtMessage 创建 @ 文本消息（同步）
//
//export OpenIM_CreateTextAtMessage
func OpenIM_CreateTextAtMessage(operationID, text, atUserList, atUsersInfo, message *C.char) *C.char {
	return C.CString(open_im_sdk.CreateTextAtMessage(C.GoString(operationID), C.GoString(text), C.GoString(atUserList), C.GoString(atUsersInfo), C.GoString(message)))
}

// OpenIM_CreateLocationMessage 创建位置消息（同步）
//
//export OpenIM_CreateLocationMessage
func OpenIM_CreateLocationMessage(operationID, description *C.char, longitude, latitude C.double) *C.char {
	return C.CString(open_im_sdk.CreateLocationMessage(C.GoString(operationID), C.GoString(description), float64(longitude), float64(latitude)))
}

// OpenIM_CreateCustomMessage 创建自定义消息（同步）
//
//export OpenIM_CreateCustomMessage
func OpenIM_CreateCustomMessage(operationID, data, extension, description *C.char) *C.char {
	return C.CString(open_im_sdk.CreateCustomMessage(C.GoString(operationID), C.GoString(data), C.GoString(extension), C.GoString(description)))
}

// OpenIM_CreateQuoteMessage 创建引用消息（同步）
//
//export OpenIM_CreateQuoteMessage
func OpenIM_CreateQuoteMessage(operationID, text, message *C.char) *C.char {
	return C.CString(open_im_sdk.CreateQuoteMessage(C.GoString(operationID), C.GoString(text), C.GoString(message)))
}

// OpenIM_CreateAdvancedQuoteMessage 创建高级引用消息（同步）
//
//export OpenIM_CreateAdvancedQuoteMessage
func OpenIM_CreateAdvancedQuoteMessage(operationID, text, message, messageEntityList *C.char) *C.char {
	return C.CString(open_im_sdk.CreateAdvancedQuoteMessage(C.GoString(operationID), C.GoString(text), C.GoString(message), C.GoString(messageEntityList)))
}

// OpenIM_CreateCardMessage 创建名片消息（同步）
//
//export OpenIM_CreateCardMessage
func OpenIM_CreateCardMessage(operationID, cardInfo *C.char) *C.char {
	return C.CString(open_im_sdk.CreateCardMessage(C.GoString(operationID), C.GoString(cardInfo)))
}

// OpenIM_CreateImageMessage 创建图片消息（同步）
//
//export OpenIM_CreateImageMessage
func OpenIM_CreateImageMessage(operationID, imagePath *C.char) *C.char {
	return C.CString(open_im_sdk.CreateImageMessage(C.GoString(operationID), C.GoString(imagePath)))
}

// OpenIM_CreateImageMessageByURL 通过 URL 创建图片消息（同步）
//
//export OpenIM_CreateImageMessageByURL
func OpenIM_CreateImageMessageByURL(operationID, sourcePath, sourcePicture, bigPicture, snapshotPicture *C.char) *C.char {
	return C.CString(open_im_sdk.CreateImageMessageByURL(C.GoString(operationID), C.GoString(sourcePath), C.GoString(sourcePicture), C.GoString(bigPicture), C.GoString(snapshotPicture)))
}

// OpenIM_CreateImageMessageFromFullPath 通过完整路径创建图片消息（同步）
//
//export OpenIM_CreateImageMessageFromFullPath
func OpenIM_CreateImageMessageFromFullPath(operationID, imageFullPath *C.char) *C.char {
	return C.CString(open_im_sdk.CreateImageMessageFromFullPath(C.GoString(operationID), C.GoString(imageFullPath)))
}

// OpenIM_CreateSoundMessage 创建语音消息（同步）
//
//export OpenIM_CreateSoundMessage
func OpenIM_CreateSoundMessage(operationID, soundPath *C.char, duration C.int64_t) *C.char {
	return C.CString(open_im_sdk.CreateSoundMessage(C.GoString(operationID), C.GoString(soundPath), int64(duration)))
}

// OpenIM_CreateSoundMessageByURL 通过 URL 创建语音消息（同步）
//
//export OpenIM_CreateSoundMessageByURL
func OpenIM_CreateSoundMessageByURL(operationID, soundBaseInfo *C.char) *C.char {
	return C.CString(open_im_sdk.CreateSoundMessageByURL(C.GoString(operationID), C.GoString(soundBaseInfo)))
}

// OpenIM_CreateSoundMessageFromFullPath 通过完整路径创建语音消息（同步）
//
//export OpenIM_CreateSoundMessageFromFullPath
func OpenIM_CreateSoundMessageFromFullPath(operationID, soundPath *C.char, duration C.int64_t) *C.char {
	return C.CString(open_im_sdk.CreateSoundMessageFromFullPath(C.GoString(operationID), C.GoString(soundPath), int64(duration)))
}

// OpenIM_CreateVideoMessage 创建视频消息（同步）
//
//export OpenIM_CreateVideoMessage
func OpenIM_CreateVideoMessage(operationID, videoPath, videoType *C.char, duration C.int64_t, snapshotPath *C.char) *C.char {
	return C.CString(open_im_sdk.CreateVideoMessage(C.GoString(operationID), C.GoString(videoPath), C.GoString(videoType), int64(duration), C.GoString(snapshotPath)))
}

// OpenIM_CreateVideoMessageByURL 通过 URL 创建视频消息（同步）
//
//export OpenIM_CreateVideoMessageByURL
func OpenIM_CreateVideoMessageByURL(operationID, videoBaseInfo *C.char) *C.char {
	return C.CString(open_im_sdk.CreateVideoMessageByURL(C.GoString(operationID), C.GoString(videoBaseInfo)))
}

// OpenIM_CreateVideoMessageFromFullPath 通过完整路径创建视频消息（同步）
//
//export OpenIM_CreateVideoMessageFromFullPath
func OpenIM_CreateVideoMessageFromFullPath(operationID, videoFullPath, videoType *C.char, duration C.int64_t, snapshotFullPath *C.char) *C.char {
	return C.CString(open_im_sdk.CreateVideoMessageFromFullPath(C.GoString(operationID), C.GoString(videoFullPath), C.GoString(videoType), int64(duration), C.GoString(snapshotFullPath)))
}

// OpenIM_CreateFileMessage 创建文件消息（同步）
//
//export OpenIM_CreateFileMessage
func OpenIM_CreateFileMessage(operationID, filePath, fileName *C.char) *C.char {
	return C.CString(open_im_sdk.CreateFileMessage(C.GoString(operationID), C.GoString(filePath), C.GoString(fileName)))
}

// OpenIM_CreateFileMessageByURL 通过 URL 创建文件消息（同步）
//
//export OpenIM_CreateFileMessageByURL
func OpenIM_CreateFileMessageByURL(operationID, fileBaseInfo *C.char) *C.char {
	return C.CString(open_im_sdk.CreateFileMessageByURL(C.GoString(operationID), C.GoString(fileBaseInfo)))
}

// OpenIM_CreateFileMessageFromFullPath 通过完整路径创建文件消息（同步）
//
//export OpenIM_CreateFileMessageFromFullPath
func OpenIM_CreateFileMessageFromFullPath(operationID, fileFullPath, fileName *C.char) *C.char {
	return C.CString(open_im_sdk.CreateFileMessageFromFullPath(C.GoString(operationID), C.GoString(fileFullPath), C.GoString(fileName)))
}

// OpenIM_CreateMergerMessage 创建合并转发消息（同步）
//
//export OpenIM_CreateMergerMessage
func OpenIM_CreateMergerMessage(operationID, messageList, title, summaryList *C.char) *C.char {
	return C.CString(open_im_sdk.CreateMergerMessage(C.GoString(operationID), C.GoString(messageList), C.GoString(title), C.GoString(summaryList)))
}

// OpenIM_CreateFaceMessage 创建表情消息（同步）
//
//export OpenIM_CreateFaceMessage
func OpenIM_CreateFaceMessage(operationID *C.char, index C.int, data *C.char) *C.char {
	return C.CString(open_im_sdk.CreateFaceMessage(C.GoString(operationID), int(index), C.GoString(data)))
}

// OpenIM_CreateForwardMessage 创建转发消息（同步）
//
//export OpenIM_CreateForwardMessage
func OpenIM_CreateForwardMessage(operationID, m *C.char) *C.char {
	return C.CString(open_im_sdk.CreateForwardMessage(C.GoString(operationID), C.GoString(m)))
}

// OpenIM_GetConversationIDBySessionType 通过会话类型获取会话ID（同步）
//
//export OpenIM_GetConversationIDBySessionType
func OpenIM_GetConversationIDBySessionType(operationID, sourceID *C.char, sessionType C.int) *C.char {
	return C.CString(open_im_sdk.GetConversationIDBySessionType(C.GoString(operationID), C.GoString(sourceID), int(sessionType)))
}

// ===================== 消息发送 API（使用 SendMsgCallBack） =====================

// OpenIM_SendMessage 发送消息（异步，进度通过 callback 返回）
//
//export OpenIM_SendMessage
func OpenIM_SendMessage(cb *C.send_msg_callback_t, operationID, message, recvID, groupID, offlinePushInfo *C.char, isOnlineOnly C.int) {
	open_im_sdk.SendMessage(newSendMsgCallback(cb), C.GoString(operationID), C.GoString(message), C.GoString(recvID), C.GoString(groupID), C.GoString(offlinePushInfo), isOnlineOnly != 0)
}

// OpenIM_SendMessageNotOss 发送消息不经过 OSS（异步）
//
//export OpenIM_SendMessageNotOss
func OpenIM_SendMessageNotOss(cb *C.send_msg_callback_t, operationID, message, recvID, groupID, offlinePushInfo *C.char, isOnlineOnly C.int) {
	open_im_sdk.SendMessageNotOss(newSendMsgCallback(cb), C.GoString(operationID), C.GoString(message), C.GoString(recvID), C.GoString(groupID), C.GoString(offlinePushInfo), isOnlineOnly != 0)
}

// ===================== 消息历史与搜索 API =====================

// OpenIM_FindMessageList 查找消息列表（异步）
//
//export OpenIM_FindMessageList
func OpenIM_FindMessageList(cb *C.base_callback_t, operationID, findMessageOptions *C.char) {
	open_im_sdk.FindMessageList(newBaseCallback(cb), C.GoString(operationID), C.GoString(findMessageOptions))
}

// OpenIM_GetAdvancedHistoryMessageList 获取高级历史消息列表（异步）
//
//export OpenIM_GetAdvancedHistoryMessageList
func OpenIM_GetAdvancedHistoryMessageList(cb *C.base_callback_t, operationID, getMessageOptions *C.char) {
	open_im_sdk.GetAdvancedHistoryMessageList(newBaseCallback(cb), C.GoString(operationID), C.GoString(getMessageOptions))
}

// OpenIM_GetAdvancedHistoryMessageListReverse 反向获取高级历史消息列表（异步）
//
//export OpenIM_GetAdvancedHistoryMessageListReverse
func OpenIM_GetAdvancedHistoryMessageListReverse(cb *C.base_callback_t, operationID, getMessageOptions *C.char) {
	open_im_sdk.GetAdvancedHistoryMessageListReverse(newBaseCallback(cb), C.GoString(operationID), C.GoString(getMessageOptions))
}

// OpenIM_SearchLocalMessages 搜索本地消息（异步）
//
//export OpenIM_SearchLocalMessages
func OpenIM_SearchLocalMessages(cb *C.base_callback_t, operationID, searchParam *C.char) {
	open_im_sdk.SearchLocalMessages(newBaseCallback(cb), C.GoString(operationID), C.GoString(searchParam))
}

// OpenIM_SearchConversation 搜索会话（异步）
//
//export OpenIM_SearchConversation
func OpenIM_SearchConversation(cb *C.base_callback_t, operationID, searchParam *C.char) {
	open_im_sdk.SearchConversation(newBaseCallback(cb), C.GoString(operationID), C.GoString(searchParam))
}

// ===================== 消息操作 API =====================

// OpenIM_RevokeMessage 撤回消息（异步）
//
//export OpenIM_RevokeMessage
func OpenIM_RevokeMessage(cb *C.base_callback_t, operationID, conversationID, clientMsgID *C.char) {
	open_im_sdk.RevokeMessage(newBaseCallback(cb), C.GoString(operationID), C.GoString(conversationID), C.GoString(clientMsgID))
}

// OpenIM_TypingStatusUpdate 输入状态更新（异步）
//
//export OpenIM_TypingStatusUpdate
func OpenIM_TypingStatusUpdate(cb *C.base_callback_t, operationID, recvID, msgTip *C.char) {
	open_im_sdk.TypingStatusUpdate(newBaseCallback(cb), C.GoString(operationID), C.GoString(recvID), C.GoString(msgTip))
}

// OpenIM_MarkConversationMessageAsRead 标记会话消息已读（异步）
//
//export OpenIM_MarkConversationMessageAsRead
func OpenIM_MarkConversationMessageAsRead(cb *C.base_callback_t, operationID, conversationID *C.char) {
	open_im_sdk.MarkConversationMessageAsRead(newBaseCallback(cb), C.GoString(operationID), C.GoString(conversationID))
}

// OpenIM_MarkAllConversationMessageAsRead 标记所有会话消息已读（异步）
//
//export OpenIM_MarkAllConversationMessageAsRead
func OpenIM_MarkAllConversationMessageAsRead(cb *C.base_callback_t, operationID *C.char) {
	open_im_sdk.MarkAllConversationMessageAsRead(newBaseCallback(cb), C.GoString(operationID))
}

// OpenIM_MarkMessagesAsReadByMsgID 按消息ID标记已读（异步）
//
//export OpenIM_MarkMessagesAsReadByMsgID
func OpenIM_MarkMessagesAsReadByMsgID(cb *C.base_callback_t, operationID, conversationID, clientMsgIDs *C.char) {
	open_im_sdk.MarkMessagesAsReadByMsgID(newBaseCallback(cb), C.GoString(operationID), C.GoString(conversationID), C.GoString(clientMsgIDs))
}

// OpenIM_DeleteMessageFromLocalStorage 从本地存储删除消息（异步）
//
//export OpenIM_DeleteMessageFromLocalStorage
func OpenIM_DeleteMessageFromLocalStorage(cb *C.base_callback_t, operationID, conversationID, clientMsgID *C.char) {
	open_im_sdk.DeleteMessageFromLocalStorage(newBaseCallback(cb), C.GoString(operationID), C.GoString(conversationID), C.GoString(clientMsgID))
}

// OpenIM_DeleteMessage 删除消息（异步）
//
//export OpenIM_DeleteMessage
func OpenIM_DeleteMessage(cb *C.base_callback_t, operationID, conversationID, clientMsgID *C.char) {
	open_im_sdk.DeleteMessage(newBaseCallback(cb), C.GoString(operationID), C.GoString(conversationID), C.GoString(clientMsgID))
}

// OpenIM_DeleteAllMsgFromLocalAndSvr 删除本地和服务端所有消息（异步）
//
//export OpenIM_DeleteAllMsgFromLocalAndSvr
func OpenIM_DeleteAllMsgFromLocalAndSvr(cb *C.base_callback_t, operationID *C.char) {
	open_im_sdk.DeleteAllMsgFromLocalAndSvr(newBaseCallback(cb), C.GoString(operationID))
}

// OpenIM_DeleteAllMsgFromLocal 删除本地所有消息（异步）
//
//export OpenIM_DeleteAllMsgFromLocal
func OpenIM_DeleteAllMsgFromLocal(cb *C.base_callback_t, operationID *C.char) {
	open_im_sdk.DeleteAllMsgFromLocal(newBaseCallback(cb), C.GoString(operationID))
}

// OpenIM_InsertSingleMessageToLocalStorage 插入单聊消息到本地存储（异步）
//
//export OpenIM_InsertSingleMessageToLocalStorage
func OpenIM_InsertSingleMessageToLocalStorage(cb *C.base_callback_t, operationID, message, recvID, sendID *C.char) {
	open_im_sdk.InsertSingleMessageToLocalStorage(newBaseCallback(cb), C.GoString(operationID), C.GoString(message), C.GoString(recvID), C.GoString(sendID))
}

// OpenIM_InsertGroupMessageToLocalStorage 插入群聊消息到本地存储（异步）
//
//export OpenIM_InsertGroupMessageToLocalStorage
func OpenIM_InsertGroupMessageToLocalStorage(cb *C.base_callback_t, operationID, message, groupID, sendID *C.char) {
	open_im_sdk.InsertGroupMessageToLocalStorage(newBaseCallback(cb), C.GoString(operationID), C.GoString(message), C.GoString(groupID), C.GoString(sendID))
}

// OpenIM_SetMessageLocalEx 设置消息本地扩展信息（异步）
//
//export OpenIM_SetMessageLocalEx
func OpenIM_SetMessageLocalEx(cb *C.base_callback_t, operationID, conversationID, clientMsgID, localEx *C.char) {
	open_im_sdk.SetMessageLocalEx(newBaseCallback(cb), C.GoString(operationID), C.GoString(conversationID), C.GoString(clientMsgID), C.GoString(localEx))
}

// OpenIM_ChangeInputStates 修改输入状态（异步）
//
//export OpenIM_ChangeInputStates
func OpenIM_ChangeInputStates(cb *C.base_callback_t, operationID, conversationID *C.char, focus C.int) {
	open_im_sdk.ChangeInputStates(newBaseCallback(cb), C.GoString(operationID), C.GoString(conversationID), focus != 0)
}

// OpenIM_GetInputStates 获取输入状态（异步）
//
//export OpenIM_GetInputStates
func OpenIM_GetInputStates(cb *C.base_callback_t, operationID, conversationID, userID *C.char) {
	open_im_sdk.GetInputStates(newBaseCallback(cb), C.GoString(operationID), C.GoString(conversationID), C.GoString(userID))
}
