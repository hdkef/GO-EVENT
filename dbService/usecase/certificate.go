package usecase

import (
	context "context"
	"dbservice/layer"
	"dbservice/repository"

	"gorm.io/gorm"
)

type CertificateService struct {
	*layer.UnimplementedCertificateLayerServer
	DB *gorm.DB
}

func (e *CertificateService) Create(ctx context.Context, payload *layer.Certificate) (*layer.Empty, error) {
	// create Certificate repo
	repo := repository.Certificate{
		DB:          e.DB,
		Certificate: payload,
	}
	//create
	err := repo.Create(&ctx)
	if err != nil {
		return &layer.Empty{}, err
	}
	return &layer.Empty{}, nil
}

func (e *CertificateService) Get(ctx context.Context, pagination *layer.Pagination) (*layer.CertificateList, error) {
	//create Certificate repo
	repo := repository.Certificate{
		DB: e.DB,
	}

	return repo.Get(&ctx, pagination)
}
