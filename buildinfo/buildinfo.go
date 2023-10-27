/*
在 go build 时候 ldflags 来注入版本号等自定义变量

参考命令：
go build -ldflags "-X github.com/virusdefender/goutils/buildinfo.GitCommit=abcdefg -X github.com/virusdefender/goutils/buildinfo.Version=1.0"
*/

package buildinfo

import (
	"fmt"
	"strings"
)

var AppName = ""

var GitCommit = ""
var GitTag = ""

var Version = ""
var InternalVersion = ""

var BuildTime = ""
var BuildID = ""

var GoVersion = ""

// Extra 系列是预留的非版本信息相关字段，有需要可以直接读取
var Extra1 = ""
var Extra2 = ""
var Extra3 = ""

type BuildInfo struct {
	AppName string

	Version         string
	InternalVersion string

	GitCommit string
	GitTag    string

	BuildTime string
	BuildID   string

	GoVersion string
}

func (b *BuildInfo) String() string {
	// only output non-empty fields
	s := make([]string, 0, 10)
	if b.AppName != "" {
		s = append(s, fmt.Sprintf("AppName: %s", b.AppName))
	}
	if b.Version != "" {
		s = append(s, fmt.Sprintf("Version: %s", b.Version))
	}
	if b.InternalVersion != "" {
		s = append(s, fmt.Sprintf("InternalVersion: %s", b.InternalVersion))
	}
	if b.GitCommit != "" {
		s = append(s, fmt.Sprintf("GitCommit: %s", b.GitCommit))
	}
	if b.GitTag != "" {
		s = append(s, fmt.Sprintf("GitTag: %s", b.GitTag))
	}
	if b.BuildTime != "" {
		s = append(s, fmt.Sprintf("BuildTime: %s", b.BuildTime))
	}
	if b.BuildID != "" {
		s = append(s, fmt.Sprintf("BuildID: %s", b.BuildID))
	}
	if b.GoVersion != "" {
		s = append(s, fmt.Sprintf("GoVersion: %s", b.GoVersion))
	}
	return strings.Join(s, "; ")
}

func Get() *BuildInfo {
	return &BuildInfo{
		AppName: AppName,

		GitCommit: GitCommit,
		GitTag:    GitTag,

		Version:         Version,
		InternalVersion: InternalVersion,

		BuildTime: BuildTime,
		BuildID:   BuildID,

		GoVersion: GoVersion,
	}
}
