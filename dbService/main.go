package main

import (
	"dbservice/driver"
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
	db, err := driver.GetDBConn()
	if err != nil {
		log.Fatalln("error db connection", err.Error())
	}

	grpcServer := grpc.NewServer()

	layer.RegisterEventLayerServer(grpcServer, &usecase.EventService{DB: db})
	layer.RegisterUserLayerServer(grpcServer, &usecase.UserService{DB: db})
	layer.RegisterParticipantLayerServer(grpcServer, &usecase.ParticipantService{DB: db})
	layer.RegisterSubscriptionLayerServer(grpcServer, &usecase.SubscriptionService{DB: db})
	layer.RegisterLikeLayerServer(grpcServer, &usecase.LikeService{DB: db})
	layer.RegisterCertificateLayerServer(grpcServer, &usecase.CertificateService{DB: db})

	fmt.Println("DB SERVICE RUNNING ON PORT " + os.Getenv("APP_PORT"))

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("error grpc", err.Error())
	}
}
