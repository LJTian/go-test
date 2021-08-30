package main

import (
	. "gin_test/WebMng-1.0/utils/log"
	"gin_test/WebMng-1.0/utils/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initWebMng() {
	LogPath := tools.GetCurrentDirectory()
	LogFile := LogPath + "/../log/WebMng.log"
	InitLog(LogFile, LOG_DEBUG)
	TlogPrintf(LOG_INFO, "日志注册成功,日志文件目录为:[%s]\n", LogFile)
}

func StartWebMng(addr string) {
	// 1.创建路由
	r := gin.Default()
	r.GET("/", index)
	r.LoadHTMLFiles(tools.GetCurrentDirectory() + "/../html/index.html")
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	v1 := r.Group("/Svr")
	// {} 是书写规范
	{
		v1.POST("/Start", SvrStart)
		v1.POST("/Stop", SvrStop)
		v1.POST("/Clear", SvrClearLog)
	}
	v2 := r.Group("/JmSvr")
	{
		v2.POST("/Start", JmSvrStart)
		v2.POST("/Stop", JmSvrStop)
		v2.POST("/Clear", JmSvrClearLog)
	}
	TlogPrintf(LOG_INFO, "WEB管理页面加载成功, 监听地址为:[%s]\n", addr)
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(addr)

}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "")
}

func main() {
	//	获取系统环境变量
	initWebMng()
	// 启动管理页面
	StartWebMng(":8006")
}
