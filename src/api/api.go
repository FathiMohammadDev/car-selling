package api

import (
	"github.com/FathiMohammadDev/car-selling/api/middlewares"
	"github.com/FathiMohammadDev/car-selling/api/validations"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func initServer() {
	cfg := config.GetConfig()
	r := gin.New()

	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("password", validations.PasswordValidator)
	}
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest(), middlewares.Cros(cfg))

	v1 := r.Group("/api/v1/")

}
