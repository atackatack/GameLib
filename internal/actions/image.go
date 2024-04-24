package actions

import (
	minioClient "GAMELIB/internal/storage/minio"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
)

func PostImage(ctx *gin.Context, img *multipart.FileHeader, bucketName string, client *minio.Client) (*minio.UploadInfo, error) {
	return minioClient.UploadImage(ctx, img, bucketName, client)
}

func GetImage(ctx *gin.Context, imgName string, bucketName string, client *minio.Client) (*minio.Object, error) {
	return minioClient.GetImage(ctx, imgName, bucketName, client)
}

func DeleteImage(ctx *gin.Context, imgName string, bucketName string, client *minio.Client) error {
	return minioClient.DeleteImage(ctx, imgName, bucketName, client)
}
