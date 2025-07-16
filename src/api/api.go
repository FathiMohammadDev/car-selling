package api

import (
	"github.com/FathiMohammadDev/car-selling/api/middlewares"
	"github.com/FathiMohammadDev/car-selling/api/validations"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func initServer() {
	r := gin.New()

	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("password", validations.PasswordValidator)
	}
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest())

	v1 := r.Group("/api/v1/")

}
