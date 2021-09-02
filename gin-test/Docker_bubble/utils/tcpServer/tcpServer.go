package tcpServer

import (
	"errors"
	"fmt"
	"net"
	"syscall"
)

//TcpInfo 连接信息
type TcpInfo struct {
	host string //ip号
	port string //端口号
}

type tReadFunc func(adder string, length int, msg []byte, err error) // 阅读函数类型
type tDeferFunc func(conn net.Conn)                                  // 下线函数类型

var maxRead int = 4095

//NewTcpInfo 初始化tcp信息
func NewTcpInfo(hostStr string, portStr string) *TcpInfo {

	tcpInfo := &TcpInfo{
		host: hostStr,
		port: portStr,
	}

	return tcpInfo
}

//InitServer 初始化网络
func (tcpInfo *TcpInfo) InitServer() (listenter net.Listener, err error) {
	hostAndPort := tcpInfo.host + ":" + tcpInfo.port
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	if err != nil {
		return nil, err
	}
	listener, err := net.ListenTCP("tcp", serverAddr)
	if err != nil {
		return nil, err
	}
	return listener, nil
}

//
//  funcName       :  ConnectionHandler 维持链接
//  参数 conn      : 连接者
//  参数 funcRead  : 消息处理函数
//  参数 funcDefer : 下线处理函数
//
func (tcpInfo *TcpInfo) ConnectionHandler(conn net.Conn, funcRead tReadFunc, funcDefer tDeferFunc) {

	if funcRead == nil {
		funcRead = ReadMsg
	}
	if funcDefer == nil {
		funcDefer = Defer
	}

	connFrom := conn.RemoteAddr().String()
	//println("Connection from: ", connFrom)

	SendMsg(conn, []byte("hello!"))

	for {
		var ibuf []byte = make([]byte, maxRead+1)
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0 // to prevent overflow
		switch err {
		case nil:
			funcRead(connFrom, length, ibuf, err)
		case syscall.EAGAIN: // try again
			continue
		default:
			goto DISCONNECT
		}
	}
DISCONNECT:
	funcDefer(conn)
	conn.Close()
}

//
//  funcName     :  SendMsg 发送信息
//  参数 conn    :  连接者
//  参数 funcTmp :  消息
//
func SendMsg(to net.Conn, Msg []byte) (err error) {

	wrote, err := to.Write(Msg)
	if err != nil {
		err = errors.New("Write: wrote " + string(wrote) + " bytes.")
	}
	return
}

//ReadMsg
func ReadMsg(adder string, length int, msg []byte, err error) {
	if length > 0 {
		print("<", length, ":")
		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
			fmt.Printf("%c", msg[i])
		}
		print(">")
		fmt.Println()
	}
}

//Defer 下线处理函数
func Defer(conn net.Conn) {
	fmt.Println("logout from: " + conn.RemoteAddr().String())
	//默认为空
}
