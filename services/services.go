package services

import (
	"github.com/aidamina/gorti/engine"
	"github.com/google/uuid"
	context "golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

type services struct {
	engine engine.Engine
}

//Services handles all RTI services
type Services interface {
	Engine() engine.Engine
}

func (s *services) Engine() engine.Engine {
	return s.engine
}

//CreateServices creates a Services instane
func CreateServices() Services {
	s := &services{}
	s.engine = engine.CreateEngine()
	return s
}
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
