package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Cfg struct {
	Logger struct {
		Mode string `yaml:"mode" default:"debug"`
	} `yaml:"logger"`
	HTTP struct {
		Host string `yaml:"host" default:":8080"`
	} `yaml:"http"`
	PostgreSQL struct {
		DSN string `yaml:"dsn" default:"postgres://postgres:postgres@postgres:5432/postgres"`
	} `yaml:"postgresql"`
}

func NewConfig() Cfg {
	cfg := Cfg{}
	if err := cleanenv.ReadConfig("./internal/config/config.yaml", &cfg); err != nil {
		log.Println("cannot read configs")
	}
	log.Println("configs getting success")
	return cfg
}

var All = NewConfig()
