package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		PG   `yaml:"postgres"`
		Nats `yaml:"nats"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name" env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}
	HTTP struct {
		Host string `env-required:"true" yaml:"host" env:"HTTP_HOST"`
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	PG struct {
		Host     string `env-required:"true" yaml:"host" env:"PG_HOST"`
		Port     string `env-required:"true" yaml:"port" env:"PG_PORT"`
		User     string `env-required:"true" yaml:"user" env:"PG_USER"`
		Password string `env-required:"true" yaml:"password" env:"PG_PASSWORD"`
		DBName   string `env-required:"true" yaml:"name" env:"PG_NAME"`
		PgDriver string `env-required:"true" yaml:"pg_driver" env:"PG_PG_DRIVER"`
	}

	Nats struct {
		Host    string `env-required:"true" yaml:"host" env:"NATS_HOST"`
		Port    string `env-required:"true" yaml:"port" env:"NATS_PORT"`
		Cluster string `env-required:"true" yaml:"cluster" env:"NATS_CLUSTER"`
		Client  string `env-required:"true" yaml:"client" env:"NATS_CLIENT"`
		Topic   string `env-required:"true" yaml:"topic" env:"NATS_TOPIC"`
	}
)

func getEnv() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Ошибка получения текущего каталога:", err)
		return
	}

	fmt.Println("Содержимое текущего каталога:")
	// Открываем текущий каталог
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Ошибка чтения каталога:", err)
		return
	}

	// Выводим имена файлов и каталогов
	for _, entry := range dirEntries {
		fmt.Println(entry.Name())
	}
}
func NewConfig() (*Config, error) {
	cfg := &Config{}
	getEnv()
	path := "bin/config/config.yml"

	err := cleanenv.ReadConfig(path, cfg)

	if err != nil {
		return nil, fmt.Errorf("Config error: %v", err)
	}

	err = cleanenv.ReadEnv(cfg)

	if err != nil {
		return nil, fmt.Errorf("Reading enviroment error: %v", err)
	}

	return cfg, nil
}
