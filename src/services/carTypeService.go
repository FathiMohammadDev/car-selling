package services

import (
	"context"

	"github.com/FathiMohammadDev/car-selling/api/dto"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/data/db"
	"github.com/FathiMohammadDev/car-selling/data/models"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
)

type CarTypeService struct {
	base *BaseService[models.CarType, dto.CreateCarTypeReq, dto.UpdateCarTypeReq, dto.CarTypeRes]
}

func NewCarTypeService(cfg *config.Config) *CarTypeService {
	return &CarTypeService{
		base: &BaseService[models.CarType, dto.CreateCarTypeReq, dto.UpdateCarTypeReq, dto.CarTypeRes]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		}}
}

func (s *CarTypeService) Create(ctx context.Context, req *dto.CreateCarTypeReq) (*dto.CarTypeRes, error) {
	return s.base.Create(ctx, req)
}

func (s *CarTypeService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CarTypeService) Update(ctx context.Context, id int, req *dto.UpdateCarTypeReq) (*dto.CarTypeRes, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CarTypeService) GetById(ctx context.Context, id int) (*dto.CarTypeRes, error) {
	return s.base.GetById(ctx, id)

}
