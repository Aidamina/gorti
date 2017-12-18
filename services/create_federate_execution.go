package services

import (
	context "golang.org/x/net/context"

	"github.com/aidamina/gorti/api"
	"google.golang.org/grpc/metadata"
)

// SayHello implements helloworld.GreeterServer
func (s *services) CreateFederateExecution(ctx context.Context, in *api.CreateFederateExecutionRequest) (*api.CreateFederateExecutionResponse, error) {
	metadata.FromIncomingContext(ctx)

	return &api.CreateFederateExecutionResponse{
		Result: api.CreateFederateExecutionResponse_SUCCESS,
	}, nil
}
