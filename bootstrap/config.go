package bootstrap

import (
	"fmt"

	"github.com/spf13/viper"
)

// NewViper is a function to load config from config.json
// You can change the implementation, for example load from env file, consul, etcd, etc
func NewViper(path string) *viper.Viper {
	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("json")
	// config.AddConfigPath(path)
	config.AddConfigPath(path)
	err := config.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config filessss: %w \n", err))
	}

	return config
}
