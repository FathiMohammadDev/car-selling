package services

import (
	"context"

	"github.com/FathiMohammadDev/car-selling/api/dto"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/data/db"
	"github.com/FathiMohammadDev/car-selling/data/models"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
)

type ColorService struct {
	base *BaseService[models.Color, dto.CreateColorReq, dto.UpdateColorReq, dto.ColorRes]
}

func NewColorService(cfg *config.Config) *ColorService {
	return &ColorService{
		base: &BaseService[models.Color, dto.CreateColorReq, dto.UpdateColorReq, dto.ColorRes]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		}}
}

func (s *ColorService) Create(ctx context.Context, req *dto.CreateColorReq) (*dto.ColorRes, error) {
	return s.base.Create(ctx, req)
}

func (s *ColorService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *ColorService) Update(ctx context.Context, id int, req *dto.UpdateColorReq) (*dto.ColorRes, error) {
	return s.base.Update(ctx, id, req)
}

func (s *ColorService) GetById(ctx context.Context, id int) (*dto.ColorRes, error) {
	return s.base.GetById(ctx, id)

}
