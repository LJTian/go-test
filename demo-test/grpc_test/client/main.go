package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
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
	_ = &pb.HelloRequest{
		Req: &pb.Msg{
			TransCode: "000001",
			Data:      "我是客户端请求",
			RespCode:  "98",
		},
	}

	//respData, err := client.SayHello(ctx, req)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(respData)

	//respData, err := client.LotsOfReplies(ctx, req)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//for {
	//	data, err := respData.Recv()
	//	if err == io.EOF {
	//		break
	//	} else if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println(data)
	//}

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

	time.Sleep(5 * time.Second)
}
