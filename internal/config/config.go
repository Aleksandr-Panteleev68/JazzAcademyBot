package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	TelegramBotToken string   `mapstructure:"TELEGRAM_BOT_TOKEN"`
	GoogleCredsPath  string   `mapstructure:"GOOGLE_CREDS_PATH"`
	DBConnString     string   `mapstructure:"DB_CONN_STRING"`
	Admins           []string `mapstructure:"ADMINS"`
	LogLevel         string   `mapstructure:"LOG_LEVEL"`
}

func Load() (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config: %w", err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if cfg.TelegramBotToken == "" {
		return nil, fmt.Errorf("TELEGRAM_BOT_TOKEN is required")
	}

	if cfg.GoogleCredsPath == "" {
		return nil, fmt.Errorf("GOOGLE_CREDS_PATH is required")
	}

	if cfg.DBConnString == "" {
		return nil, fmt.Errorf("DB_CONN_STRING is required")
	}

	return &cfg, nil
}
