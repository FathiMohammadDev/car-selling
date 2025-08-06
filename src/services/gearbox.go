package services

import (
	"context"

	"github.com/FathiMohammadDev/car-selling/api/dto"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/data/db"
	"github.com/FathiMohammadDev/car-selling/data/models"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
)

type GearboxService struct {
	base *BaseService[models.Gearbox, dto.CreateGearboxReq, dto.UpdateGearboxReq, dto.GearboxRes]
}

func NewGearboxService(cfg *config.Config) *GearboxService {
	return &GearboxService{
		base: &BaseService[models.Gearbox, dto.CreateGearboxReq, dto.UpdateGearboxReq, dto.GearboxRes]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		}}
}

func (s *GearboxService) Create(ctx context.Context, req *dto.CreateGearboxReq) (*dto.GearboxRes, error) {
	return s.base.Create(ctx, req)
}

func (s *GearboxService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *GearboxService) Update(ctx context.Context, id int, req *dto.UpdateGearboxReq) (*dto.GearboxRes, error) {
	return s.base.Update(ctx, id, req)
}

func (s *GearboxService) GetById(ctx context.Context, id int) (*dto.GearboxRes, error) {
	return s.base.GetById(ctx, id)

}
