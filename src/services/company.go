package services

import (
	"context"

	"github.com/FathiMohammadDev/car-selling/api/dto"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/data/db"
	"github.com/FathiMohammadDev/car-selling/data/models"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
)

type CompanyService struct {
	base *BaseService[models.Company, dto.CreateCompanyReq, dto.UpdateCompanyReq, dto.CompanyRes]
}

func NewCompanyService(cfg *config.Config) *CompanyService {
	return &CompanyService{
		base: &BaseService[models.Company, dto.CreateCompanyReq, dto.UpdateCompanyReq, dto.CompanyRes]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		}}
}

func (s *CompanyService) Create(ctx context.Context, req *dto.CreateCompanyReq) (*dto.CompanyRes, error) {
	return s.base.Create(ctx, req)
}

func (s *CompanyService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CompanyService) Update(ctx context.Context, id int, req *dto.UpdateCompanyReq) (*dto.CompanyRes, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CompanyService) GetById(ctx context.Context, id int) (*dto.CompanyRes, error) {
	return s.base.GetById(ctx, id)

}
