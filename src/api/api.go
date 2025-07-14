package api

import "github.com/gin-gonic/gin"

func initServer() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1/")

}
