package main

import (
	"dbservice/layer"
	"dbservice/usecase"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalln("error listening", err.Error())
	}
	grpcServer := grpc.NewServer()

	layer.RegisterEventLayerServer(grpcServer, &usecase.EventService{})
	layer.RegisterUserLayerServer(grpcServer, &usecase.UserService{})
	layer.RegisterParticipantLayerServer(grpcServer, &usecase.ParticipantService{})
	layer.RegisterSubscriptionLayerServer(grpcServer, &usecase.SubscriptionService{})
	layer.RegisterLikeLayerServer(grpcServer, &usecase.LikeService{})
	layer.RegisterCertificateLayerServer(grpcServer, &usecase.CertificateService{})

	fmt.Println("DB SERVICE RUNNING ON PORT " + os.Getenv("APP_PORT"))

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("error grpc", err.Error())
	}
}
