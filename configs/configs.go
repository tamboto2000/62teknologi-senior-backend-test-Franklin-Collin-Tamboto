package configs

import (
	"github.com/joho/godotenv"
)

type Config struct {
	BusinessUrlBase string `yaml:"business_url_base"`
	AssetsUrlBase   string `yaml:"assets_url_base"`
}

func LoadLocalEnv() error {
	return godotenv.Load()
}
