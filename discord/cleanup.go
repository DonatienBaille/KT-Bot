package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/samber/lo"
)

func cleanupMessages() {
	log.Println("Cleaning up channel messages...")

	messages, err := botSession.ChannelMessages(channelId, 100, "", "", "")
	handleErr(err, "cleaning up channel messages")
	messagesIds := lo.Map(messages, func(message *discordgo.Message, i int) string { return message.ID })
	botSession.ChannelMessagesBulkDelete(channelId, messagesIds)

	log.Println("Cleaning up channel messages finished.")
}
