package services

import (
	"log"

	"github.com/aidamina/gorti/engine"
	"github.com/google/uuid"

	context "golang.org/x/net/context"

	"github.com/aidamina/gorti/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	// ConnectionIDKey contains the Connection Id Metadata Key
	ConnectionIDKey string = "connection_id"
)

func GetConnectionID(md metadata.MD) engine.ConnectionID {
	conns := md[ConnectionIDKey]
	if conns != nil && len(conns) > 0 {
		id, err := uuid.Parse(conns[0])
		if err == nil {
			return engine.ConnectionID(id)
		}
	}
	return nil
}

func (s *services) GetConnection(ctx context.Context) engine.Connection {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		connectionID := GetConnectionID(md)
		if connectionID != nil {
			return s.Engine().ConnectionManager().GetConnection(connectionID)
		}
	}
	return nil
}

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
	header := metadata.Pairs(ConnectionIDKey, "test_set_by_server")
	grpc.SendHeader(ctx, header)
	return &api.ConnectResponse{
		Result: api.ConnectResponse_SUCCESS,
	}, nil
}
