package main

import (
	"context"
	"log"
	"time"

	pb "grpc-demo/proto"
)

func callGetPlayById(client pb.PlayServiceClient){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetPlayById(ctx, &pb.GetPlayByIdRequest{ Id: "1" })
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("%s", res.Name)
}

func callGetAllPlays(client pb.PlayServiceClient){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetAllPlays(ctx, &pb.GetAllPlaysRequest{})
	if err != nil {
		log.Fatalf("could not get all user: %v", err)
	}
	log.Printf("%v", res.Plays)
}

func callCreatePLay(client pb.PlayServiceClient){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.CreatePlay(ctx, &pb.CreatePlayRequest{ Name: "Dana" })
	if err != nil {
		log.Fatalf("could not get all user: %v", err)
	}
	log.Printf("%v", res)
}

