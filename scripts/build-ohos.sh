#!/bin/bash
# ==============================================================================
# 鸿蒙(HarmonyOS) C 共享库编译脚本
#
# 前置条件：
#   1. 已安装 ohos_golang_go（OpenHarmony 官方 Go fork）
#      下载地址: https://gitcode.com/openharmony-sig/ohos_golang_go
#      编译方法: 下载 Go 1.20.7+ 作为 GOROOT_BOOTSTRAP，然后编译 ohos_golang_go 源码
#   2. 已下载鸿蒙 Command Line Tools（含 OHOS SDK / NDK）
#      下载地址: https://developer.huawei.com/consumer/cn/download/
#   3. 编译环境: Ubuntu 18.04+ (glibc 2.27+)
#
# 使用方法：
#   方式1: 直接运行脚本，通过环境变量指定 SDK 路径
#     OHOS_SDK_PATH=/path/to/openharmony/native bash scripts/build-ohos.sh
#   方式2: 通过 Makefile 运行
#     make ohos OHOS_SDK_PATH=/path/to/openharmony/native
#
# 备选方案: 若 GOOS=openharmony 的 c-shared 模式不支持，
#           可设置 USE_LINUX=1 使用 GOOS=linux GOARCH=arm64 方案
#     USE_LINUX=1 bash scripts/build-ohos.sh
# ==============================================================================

set -e

# 项目根目录
PROJECT_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$PROJECT_ROOT"

# 鸿蒙 NDK 工具链路径（需用户根据实际路径修改）
OHOS_SDK_PATH="${OHOS_SDK_PATH:-/path/to/command-line-tools/sdk/default/openharmony/native}"

# 输出目录
OUTPUT_DIR="build/ohos/arm64-v8a"
OUTPUT_LIB="${OUTPUT_DIR}/libopenim_sdk.so"

# 创建输出目录
mkdir -p "$OUTPUT_DIR"

echo "===========> Building HarmonyOS C shared library"

# 检查 Go 版本
if ! command -v go &> /dev/null; then
    echo "ERROR: Go is not installed. Please install ohos_golang_go first."
    echo "  https://gitcode.com/openharmony-sig/ohos_golang_go"
    exit 1
fi

GO_VERSION=$(go version)
echo "Go version: $GO_VERSION"

# 检查鸿蒙 NDK 工具链
if [ ! -f "${OHOS_SDK_PATH}/llvm/bin/clang" ]; then
    echo "ERROR: HarmonyOS NDK not found at: ${OHOS_SDK_PATH}"
    echo "Please set OHOS_SDK_PATH to the correct path."
    echo "  export OHOS_SDK_PATH=/path/to/command-line-tools/sdk/default/openharmony/native"
    exit 1
fi

# 设置鸿蒙 NDK 交叉编译工具链
export CC="${OHOS_SDK_PATH}/llvm/bin/aarch64-unknown-linux-ohos-clang"
export CXX="${OHOS_SDK_PATH}/llvm/bin/aarch64-unknown-linux-ohos-clang++"
export AR="${OHOS_SDK_PATH}/llvm/bin/llvm-ar"

# 验证工具链文件存在
for tool in "$CC" "$CXX" "$AR"; do
    if [ ! -f "$tool" ]; then
        echo "ERROR: Tool not found: $tool"
        exit 1
    fi
done

# 设置 cgo 编译标志
export CGO_ENABLED=1
export CGO_CFLAGS="-g -O2 --target=aarch64-linux-ohos --sysroot=${OHOS_SDK_PATH}/sysroot"
export CGO_LDFLAGS="--target=aarch64-linux-ohos -fuse-ld=lld"

# 选择编译目标平台
if [ "${USE_LINUX:-0}" = "1" ]; then
    # 备选方案: 使用 GOOS=linux GOARCH=arm64
    echo "Using GOOS=linux GOARCH=arm64 (fallback mode)"
    export GOOS=linux
    export GOARCH=arm64
else
    # 默认方案: 使用 ohos_golang_go 的 GOOS=openharmony
    echo "Using GOOS=openharmony GOARCH=arm64"
    export GOOS=openharmony
    export GOARCH=arm64
fi

echo "Build configuration:"
echo "  GOOS=$GOOS"
echo "  GOARCH=$GOARCH"
echo "  CGO_ENABLED=$CGO_ENABLED"
echo "  CC=$CC"
echo "  Output: $OUTPUT_LIB"
echo ""

# 编译为 c-shared 动态库
# -buildmode=c-shared: 生成 .so 动态库 + .h 头文件
# -trimpath: 移除编译路径信息
# -ldflags "-s -w": 去除调试信息，减小体积
go build -buildmode=c-shared -trimpath -ldflags "-s -w" \
    -o "$OUTPUT_LIB" \
    ./cmd/ohos/

echo ""
echo "===========> Build completed successfully!"
echo "Output files:"
echo "  Library: $OUTPUT_LIB"
echo "  Header:  ${OUTPUT_DIR}/libopenim_sdk.h"
echo ""
echo "To use in HarmonyOS project:"
echo "  1. Copy libopenim_sdk.so to your project's libs/arm64-v8a/ directory"
echo "  2. Copy libopenim_sdk.h to your project's include/ directory"
echo "  3. In CMakeLists.txt:"
echo "     target_link_libraries(entry PRIVATE \$\{NATIVERENDER_ROOT_PATH\}/../../../libs/arm64-v8a/libopenim_sdk.so)"
echo "     target_include_directories(entry PRIVATE \$\{NATIVERENDER_ROOT_PATH\}/../../../include)"
