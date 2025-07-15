package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"os"
)

type Config struct {
	env       string    `yaml:"env" default:"dev"`
	appName   string    `yaml:"appName"`
	Postgres  Storage   `yaml:"postgres"`
	Tarantool Tarantool `yaml:"tarantool"`
}

type Storage struct {
	Host   string `yaml:"POSTGRES_HOST"`
	UserDb string `yaml:"POSTGRES_USER"`
	DbName string `yaml:"POSTGRES_DB"`
	PassDb string `yaml:"POSTGRES_PASSWORD"`
	PortDb string `yaml:"POSTGRES_PORT"`
}

type Tarantool struct {
	UserDb string `yaml:"TARANTOOL_USER"`
	PassDb string `yaml:"TARANTOOL_PASS"`
	PortDb string `yaml:"TARANTOOL_PORT"`
	HostDb string `yaml:"TARANTOOL_HOST"`
}

func MustLoad() *Config {
	fetchPath := fetchConfigPath()
	if fetchPath == "" {
		panic("Пустой файл конфигурации")
	}

	return MustLoadByPath(fetchPath)
}

func MustLoadByPath(fetchPath string) *Config {
	if _, err := os.Stat(fetchPath); os.IsNotExist(err) {
		panic("Не существует файл конфигурации: " + fetchPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(fetchPath, &cfg); err != nil {
		panic("Ошибка чтения конфига" + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}

func SetupLoger(env Config) *slog.Logger {
	var log *slog.Logger

	switch env.env {
	case "dev":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "prod":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
