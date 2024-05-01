package config

import "github.com/spf13/viper"

const dockerHostKey = "DOCKER_HOST"

func configureDockerKeys(v *viper.Viper) {
	v.BindEnv(dockerHostKey)
}

func addDefaultDockerValues(v *viper.Viper) {
	v.SetDefault(dockerHostKey, "/var/run/docker.sock")
}
