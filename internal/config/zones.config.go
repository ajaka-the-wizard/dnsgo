package config

import (
	"github.com/spf13/viper"
)

type Zones struct {
	allZones map[string][]Record `mapstructure:"zones"`
}
type Record struct {
	Name  string `mapstructure:"name"`
	Type  string `mapstructure:"type"`
	Value string `mapstructure:"value"`
	TTL   int    `mapstructure:"ttl"`
}

func LoadZones() (*Zones, error) {
	viper.SetConfigName("zones")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Zones
	if err := viper.UnmarshalExact(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
