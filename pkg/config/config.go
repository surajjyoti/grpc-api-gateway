package config

import "github.com/spf13/viper"

type Config struct {
	Port              string `mapstructure:"PORT"`
	AuthServiceUrl    string `mapstructure:"AUTH_SERVICE_URL"`
	ProductServiceUrl string `mapstructure:"PRODUCT_SERVICE_URL"`
	OrderServiceUrl   string `mapstructure:"ORDER_SERVICE_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
