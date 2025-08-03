package services

import (
	"context"

	"github.com/FathiMohammadDev/car-selling/api/dto"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/data/db"
	"github.com/FathiMohammadDev/car-selling/data/models"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
)

type CityService struct {
	base *BaseService[models.City, dto.CreateCityReq, dto.UpdateCityReq, dto.CityRes]
}

func NewCityService(cfg *config.Config) *CityService {
	return &CityService{
		base: &BaseService[models.City, dto.CreateCityReq, dto.UpdateCityReq, dto.CityRes]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		}}
}

func (s *CityService) Create(ctx context.Context, req *dto.CreateCityReq) (*dto.CityRes, error) {
	return s.base.Create(ctx, req)
}

func (s *CityService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CityService) Update(ctx context.Context, id int, req *dto.UpdateCityReq) (*dto.CityRes, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CityService) GetById(ctx context.Context, id int) (*dto.CityRes, error) {
	return s.base.GetById(ctx, id)

}
// func (u *CityService) GetByFilter(ctx context.Context, req dto.PaginationInputWithFilter) (*dto.PagedList[dto.City], error) {
// 	return u.base.GetByFilter(ctx, req)
// }
