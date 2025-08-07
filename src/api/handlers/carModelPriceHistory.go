package handlers

import (
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/services"
	"github.com/gin-gonic/gin"
)

type CarModelPriceHistoryHandler struct {
	service services.CarModelPriceHistoryService
}

func NewCarModelPriceHistoryHandler(cfg *config.Config) *CarModelPriceHistoryHandler {
	return &CarModelPriceHistoryHandler{
		*services.NewCarModelPriceHistoryService(cfg),
	}
}

// CreateCarModelPriceHistory godoc
// @Summary Create new CarModelPriceHistory
// @Description Create new CarModelPriceHistory
// @Tags CarModelPriceHistory
// @Accept  json
// @Produce  json
// @Param Request body dto.CreateCarModelPriceHistoryReq true "CreateCarModelPriceHistoryReq"
// @Router /v1/car-model-price-history/ [post]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// CreateCarModelPriceHistory godoc
// @Summary Update CarModelPriceHistory
// @Description Update CarModelPriceHistory
// @Tags CarModelPriceHistory
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelPriceHistoryReq true "UpdateCarModelPriceHistoryReq"
// @Router /v1/car-model-price-history/{id} [put]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteCarModelPriceHistory godoc
// @Summary delete CarModelPriceHistory
// @Description delete CarModelPriceHistory
// @Tags CarModelPriceHistory
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/car-model-price-history/{id} [delete]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)

}

// GetCty godoc
// @Summary Get a CarModelPriceHistory
// @Description Get a CarModelPriceHistory
// @Tags CarModelPriceHistory
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/car-model-price-history/{id} [get]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}
