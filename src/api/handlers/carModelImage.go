package handlers

import (
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/services"
	"github.com/gin-gonic/gin"
)

type CarModelImageHandler struct {
	service services.CarModelImageService
}

func NewCarModelImageHandler(cfg *config.Config) *CarModelImageHandler {
	return &CarModelImageHandler{
		*services.NewCarModelImageService(cfg),
	}
}

// CreateCarModelImage godoc
// @Summary Create new CarModelImage
// @Description Create new CarModelImage
// @Tags CarModelImage
// @Accept  json
// @Produce  json
// @Param Request body dto.CreateCarModelImageReq true "CreateCarModelImageReq"
// @Router /v1/car-model-image/ [post]
// @Security AuthBearer
func (h *CarModelImageHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// CreateCarModelImage godoc
// @Summary Update CarModelImage
// @Description Update CarModelImage
// @Tags CarModelImage
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelImageReq true "UpdateCarModelImageReq"
// @Router /v1/car-model-image/{id} [put]
// @Security AuthBearer
func (h *CarModelImageHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteCarModelImage godoc
// @Summary delete CarModelImage
// @Description delete CarModelImage
// @Tags CarModelImage
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/car-model-image/{id} [delete]
// @Security AuthBearer
func (h *CarModelImageHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)

}

// GetCty godoc
// @Summary Get a CarModelImage
// @Description Get a CarModelImage
// @Tags CarModelImage
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/car-model-image/{id} [get]
// @Security AuthBearer
func (h *CarModelImageHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}
