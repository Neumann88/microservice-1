package main

import (
	"net"

	pb "github.com/Neumann88/microservice-1/internal/gen/proto/microservice-1/v1"
	grpcserver "github.com/Neumann88/microservice-1/internal/grpc/microservice-1/v1"
	"google.golang.org/grpc"
)

func main() {
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

	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
