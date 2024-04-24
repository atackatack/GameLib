package web

import "github.com/spf13/viper"

type ServerConfig struct {
	Port string `config:"port" yaml:"port" evn:"SERVER_PORT"`
}

func InitServerConfig() ServerConfig {
	return ServerConfig{
		Port: viper.GetString("server.port"),
	}
}
