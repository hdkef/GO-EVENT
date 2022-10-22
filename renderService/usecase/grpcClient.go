package usecase

import "renderService/layer"

type GRPCClient struct {
	Event        layer.EventLayerClient
	User         layer.UserLayerClient
	Participant  layer.ParticipantLayerClient
	Subscription layer.SubscriptionLayerClient
	Certificate  layer.CertificateLayerClient
	Like         layer.LikeLayerClient
}
