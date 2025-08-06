package handlers

import (
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/services"
	"github.com/gin-gonic/gin"
)

type CarModelYearlHandler struct {
	service services.CarModelYearService
}

func NewCarModelYearlHandler(cfg *config.Config) *CarModelYearlHandler {
	return &CarModelYearlHandler{
		*services.NewCarModelYearService(cfg),
	}
}

// CreateCarModelYearl godoc
// @Summary Create new CarModelYear
// @Description Create new CarModelYear
// @Tags CarModelYear
// @Accept  json
// @Produce  json
// @Param Request body dto.CreateCarModelYearReq true "CreateCarModelYearReq"
// @Router /v1/car-model-year/ [post]
// @Security AuthBearer
func (h *CarModelYearlHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// CreateCarModelYearl godoc
// @Summary Update CarModelYear
// @Description Update CarModelYear
// @Tags CarModelYear
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelYearReq true "UpdateCarModelYearReq"
// @Router /v1/car-model-year/{id} [put]
// @Security AuthBearer
func (h *CarModelYearlHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteCarModelYear godoc
// @Summary delete CarModelYear
// @Description delete CarModelYear
// @Tags CarModelYear
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/car-model-year/{id} [delete]
// @Security AuthBearer
func (h *CarModelYearlHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)

}

// GetCty godoc
// @Summary Get a CarModelYear
// @Description Get a CarModelYear
// @Tags CarModelYear
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/car-model-year/{id} [get]
// @Security AuthBearer
func (h *CarModelYearlHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}
