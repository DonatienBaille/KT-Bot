package discord

import (
	"kaki-tech/kt-bot/discord/helpers"
	"kaki-tech/kt-bot/models"
	"log"

	"github.com/bwmarrin/discordgo"
)

var containerMessages = make(map[string]*discordgo.Message)

func addInitialState() {
	containers := dockerClient.GetContainers()

	for _, c := range containers {
		msg, err := botSession.ChannelMessageSendComplex(channelId, getMessageForContainer(c))

		if err != nil {
			log.Fatalf("Unable to create initial messages: %v", err)
		}

		containerMessages[c.Name] = msg
	}

	botSession.ChannelMessageSendComplex(channelId, &discordgo.MessageSend{
		Components: []discordgo.MessageComponent{
			&discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "Reload states",
						CustomID: "reload_state",
						Style:    discordgo.SuccessButton,
					},
				},
			},
		},
	})
}

func getMessageForContainer(container *models.KtContainer) *discordgo.MessageSend {
	return &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{helpers.GetEmbedForContainer(container)},
		Components: []discordgo.MessageComponent{
			helpers.GetComponentsForContainer(container),
		},
	}
}
