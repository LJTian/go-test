package log

import (
	"fmt"
	"testing"
)

func TestInitLog(t *testing.T) {

	fmt.Println("开始测试！")
	InitLog("E:/main/goland/TYServer/utils/log/test.log", 2)

	TlogPrintln(LOG_INFO, "信息描述")
	TlogPrintln(LOG_ERROR, "错误信息")
	TlogPrintln(LOG_DEBUG, "调试信息")

	TlogPrintf(LOG_INFO, "%d%s\n", 100, "函数进入")
	TlogPrintf(LOG_ERROR, "%d%s\n", 200, "错误信息")
	TlogPrintf(LOG_DEBUG, "%d%s\n", 300, "函数进入")

	fmt.Println("结束测试！")
}
