package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

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
