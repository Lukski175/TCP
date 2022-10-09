package server

import (
	"fmt"
	"log"
	"net"

	"github.com/Lukski175/TCP/time"

	"google.golang.org/grpc"
)

type Server struct {
	time.UnimplementedGettcpServer
}

var thisSeq = 1

func (s *Server) GetSynack(sy *time.syn) (*time.synack, error) {
	fmt.Printf("Server Received first handshake. Sending second handshake")
	return &time.synack{synSeq: sy.seq + 1, ackSeq: thisSeq + 1}, nil
}

func main() {
	// Create listener tcp on port 9080
	list, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}
	grpcServer := grpc.NewServer()
	time.RegistertcpServer(grpcServer, &Server{})

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
