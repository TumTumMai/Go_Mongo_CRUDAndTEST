package config

import (
	"strings"

	"github.com/spf13/viper"
)

type AppConfiguration struct {
	Port int
}

type MongoConfiguration struct {
	Username string
	Password string
}

type Configurations struct {
	Environment string
	App         AppConfiguration
	Mongo       MongoConfiguration
}

func initialConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	// read config from ENV
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}

func GetConfig() Configurations {
	initialConfig()
	config := Configurations{}
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return config
}
