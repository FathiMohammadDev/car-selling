package services

import (
	"errors"
	"time"

	"github.com/FathiMohammadDev/car-selling/api/dto"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/constants"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	cfg    *config.Config
	Logger logging.Logger
}

type tokenDto struct {
	UserId       int
	FirstName    string
	LastName     string
	UserName     string
	MobileNumber string
	Email        string
	Roles        []string
}

func NewTokenService(cfg *config.Config) *TokenService {
	logger := logging.NewLogger(cfg)
	return &TokenService{
		cfg,
		logger,
	}
}

func (s *TokenService) GenerateToken(tokenClaims *tokenDto) (*dto.TokenDetail, error) {
	tokenDetails := &dto.TokenDetail{}
	tokenDetails.AccessTokenExpireTime = time.Now().Add(s.cfg.JWT.AccessTokenExpireDuration * time.Minute).Unix()
	tokenDetails.RefreshTokenExpireTime = time.Now().Add(s.cfg.JWT.RefreshTokenExpireDuration * time.Minute).Unix()

	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims[constants.UserIdKey] = tokenClaims.UserId
	accessTokenClaims[constants.FirstNameKey] = tokenClaims.FirstName
	accessTokenClaims[constants.LastNameKey] = tokenClaims.LastName
	accessTokenClaims[constants.UsernameKey] = tokenClaims.UserName
	accessTokenClaims[constants.MobileNumberKey] = tokenClaims.MobileNumber
	accessTokenClaims[constants.EmailKey] = tokenClaims.Email
	accessTokenClaims[constants.RolesKey] = tokenClaims.Roles
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	var err error

	tokenDetails.AccessToken, err = accessToken.SignedString([]byte(s.cfg.JWT.Secret))
	if err != nil {
		return nil, err
	}

	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims[constants.UserIdKey] = tokenClaims.UserId
	refreshTokenClaims[constants.FirstNameKey] = tokenClaims.FirstName
	refreshTokenClaims[constants.LastNameKey] = tokenClaims.LastName
	refreshTokenClaims[constants.UsernameKey] = tokenClaims.UserName
	refreshTokenClaims[constants.MobileNumberKey] = tokenClaims.MobileNumber
	refreshTokenClaims[constants.EmailKey] = tokenClaims.Email
	refreshTokenClaims[constants.RolesKey] = tokenClaims.Roles

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	tokenDetails.RefreshToken, err = refreshToken.SignedString([]byte(s.cfg.JWT.Secret))
	if err != nil {
		return nil, err
	}

	return tokenDetails, nil
}

func (u *TokenService) VerifyToken(token string) (*jwt.Token, error) {
	at, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("signmethod problem")
		}
		return []byte(u.cfg.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return at, nil
}

func (u *TokenService) GetClaims(token string) (claimMap map[string]interface{}, err error) {
	claimMap = map[string]interface{}{}

	verifyToken, err := u.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, errors.New("claims not found")
}
