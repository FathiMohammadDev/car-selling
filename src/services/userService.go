package services

import (
	"context"

	"github.com/FathiMohammadDev/car-selling/api/dto"
	"github.com/FathiMohammadDev/car-selling/common"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/data/db"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
	"gorm.io/gorm"
)

type UserService struct {
	cfg        *config.Config
	database   *gorm.DB
	logger     logging.Logger
	otpService *OtpService
}

func NewUserService(cfg *config.Config) *UserService {
	logger := logging.NewLogger(cfg)
	database := db.GetDb()

	return &UserService{
		cfg,
		database,
		logger,
		NewOtpService(cfg),
	}
}

func (s *UserService) SendOtp(req dto.GetOtpRequest, ctx context.Context) error {
	otp, err := common.GenerateOtp()
	if err != nil {
		return err
	}
	err = s.otpService.SetOtp(req.MobileNumber, otp, ctx)
	if err != nil {
		return err
	}
	return nil
}
