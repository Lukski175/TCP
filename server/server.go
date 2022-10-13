package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	t "github.com/Lukski175/TCP/time"

	"google.golang.org/grpc"
)

type Server struct {
	t.UnimplementedTcpServer
}

// Server seq number
var thisSeq int64 = 5

func (s *Server) GetSynack(ctx context.Context, sy *t.Syn) (*t.Synack, error) {
	fmt.Printf("Server Received first handshake. Syn: %v \nSending second handshake.\n", sy.Seq)
	return &t.Synack{SynSeq: thisSeq, AckSeq: sy.Seq + 1}, nil
}

func (s *Server) SendAck(ctx context.Context, sy *t.Ack) (*t.Ack, error) {
	fmt.Printf("Server Received third handshake. Syn: %v - Ack: %v - Data: %s\n", sy.Seq, sy.Ack, sy.ClientData)
	return nil, nil
}

func main() {
	// Create listener tcp on port 9080
	list, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}
	grpcServer := grpc.NewServer()
	t.RegisterTcpServer(grpcServer, &Server{})

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to serve %v", err)
	}

	time.Sleep(3 * time.Second)
}
