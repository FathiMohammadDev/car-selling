package routers

import (
	"github.com/FathiMohammadDev/car-selling/api/handlers"
	"github.com/FathiMohammadDev/car-selling/api/middlewares"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/gin-gonic/gin"
)

func Users(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)
	router.POST("/send-otp", middlewares.OtpLimiter(cfg), h.SendOtp)
	router.POST("/register-by-username", h.RegisterByUserName)
	router.POST("/login-by-username", h.LoginByUserName)
}
