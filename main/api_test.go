package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/aidamina/gorti/api"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func TestGrpc(t *testing.T) {
	Setup()

	data, _ := json.Marshal("")

	fmt.Println("json: " + string(data))

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewConnectServiceClient(conn)

	r, err := c.Connect(context.Background(), &api.ConnectRequest{CallbackModel: api.CallbackModel_EVOKED, LocalSettings: ""})
	if err != nil {
		log.Printf("could not connect: %v", err)
	}
	r, err = c.Connect(context.Background(), &api.ConnectRequest{CallbackModel: api.CallbackModel_IMMEDIATE, LocalSettings: ""})
	if err != nil {
		log.Printf("could not connect: %v", err)
	}
	log.Printf("Connecting: %s", r.GetOptionalError())

}
