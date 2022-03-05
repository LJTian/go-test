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
	req := &pb.HelloRequest{
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
		}
		fmt.Println(data)
	}

}
