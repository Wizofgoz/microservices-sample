package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	pb "github.com/wizofgoz/microservices-sample/vessel-service/proto/vessel"
	"log"
	"os"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	srv := micro.NewService(
		micro.Name("microservices.service.vessel"),
	)

	srv.Init()

	uri := os.Getenv("DB_HOST")

	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(uri)

	if err != nil {
		log.Panic(err)
	}

	defer client.Disconnect(context.TODO())

	vesselCollection := client.Database("microservices-sample").Collection("vessels")
	repository := &VesselRepository{
		vesselCollection,
	}


	// Register our implementation with
	pb.RegisterVesselServiceHandler(srv.Server(), &handler{repository})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}