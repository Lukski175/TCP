package main

import (
	"fmt"
	"log"

	time "github.com/Lukski175/TCP/time"
	"google.golang.org/grpc"
)

var thisSeq = 1

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
	c := time.NewTcpClient(conn)

	response, err := c.GetSynack(thisSeq)
	if err != nil {
		log.Fatalf("Error when calling GetTime: %s", err)
	}

	fmt.Printf("Current time right now: %s %d", response.SynSeq, response.AckSeq)
}
