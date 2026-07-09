# Windows 原生编译鸿蒙 C 库适配计划

## 概述

将现有的鸿蒙编译流程从 Linux/bash 适配到 Windows 原生环境。创建 PowerShell 编译脚本，更新 Makefile 支持 Windows，并提供完整的 Windows 环境搭建指南。

## 当前状态分析

### 已有代码（无需修改）

* `cmd/ohos/*.go`（11 个文件）- cgo 导出层，Go 代码本身是平台无关的，无需改动

* `scripts/build-ohos.sh` - Linux bash 编译脚本，保留不动

* Makefile 中的 `ohos` 目标 - 调用 `bash scripts/build-ohos.sh`，需增加 Windows 分支

### 需要解决的问题

1. **编译脚本**：现有 `build-ohos.sh` 是 bash 脚本，Windows 原生无 bash
2. **Makefile**：`SHELL := /bin/bash` 和 `@bash scripts/build-ohos.sh` 在 Windows 上不可用
3. **NDK 工具链路径**：Windows 上 NDK 工具带 `.exe` 后缀，路径格式不同
4. **环境搭建指南**：ohos\_golang\_go 的安装指南是 Linux 版，需要 Windows 版步骤

## 实施计划

### 步骤 1：创建 PowerShell 编译脚本

**新建文件**：`scripts/build-ohos.ps1`

这是 `build-ohos.sh` 的 Windows PowerShell 等价脚本，功能完全相同：

```powershell
# 关键差异点：
# 1. 环境变量用 $env: 前缀设置
# 2. NDK 工具路径带 .exe 后缀
# 3. 路径分隔符用反斜杠（但 CGO_CFLAGS 中的 --sysroot 用正斜杠，clang 兼容）
# 4. mkdir 用 New-Item -ItemType Directory -Force
# 5. 文件存在检查用 Test-Path
```

脚本结构：

1. 接收 `OHOS_SDK_PATH` 参数（默认尝试从 DevEco Studio 安装路径推断）
2. 检查 go 命令是否可用
3. 检查 NDK 工具链文件是否存在（`aarch64-unknown-linux-ohos-clang.exe`）
4. 设置环境变量：`CC`、`CXX`、`AR`、`CGO_ENABLED`、`CGO_CFLAGS`、`CGO_LDFLAGS`、`GOOS`、`GOARCH`
5. 执行 `go build -buildmode=c-shared`
6. 输出结果路径和使用说明

关键环境变量设置：

```powershell
$env:CC = "$OHOS_SDK_PATH\llvm\bin\aarch64-unknown-linux-ohos-clang.exe"
$env:CXX = "$OHOS_SDK_PATH\llvm\bin\aarch64-unknown-linux-ohos-clang++.exe"
$env:AR = "$OHOS_SDK_PATH\llvm\bin\llvm-ar.exe"
$env:CGO_ENABLED = "1"
# --sysroot 路径用正斜杠避免转义问题，clang 在 Windows 上兼容正斜杠
$env:CGO_CFLAGS = "-g -O2 --target=aarch64-linux-ohos --sysroot=$($OHOS_SDK_PATH -replace '\\','/')/sysroot"
$env:CGO_LDFLAGS = "--target=aarch64-linux-ohos -fuse-ld=lld"
$env:GOOS = "openharmony"
$env:GOARCH = "arm64"
```

支持 `USE_LINUX` 参数作为备选方案（`GOOS=linux`）。

### 步骤 2：更新 Makefile

修改 `ohos` 目标，使其自动检测操作系统并调用对应脚本：

```makefile
## ohos: Build the HarmonyOS shared library (.so + .h)
.PHONY: ohos
ohos:
ifeq ($(OS),Windows_NT)
	@powershell -ExecutionPolicy Bypass -File scripts/build-ohos.ps1
else
	@bash scripts/build-ohos.sh
endif
```

注意：Makefile 的 `ifeq` 是 GNU Make 的条件判断，Windows 上如果安装了 GNU Make（如通过 Chocolatey 或 MSYS2），`OS` 环境变量在 Windows 上自动为 `Windows_NT`。

### 步骤 3：Go 代码无需修改

`cmd/ohos/` 目录下的所有 Go 源文件（cgo 导出层）是平台无关的。cgo 的 C 代码只使用了标准 C 库（`stdint.h`、`stdlib.h`），在 Windows 和 Linux 上都能编译。`go build -buildmode=c-shared` 在 Windows 上同样生成 `.so` + `.h`（Go 在 Windows 上的 c-shared 模式生成 `.so` 而非 `.dll`，因为目标平台是 `openharmony/arm64` 而非 `windows`）。

## 假设与决策

1. **ohos\_golang\_go 支持 Windows**：它是 Go 的 fork，Go 官方支持 Windows，fork 保留了 `src/make.bat`。用户需在 Windows 上用 `make.bat` 编译 ohos\_golang\_go
2. **鸿蒙 NDK 有 Windows 版本**：DevEco Studio 支持 Windows，内置 HarmonyOS SDK 包含 NDK 工具链（`aarch64-unknown-linux-ohos-clang.exe`）
3. **NDK 路径推断**：PowerShell 脚本默认尝试从 DevEco Studio 安装路径（`C:\Program Files\DevEco Studio\sdk`）推断 SDK 路径，用户也可通过参数覆盖
4. **CGO\_CFLAGS 中 sysroot 路径用正斜杠**：LLVM/Clang 在 Windows 上兼容正斜杠路径，避免反斜杠转义问题
5. **路径含空格处理**：若 NDK 安装在 `C:\Program Files\...` 等含空格路径，PowerShell 脚本中变量引用需正确处理
6. **不修改现有 build-ohos.sh**：保留 Linux 脚本不变，新增 Windows 脚本

## 产出文件清单

| 文件                           | 说明                              |
| ---------------------------- | ------------------------------- |
| `scripts/build-ohos.ps1`（新建） | Windows PowerShell 编译脚本         |
| `Makefile`（修改）               | ohos 目标增加 OS 检测，Windows 下调用 ps1 |

## 验证步骤

1. **环境验证**：在 Windows PowerShell 中执行 `go version`，确认使用的是 ohos\_golang\_go 编译的 Go
2. **NDK 验证**：确认 `aarch64-unknown-linux-ohos-clang.exe` 存在于 SDK 路径下
3. **编译执行**：`powershell -ExecutionPolicy Bypass -File scripts/build-ohos.ps1 -OHOS_SDK_PATH "C:\path\to\native"`
4. **产物验证**：确认生成 `build/ohos/arm64-v8a/libopenim_sdk.so` 和 `libopenim_sdk.h`
5. **头文件检查**：检查 `.h` 文件包含所有 `OpenIM_*` 导出函数声明

