package v1

import (
	"context"
	"time"

	pb "github.com/Neumann88/microservice-1/internal/gen/proto/microservice-1/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MicroserviceOneGrpcServer struct {
	pb.UnimplementedMicroserviceOneExampleServer
}

func NewMicroserviceOneGrpcServer() *MicroserviceOneGrpcServer {
	return &MicroserviceOneGrpcServer{}
}

func (h *MicroserviceOneGrpcServer) GetExampleMessage(ctx context.Context, in *pb.ExampleMessageRequest) (*pb.ExampleMessageResponse, error) {
	// validation
	if in.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid id")
	}

	// call some service func
	time.Sleep(1 * time.Second)

	// response
	return &pb.ExampleMessageResponse{
		Value: "example",
	}, nil
}
