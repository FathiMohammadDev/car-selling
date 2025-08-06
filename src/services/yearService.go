package services

import (
	"context"

	"github.com/FathiMohammadDev/car-selling/api/dto"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/data/db"
	"github.com/FathiMohammadDev/car-selling/data/models"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
)

type YearService struct {
	base *BaseService[models.PersianYear, dto.CreateYearReq, dto.UpdateYearReq, dto.YearRes]
}

func NewYearService(cfg *config.Config) *YearService {
	return &YearService{
		base: &BaseService[models.PersianYear, dto.CreateYearReq, dto.UpdateYearReq, dto.YearRes]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		}}
}

func (s *YearService) Create(ctx context.Context, req *dto.CreateYearReq) (*dto.YearRes, error) {
	return s.base.Create(ctx, req)
}

func (s *YearService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *YearService) Update(ctx context.Context, id int, req *dto.UpdateYearReq) (*dto.YearRes, error) {
	return s.base.Update(ctx, id, req)
}

func (s *YearService) GetById(ctx context.Context, id int) (*dto.YearRes, error) {
	return s.base.GetById(ctx, id)

}
