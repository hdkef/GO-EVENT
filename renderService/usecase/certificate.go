package usecase

import (
	"context"
	"fmt"
	"renderService/layer"
	"renderService/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type CertificateService struct{}

func (u *CertificateService) Create(c *gin.Context) error {
	//get event-id
	eventID, err := utils.GetParamEventID(c)
	if err != nil {
		return err
	}

	//decode payload
	var certificate layer.Certificate
	certificate.Event_ID = eventID

	_, err = utils.DecodeCertificate(c, &certificate)
	if err != nil {
		return err
	}

	//get grpc client
	grpc := GetGRPCClient(c)

	//generate certificate [TODO]
	fileURL := "/public/certificate/hadekha10122022.pdf"
	certificate.FileUrl = &fileURL

	//validate certificate
	err = utils.ValidateCertificate(&certificate, utils.VALIDATE_TYPE_CREATE)
	if err != nil {
		//delete certificate file
		return err
	}

	//create
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err = grpc.Certificate.Create(ctx, &certificate)
	if err != nil {
		//delete certificate file
		return err
	}
	return nil
}

func (u *CertificateService) GetByUserID(c *gin.Context, userID *uint32) (*layer.CertificateList, error) {
	//get pagination
	lastID, limit, _, err := utils.GetPagination(c)
	if err != nil {
		return nil, err
	}

	//get grpc client
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	grpc := GetGRPCClient(c)

	//query
	query := fmt.Sprintf("user_id = %d", *userID)

	//get by id
	data, err := grpc.Certificate.Get(ctx, &layer.Pagination{
		LastID: lastID,
		Limit:  limit,
		Query:  &query,
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}
