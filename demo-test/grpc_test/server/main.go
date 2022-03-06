package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	pb "go_test/grpc_test/protocol_out"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedHelloServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {

	fmt.Printf("Received: %s[%s]:%s\n", in.GetReq().TransCode, in.GetReq().RespCode, in.GetReq().Data)

	resp := &pb.HelloResponse{
		Resp: &pb.Msg{
			TransCode: "000002",
			Data:      "我是服务器回应",
			RespCode:  "00",
		},
	}
	return resp, nil
}

// SayHello implements helloworld.GreeterServer
func (s *server) LotsOfReplies(in *pb.HelloRequest, stream pb.HelloService_LotsOfRepliesServer) error {

	fmt.Printf("Received: %s[%s]:%s\n", in.GetReq().TransCode, in.GetReq().RespCode, in.GetReq().Data)

	for i := 0; i < 10; i++ {
		resp := &pb.HelloResponseStream{
			Resp: &pb.Msg{
				TransCode: fmt.Sprintf("%06d", i+1),
				Data:      fmt.Sprintf("我是服务器回应%d", i),
				RespCode:  "00",
			},
		}
		stream.Send(resp)
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (s *server) LotsOfGreetings(stream pb.HelloService_LotsOfGreetingsServer) error {

	for {
		data, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("断开链接")
			break
		} else if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(data)
	}

	return nil
}

func (s *server) BidiHello(stream pb.HelloService_BidiHelloServer) error {

	// 建立链接
	fmt.Println("建立链接")
	clientClossFlag := make(chan int, 0)

	// 接受线程
	go func(<-chan int) {
		for {
			data, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("断开链接")
				clientClossFlag <- 1
				return
			} else if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(data)
		}
		time.Sleep(1 * time.Second)
	}(clientClossFlag)

	//发送处理
	var i int

for1:
	for {
		select {
		case <-clientClossFlag:
			fmt.Println("退出")
			break for1
		default:
			resp := &pb.HelloResponseStream{
				Resp: &pb.Msg{
					TransCode: fmt.Sprintf("%06d", i+1),
					Data:      fmt.Sprintf("我是服务器回应%d", i),
					RespCode:  "00",
				},
			}
			err := stream.Send(resp)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
		i++
		time.Sleep(1 * time.Second)
	}
	time.Sleep(5 * time.Second)

	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
