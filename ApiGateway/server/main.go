package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"item_api/api"

	"log"
)

func main() {
	conn1, err := grpc.NewClient(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error")
	}
	conn2, err := grpc.NewClient(":8020", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error")
	}
	
	router := api.RouterApi(conn1, conn2)
	err = router.Run(":8080")
	if err != nil {
		log.Fatal("error is api get way connection port")
	}

}
