package main

import (
	"item_ser/config"
	pb "item_ser/genproto"
	"item_ser/service"
	"item_ser/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cnf := config.Load()

	db, err := postgres.ConnectionDb()
	if err != nil {
		log.Fatal("error-> connection db", err.Error())
	}

	listen, err := net.Listen("tcp", cnf.HTTPPort)
	if err != nil {
		log.Fatal("error is listen up tcp connection in port->", err.Error())
	}
	grpcServer := grpc.NewServer()
	
	pb.RegisterEcoServiceServer(grpcServer, service.NewItemService(
		postgres.NewItemRepository(db),
		postgres.NewSwapsRepository(db),
		postgres.NewRecyclingCentersRepository(db),
		postgres.NewRecyclingSubmissionsRepository(db),
		postgres.NewUserRatingRepository(db),
		postgres.NewItemCategoryRepository(db),
		postgres.NewEcoChallengesRepository(db),
		postgres.NewChallengePrtisipationsRepository(db),
		postgres.NewEcoTipsRepository(db),
	))
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("error-> connection ->%s", err.Error())
	}
}
