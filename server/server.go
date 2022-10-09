package tcp

import (
	"log"
	"net"

	"github.com/Lukski175/TCP/time"

	"google.golang.org/grpc"
)

type Server struct {
	time.UnimplementedTimeServiceServer
}

func main() {
	// Create listener tcp on port 9080
	list, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}
	grpcServer := grpc.NewServer()
	time.RegisterTimeServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
