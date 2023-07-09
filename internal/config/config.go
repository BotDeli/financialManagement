package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	HttpServer HTTPServerConfig `yaml:"http-HTTPServer" env-required:"true"`
	Logger     LoggerConfig     `yaml:"logger" env-required:"true"`
	Postgres   PostgresConfig   `yaml:"postgres" env-required:"true"`
}

type HTTPServerConfig struct {
	Address  string        `yaml:"address" env-default:"localhost:8080"`
	Timeout  time.Duration `yaml:"timeout" env-default:"4s"`
	Lifetime time.Duration `yaml:"lifetime" env-default:"60s"`
}

type LoggerConfig struct {
	Level  string `yaml:"level" env-default:"info"`
	Format string `yaml:"format" env-default:"text"`
}

type PostgresConfig struct {
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-default:""`
	Host     string `yaml:"host" env-default:"localhost"`
	Dbname   string `yaml:"dbname" env-required:"true"`
	Sslmode  string `yaml:"sslmode" env-default:"disable"`
}

func (config *PostgresConfig) GetDataSourceName() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		config.User,
		config.Password,
		config.Host,
		config.Dbname,
		config.Sslmode)
}

const (
	errPathIsNotSet   = "config file path is not set"
	errFileIsNotExist = "config file is not exist"
	errReadConfigFile = "error read config file: %s"
)

func MustGetOSConfig() Config {
	path := os.Getenv("ConfigPath")
	cfg := MustGetConfig(path)
	return cfg
}

func MustGetConfig(path string) (cfg Config) {
	checkSetPath(path)
	checkFileIsExist(path)
	readConfig(path, &cfg)
	return
}

func checkSetPath(path string) {
	if path == "" {
		log.Fatal(errPathIsNotSet)
	}
}

func checkFileIsExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal(errFileIsNotExist)
	}
}

func readConfig(path string, cfg *Config) {
	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		log.Fatalf(errReadConfigFile, err)
	}
}
