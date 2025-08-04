package handlers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/FathiMohammadDev/car-selling/api/dto"
	"github.com/FathiMohammadDev/car-selling/api/helpers"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
	"github.com/FathiMohammadDev/car-selling/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FileHandler struct {
	service *services.FileService
}

var logger = logging.NewLogger(config.GetConfig())

func NewFileHandler(cfg *config.Config) *FileHandler {
	return &FileHandler{
		service: services.NewFileService(cfg),
	}
}

// CreateFile godoc
// @Summary Create a file
// @Description Create a file
// @Tags Files
// @Accept x-www-form-urlencoded
// @produces json
// @Param file formData dto.UploadFileRequest true "Create a file"
// @Param file formData file true "Create a file"
// @Router /v1/files/ [post]
// @Security AuthBearer
func (h *FileHandler) Create(c *gin.Context) {
	upload := dto.UploadFileRequest{}
	err := c.ShouldBind(&upload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResWithValidationError(nil, false, 121, err))
		return
	}
	req := dto.CreateFileRequest{}
	req.Description = upload.Description
	//req.Name = upload.File.Filename
	req.MimeType = upload.File.Header.Get("Content-Type")
	req.Directory = "uploads"
	req.Name, err = saveUploadFile(upload.File, req.Directory)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helpers.GenerateBaseResWithErr(nil, false, 122, err))
		return
	}

	res, err := h.service.Create(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helpers.GenerateBaseResWithErr(nil, false, 123, err))
		return
	}
	c.JSON(http.StatusCreated, helpers.GenerateBaseRes(res, true, 0))

}

// DeleteFile godoc
// @Summary Delete a file
// @Description Delete a file
// @Tags Files
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Router /v1/files/{id} [delete]
// @Security AuthBearer
func (h *FileHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helpers.GenerateBaseRes(nil, false, 121))
		return
	}
	file, err := h.service.GetById(c, id)
	if err != nil {
		logger.Error(logging.IO, logging.RemoveFile, err.Error(), nil)
		c.AbortWithStatusJSON(http.StatusNotFound,
			helpers.GenerateBaseRes(nil, false, 122))
		return
	}
	err = os.Remove(fmt.Sprintf("%s/%s", file.Directory, file.Name))
	if err != nil {
		logger.Error(logging.IO, logging.RemoveFile, err.Error(), nil)
		c.AbortWithStatusJSON(http.StatusNotFound,
			helpers.GenerateBaseRes(nil, false, 123))
		return
	}
	err = h.service.Delete(c, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helpers.GenerateBaseResWithErr(nil, false, 124, err))
		return
	}
	c.JSON(http.StatusOK, helpers.GenerateBaseRes(nil, true, 0))
}

// UpdateFile godoc
// @Summary Update a file
// @Description Update a file
// @Tags Files
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateFileRequest true "Update a file"
// @Router /v1/files/{id} [put]
// @Security AuthBearer
func (h *FileHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	req := dto.UpdateFileRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResWithValidationError(nil, false, 121, err))
		return
	}

	res, err := h.service.Update(c, id, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helpers.GenerateBaseResWithErr(nil, false, 121, err))
		return
	}
	c.JSON(http.StatusOK, helpers.GenerateBaseRes(res, true, 0))
}

// GetFile godoc
// @Summary Get a file
// @Description Get a file
// @Tags Files
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Router /v1/files/{id} [get]
// @Security AuthBearer
func (h *FileHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helpers.GenerateBaseRes(nil, false, 121))
		return
	}

	res, err := h.service.GetById(c, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helpers.GenerateBaseResWithErr(nil, false, 121, err))
		return
	}
	c.JSON(http.StatusOK, helpers.GenerateBaseRes(res, true, 0))
}

func saveUploadFile(file *multipart.FileHeader, directory string) (string, error) {
	randFileName := uuid.New()
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return "", err
	}
	fileName := file.Filename
	fileNameArr := strings.Split(fileName, ".")
	fileExt := fileNameArr[len(fileNameArr)-1]
	fileName = fmt.Sprintf("%s.%s", randFileName, fileExt)
	dst := fmt.Sprintf("%s/%s", directory, fileName)

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return "", err
	}
	return fileName, nil
}
