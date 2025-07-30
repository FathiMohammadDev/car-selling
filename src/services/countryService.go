package services

import (
	"context"

	"github.com/FathiMohammadDev/car-selling/api/dto"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/data/db"
	"github.com/FathiMohammadDev/car-selling/data/models"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
)

type CountryService struct {
	base *BaseService[models.Country, dto.CreateUpdateCountryReq, dto.CreateUpdateCountryReq, dto.CountryResponse]
}

func NewCountryService(cfg *config.Config) *CountryService {
	return &CountryService{
		base: &BaseService[models.Country, dto.CreateUpdateCountryReq, dto.CreateUpdateCountryReq, dto.CountryResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

func (s *CountryService) Create(ctx context.Context, req *dto.CreateUpdateCountryReq) (*dto.CountryResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *CountryService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CountryService) Update(ctx context.Context, id int, req *dto.CreateUpdateCountryReq) (*dto.CountryResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CountryService) GetById(ctx context.Context, id int) (*dto.CountryResponse, error) {
	return s.base.GetById(ctx, id)

}
