package api

import (
	"github.com/FathiMohammadDev/car-selling/api/middlewares"
	"github.com/gin-gonic/gin"
)

func initServer() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest())

	v1 := r.Group("/api/v1/")

}
