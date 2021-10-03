package main

import (
	"context"
	"fmt"

	"github.com/uh-zz/grpc-sample/pinger/pinger"
	"google.golang.org/grpc"
)

const (
	port = ":5300"
)

func main() {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("grpc.Dial: %v\n", err)
	}

	defer conn.Close()

	client := pinger.NewPingerClient(conn)
	req := &pinger.Empty{}

	pong, err := client.Ping(context.Background(), req)
	if err != nil {
		fmt.Printf("Ping: %v\n", err)
	}

	fmt.Printf("Pong: %v\n", pong)
}
