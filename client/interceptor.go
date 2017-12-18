package client

import (
	"context"

	"github.com/aidamina/gorti/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type connectionClientInterceptor struct {
	connectionID string
}

type ConnectionClientInterceptor interface {
	Handle(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error
}

func (interceptor *connectionClientInterceptor) Handle(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md, ok := metadata.FromOutgoingContext(ctx)
	//grpc.Header()
	if ok {
		conns := md[services.ConnectionIDKey]
		if conns != nil && len(conns) > 0 {
			interceptor.connectionID = conns[0]
		}
	}
	return nil
}

func CreateConnectionClientInterceptor() ConnectionClientInterceptor {
	return &connectionClientInterceptor{}
}
