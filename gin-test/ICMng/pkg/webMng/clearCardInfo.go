package webMng

import (
	"fmt"
	. "github.com/LJTian/Tools/log"
	"github.com/LJTian/Tools/tools"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os/exec"
	"syscall"
)

var maxRead int = 20480

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
		shellHead.POST("/bootStart", bootStart)
		shellHead.POST("/bootStop", bootStop)
	}
	MsgHead := r.Group("/Msg")
	{
		MsgHead.POST("/SendMsg", sendMsg)
	}
	r.StaticFS("/html", http.Dir(tools.GetCurrentDirectory()+"/../../html"))
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

func bootStart(c *gin.Context) {
	TlogPrintln(LOG_DEBUG, "boot.sh"+" start")
	cmd := exec.Command("boot.sh", "start")
	err := cmd.Run()
	if err != nil {
		c.String(200, fmt.Sprintf("%v", err.Error()))
		TlogPrintln(LOG_ERROR, "err : ", err.Error())
	} else {
		c.String(http.StatusOK, "ok")
	}
}

func bootStop(c *gin.Context) {
	TlogPrintln(LOG_DEBUG, "boot.sh"+" stop")
	cmd := exec.Command("boot.sh", "stop")
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

//handleMsg 读取器信息
func handleMsg(length int, err error, msg []byte) []byte {
	var strBuff string
	if length > 0 {
		strBuff = fmt.Sprintf("%s", string(msg[2:length]))
	}
	return []byte(strBuff)
}

func sendMsg(c *gin.Context) {

	msg, ok := c.GetPostForm("msg")
	if !ok {
		TlogPrintln(LOG_ERROR, "err : GetPostForm Err")
		c.String(200, fmt.Sprintf("err : GetPostForm Err"))
	}

	var resBuf []byte
	ip, _ := c.GetPostForm("ip")
	port, _ := c.GetPostForm("port")
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}

	// 添加报文长度
	_, err = conn.Write(tools.StatisticalLen(msg, 2))
	if err != nil {
		fmt.Println("conn.Write err=", err)
	}

	for {
		var ibuf []byte = make([]byte, maxRead+1)
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0 // to prevent overflow
		switch err {
		case nil:
			resBuf = handleMsg(length, err, ibuf)
			TlogPrintln(LOG_DEBUG, "收到的报文为:\n", string(resBuf))
			c.String(http.StatusOK, string(resBuf))
			conn.Close()
			break
		case syscall.EAGAIN: // try again
			continue
		default:
			TlogPrintln(LOG_ERROR, "未收到报文")
			return
		}
	}
}
