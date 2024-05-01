package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// When the bot is ready
func onSessionReady(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Bot is up and running for version %v!", r.Version)
	cleanupMessages()
	addInitialState()
	go listenToDockerEvents()
}
