package config

import (
	"log"

	"github.com/spf13/viper"
)

const DiscordApiToken = "KT_BOT_DISCORD_TOKEN"
const DiscordGuildId = "KT_GUILD_ID"
const DiscordChannelId = "KT_CHANNEL_ID"

func configureDiscordKeys(v *viper.Viper) {
	v.BindEnv(DiscordApiToken)
	v.BindEnv(DiscordGuildId)
	v.BindEnv(DiscordChannelId)
}

func addDefaultDiscordValues(v *viper.Viper) {
	if !v.IsSet(DiscordApiToken) {
		log.Fatalf("The Discord Api Token must be define via environment variable: %v", DiscordApiToken)
	}
	if !v.IsSet(DiscordGuildId) {
		log.Fatalf("The guild id must be define via environment variable: %v", DiscordGuildId)
	}
	if !v.IsSet(DiscordChannelId) {
		log.Fatalf("The channel id Token must be define via environment variable: %v", DiscordChannelId)
	}
}
