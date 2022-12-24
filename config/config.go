package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost          string `mapstructure:"DB_HOST"`
	DBPort          string `mapstructure:"DB_PORT"`
	DBUsername      string `mapstructure:"DB_USERNAME"`
	DBPassword      string `mapstructure:"DB_PASSWORD"`
	DBName          string `mapstructure:"DB_NAME"`
	JWTSecretToken  string `mapstructure:"JWT_Secret_Token"`
	EmailFrom       string `mapstructure:"EMAIL_FROM"`
	SMTPAppPassword string `mapstructure:"SMTP_APP_PASSWORD"`
	SMTPHost        string `mapstructure:"SMTP_HOST"`
	SMTPPort        int    `mapstructure:"SMTP_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
