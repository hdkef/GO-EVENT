package main

import (
	"fmt"
	"log"
	"os"
	"renderService/delivery"
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
	eventService := usecase.NewEventLayerClient(dbServiceConn)
	userService := usecase.NewUserLayerClient(dbServiceConn)

	r := gin.New()

	//middleware grpc client
	r.Use(delivery.GRPCMiddleware(eventService, userService))
	//initialize routes
	delivery.Routes(r)
	//run server

	r.Run(fmt.Sprintf(":%s", PORT))

}
