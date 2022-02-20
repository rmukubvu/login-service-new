package config

import (
	"github.com/spf13/viper"
)

const (
	mongoUrl   = "mongodatabase.uri"
	serverPort = "webserver.port"
)

var (
	mongoConfig  MongoConfig
	serverConfig WebServerConfig
)

type MongoConfig struct {
	Url string
}

type WebServerConfig struct {
	Port int
}

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
	mongoConfig.Url = viper.GetString(mongoUrl)
	serverConfig.Port = viper.GetInt(serverPort)
}

func MongoUrlConfig() *MongoConfig {
	return &mongoConfig
}

func WebServerPort() int {
	return serverConfig.Port
}
