package main

import "go_test/internal_test/mainDir1/pkg"

func main() {
	// internal.PrintfLog() 失败！语言限制无法调用
	pkg.PrintfLog()

}
