package main

import (
	"context"
	"log"
	"time"

	hi "PYTHONGRPC/go_protos/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := hi.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &hi.HelloRequest{Name: "Go Client"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response from Python server: %s", r.GetMessage())
}
