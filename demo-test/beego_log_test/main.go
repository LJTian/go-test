package main

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
)

func main() {
	log := logs.NewLogger()
	log1 := logs.NewLogger()

	fileName := fmt.Sprintf("{\"filename\":\"%s\",\"level\":7}", "tianlj.logs")
	fileName1 := fmt.Sprintf("{\"filename\":\"%s\",\"level\":7}", "tianlj1.logs")

	log.SetLogger(logs.AdapterFile, fileName)

	log1.SetLogger(logs.AdapterFile, fileName1)
	log.Debug("this is a debug message")
	log1.Debug("111")
}
