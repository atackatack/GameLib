package minio

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
)

func UploadImage(ctx *gin.Context, img *multipart.FileHeader, bucketName string, client *minio.Client) (*minio.UploadInfo, error) {
	buffer, err := img.Open()
	if err != nil {
		return nil, err
	}

	info, err := client.PutObject(ctx, bucketName, img.Filename, buffer, img.Size,
		minio.PutObjectOptions{
			ContentType: img.Header.Get("Content-Type"),
		})

	if err != nil {
		return nil, err
	}
	return &info, err
}

func GetImage(ctx *gin.Context, imgName string, bucketName string, client *minio.Client) (*minio.Object, error) {
	object, err := client.GetObject(ctx, bucketName, imgName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return object, err
}

func DeleteImage(ctx *gin.Context, imgName string, bucketName string, client *minio.Client) error {
	err := client.RemoveObject(ctx, bucketName, imgName, minio.RemoveObjectOptions{})
	return err
}
