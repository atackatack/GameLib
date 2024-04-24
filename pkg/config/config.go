package config

import (
	"GAMELIB/internal/storage/minio"
	"GAMELIB/internal/storage/postgres"
	"GAMELIB/pkg/web"
	"github.com/spf13/viper"
)

type Config struct {
	Server  web.ServerConfig
	Storage postgres.StorageConfig
	Minio   minio.MinioConfig
}

func ParseConfig(env string) error {
	viper.AddConfigPath("configs/" + env)
	viper.SetConfigName("common")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
