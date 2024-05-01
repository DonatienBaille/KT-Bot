package helpers

import (
	"fmt"
	"kaki-tech/kt-bot/models"

	"github.com/bwmarrin/discordgo"
)

func GetComponentsForContainer(c *models.KtContainer) *discordgo.ActionsRow {
	return &discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			discordgo.Button{
				Label:    "Start",
				CustomID: fmt.Sprintf("start_%v", c.Name),
				Style:    discordgo.SuccessButton,
				Disabled: c.State == "running" || c.State == "restarting",
			},
			discordgo.Button{
				Label:    "Stop",
				CustomID: fmt.Sprintf("stop_%v", c.Name),
				Style:    discordgo.DangerButton,
				Disabled: c.State != "running",
			},
			discordgo.Button{
				Label:    "Restart",
				CustomID: fmt.Sprintf("restart_%v", c.Name),
				Style:    discordgo.PrimaryButton,
				Disabled: c.State == "restarting",
			},
			discordgo.Button{
				Label:    "Update",
				CustomID: fmt.Sprintf("update_%v", c.Name),
				Style:    discordgo.SecondaryButton,
				// WatchTower, by default, doesn't update stopped container
				// so we disable the button if the container is not running
				Disabled: c.State != "running",
			},
		},
	}
}
