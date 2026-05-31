package config

import "github.com/spf13/viper"

type Env struct {
	ADDR string `mapstructure:"ADDR"`
}

func LoadEnv() (*Env, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var env Env
	if err := viper.UnmarshalExact(&env); err != nil {
		return nil, err
	}

	return &env, nil
}
