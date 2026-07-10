# ==============================================================================
# HarmonyOS C Shared Library Build Script (Windows PowerShell)
#
# Prerequisites:
#   1. Standard Go 1.20+ (no need for ohos_golang_go)
#      https://go.dev/dl/
#   2. HarmonyOS DevEco Studio or Command Line Tools (with NDK)
#      https://developer.huawei.com/consumer/cn/download/
#   3. Windows 10/11 64-bit
#
# Usage:
#   powershell -ExecutionPolicy Bypass -File scripts\build-ohos.ps1
#   powershell -ExecutionPolicy Bypass -File scripts\build-ohos.ps1 -OHOS_SDK_PATH "C:\path\to\native"
#   $env:OHOS_SDK_PATH = "C:\path\to\native"; powershell -ExecutionPolicy Bypass -File scripts\build-ohos.ps1
#   make ohos
#
# Optional: use GOOS=openharmony (requires ohos_golang_go, c-shared may not work):
#   powershell -ExecutionPolicy Bypass -File scripts\build-ohos.ps1 -USE_OHOS
# ==============================================================================

param(
    [string]$OHOS_SDK_PATH = "",
    [switch]$USE_OHOS
)

$ErrorActionPreference = "Stop"

$PROJECT_ROOT = Resolve-Path "$PSScriptRoot\.."
Set-Location $PROJECT_ROOT

# Get OHOS_SDK_PATH from env var if not specified
if ([string]::IsNullOrEmpty($OHOS_SDK_PATH)) {
    $OHOS_SDK_PATH = $env:OHOS_SDK_PATH
}

# Auto-detect NDK path from common install locations
if ([string]::IsNullOrEmpty($OHOS_SDK_PATH)) {
    $devecoPaths = @(
        "${env:LOCALAPPDATA}\OpenHarmony\command-line-tools\sdk\default\openharmony\native",
        "${env:LOCALAPPDATA}\Huawei\Sdk\default\openharmony\native",
        "C:\Program Files\DevEco Studio\sdk\default\openharmony\native",
        "C:\Program Files (x86)\DevEco Studio\sdk\default\openharmony\native"
    )
    foreach ($path in $devecoPaths) {
        if (Test-Path "$path\llvm\bin\clang.exe") {
            $OHOS_SDK_PATH = $path
            break
        }
    }
}

if ([string]::IsNullOrEmpty($OHOS_SDK_PATH)) {
    Write-Host "ERROR: HarmonyOS NDK path not found." -ForegroundColor Red
    Write-Host "Please specify via parameter or environment variable:"
    Write-Host '  powershell -File scripts\build-ohos.ps1 -OHOS_SDK_PATH "C:\path\to\native"'
    Write-Host '  $env:OHOS_SDK_PATH = "C:\path\to\native"; powershell -File scripts\build-ohos.ps1'
    exit 1
}

# Output directory
$OUTPUT_DIR = "build\ohos\arm64-v8a"
$OUTPUT_LIB = "$OUTPUT_DIR\libopenim_sdk.so"

if (-not (Test-Path $OUTPUT_DIR)) {
    New-Item -ItemType Directory -Path $OUTPUT_DIR -Force | Out-Null
}

Write-Host "===========> Building HarmonyOS C shared library" -ForegroundColor Cyan

# Check Go installation
$goCmd = Get-Command go -ErrorAction SilentlyContinue
if (-not $goCmd) {
    Write-Host "ERROR: Go is not installed or not in PATH." -ForegroundColor Red
    Write-Host "Please install ohos_golang_go first:"
    Write-Host "  https://gitcode.com/openharmony-sig/ohos_golang_go"
    exit 1
}

$GO_VERSION = & go version
Write-Host "Go version: $GO_VERSION"

# Check HarmonyOS NDK clang
# On Windows, the NDK provides clang.exe (not aarch64-unknown-linux-ohos-clang.exe).
# The target is specified via --target flag in CGO_CFLAGS.
$CLANG_PATH = "$OHOS_SDK_PATH\llvm\bin\clang.exe"
if (-not (Test-Path $CLANG_PATH)) {
    Write-Host "ERROR: HarmonyOS NDK clang not found at: $CLANG_PATH" -ForegroundColor Red
    Write-Host "Please verify OHOS_SDK_PATH is correct."
    exit 1
}

