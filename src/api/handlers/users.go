package handlers

import (
	"net/http"

	"github.com/FathiMohammadDev/car-selling/api/dto"
	"github.com/FathiMohammadDev/car-selling/api/helpers"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/services"
	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	service *services.UserService
}

func NewUsersHandler(cfg *config.Config) *UsersHandler {
	return &UsersHandler{service: services.NewUserService(cfg)}
}

// SendOtp godoc
// @Summary Send otp to user
// @Description Send otp to user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.GetOtpRequest true "GetOtpRequest"
// @Router /v1/users/send-otp [post]
func (h *UsersHandler) SendOtp(ctx *gin.Context) {
	req := new(dto.GetOtpRequest)
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResWithValidationError(nil, false, 409, err))
	}
	err = h.service.SendOtp(*req, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResWithErr(nil, false, 409, err))
	}

	ctx.JSON(http.StatusCreated, helpers.GenerateBaseRes(nil, true, 0))
}

// RegisterByUserName godoc
// @Summary Registe user by userName
// @Description Registe user by userName
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.RegisterUserByUsernameRequest true "RegisterUserByUsernameRequest"
// @Router /v1/users/register-by-username [post]
func (h *UsersHandler) RegisterByUserName(ctx *gin.Context) {
	req := new(dto.RegisterUserByUsernameRequest)
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResWithValidationError(nil, false, 409, err))
	}

	err = h.service.RegisterByUserName(*req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResWithErr(nil, false, 409, err))
	}

	ctx.JSON(http.StatusCreated, helpers.GenerateBaseRes(nil, true, 0))
}

// LoginByUserName godoc
// @Summary Login user by userName
// @Description Login user by userName
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.LoginByUsernameRequest true "LoginByUsernameRequest"
// @Router /v1/users/login-by-username [post]
func (h *UsersHandler) LoginByUserName(ctx *gin.Context) {
	req := new(dto.LoginByUsernameRequest)
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResWithValidationError(nil, false, 409, err))
	}

	token, err := h.service.LoginByUserName(*req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResWithErr(nil, false, 409, err))
	}

	ctx.JSON(http.StatusCreated, helpers.GenerateBaseRes(token, true, 0))
}
