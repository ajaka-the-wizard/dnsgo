package config

import (
	"log"

	"github.com/spf13/viper"
)

type Server struct {
	Operator string `mapstructure:"operator"`
	TTL      int    `mapstructure:"ttl"`
	IPv4     string `mapstructure:"ipv4"`
	IPv6     string `mapstructure:"ipv6"`
}

type Roots struct {
	RootServers map[string]Server `mapstructure:"root_servers"`
}

func LoadRoots() *Roots {
	v := viper.New()
	v.SetConfigName("roots")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read root servers : %v", err)
	}

	var roots Roots
	if err := v.UnmarshalExact(&roots); err != nil {
		log.Fatalf("Failed to map root servers : %v", err)
	}
	log.Println("Roots loaded successfully")
	return &roots
}
