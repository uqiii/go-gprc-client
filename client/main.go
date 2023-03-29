package main

import (
	"log"

	pb "grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", port)
	}
	defer conn.Close()

	client := pb.NewPlayServiceClient(conn)

	callGetAllPlays(client)
}