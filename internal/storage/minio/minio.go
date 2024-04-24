package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

type Client struct {
	DataBase *minio.Client
	Config   MinioConfig
}

func NewStorage(cfg MinioConfig) (Client, error) {
	ctx := context.Background()
	client, err := minio.New(cfg.Host+":"+cfg.Port, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.SSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	CreateImageBucket(ctx, client, cfg.ImageBucketName)

	return Client{
		DataBase: client,
		Config:   cfg,
	}, err
}

func CreateImageBucket(ctx context.Context, client *minio.Client, bucketName string) {
	location := "us-east-1"
	err := client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := client.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("[Minio] We already own %s: %s\n", bucketName, err)
		}
	}
}
