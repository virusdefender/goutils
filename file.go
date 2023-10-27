package goutils

import "os"

// 仅供简单判断，不处理符号链接、各种错误类型等问题
// XXXExists 函数其取反的结果不一定代表不存在
// 比如可能因为权限问题，无法访问某个路径，但是这个路径确实存在的

// 任意路径存在
func PathExists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

// 路径存在，而且是文件类型
func FileExists(name string) bool {
	stat, err := os.Stat(name)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}
