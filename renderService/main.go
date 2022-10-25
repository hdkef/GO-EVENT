package main

import (
	"fmt"
	"log"
	"os"
	"renderService/delivery"
	"renderService/layer"
	"renderService/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var opts []grpc.DialOption
var DBSERVICEHOST string
var PORT string

func init() {
	_ = godotenv.Load()
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	DBSERVICEHOST = os.Getenv("GRPC_DB_HOST")
	PORT = os.Getenv("APP_PORT")
}

func main() {
	//dial dbService
	dbServiceConn, err := grpc.Dial(DBSERVICEHOST, opts...)
	if err != nil {
		log.Fatalln("grpc dial fail", err.Error())
	}
	defer dbServiceConn.Close()

	//register service client
	GRPCClient := usecase.GRPCClient{}

	GRPCClient.Event = layer.NewEventLayerClient(dbServiceConn)
	GRPCClient.User = layer.NewUserLayerClient(dbServiceConn)
	GRPCClient.Participant = layer.NewParticipantLayerClient(dbServiceConn)
	GRPCClient.Like = layer.NewLikeLayerClient(dbServiceConn)
	GRPCClient.Subscription = layer.NewSubscriptionLayerClient(dbServiceConn)
	GRPCClient.Certificate = layer.NewCertificateLayerClient(dbServiceConn)
	GRPCClient.Register = layer.NewRegisterLayerClient(dbServiceConn)

	r := gin.New()

	//middleware
	r.Use(delivery.GRPCMiddleware(GRPCClient))
	//initialize routes
	delivery.Routes(r)
	//run server

	r.Run(fmt.Sprintf(":%s", PORT))

}
