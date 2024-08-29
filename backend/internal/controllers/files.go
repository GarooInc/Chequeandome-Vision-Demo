package controllers

import (
	"Chequeandome-Vision-Demo/internal/configs"
	"Chequeandome-Vision-Demo/internal/database"
	"Chequeandome-Vision-Demo/internal/responses"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

type Form struct {
	Files []*multipart.FileHeader `form:"files" binding:"required"`
}

// Upload sube un archivo a S3 y lo asocia al usuario que lo subió
// @Summary Sube un archivo a S3 y lo asocia al usuario que lo subió
// @Description Sube un archivo a S3 y lo asocia al usuario que lo subió.
// @Tags Files
// @Accept  mpfd
// @Produce  json
// @Param files formData file true "Archivo a subir"
// @Success 200 {object} responses.UploadSuccessResponse
// @Failure 400 {object} responses.StandardResponse
// @Failure 500 {object} responses.StandardResponse
// @Router / [post]
func Upload(c *gin.Context) {
	var form Form

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, responses.StandardResponse{
			Code:    http.StatusBadRequest,
			Message: "File is required",
		})
		return
	}

	paths := make([]string, 0)

	for _, file := range form.Files {
		// Upload to S3

		path := fmt.Sprintf("%s", file.Filename)

		err := database.Instance.Insert(path, file)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: fmt.Sprintf("Error uploading file: %s", err.Error()),
			})
			return
		}

		paths = append(paths, fmt.Sprintf("%s/files/%s", configs.URL, file.Filename))
	}

	c.JSON(http.StatusOK, responses.UploadSuccessResponse{
		StandardResponse: responses.StandardResponse{
			Code:    http.StatusOK,
			Message: "File uploaded successfully",
		},
		Paths: paths,
	})
}

// GetFile obtiene un archivo de S3
// @Summary Obtiene un archivo de S3
// @Description Obtiene un archivo de S3
// @Tags Files
// @Accept  json
// @Produce  mpfd
// @Param file path string true "Nombre del archivo"
// @Success 200 {object} responses.StandardResponse
// @Failure 400 {object} responses.StandardResponse
// @Failure 500 {object} responses.StandardResponse
// @Router /files/:file [get]
func GetFile(c *gin.Context) {
	name := c.Param("file")

	if name == "" {
		c.JSON(http.StatusBadRequest, responses.StandardResponse{
			Code:    http.StatusBadRequest,
			Message: "Directory and name are required",
		})
		return
	}

	path := fmt.Sprintf("%s", name)

	file, err := database.Instance.GetFile(path)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.StandardResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Error getting file: %s", err.Error()),
		})
		return
	}

	contentType := http.DetectContentType(file.Bytes())

	//c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))
	c.Header("Content-Type", contentType)
	c.Header("Content-Length", fmt.Sprintf("%d", len(file.Bytes())))
	c.Data(http.StatusOK, contentType, file.Bytes())
}
