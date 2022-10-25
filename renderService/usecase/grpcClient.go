package usecase

import (
	"log"
	"renderService/layer"

	"github.com/gin-gonic/gin"
)

type GRPCClient struct {
	Event        layer.EventLayerClient
	User         layer.UserLayerClient
	Participant  layer.ParticipantLayerClient
	Subscription layer.SubscriptionLayerClient
	Certificate  layer.CertificateLayerClient
	Like         layer.LikeLayerClient
	Register     layer.RegisterLayerClient
}

func GetGRPCClient(c *gin.Context) GRPCClient {
	value, exist := c.Get("GRPCClient")
	if !exist {
		log.Fatalln("no eventService in gin context")
	}
	return value.(GRPCClient)
}
