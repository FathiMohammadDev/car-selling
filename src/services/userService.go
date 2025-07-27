package services

import (
	"context"
	"errors"

	"github.com/FathiMohammadDev/car-selling/api/dto"
	"github.com/FathiMohammadDev/car-selling/common"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/constants"
	"github.com/FathiMohammadDev/car-selling/data/db"
	"github.com/FathiMohammadDev/car-selling/data/models"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
	"gorm.io/gorm"
)

type UserService struct {
	cfg          *config.Config
	database     *gorm.DB
	logger       logging.Logger
	otpService   *OtpService
	tokenService *TokenService
}

func NewUserService(cfg *config.Config) *UserService {
	logger := logging.NewLogger(cfg)
	database := db.GetDb()

	return &UserService{
		cfg,
		database,
		logger,
		NewOtpService(cfg),
		NewTokenService(cfg),
	}
}

func (s *UserService) RegisterByUserName(req dto.RegisterUserByUsernameRequest) error {
	user := models.User{Username: req.Username, FirstName: req.FirstName, LastName: req.LastName, Email: req.Email, Password: req.Password}

	exists, err := s.existByEmail(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email exists")
	}

	exists, err = s.existByUserName(user.Username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("userName exists")
	}

	hashedPassword, err := common.HashPassword(user.Password)
	if err != nil {
		s.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return err
	}
	user.Password = hashedPassword

	roleId, err := s.getDefaultRole()
	if err != nil {
		s.logger.Error(logging.Postgres, logging.DefaultRoleNotFound, err.Error(), nil)
		return err
	}

	tx := s.database.Begin()
	err = tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}
	err = tx.Create(&models.UserRole{RoleId: roleId, UserId: user.Id}).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}
	tx.Commit()

	return nil
}

func (s *UserService) LoginByUserName(req dto.LoginByUsernameRequest) (*dto.TokenDetail, error) {
	var user models.User

	err := s.database.
		Model(&models.User{}).
		Where("username = ?", req.Username).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).
		Find(&user).Error

	if err != nil {
		return nil, err
	}

	isCorect := common.ComparePassword(req.Password, user.Password)
	if !isCorect {
		return nil, errors.New("password is not corect")
	}

	tokenClaims := tokenDto{UserId: user.Id, FirstName: user.FirstName, LastName: user.LastName,
		Email: user.Email, MobileNumber: user.MobileNumber}
	if len(*user.UserRoles) > 0 {
		for _, ur := range *user.UserRoles {
			tokenClaims.Roles = append(tokenClaims.Roles, ur.Role.Name)
		}
	}

	token, err := s.tokenService.GenerateToken(&tokenClaims)
	if err != nil {
		return nil, err
	}
	return token, nil

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

func (s *UserService) existByEmail(email string) (bool, error) {
	var exists bool
	err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (s *UserService) existByMobileNum(mobileNum string) (bool, error) {
	var exists bool
	err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("mobile_number = ?", mobileNum).
		Find(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (s *UserService) existByUserName(userName string) (bool, error) {
	var exists bool
	err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("username = ?", userName).
		Find(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (s *UserService) getDefaultRole() (roleId int, err error) {
	err = s.database.Model(&models.Role{}).
		Select("id").
		Where("name = ?", constants.DefaultRoleName).
		Find(&roleId).Error
	if err != nil {
		return 0, err
	}
	return roleId, nil
}
