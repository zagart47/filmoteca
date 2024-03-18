package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Cfg struct {
	Logger struct {
		Mode string `yaml:"mode"`
	} `yaml:"logger"`
	HTTP struct {
		Host string `yaml:"host"`
	} `yaml:"http"`
	PostgreSQL struct {
		DSN string `yaml:"dsn" default:"localhost"`
	} `yaml:"postgresql"`
}

func NewConfig() Cfg {
	cfg := Cfg{}
	//if err := cleanenv.ReadConfig("./internal/config/config.yaml", &cfg); err != nil {
	if err := cleanenv.ReadConfig("E:\\Projects\\filmoteca\\internal\\config\\config.yaml", &cfg); err != nil {
		log.Println("cannot read configs")
		os.Exit(1)
	}
	log.Println("configs getting success")
	return cfg
}

var All = NewConfig()
