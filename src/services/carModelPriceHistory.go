package services

import (
	"context"

	"github.com/FathiMohammadDev/car-selling/api/dto"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/data/db"
	"github.com/FathiMohammadDev/car-selling/data/models"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
)

type CarModelPriceHistoryService struct {
	base *BaseService[models.CarModelPriceHistory, dto.CreateCarModelPriceHistoryReq, dto.UpdateCarModelPriceHistoryReq, dto.CarModelPriceHistoryRes]
}

func NewCarModelPriceHistoryService(cfg *config.Config) *CarModelPriceHistoryService {
	return &CarModelPriceHistoryService{
		base: &BaseService[models.CarModelPriceHistory, dto.CreateCarModelPriceHistoryReq, dto.UpdateCarModelPriceHistoryReq, dto.CarModelPriceHistoryRes]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		}}
}

func (s *CarModelPriceHistoryService) Create(ctx context.Context, req *dto.CreateCarModelPriceHistoryReq) (*dto.CarModelPriceHistoryRes, error) {
	return s.base.Create(ctx, req)
}

func (s *CarModelPriceHistoryService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CarModelPriceHistoryService) Update(ctx context.Context, id int, req *dto.UpdateCarModelPriceHistoryReq) (*dto.CarModelPriceHistoryRes, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CarModelPriceHistoryService) GetById(ctx context.Context, id int) (*dto.CarModelPriceHistoryRes, error) {
	return s.base.GetById(ctx, id)

}
