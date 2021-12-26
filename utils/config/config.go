package config

import (
	"log"

	"github.com/spf13/viper"
)

// Parse receives the filename and parses the configuration into out struct
// out must be a pointer to the output struct
func Parse(file string, out interface{}) error {
	viper.SetConfigName(file)
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	if err := viper.Unmarshal(out); err != nil {
		log.Fatalf("unable to decode into struct: %v", err)
	}

	return nil
}
