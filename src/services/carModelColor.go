package services

import (
	"context"

	"github.com/FathiMohammadDev/car-selling/api/dto"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/data/db"
	"github.com/FathiMohammadDev/car-selling/data/models"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
)

type CarModelColorService struct {
	base *BaseService[models.CarModelColor, dto.CreateCarModelColorReq, dto.UpdateCarModelColorReq, dto.CarModelColorRes]
}

func NewCarModelColorService(cfg *config.Config) *CarModelColorService {
	return &CarModelColorService{
		base: &BaseService[models.CarModelColor, dto.CreateCarModelColorReq, dto.UpdateCarModelColorReq, dto.CarModelColorRes]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		}}
}

func (s *CarModelColorService) Create(ctx context.Context, req *dto.CreateCarModelColorReq) (*dto.CarModelColorRes, error) {
	return s.base.Create(ctx, req)
}

func (s *CarModelColorService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CarModelColorService) Update(ctx context.Context, id int, req *dto.UpdateCarModelColorReq) (*dto.CarModelColorRes, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CarModelColorService) GetById(ctx context.Context, id int) (*dto.CarModelColorRes, error) {
	return s.base.GetById(ctx, id)

}
