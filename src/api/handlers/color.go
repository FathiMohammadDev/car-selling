package handlers

import (
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/services"
	"github.com/gin-gonic/gin"
)

type ColorHandler struct {
	service services.ColorService
}

func NewColorHandler(cfg *config.Config) *ColorHandler {
	return &ColorHandler{
		*services.NewColorService(cfg),
	}
}

// CreateColor godoc
// @Summary Create new Color
// @Description Create new Color
// @Tags Color
// @Accept  json
// @Produce  json
// @Param Request body dto.CreateColorReq true "CreateColorReq"
// @Router /v1/color/ [post]
// @Security AuthBearer
func (h *ColorHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// CreateColor godoc
// @Summary Update Color
// @Description Update Color
// @Tags Color
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Param Request body dto.UpdateColorReq true "UpdateColorReq"
// @Router /v1/color/{id} [put]
// @Security AuthBearer
func (h *ColorHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteColor godoc
// @Summary delete Color
// @Description delete Color
// @Tags Color
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/color/{id} [delete]
// @Security AuthBearer
func (h *ColorHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)

}

// GetCty godoc
// @Summary Get a Color
// @Description Get a Color
// @Tags Color
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/color/{id} [get]
// @Security AuthBearer
func (h *ColorHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}
