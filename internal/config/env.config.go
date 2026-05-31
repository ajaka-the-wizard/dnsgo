package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	ADDR       string `mapstructure:"ADDR"`
	PRODUCTION bool   `mapstructure:"PRODUCTION"`
}

func LoadEnv() *Env {
	v := viper.New()
	v.SetConfigFile(".env")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read env config: %v", err)
	}

	var env Env
	if err := v.UnmarshalExact(&env); err != nil {
		log.Fatalf("Failed to map env config: %v", err)
	}
	log.Println("Env loaded successfully")
	return &env
}
