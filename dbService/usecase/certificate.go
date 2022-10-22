package usecase

import (
	context "context"
	"dbservice/layer"

	"gorm.io/gorm"
)

type CertificateService struct {
	*layer.UnimplementedCertificateLayerServer
	DB *gorm.DB
}

func (e *CertificateService) Get(ctx context.Context, pagination *layer.Pagination) (*layer.CertificateList, error) {
	return nil, nil
}

func (e *CertificateService) Delete(ctx context.Context, id *layer.IDPayload) (*layer.Empty, error) {
	return nil, nil
}
