package tcpProess

import (
	"TYServer/model/message"
	"TYServer/model/userInfo"
	"TYServer/pkg/detailed"
	"TYServer/pkg/peripheral"
	"TYServer/utils/conf"
	. "TYServer/utils/log"
	"TYServer/utils/tcpServer"
	"TYServer/utils/uuid"
	"fmt"
	"net"
)

//InitTcpProess 初始化tcp 进程
func MainTcpProess(ip string, port string) (err error) {

	//注册所有函数
	detailed.MainDetailed()

	//初始化管道
	peripheral.ReturnMsg = make(chan *message.PackInfo, conf.CFG.FIFOLen)
	//初始化发送函数携程池
	for i := 0; i <= conf.CFG.SendThreadNum; i++ {
		go peripheral.SendMsgFunc()
	}

	//初始化网络
	tcp := tcpServer.NewTcpInfo(ip, port)
	listener, err := tcp.InitServer()
	if err != nil {
		TlogPrintln(LOG_ERROR, err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			TlogPrintln(LOG_ERROR, string(err.Error()))
		}
		//登记连接信息
		SetUserInfoToList(conn)
		//开始监听网络
		go tcp.ConnectionHandler(conn, ReadMsg, Defer)
	}
}

//ReadMsg 消息处理函数
func ReadMsg(adder string, length int, msg []byte, err error) {
	var strBuff string
	if length > 0 {
		strBuff = fmt.Sprintf("%s<%06d:", strBuff, length)
		strBuff = fmt.Sprintf("%s%s>", strBuff, string(msg[:length]))
	}
	TlogPrintf(LOG_DEBUG, "从客户端收到的数据为[%s] \n", strBuff)
	inMsg := msg[:length]

	//读包
	peripheral.MainPeripheral(&inMsg, adder)

}

//Defer 下线处理函数
func Defer(conn net.Conn) {
	TlogPrintln(LOG_DEBUG, "logout from: "+conn.RemoteAddr().String())
	key := conn.RemoteAddr().String()
	TlogPrintf(LOG_DEBUG, "IP :[%v] 断开建立连接\n", key)

	userInfo.OnlineUser.DeleteNode2List(key)
}

//SetUserInfoToList 将用户信息添加到在线列表中
func SetUserInfoToList(conn net.Conn) (err error) {

	uuidStr := uuid.CreateUUID()
	ip := conn.RemoteAddr().String()
	userInfoNode := userInfo.NewUserInfo(uuidStr, "", ip, "", conn, 1)
	TlogPrintf(LOG_DEBUG, "IP :[%v] 建立连接\n", ip)
	userInfo.OnlineUser.AddNode2List(userInfoNode)

	return nil
}
