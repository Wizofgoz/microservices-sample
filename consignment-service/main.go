package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	pb "github.com/wizofgoz/microservices-sample/consignment-service/proto/consignment"
	vesselProto "github.com/wizofgoz/microservices-sample/vessel-service/proto/vessel"
	"log"
	"os"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	// Set-up micro instance
	srv := micro.NewService(
		micro.Name("microservices.service.consignment"),
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

	consignmentCollection := client.Database("microservices-sample").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselServiceClient("microservices.service.vessel", srv.Client())
	h := &handler{repository, vesselClient}

	// Register handlers
	pb.RegisterShippingServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}