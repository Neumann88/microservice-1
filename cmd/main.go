package main

import (
	"context"
	"net"
	"os"
	"os/signal"

	pb "github.com/Neumann88/microservice-1/internal/gen/proto/microservice-1/v1"
	grpcserver "github.com/Neumann88/microservice-1/internal/grpc/microservice-1/v1"
	"google.golang.org/grpc"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-c
		cancel()
	}()

	listen, err := net.Listen("tcp", ":5010")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	gSrv := grpcserver.NewMicroserviceOneGrpcServer()
	pb.RegisterMicroserviceOneExampleServer(s, gSrv)

	// if dev we can use reflection for postman grpc calls
	// if cfg.Dev {
	// 	reflection.Register(s)
	// }

	go func() {
		if err := s.Serve(listen); err != nil {
			panic(err)
		}
	}()

	<-ctx.Done()
}
