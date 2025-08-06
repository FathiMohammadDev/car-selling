package services

import (
	"context"

	"github.com/FathiMohammadDev/car-selling/api/dto"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/data/db"
	"github.com/FathiMohammadDev/car-selling/data/models"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
)

type CarModelYearService struct {
	base *BaseService[models.CarModelYear, dto.CreateCarModelYearReq, dto.UpdateCarModelYearReq, dto.CarModelYearRes]
}

func NewCarModelYearService(cfg *config.Config) *CarModelYearService {
	return &CarModelYearService{
		base: &BaseService[models.CarModelYear, dto.CreateCarModelYearReq, dto.UpdateCarModelYearReq, dto.CarModelYearRes]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		}}
}

func (s *CarModelYearService) Create(ctx context.Context, req *dto.CreateCarModelYearReq) (*dto.CarModelYearRes, error) {
	return s.base.Create(ctx, req)
}

func (s *CarModelYearService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CarModelYearService) Update(ctx context.Context, id int, req *dto.UpdateCarModelYearReq) (*dto.CarModelYearRes, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CarModelYearService) GetById(ctx context.Context, id int) (*dto.CarModelYearRes, error) {
	return s.base.GetById(ctx, id)

}
