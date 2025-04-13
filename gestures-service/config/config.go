package config

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config struct
type Config struct {
	App        App        `yaml:"app"`
	PostgreSQL PostgreSQL `yaml:"postgres"`
}

// App config struct
type App struct {
	Env          string `yaml:"env" env:"APP_ENV" env-default:"development"`
	GRPC         GRPC   `yaml:"grpc"`
	JWTSecretKey string `yaml:"jwt_secret_key" env-required:"true"`
}

// GRPC server config struct
type GRPC struct {
	Port    int64         `yaml:"port" env:"GRPC_PORT" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env:"GRPC_TIMEOUT" env-required:"true"`
}

// PostgreSQL config struct
type PostgreSQL struct {
	Host        string        `env:"POSTGRES_HOST" env-required:"true" env-default:"localhost"`
	Port        int64         `env:"POSTGRES_PORT" env-required:"true"`
	User        string        `env:"POSTGRES_USER" env-required:"true"`
	Password    string        `env:"POSTGRES_PASSWORD" env-required:"true"`
	DBName      string        `env:"POSTGRES_DB" env-required:"true"`
	SSLMode     string        `env:"POSTGRES_SSLMODE" env-required:"true"`
	MaxPoolSize int32         `yaml:"max_pool_size" env:"POSTGRES_MAX_POOL_SIZE" env-required:"true"`
	ConnTimeout time.Duration `yaml:"conn_timeout" env:"POSTGRES_CONN_TIMEOUT" env-required:"true"`
	Driver      string        `yaml:"driver" env:"POSTGRES_DRIVER" env-required:"true"`
}

// Load config file from given path and env variables
func LoadConfig(filename string) (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(filepath.Join(".", filename), &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to read env variables: %w", err)
	}

	return &cfg, nil
}
