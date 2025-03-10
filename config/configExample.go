package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	ServerHost  string `mapstructure:"SERVER_HOST"`
	ServerPort  int    `mapstructure:"SERVER_PORT"`

	// Database
	DbDsn string `mapstructure:"DB_DSN"`

	// Email
	SmtpSender string `mapstructure:"SMTP_SENDER"`
	SmtpHost   string `mapstructure:"SMTP_HOST"`
	SmtpPort   int    `mapstructure:"SMTP_PORT"`
	SmtpUser   string `mapstructure:"SMTP_USER"`
	SmtpPass   string `mapstructure:"SMTP_PASS"`
	StarTLS    bool   `mapstructure:"STARTLS"`
}

func LoadConfig(path string) (config Config) {
	v := viper.New()

	// Default Config
	v.SetDefault("ENVIRONMENT", "development")
	v.SetDefault("SERVER_HOST", "localhost")
	v.SetDefault("SERVER_PORT", 8000)

	v.SetConfigFile(".env")
	v.SetConfigType("env")
	v.AddConfigPath(path)
	v.AutomaticEnv()

	_ = v.ReadInConfig()
	_ = v.Unmarshal(&config)
	return
}
