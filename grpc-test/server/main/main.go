//go:generate protoc -I ../protocol --go_out=plugins=grpc:../protocol ../protocol/protocol.proto
// 参数文件生成处

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/golang/protobuf/proto"

	"google.golang.org/grpc"
	pb "grpc_test/protocol" //传输文件
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	//pb.UnimplementedGreeterServer
	pb.UnimplementedRouteGuideServer
	savedFeatures []*pb.Feature // read-only after initialized

	mu sync.Mutex // protects routeNotes
	//routeNotes map[string][]*pb.RouteNote
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// 可以被远程调用的函数
// HelloReplyTian implements helloworld.GreeterServer
func (s *server) SayTian(ctx context.Context, in *pb.HelloRequestTian) (*pb.HelloReplyTian, error) {
	log.Printf("Received: %v Say [ %v ] ", in.GetName(), in.GetSayMsg())
	return &pb.HelloReplyTian{Message: "Hello " + in.GetName(), Name: "server"}, nil
}

// 测试简单rpc
func (s *server) GetFeature(tx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range s.savedFeatures {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	return &pb.Feature{Name: "NULL", Location: &pb.Point{Latiude: 0, Longitude: 0}}, nil
}

// 主函数
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//pb.RegisterGreeterServer(s, &server{})
	pb.RegisterRouteGuideServer(s, &server{})

	// 赋值
	var TmpServer server
	TmpServer.savedFeatures[0] = &pb.Feature{
		Name: "",
		Location: &pb.Point{
			Latiude:   1,
			Longitude: 2,
		},
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
