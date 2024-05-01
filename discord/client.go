package discord

import (
	"kaki-tech/kt-bot/config"
	"kaki-tech/kt-bot/docker"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var botSession *discordgo.Session
var channelId string
var dockerClient docker.DockerClient

func StartBot() {
	var err error

	apiToken := config.GetVariable(config.DiscordApiToken)
	channelId = config.GetVariable(config.DiscordChannelId)

	if !strings.HasPrefix(apiToken, "Bot ") {
		log.Fatalf("Discord Bot Token must be prefixed by 'Bot '. Value read: %v", apiToken)
	}

	dockerClient = docker.GetClient()

	botSession, err = discordgo.New(apiToken)

	if err != nil {
		log.Panicf("Unable to connect to Discord %v", err)
	}

	botSession.AddHandler(onSessionReady)
	botSession.AddHandler(onUserInteraction)

	err = botSession.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
}

func StopBot() {
	cancelFunc()
	cleanupMessages()
	botSession.Close()
	log.Println("Bot stopped gracefully")
}

func handleErr(err error, context string) {
	if err != nil {
		log.Fatalf("Error occured while %v: %v", context, err)
	}
}
