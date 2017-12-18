package services

import (
	context "golang.org/x/net/context"

	"github.com/aidamina/gorti/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	// ConnectionIDKey contains the Connection Id Metadata Key
	ConnectionIDKey string = "connection_id"
)

// Connect implements api.ConnectService
func (s *services) Connect(ctx context.Context, in *api.ConnectRequest) (*api.ConnectResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		conns := md[ConnectionIDKey]
		if conns != nil && len(conns) > 0 {
			return &api.ConnectResponse{
				Result: api.ConnectResponse_ALREADY_CONNECTED,
			}, nil
		}
	}
	if in.GetCallbackModel() == api.CallbackModel_EVOKED {
		header := metadata.Pairs("error", "whoopsy!")
		grpc.SendHeader(ctx, header)
		return &api.ConnectResponse{
			Result: api.ConnectResponse_UNSUPPORTED_CALLBACK_MODEL,
		}, nil
	}
	header := metadata.Pairs(ConnectionIDKey, "test_set_by_server")
	grpc.SendHeader(ctx, header)
	return &api.ConnectResponse{
		Result: api.ConnectResponse_SUCCESS,
	}, nil
}
