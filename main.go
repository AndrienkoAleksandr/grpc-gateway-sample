package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "example.com/m/gen/go/your/service/v1"

	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint",  "localhost:50051", "gRPC server endpoint")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedYourServiceServer
}

func (server) Echo(ctx context.Context, msg *pb.StringMessage) (*pb.StringMessage, error) {
	log.Printf("Echo: %s", msg.Value)
	return msg, nil
}

func runGateWayProxy() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterYourServiceHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	fmt.Println("Run gRPC gateway...")
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":3001", mux)
}

func runGrpcService() {
	// todo hardcode port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("server listening at %v\n", lis.Addr())

	s := grpc.NewServer()

	// Register Greeter on the server.
	pb.RegisterYourServiceServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		fmt.Println("Run gRPC service...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}

func main() {
	flag.Parse()
	defer glog.Flush()

	runGrpcService()
	if err := runGateWayProxy(); err != nil {
		glog.Fatal(err)
	}
}
