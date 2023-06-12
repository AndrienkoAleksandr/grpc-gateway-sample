package main

import (
	"context"
	"log"

	pb "example.com/m/gen/go/your/service/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedYourServiceServer
}

func (server) Echo(ctx context.Context, msg *pb.StringMessage) (*pb.StringMessage, error) {
	log.Printf("Echo: ", msg.Value)
	return msg, nil
}

func main() {

}