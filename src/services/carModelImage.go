package services

import (
	"context"

	"github.com/FathiMohammadDev/car-selling/api/dto"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/data/db"
	"github.com/FathiMohammadDev/car-selling/data/models"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
)

type CarModelImageService struct {
	base *BaseService[models.CarModelImage, dto.CreateCarModelImageReq, dto.UpdateCarModelImageReq, dto.CarModelImageRes]
}

func NewCarModelImageService(cfg *config.Config) *CarModelImageService {
	return &CarModelImageService{
		base: &BaseService[models.CarModelImage, dto.CreateCarModelImageReq, dto.UpdateCarModelImageReq, dto.CarModelImageRes]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		}}
}

func (s *CarModelImageService) Create(ctx context.Context, req *dto.CreateCarModelImageReq) (*dto.CarModelImageRes, error) {
	return s.base.Create(ctx, req)
}

func (s *CarModelImageService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CarModelImageService) Update(ctx context.Context, id int, req *dto.UpdateCarModelImageReq) (*dto.CarModelImageRes, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CarModelImageService) GetById(ctx context.Context, id int) (*dto.CarModelImageRes, error) {
	return s.base.GetById(ctx, id)

}
