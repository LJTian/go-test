package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"

	pb "go_test/grpc_test/protocol_out"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

//一元函数
func TestSayHello(ctx context.Context, client pb.HelloServiceClient) {

	req := &pb.HelloRequest{
		Req: &pb.Msg{
			TransCode: "000001",
			Data:      "我是客户端请求",
			RespCode:  "98",
		},
	}

	respData, err := client.SayHello(ctx, req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(respData)
}

// 服务端流式函数
func TestLotsOfReplies(ctx context.Context, client pb.HelloServiceClient) {

	req := &pb.HelloRequest{
		Req: &pb.Msg{
			TransCode: "000001",
			Data:      "我是客户端请求",
			RespCode:  "98",
		},
	}

	respData, err := client.LotsOfReplies(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	for {
		data, err := respData.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(data)
	}
}

// 客户端流式函数
func TestLotsOfGreetings(ctx context.Context, client pb.HelloServiceClient) {

	respData, err := client.LotsOfGreetings(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer respData.CloseSend()
	for i := 0; i < 10; i++ {
		req := &pb.HelloRequestStream{
			Req: &pb.Msg{
				TransCode: fmt.Sprintf("%06d", i+1),
				Data:      fmt.Sprintf("我是客户端发送%d", i),
				RespCode:  "00",
			},
		}
		fmt.Println(req)
		err = respData.Send(req)
		if err != nil {
			fmt.Println(err)
			break
		}
		time.Sleep(1 * time.Second)
	}
}

// 双流式
func TestBidiHello(ctx context.Context, client pb.HelloServiceClient) {

	req, err := client.BidiHello(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer req.CloseSend()

	go func() {
		for {
			data, err := req.Recv()
			if err == io.EOF {
				fmt.Println("断开链接")
				return
			} else if err != nil {
				fmt.Println()
				return
			}
			fmt.Println(data.Resp)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < 10; i++ {
		reqDta := &pb.HelloRequestStream{
			Req: &pb.Msg{
				TransCode: fmt.Sprintf("%06d", i+1),
				Data:      fmt.Sprintf("我是客户端发送%d", i),
				RespCode:  "00",
			},
		}
		//fmt.Println(req)
		err = req.Send(reqDta)
		if err != nil {
			fmt.Println(err)
			break
		}
		time.Sleep(1 * time.Second)
	}

}

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	TestSayHello(ctx, client)
	TestLotsOfReplies(ctx, client)
	TestLotsOfGreetings(ctx, client)
	TestBidiHello(ctx, client)

	time.Sleep(5 * time.Second)
}
