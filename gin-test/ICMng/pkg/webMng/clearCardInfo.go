package webMng

import (
	"fmt"
	. "github.com/LJTian/Tools/log"
	"github.com/LJTian/Tools/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

func InitWebMng() {
	LogPath := tools.GetCurrentDirectory()
	LogFile := LogPath + "/../../log/WebMng.log"
	InitLog(LogFile, LOG_DEBUG)
	TlogPrintf(LOG_INFO, "日志注册成功,日志文件目录为:[%s]\n", LogFile)
}

func StartWebMng(addr string) {
	// 1.创建路由
	r := gin.Default()
	r.GET("/", index)
	shellHead := r.Group("/shell")
	{
		shellHead.POST("/clearCardInfo", clearCardInfo)
		shellHead.POST("/genCard", genCard)
	}
	r.LoadHTMLFiles(tools.GetCurrentDirectory() + "/../../html/index.html")
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response

	TlogPrintf(LOG_INFO, "WEB管理页面加载成功, 监听地址为:[%s]\n", addr)
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(addr)

}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "")
}

func clearCardInfo(c *gin.Context) {
	shellPwd := "/home/card_saas/shell/card_clear/"
	shellName := "card_clear.sh"
	TlogPrintln(LOG_DEBUG, "sh "+shellPwd+shellName)
	cmd := exec.Command("sh", shellPwd+shellName)
	err := cmd.Run()
	if err != nil {
		c.String(200, fmt.Sprintf("%v", err.Error()))
		TlogPrintln(LOG_ERROR, "err : ", err.Error())
	} else {
		c.String(http.StatusOK, "ok")
	}
}

func genCard(c *gin.Context) {
	shellPwd := "/home/card_saas/shell/genCard/"
	shellName := "genCard_Web.sh"
	TlogPrintln(LOG_DEBUG, "sh "+shellPwd+shellName)
	cmd := exec.Command("sh", shellPwd+shellName)
	err := cmd.Run()
	if err != nil {
		c.String(200, fmt.Sprintf("%v", err.Error()))
		TlogPrintln(LOG_ERROR, "err : ", err.Error())
	} else {
		c.String(http.StatusOK, "ok")
	}
}
