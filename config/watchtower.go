package config

import "github.com/spf13/viper"

const WatchtowerApiUrlKey = "KT_WATCHTOWER_API_URL"
const WatchtowerApiTokenKey = "KT_WATCHTOWER_API_TOKEN"

func configureWatchtowerKeys(v *viper.Viper) {
	v.BindEnv(WatchtowerApiUrlKey)
	v.BindEnv(WatchtowerApiTokenKey)
}

func addDefaultWatchtowerValues(v *viper.Viper) {
	v.SetDefault(WatchtowerApiUrlKey, "http://127.0.0.1:8080")
	v.SetDefault(WatchtowerApiTokenKey, nil)
}
