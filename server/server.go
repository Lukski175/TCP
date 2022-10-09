package server

import (
	"fmt"
	"log"
	"net"

	"github.com/Lukski175/GO-Exercise5/time"
)

func main() {
	fmt.Println("Input port to listen on...")
	var port string
	fmt.Scan(&port)

	// Create listener tcp on port 9080
	fmt.Printf("Listening on port %s", port)
	list, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}
	grpcServer := grpc.NewServer()
	time.RegisterTimeServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
