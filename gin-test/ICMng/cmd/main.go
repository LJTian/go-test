package main

import (
	"fmt"
	"gin_test/ICMng/pkg/webMng"
)

func main() {
	// 启动程序
	fmt.Println("启动")
	webMng.InitWebMng()
	webMng.StartWebMng(":8009")
}
