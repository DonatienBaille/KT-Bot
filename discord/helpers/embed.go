package helpers

import (
	"kaki-tech/kt-bot/discord/colors"
	"kaki-tech/kt-bot/models"

	"github.com/bwmarrin/discordgo"
)

func GetEmbedForContainer(c *models.KtContainer) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: c.Name,
		Color: func() int {
			switch c.State {
			case "created":
				return colors.Purple
			case "restarting":
				return colors.Blue
			case "running":
				return colors.Green
			case "pause":
				return colors.Gold
			default:
				return colors.Red
			}
		}(),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "State",
				Value: c.State,
			},
			{
				Name:  "Image",
				Value: c.Image,
			},
		},
	}
}
