package postgres

import "github.com/spf13/viper"

type StorageConfig struct {
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslMode"`
}

func InitStorageConfig() StorageConfig {
	return StorageConfig{
		Name:     viper.GetString("storage.name"),
		Password: viper.GetString("storage.password"),
		Host:     viper.GetString("storage.host"),
		Port:     viper.GetString("storage.port"),
		DBName:   viper.GetString("storage.dbname"),
		SSLMode:  viper.GetString("storage.sslMode"),
	}
}
