package services

import (
	"context"

	"github.com/FathiMohammadDev/car-selling/api/dto"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/data/db"
	"github.com/FathiMohammadDev/car-selling/data/models"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
)

type CarModelPropertyService struct {
	base *BaseService[models.CarModelProperty, dto.CreateCarModelPropertyReq, dto.UpdateCarModelPropertyReq, dto.CarModelPropertyRes]
}

func NewCarModelPropertyService(cfg *config.Config) *CarModelPropertyService {
	return &CarModelPropertyService{
		base: &BaseService[models.CarModelProperty, dto.CreateCarModelPropertyReq, dto.UpdateCarModelPropertyReq, dto.CarModelPropertyRes]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		}}
}

func (s *CarModelPropertyService) Create(ctx context.Context, req *dto.CreateCarModelPropertyReq) (*dto.CarModelPropertyRes, error) {
	return s.base.Create(ctx, req)
}

func (s *CarModelPropertyService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CarModelPropertyService) Update(ctx context.Context, id int, req *dto.UpdateCarModelPropertyReq) (*dto.CarModelPropertyRes, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CarModelPropertyService) GetById(ctx context.Context, id int) (*dto.CarModelPropertyRes, error) {
	return s.base.GetById(ctx, id)

}
