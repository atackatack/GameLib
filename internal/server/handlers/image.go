package handlers

import (
	"GAMELIB/internal/actions"
	"GAMELIB/pkg/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetMinioImage(ctx *gin.Context) {
	imgName := ctx.Param("name")
	img, err := actions.GetImage(ctx, imgName, h.Minio.Config.ImageBucketName, h.Minio.DataBase)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	data, _ := img.Stat()
	if data.Size == 0 {
		ctx.JSON(http.StatusNoContent, gin.H{"error": "File not found"})
		return
	}
	extraHeaders := map[string]string{
		"Content-Disposition": "inline; filename=" + imgName,
	}
	ctx.DataFromReader(http.StatusOK, data.Size, "image/png", img, extraHeaders)
}

func (h *Handler) PostMinioImage(ctx *gin.Context) {
	img, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	info, err := actions.PostImage(ctx, img, h.Minio.Config.ImageBucketName, h.Minio.DataBase)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"filepath": info.Key,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"filepath": info.Key,
	})
}

func (h *Handler) DeleteMinioImage(ctx *gin.Context) {
	name := ctx.Param("name")
	err := actions.DeleteImage(ctx, name, h.Minio.Config.ImageBucketName, h.Minio.DataBase)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.ErrorResponse(err))
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "image deleted",
	})
}
