package handlers

import (
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/services"
	"github.com/gin-gonic/gin"
)

type YearHandler struct {
	service services.YearService
}

func NewYearHandler(cfg *config.Config) *YearHandler {
	return &YearHandler{
		*services.NewYearService(cfg),
	}
}

// CreateYear godoc
// @Summary Create new Year
// @Description Create new Year
// @Tags Year
// @Accept  json
// @Produce  json
// @Param Request body dto.CreateYearReq true "CreateYearReq"
// @Router /v1/year/ [post]
// @Security AuthBearer
func (h *YearHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// CreateYear godoc
// @Summary Update Year
// @Description Update Year
// @Tags Year
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Param Request body dto.UpdateYearReq true "UpdateYearReq"
// @Router /v1/year/{id} [put]
// @Security AuthBearer
func (h *YearHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteYear godoc
// @Summary delete Year
// @Description delete Year
// @Tags Year
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/year/{id} [delete]
// @Security AuthBearer
func (h *YearHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)

}

// GetCty godoc
// @Summary Get a Year
// @Description Get a Year
// @Tags Year
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/year/{id} [get]
// @Security AuthBearer
func (h *YearHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}