# Set NDK cross-compile toolchain
# Use clang.exe directly; --target=aarch64-linux-ohos is set in CGO_CFLAGS/CGO_LDFLAGS
$env:CC = "$OHOS_SDK_PATH\llvm\bin\clang.exe"
$env:CXX = "$OHOS_SDK_PATH\llvm\bin\clang++.exe"
$env:AR = "$OHOS_SDK_PATH\llvm\bin\llvm-ar.exe"

# Verify toolchain files exist
$tools = @($env:CC, $env:CXX, $env:AR)
foreach ($tool in $tools) {
    if (-not (Test-Path $tool)) {
        Write-Host "ERROR: Tool not found: $tool" -ForegroundColor Red
        exit 1
    }
}

# Set cgo flags (use forward slashes in sysroot path for clang compatibility)
$env:CGO_ENABLED = "1"
$sysrootPath = ($OHOS_SDK_PATH -replace '\\', '/') + "/sysroot"
$env:CGO_CFLAGS = "-g -O2 --target=aarch64-linux-ohos --sysroot=$sysrootPath"
$env:CGO_LDFLAGS = "--target=aarch64-linux-ohos -fuse-ld=lld"

# Select target platform
# Default: GOOS=linux (works with standard Go, community verified)
# Optional: -USE_OHOS flag uses GOOS=openharmony (requires ohos_golang_go, c-shared may not work)
if ($USE_OHOS) {
    Write-Host "Using GOOS=openharmony GOARCH=arm64 (requires ohos_golang_go)"
    $env:GOOS = "openharmony"
    $env:GOARCH = "arm64"
} else {
    Write-Host "Using GOOS=linux GOARCH=arm64 (default, works with standard Go)"
    $env:GOOS = "linux"
    $env:GOARCH = "arm64"
}

Write-Host "Build configuration:"
Write-Host "  GOOS=$env:GOOS"
Write-Host "  GOARCH=$env:GOARCH"
Write-Host "  CGO_ENABLED=$env:CGO_ENABLED"
Write-Host "  CC=$env:CC"
Write-Host "  Output: $OUTPUT_LIB"
Write-Host ""

# Build c-shared library (.so + .h)
& go build -buildmode=c-shared -trimpath -ldflags "-s -w" `
    -o $OUTPUT_LIB `
    ./cmd/ohos/

if ($LASTEXITCODE -ne 0) {
    Write-Host "ERROR: Build failed with exit code $LASTEXITCODE" -ForegroundColor Red
    exit $LASTEXITCODE
}

# Post-process: inline callback_types.h into the generated header
# The generated libopenim_sdk.h contains #include "callback_types.h" which is not
# distributed alongside it. Replace with actual content to make the header self-contained.
$generatedHeader = "$OUTPUT_DIR\libopenim_sdk.h"
$callbackTypesPath = "$PROJECT_ROOT\cmd\ohos\callback_types.h"
if ((Test-Path $generatedHeader) -and (Test-Path $callbackTypesPath)) {
    $headerContent = [System.IO.File]::ReadAllText((Resolve-Path $generatedHeader))
    $callbackContent = [System.IO.File]::ReadAllText((Resolve-Path $callbackTypesPath))
    $headerContent = $headerContent.Replace('#include "callback_types.h"', $callbackContent)
    [System.IO.File]::WriteAllText((Resolve-Path $generatedHeader), $headerContent)
    Write-Host "  Inlined callback_types.h into generated header"
}

Write-Host ""
Write-Host "===========> Build completed successfully!" -ForegroundColor Green
Write-Host "Output files:"
Write-Host "  Library: $OUTPUT_LIB"
Write-Host "  Header:  $OUTPUT_DIR\libopenim_sdk.h (self-contained)"
Write-Host ""
Write-Host "To use in HarmonyOS project:"
Write-Host "  1. Copy libopenim_sdk.so to your project's libs/arm64-v8a/ directory"
Write-Host "  2. Copy libopenim_sdk.h to your project's include/ directory"
Write-Host "  3. In CMakeLists.txt:"
Write-Host '     target_link_libraries(entry PRIVATE ${NATIVERENDER_ROOT_PATH}/../../../libs/arm64-v8a/libopenim_sdk.so)'
Write-Host '     target_include_directories(entry PRIVATE ${NATIVERENDER_ROOT_PATH}/../../../include)'
