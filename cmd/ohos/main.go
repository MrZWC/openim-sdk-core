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

// Package main 是鸿蒙(HarmonyOS) C 共享库的入口包。
// 通过 go build -buildmode=c-shared 编译为 .so 动态库 + .h 头文件。
package main

import "C"

// main 函数是 c-shared 编译模式必需的入口，不需要做任何操作。
// Go 运行时会在库被加载时自动初始化（包括 open_im_sdk 包的 init 函数）。
func main() {}
