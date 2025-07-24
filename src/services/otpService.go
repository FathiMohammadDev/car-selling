package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/constants"
	"github.com/FathiMohammadDev/car-selling/data/cache"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
	"github.com/redis/go-redis/v9"
)

type OtpService struct {
	cfg    *config.Config
	redis  *redis.Client
	logger logging.Logger
}

type OtpDto struct {
	value string
	used  bool
}

func NewOtpService(cfg *config.Config) *OtpService {
	logger := logging.NewLogger(cfg)
	redis := cache.GetRedis()
	return &OtpService{
		cfg,
		redis,
		logger,
	}
}

func (s *OtpService) SetOtp(mobileNumber, otp string, ctx context.Context) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOTPDefaultKey, mobileNumber)
	value := &OtpDto{
		value: otp,
		used:  false,
	}

	res, err := cache.Get[OtpDto](s.redis, ctx, key)
	if err == nil && !res.used {
		return errors.New("otp exist")
	} else if err == nil && res.used {
		return errors.New("otp used")
	}
	err = cache.Set(s.redis, ctx, key, value, s.cfg.Otp.ExpireTime)
	if err != nil {
		return err
	}
	return nil
}

func (s *OtpService) ValidOtp(mobileNumber, otp string, ctx context.Context) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOTPDefaultKey, mobileNumber)
	res, err := cache.Get[OtpDto](s.redis, ctx, key)

	if err != nil {
		return err
	} else if err == nil && res.used {
		return errors.New("otp used")
	} else if err == nil && !res.used && res.value != otp {
		return errors.New("otp unvalid")
	} else if err == nil && !res.used && res.value == otp {
		res.used = true
		cache.Set(s.redis, ctx, key, res, s.cfg.Otp.ExpireTime)
	}
	return nil

}
