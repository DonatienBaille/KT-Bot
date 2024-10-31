package discord

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// When the user interact with the button
func onUserInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionMessageComponent:
		customId := i.MessageComponentData().CustomID
		containerName := strings.SplitAfterN(customId, "_", 2)[1]

		var err error
		switch {
		case strings.HasPrefix(customId, "start"):
			notifyInteractionState(s, i, fmt.Sprintf("Starting server %v...", containerName))
			err = dockerClient.StartContainer(containerName)
		case strings.HasPrefix(customId, "stop"):
			notifyInteractionState(s, i, fmt.Sprintf("Stopping server %v...", containerName))
			err = dockerClient.StopContainer(containerName)
		case strings.HasPrefix(customId, "restart"):
			notifyInteractionState(s, i, fmt.Sprintf("Restarting server %v...", containerName))
			err = dockerClient.RestartContainer(containerName)
		case strings.HasPrefix(customId, "update"):
			notifyInteractionState(s, i, fmt.Sprintf("Updating server %v...", containerName))
			err = dockerClient.UpdateContainer(containerName)
		case customId == "reload_state":
			notifyInteractionState(s, i, "Reloading state of servers...")
			cleanupMessages()
			addInitialState()
			err = nil
		default:
			err = fmt.Errorf("unknown interaction: %v", customId)
		}

		if err != nil {
			log.Printf("Error reacting to interaction %v: %v", customId, err)

			msg, respErr := s.ChannelMessageSendComplex(channelId, &discordgo.MessageSend{
				Content: "Error handling interaction. Please contact administrator.",
				Flags:   discordgo.MessageFlagsEphemeral,
			})

			if respErr != nil {
				log.Printf("Unable to send response to interaction: %v", respErr)
			} else {
				go func() {
					time.Sleep(5 * time.Second)
					s.ChannelMessageDelete(msg.ChannelID, msg.ID)
				}()
			}
		}
	}
}

func notifyInteractionState(s *discordgo.Session, i *discordgo.InteractionCreate, message string) {
	log.Printf("Handling interaction: %v", message)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})
}
