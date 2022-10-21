package main

import (
	"dbservice/usecase"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var PORT string

func init() {
	_ = godotenv.Load()
	PORT = os.Getenv("PORT")
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatalln("error listening", err.Error())
	}
	grpcServer := grpc.NewServer()

	usecase.RegisterEventLayerServer(grpcServer, &usecase.EventService{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("error grpc", err.Error())
	}
}
