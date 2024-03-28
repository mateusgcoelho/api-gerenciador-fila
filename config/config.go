package config

import "github.com/spf13/viper"

type appConfig struct {
	ServerPort       int    `mapstructure:"SERVER_PORT"`
	DatabaseHost     string `mapstructure:"DATABASE_HOST"`
	DatabasePort     int    `mapstructure:"DATABASE_PORT"`
	DatabaseUsername string `mapstructure:"DATABASE_USERNAME"`
	DatabaseName     string `mapstructure:"DATABASE_NAME"`
	JwtSecretKey     string `mapstructure:"JWT_SECRET_KEY"`
}

var (
	AppConfig appConfig
)

func loadConfig(path string) (appConfig appConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&appConfig)
	return
}

func InitAppConfig(path string) {
	config, err := loadConfig(path)
	if err != nil {
		panic(err)
	}

	AppConfig = config
}
