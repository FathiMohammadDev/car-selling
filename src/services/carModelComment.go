package services

import (
	"context"

	"github.com/FathiMohammadDev/car-selling/api/dto"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/data/db"
	"github.com/FathiMohammadDev/car-selling/data/models"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
)

type CarModelCommentService struct {
	base *BaseService[models.CarModelComment, dto.CreateCarModelCommentReq, dto.UpdateCarModelCommentReq, dto.CarModelCommentRes]
}

func NewCarModelCommentService(cfg *config.Config) *CarModelCommentService {
	return &CarModelCommentService{
		base: &BaseService[models.CarModelComment, dto.CreateCarModelCommentReq, dto.UpdateCarModelCommentReq, dto.CarModelCommentRes]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		}}
}

func (s *CarModelCommentService) Create(ctx context.Context, req *dto.CreateCarModelCommentReq) (*dto.CarModelCommentRes, error) {
	return s.base.Create(ctx, req)
}

func (s *CarModelCommentService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CarModelCommentService) Update(ctx context.Context, id int, req *dto.UpdateCarModelCommentReq) (*dto.CarModelCommentRes, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CarModelCommentService) GetById(ctx context.Context, id int) (*dto.CarModelCommentRes, error) {
	return s.base.GetById(ctx, id)

}
