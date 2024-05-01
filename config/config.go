package config

import (
	"github.com/spf13/viper"
)

var v *viper.Viper

// Apply default configuration for environment variables
func ConfigureEnvironmentVariables() {
	v = viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")

	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
	}

	configureDockerKeys(v)
	configureDiscordKeys(v)
	configureWatchtowerKeys(v)

	v.AutomaticEnv()

	addDefaultDockerValues(v)
	addDefaultDiscordValues(v)
	addDefaultWatchtowerValues(v)
}

func GetVariable(name string) string {
	return v.GetString(name)
}
