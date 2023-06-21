package main

import (
	"context"
	"fmt"
	"net"
	"path"

	pb "example.com/m/gen/go/your/service/v1"

	"flag"
	"net/http"

	"github.com/golang/glog"
	"log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint",  "localhost:50051", "gRPC server endpoint")
)

const (
	tlsPath = "/etc/tls"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedYourServiceServer
}

func (server) Echo(ctx context.Context, msg *pb.StringMessage) (*pb.StringMessage, error) {
	log.Printf("Echo: %s", msg.Value)
	return msg, nil
}

func runGateWayProxy(cred credentials.TransportCredentials) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	// opts := []grpc.DialOption{grpc.WithTransportCredentials(cred)}
	// fmt.Println("Gateway proxy will connect to: " + *grpcServerEndpoint)
	// err := pb.RegisterYourServiceHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts) //
	// if err != nil {
	// 	return err
	// }


	err := mux.HandlePath("GET", "/", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Write([]byte("this is homepage"))
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Run gRPC gateway...")
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServeTLS(":3001", path.Join(tlsPath, "tls.crt"), path.Join(tlsPath, "tls.key"), mux)
}

func runGrpcService(cred credentials.TransportCredentials) {
	// todo hardcode port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("server listening at %v\n", lis.Addr())

	s := grpc.NewServer(grpc.Creds(cred))

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

func loadTlsServerCert() credentials.TransportCredentials {
	fmt.Println("Load server certificates")
	// Load TLS cert for server
	creds, tlsError := credentials.NewServerTLSFromFile(path.Join(tlsPath, "tls.crt"), path.Join(tlsPath, "tls.key", ""))
	if tlsError != nil {
		fmt.Printf("Error loading TLS key pair for server: %v", tlsError)
		fmt.Printf("Creating server without TLS")
		creds = insecure.NewCredentials()
	}
	return creds
}

func loadTlsClientCert() credentials.TransportCredentials {
	fmt.Println("Load client certificates")
	creds, err := credentials.NewClientTLSFromFile(path.Join(tlsPath, "tls.crt"), "")
	if err != nil {
		log.Fatalf("Error loading TLS certificate for REST: %v", err)
	}
	return creds
}

func main() {
	flag.Parse()
	defer glog.Flush()
	fmt.Println("=== 0 ===")

	serverCred := loadTlsServerCert()
	runGrpcService(serverCred)

	clientCred := loadTlsClientCert()
	if err := runGateWayProxy(clientCred); err != nil {
		glog.Fatal(err)
	}
}
