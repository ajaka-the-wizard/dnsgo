package config

import "github.com/spf13/viper"

type Env struct {
	ADDR       string `mapstructure:"ADDR"`
	PRODUCTION bool   `mapstructure:"PRODUCTION"`
}

func LoadEnv() (*Env, error) {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var env Env
	if err := viper.UnmarshalExact(&env); err != nil {
		return nil, err
	}

	return &env, nil
}
