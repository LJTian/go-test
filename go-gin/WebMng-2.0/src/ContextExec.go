package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "gotest/utils/log"
	"gotest/utils/tools"
	"net/http"
	"os"
)

func SvrStart(c *gin.Context) {

	SvrName := "MISWinService_go"
	TlogPrintln( LOG_DEBUG, "net start" + SvrName )
	ok := tools.OsExe( "net","start", SvrName )
	if ok != nil {
		c.String(400,fmt.Sprintf("%v", ok.Error()))
	}else{
		c.String(http.StatusOK,"ok")
	}
}

func SvrStop(c *gin.Context) {

	SvrName := "MISWinService_go"
	TlogPrintln( LOG_DEBUG, "net stop" + SvrName )
	ok := tools.OsExe( "net","stop", SvrName )
	if ok != nil {
		c.String(400,fmt.Sprintf("%v", ok.Error()))
	}else{
		c.String(http.StatusOK,"ok")
	}
}

func SvrClearLog(c *gin.Context) {

	LogNamePath := os.Getenv("WORKDIR_GO") + "\\log\\"
	TlogPrintf( LOG_DEBUG, "%s 日志清理 成功", LogNamePath )
	ok := os.RemoveAll( LogNamePath )
	if ok != nil {
		c.String(400,fmt.Sprintf("%v", ok.Error()))
	}else{
		c.String(http.StatusOK,"ok")
	}
}

func JmSvrStart(c *gin.Context) {

	SvrName := "MISWinService"
	TlogPrintln( LOG_DEBUG, "net start" + SvrName )
	ok := tools.OsExe( "net","start", SvrName )
	if ok != nil {
		c.String(400,fmt.Sprintf("%v", ok.Error()))
	}else{
		c.String(http.StatusOK,"ok")
	}
}

func JmSvrStop(c *gin.Context) {

	SvrName := "MISWinService"
	TlogPrintln( LOG_DEBUG, "net stop" + SvrName )
	ok := tools.OsExe( "net","stop", SvrName )
	if ok != nil {
		c.String(400,fmt.Sprintf("%v", ok.Error()))
	}else{
		c.String(http.StatusOK,"ok")
	}
}

func JmSvrClearLog(c *gin.Context) {

	LogNamePath := os.Getenv("WORKDIR_NEWMIS") + "\\log\\"
	TlogPrintf( LOG_DEBUG, "%s 日志清理 成功", LogNamePath )
	ok := os.RemoveAll( LogNamePath )
	if ok != nil {
		c.String(400,fmt.Sprintf("%v", ok.Error()))
	}else{
		c.String(http.StatusOK,"ok")
	}
}