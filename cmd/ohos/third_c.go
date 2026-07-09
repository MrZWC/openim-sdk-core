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

// OpenIM_UpdateFcmToken 更新 FCM 推送 Token（异步）
//
//export OpenIM_UpdateFcmToken
func OpenIM_UpdateFcmToken(cb *C.base_callback_t, operationID, fcmToken *C.char, expireTime C.int64_t) {
	open_im_sdk.UpdateFcmToken(newBaseCallback(cb), C.GoString(operationID), C.GoString(fcmToken), int64(expireTime))
}

// OpenIM_SetAppBadge 设置应用角标未读数（异步）
//
//export OpenIM_SetAppBadge
func OpenIM_SetAppBadge(cb *C.base_callback_t, operationID *C.char, appUnreadCount C.int32_t) {
	open_im_sdk.SetAppBadge(newBaseCallback(cb), C.GoString(operationID), int32(appUnreadCount))
}

// OpenIM_UploadLogs 上传日志（异步）
// 参数:
//   - progress: 日志上传进度回调结构体指针
//
//export OpenIM_UploadLogs
func OpenIM_UploadLogs(cb *C.base_callback_t, operationID *C.char, line C.int, ex *C.char, progress *C.upload_log_progress_t) {
	open_im_sdk.UploadLogs(newBaseCallback(cb), C.GoString(operationID), int(line), C.GoString(ex), newUploadLogProgress(progress))
}

// OpenIM_Logs 写入日志（异步）
//
//export OpenIM_Logs
func OpenIM_Logs(cb *C.base_callback_t, operationID *C.char, logLevel C.int, file *C.char, line C.int, msgs, err, keyAndValue *C.char) {
	open_im_sdk.Logs(newBaseCallback(cb), C.GoString(operationID), int(logLevel), C.GoString(file), int(line), C.GoString(msgs), C.GoString(err), C.GoString(keyAndValue))
}

// OpenIM_UploadFile 上传文件（异步）
// 参数:
//   - progress: 文件上传回调结构体指针
//
//export OpenIM_UploadFile
func OpenIM_UploadFile(cb *C.base_callback_t, operationID, req *C.char, progress *C.upload_file_callback_t) {
	// 将 C 上传回调结构体包装为 Go 接口实现
	open_im_sdk.UploadFile(newBaseCallback(cb), C.GoString(operationID), C.GoString(req), newUploadFileCallback(progress))
}
