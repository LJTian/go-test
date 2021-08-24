package tools

import (
	"os"
	"path/filepath"
	"time"
)

// LocalTime 获取时间
func LocalTime() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

// GetCurrentDirectory 获取执行目录
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return dir
}
