package main

import (
	"log"
	"net"

	api "github.com/aidamina/gorti/api"
	svc "github.com/aidamina/gorti/services"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// Services exposes all RTI services
type Services interface{}

// server is used to implement helloworld.GreeterServer.
type services struct{}

//Setup should
func Setup() (bool, error) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return false, err
	}
	server := grpc.NewServer()
	services := svc.CreateServices()
	api.RegisterConnectServiceServer(server, services.(api.ConnectServiceServer))
	api.RegisterCreateFederateExecutionServiceServer(server, services.(api.CreateFederateExecutionServiceServer))
	// Register reflection service on gRPC server.
	reflection.Register(server)

	up := make(chan bool)
	go func() {
		up <- true
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return <-up, nil
}
