package config

import "github.com/spf13/viper"

type Config struct {
	Port       string `json:"PORT"`
	DBHost     string `json:"DB_HOST"`
	DBUser     string `json:"DB_USER"`
	DBPassword string `json:"DB_PASS"`
	DBName     string `json:"DB_NAME"`
	DBPort     string `json:"DB_PORT"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)
	return
}
