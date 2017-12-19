package services

import (
	"log"

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
	log.Printf("Connecting...")
	connection := s.GetConnection(ctx)
	if connection != nil {
		return &api.ConnectResponse{
			Result: api.ConnectResponse_ALREADY_CONNECTED,
		}, nil
	}
	if in.GetCallbackModel() == api.CallbackModel_EVOKED {
		header := metadata.Pairs("error", "whoopsy!")
		grpc.SendHeader(ctx, header)
		return &api.ConnectResponse{
			Result: api.ConnectResponse_UNSUPPORTED_CALLBACK_MODEL,
		}, nil
	}
	connection = s.Engine().ConnectionManager().CreateConnection()
	header := metadata.Pairs(ConnectionIDKey, connection.ID().String())
	grpc.SendHeader(ctx, header)
	return &api.ConnectResponse{
		Result: api.ConnectResponse_SUCCESS,
	}, nil
}
