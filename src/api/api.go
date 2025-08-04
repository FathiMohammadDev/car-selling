package api

import (
	"fmt"

	"github.com/FathiMohammadDev/car-selling/api/middlewares"
	"github.com/FathiMohammadDev/car-selling/api/routers"
	"github.com/FathiMohammadDev/car-selling/api/validations"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	registerValidator()
	RegisterSwagger(r, cfg)
	r.Use(middlewares.DefaultStructurdLogger(cfg))
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest(), middlewares.Cros(cfg))

	registerRoutes(r, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
}

func registerRoutes(r *gin.Engine, cfg *config.Config) {

	v1 := r.Group("/api/v1/")
	{
		users := v1.Group("/users")
		country := v1.Group("/countries", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		city := v1.Group("/cities", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		file := v1.Group("/files", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))

		routers.File(file, cfg)
		routers.City(city, cfg)
		routers.Country(country, cfg)
		routers.Users(users, cfg)
	}

}

func registerValidator() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("password", validations.PasswordValidator)
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.ExternalPort)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
