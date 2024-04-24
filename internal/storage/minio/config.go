package minio

import "github.com/spf13/viper"

type MinioConfig struct {
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	AccessKey       string `yaml:"access_key"`
	SecretKey       string `yaml:"secret_key"`
	ImageBucketName string `yaml:"image_bucket_name"`
	SSL             bool   `yaml:"ssl"`
}

func InitMinioConfig() MinioConfig {
	return MinioConfig{
		Host:            viper.GetString("minio.host"),
		Port:            viper.GetString("minio.port"),
		AccessKey:       viper.GetString("minio.access_key"),
		SecretKey:       viper.GetString("minio.secret_key"),
		ImageBucketName: viper.GetString("minio.image_bucket_name"),
		SSL:             viper.GetBool("minio.SSL"),
	}
}
