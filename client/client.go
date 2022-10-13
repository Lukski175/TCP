package main

import (
	"context"
	"fmt"
	"log"

	t "github.com/Lukski175/TCP/time"
	"google.golang.org/grpc"
)

// Client seq number
var thisSeq int64 = 1

func main() {

	// Creat a virtual RPC Client Connection on port  9080 WithInsecure (because  of http)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	// Defer means: When this function returns, call this method (meaing, one main is done, close connection)
	defer conn.Close()

	//  Create new Client from generated gRPC code from proto
	c := t.NewTcpClient(conn)

	fmt.Printf("Client sending first handshake. Seq: %v\n", thisSeq)
	response, err := c.GetSynack(context.Background(), &t.Syn{Seq: thisSeq})
	if err != nil {
		log.Fatalf("Error when calling GetSynack: %s", err)
	}

	fmt.Printf("Client Received second handshake. Syn: %v - Ack: %v \nSending third handshake.\n", response.SynSeq, response.AckSeq)
	something, err := c.SendAck(context.Background(), &t.Ack{Seq: response.AckSeq + 1, Ack: thisSeq + 1, ClientData: "Some client data"})
	if err != nil {
		log.Fatalf("Error when calling SendAck: %s", err)
	}
	_ = something
}
