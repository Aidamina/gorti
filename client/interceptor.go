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
	if interceptor.connectionID != "" {
		ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(services.ConnectionIDKey, interceptor.connectionID))
	}
	var header metadata.MD
	opts = append(opts, grpc.Header(&header))
	err := invoker(ctx, method, req, reply, cc, opts...)

	conns := header[services.ConnectionIDKey]
	if conns != nil && len(conns) > 0 {
		interceptor.connectionID = conns[0]
	}

	return err
}

func CreateConnectionClientInterceptor() ConnectionClientInterceptor {
	return &connectionClientInterceptor{}
}
