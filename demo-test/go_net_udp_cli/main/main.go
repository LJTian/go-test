package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	address := "10.12.17.37:30000"
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	socket, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()

	for true {
		sendData := []byte("Hello server")
		_, err = socket.Write(sendData) // 发送数据
		if err != nil {
			fmt.Println("发送数据失败，err:", err)
			return
		}
		data := make([]byte, 4096)
		n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
		if err != nil {
			fmt.Println("接收数据失败，err:", err)
			return
		}
		fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
		time.Sleep(1 * time.Second)
	}
}
