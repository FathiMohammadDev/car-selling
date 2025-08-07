package handlers

import (
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/services"
	"github.com/gin-gonic/gin"
)

type CarModelCommentHandler struct {
	service services.CarModelCommentService
}

func NewCarModelCommentHandler(cfg *config.Config) *CarModelCommentHandler {
	return &CarModelCommentHandler{
		*services.NewCarModelCommentService(cfg),
	}
}

// CreateCarModelComment godoc
// @Summary Create new CarModelComment
// @Description Create new CarModelComment
// @Tags CarModelComment
// @Accept  json
// @Produce  json
// @Param Request body dto.CreateCarModelCommentReq true "CreateCarModelCommentReq"
// @Router /v1/car-model-comment/ [post]
// @Security AuthBearer
func (h *CarModelCommentHandler) Create(ctx *gin.Context) {
	Create(ctx, h.service.Create)
}

// CreateCarModelComment godoc
// @Summary Update CarModelComment
// @Description Update CarModelComment
// @Tags CarModelComment
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelCommentReq true "UpdateCarModelCommentReq"
// @Router /v1/car-model-comment/{id} [put]
// @Security AuthBearer
func (h *CarModelCommentHandler) Update(ctx *gin.Context) {
	Update(ctx, h.service.Update)
}

// DeleteCarModelComment godoc
// @Summary delete CarModelComment
// @Description delete CarModelComment
// @Tags CarModelComment
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/car-model-comment/{id} [delete]
// @Security AuthBearer
func (h *CarModelCommentHandler) Delete(ctx *gin.Context) {
	Delete(ctx, h.service.Delete)

}

// GetCty godoc
// @Summary Get a CarModelComment
// @Description Get a CarModelComment
// @Tags CarModelComment
// @Accept  json
// @Produce  json
// @Param id path int true "Id"
// @Router /v1/car-model-comment/{id} [get]
// @Security AuthBearer
func (h *CarModelCommentHandler) GetById(ctx *gin.Context) {
	GetById(ctx, h.service.GetById)
}
