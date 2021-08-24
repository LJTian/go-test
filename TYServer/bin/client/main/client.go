package main

import (
	"TYServer/model/message"
	"TYServer/utils/tools"
	_ "bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	_ "os"
	_ "strings"
	"syscall"
	_ "time"
)

var IP string
var maxRead int = 20480

func init() {
	flag.StringVar(&IP, "IP", "default", "log in user")
}

func main() {

	ch := make(chan int)

	flag.Parse() //暂停获取参数

	var messageOb message.PackInfo

	messageOb.Code = "SYS0001"
	messageOb.PackType = "12332121"
	messageOb.Data = []byte("10002")
	messageOb.Lenth = 53
	messageOb.Receer = "10001"
	messageOb.Sender = "10002"
	messageOb.UUID = "123321"
	messageOb.Time = tools.LocalTime()

	data, err := json.Marshal(messageOb)
	if err != nil {
		//fmt.Println("json.Marshal err=", err)
	}
	go func() {
		for i := 0; i < 20; i++ {
			url := "127.0.0.1:14333"
			//fmt.Println(url)
			conn, err := net.Dial("tcp", url)
			if err != nil {
				fmt.Println("client dial err=", err)
				return
			}
			go sendMsg(conn, data)
		}
	}()

	<-ch
	// }
}
func sendMsg(conn net.Conn, data []byte) {

	_, err := conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write err=", err)
	}

	for {
		var ibuf []byte = make([]byte, maxRead+1)
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0 // to prevent overflow
		switch err {
		case nil:
			handleMsg(length, err, ibuf)
			//conn.Close()
		case syscall.EAGAIN: // try again
			continue
		default:
			return
		}
	}
}

//handleMsg 读取器信息
func handleMsg(length int, err error, msg []byte) {
	var strBuff string
	if length > 0 {
		strBuff = fmt.Sprintf("%s<%06d:", strBuff, length)
		strBuff = fmt.Sprintf("%s%s>", strBuff, string(msg[:length]))
	}

	//fmt.Println(strBuff)
}
