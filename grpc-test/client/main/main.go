// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "grpc_test/protocol"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRouteGuideClient(conn)

	// Contact the server and print out its response.
	//name := defaultName
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetFeature(ctx, &pb.Point{Latiude: 1, Longitude: 2})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Latiude: %v Longitude %v ", r.GetLocation().Latiude, r.GetLocation().Longitude)
}
