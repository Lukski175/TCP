package tcp

import (
	"context"
	"fmt"
	"log"

	"github.com/Lukski175/TCP/time"
	"google.golang.org/grpc"
)

type Server struct {
	time.UnimplementedTimeServiceServer
}

func main() {

	//Sets up GRPC connection, which we use to simulate TCP on
	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	c := time.NewTimeServiceClient(conn)

	message := time.TimeRequest{}

	response, err := c.GetTime(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling GetTime: %s", err)
	}

	fmt.Printf("Current time right now: %s\n", response.Reply)
}
