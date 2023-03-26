package config

import (
	"github.com/spf13/viper"
	"github.com/surajkumarsinha/kafka-go-poc/providers"
)

type Config map[string]any

func App() map[string]any {
	return Config{
		"environment": viper.GetString("APP_ENV"),
		"providers": []func(){
			providers.SystemProviders,
			providers.EndpointProviders,
		},
	}
}
