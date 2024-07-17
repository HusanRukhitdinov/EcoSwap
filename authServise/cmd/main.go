package main

import (
	api "eco_system/api"
	"eco_system/config"
	pb "eco_system/genproto"
	"eco_system/service"
	postgres "eco_system/storage/postgres"
	"log"
	"net"
	"os"
	"os/signal"

	"syscall"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectionDb()
	if err != nil {
		log.Fatal("error connection for database ")
	}
	cnf := config.Load()

	go func() {
		listen, err := net.Listen("tcp", cnf.HTTPPort)
		if err != nil {
			log.Fatalf("error-> tcp connection error->%s", err.Error())
		}
		grpcServer := grpc.NewServer()
		pb.RegisterAuthServiceServer(grpcServer, service.NewUserService(postgres.NewUserRepository(db)))
		err = grpcServer.Serve(listen)
		if err != nil {
			log.Fatalf("error-> api_get_way connection error->%s", err.Error())
		}
	}()
	go func() {
		router := api.RouterApi(service.NewUserService(postgres.NewUserRepository(db)))
		if err := router.Run(":8087"); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGINT)
	<-quit
}
