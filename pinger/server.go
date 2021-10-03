package main

import (
	"context"
	"log"
	"net"

	"github.com/uh-zz/grpc-sample/pinger/pinger"
	"google.golang.org/grpc"
)

const (
	port = ":5300"
)

type server struct {
	pinger.UnimplementedPingerServer
}

func (s *server) Ping(ctx context.Context, req *pinger.Empty) (*pinger.Pong, error) {
	pong := &pinger.Pong{
		Text: "pong",
	}
	return pong, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return
	}

	grpcServer := grpc.NewServer()

	pinger.RegisterPingerServer(grpcServer, &server{})
	log.Printf("Pinger listening at %v", listener.Addr())

	grpcServer.Serve(listener)
}
