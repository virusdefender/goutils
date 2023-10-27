/*
将 GOOS 和 GOARCH 的值转换为更易读的布尔值
*/

package goutils

import "runtime"

const (
	IsLinux   = runtime.GOOS == "linux"
	IsWindows = runtime.GOOS == "windows"
	IsDarwin  = runtime.GOOS == "darwin"

	Is64Bit = runtime.GOARCH == "amd64" || runtime.GOARCH == "arm64"
	Is32Bit = runtime.GOARCH == "386" || runtime.GOARCH == "arm"

	IsArm   = runtime.GOARCH == "arm" || runtime.GOARCH == "arm64"
	IsArm64 = runtime.GOARCH == "arm64"
	IsAmd64 = runtime.GOARCH == "amd64"
	Is386   = runtime.GOARCH == "386"
)
