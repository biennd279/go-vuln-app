package config

import (
	"errors"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("config",
	fx.Provide(
		provideConfig,
	),
)

func provideConfig() *Config {
	logger := zap.NewExample()

	cfg := viper.New()
	cfg.SetConfigFile(".env")
	cfg.AllowEmptyEnv(true)

	cfg.SetDefault("isDev", true)
	cfg.SetDefault("server.host", "0.0.0.0")
	cfg.SetDefault("server.port", "8080")

	var config Config

	if err := cfg.Unmarshal(&config); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			logger.Warn("Config file not found", zap.Error(err))
		}
	}

	return &config
}
