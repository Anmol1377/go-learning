package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string
}

type Config struct {
	ENV         string `yaml:"env" envconfig:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var ConfigPath string

	// ConfigPath = os.Getenv("CONFIG_PATH")
	ConfigPath = "/Users/inc42/Desktop/goo/Go-prod/config/local.yaml"

	if ConfigPath == "" {
		flags := flag.String("config", "", "path to config file")
		flag.Parse()
		ConfigPath = *flags

		if ConfigPath == "" {
			log.Fatal("config not set")
		}

	}
	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist %s", ConfigPath)
	}
	var cfg Config
	err := cleanenv.ReadConfig(ConfigPath, &cfg)
	if err != nil {
		log.Fatalf("cant readf config %s", err.Error())
	}
	return &cfg
}
