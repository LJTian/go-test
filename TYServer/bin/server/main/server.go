package main

import (
	"TYServer/proess/tcpProess"
	"TYServer/utils/conf"
	. "TYServer/utils/log"
	"TYServer/utils/tools"
	_ "fmt"
	_ "net"
	"runtime"
)

func init() {

	//判断系统类型加载配置文件
	sysType := runtime.GOOS
	switch sysType {
	case "linux":
		conf.NewConf(tools.GetCurrentDirectory() + "/../conf/TYserver.ini")
	case "windows":
		conf.NewConf(tools.GetCurrentDirectory() + "\\..\\conf\\TYserver.ini")
	}
	//zap.InitLogger(conf.CFG.LogFilePath)
	//defer zap.CloseLogger()
	//加载初始化日志文件
	InitLog(conf.CFG.LogFilePath, conf.CFG.LogLevel)
}
func main() {

	//初始化http进程
	//启动Tcp进程
	tcpProess.MainTcpProess(conf.CFG.Ip, conf.CFG.ServerPort)
}
