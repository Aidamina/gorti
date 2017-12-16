package services

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"

	"github.com/aidamina/gorti/api"
	"google.golang.org/grpc/codes"
)

// SayHello implements helloworld.GreeterServer
func (s *services) Connect(ctx context.Context, in *api.ConnectRequest) (*api.ConnectResponse, error) {

	if in.GetCallbackModel() == api.CallbackModel_EVOKED {
		return &api.ConnectResponse{
				OptionalError: &api.ConnectResponse_Error{
					Error: api.ConnectResponse_UNSUPPORTED_CALLBACK_MODEL,
				},
			},
			grpc.Errorf(codes.Unimplemented, "CallbackModel EVOKED not implemented")
	}

	return &api.ConnectResponse{
		OptionalError: &api.ConnectResponse_Error{
			Error: api.ConnectResponse_UNSUPPORTED_CALLBACK_MODEL,
		},
	}, nil
}
