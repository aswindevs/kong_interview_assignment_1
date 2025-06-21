package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		HTTP   `yaml:"http"`
		Logger `yaml:"logger"`
		PG     `yaml:"postgres"`
		Auth   `yaml:"auth"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Logger struct {
		Level  string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
		Format string `env-required:"true" yaml:"log_format"  env:"LOG_FORMAT"`
	}

	// PG -.
	PG struct {
		URL     string `env-required:"true"                 env:"PG_URL"`
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		Level   string `env-required:"true" yaml:"log_level" env:"PG_LOG_LEVEL"`
	}

	Auth struct {
		SecretKey           string   `env-required:"true" yaml:"secret_key" env:"AUTH_SECRET_KEY"`
		TokenExpirationTime int      `env-required:"true" yaml:"token_expiration_time" env:"AUTH_TOKEN_EXPIRATION_TIME"`
		Domain              string   `env-required:"true" yaml:"domain" env:"AUTH_DOMAIN"`
		CookieName          string   `env-required:"true" yaml:"cookie_name" env:"AUTH_COOKIE_NAME"`
		ExcludePaths        []string `env-required:"true" yaml:"exclude_paths" env:"AUTH_EXCLUDE_PATHS"`
	}
)

func NewConfig() (*Config, error) {
	return NewConfigWithPath("config/config.yaml")
}

func NewConfigWithPath(configPath string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, fmt.Errorf("Config Error: %w", err)
	}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
