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
#include <stdlib.h>
*/
import "C"
import "unsafe"

// OpenIM_FreeString 释放由 SDK 返回的 C 字符串内存。
// 鸿蒙侧调用 SDK 返回字符串的函数后，使用完毕需调用此函数释放内存，避免内存泄漏。
//
//export OpenIM_FreeString
func OpenIM_FreeString(s *C.char) {
	if s != nil {
		C.free(unsafe.Pointer(s))
	}
}
