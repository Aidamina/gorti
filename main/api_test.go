package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/aidamina/gorti/api"
	"github.com/aidamina/gorti/client"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	address = "localhost:50051"
)

func TestGrpc(t *testing.T) {
	Setup()

	//md := metadata.Pairs(svc.ConnectionIDKey, "test")
	data, _ := json.Marshal("")
	cci := client.CreateConnectionClientInterceptor()

	fmt.Println("json: " + string(data))
	interceptor := grpc.WithUnaryInterceptor(grpc.UnaryClientInterceptor(grpc_middleware.ChainUnaryClient(cci.Handle)))
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), interceptor)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewConnectServiceClient(conn)

	ctx := context.Background() //metadata.NewOutgoingContext(context.Background(), md)

	r, err := c.Connect(ctx, &api.ConnectRequest{CallbackModel: api.CallbackModel_EVOKED, LocalSettings: ""})
	if err != nil {
		log.Printf("could not connect: %v", err)
	}

	var header, trailer metadata.MD
	r, err = c.Connect(ctx, &api.ConnectRequest{CallbackModel: api.CallbackModel_IMMEDIATE, LocalSettings: ""}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		log.Printf("could not connect: %v", err)
	}
	log.Printf("Connecting: %s", r.GetResult())

}
