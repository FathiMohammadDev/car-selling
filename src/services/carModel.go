package services

import (
	"context"

	"github.com/FathiMohammadDev/car-selling/api/dto"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/data/db"
	"github.com/FathiMohammadDev/car-selling/data/models"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
)

type CarModelService struct {
	base *BaseService[models.CarModel, dto.CreateCarModelReq, dto.UpdateCarModelReq, dto.CarModelRes]
}

func NewCarModelService(cfg *config.Config) *CarModelService {
	return &CarModelService{
		base: &BaseService[models.CarModel, dto.CreateCarModelReq, dto.UpdateCarModelReq, dto.CarModelRes]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		}}
}

func (s *CarModelService) Create(ctx context.Context, req *dto.CreateCarModelReq) (*dto.CarModelRes, error) {
	return s.base.Create(ctx, req)
}

func (s *CarModelService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CarModelService) Update(ctx context.Context, id int, req *dto.UpdateCarModelReq) (*dto.CarModelRes, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CarModelService) GetById(ctx context.Context, id int) (*dto.CarModelRes, error) {
	return s.base.GetById(ctx, id)

}
