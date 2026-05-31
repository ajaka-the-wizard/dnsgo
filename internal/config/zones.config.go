package config

import (
	"log"

	"github.com/spf13/viper"
)

type Zones struct {
	AllZones map[string][]Record `mapstructure:"zones"`
}
type Record struct {
	Name  string `mapstructure:"name"`
	Type  string `mapstructure:"type"`
	Value string `mapstructure:"value"`
	TTL   int    `mapstructure:"ttl"`
}

func LoadZones() *Zones {
	v := viper.NewWithOptions(viper.KeyDelimiter("::"))
	v.SetConfigName("zones")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read zones : %v", err)
	}

	var zones Zones
	if err := v.UnmarshalExact(&zones); err != nil {
		log.Fatalf("Failed to map zones : %v", err)
	}
	log.Println("Zones loaded successfully")
	return &zones
}
