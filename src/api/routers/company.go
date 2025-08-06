package routers

import (
	"github.com/FathiMohammadDev/car-selling/api/handlers"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/gin-gonic/gin"
)

func Company(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCompanyHandler(cfg)
	router.POST("/", h.Create)
	router.PUT("/:id", h.Update)
	router.DELETE("/:id", h.Delete)
	router.GET("/:id", h.GetById)
}
