package grpcserver

import (
	"log"
	"google.golang.org/grpc"
)

func Connect(address string) (*grpc.ClientConn) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Printf("Error: Could not connect to grpc server: %v", err)
		return nil
	}

	return conn
}
