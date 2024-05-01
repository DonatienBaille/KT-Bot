package discord

import (
	"context"
	"kaki-tech/kt-bot/docker"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
)

var cancellableCtx, cancelFunc = context.WithCancel(context.Background())

func listenToDockerEvents() {
	filter := docker.Filter.Clone()
	filter.Add("type", "container")

	msgs, err := dockerClient.Events(cancellableCtx, types.EventsOptions{
		Filters: filter,
	})

	for {
		select {
		case err := <-err:
			log.Panicf("Error while reading events from Docker daemon: %v", err)
		case msg := <-msgs:
			go log.Printf("Event: %v", msg)
			updateContainerStatus(msg.Actor.ID, msg.Action)
		}
	}
}

func updateContainerStatus(id string, action events.Action) {
	container := dockerClient.GetContainer(id)

	// The state of the container is not accurate at this point
	// so we take from the event and we pray that it's good?
	container.State = getStateFromAction(action)

	msg := getMessageForContainer(container)

	if existingMsg, ok := containerMessages[container.Name]; ok {
		edit := discordgo.NewMessageEdit(existingMsg.ChannelID, existingMsg.ID)
		edit.Components = &msg.Components
		edit.Embeds = &msg.Embeds
		botSession.ChannelMessageEditComplex(edit)
	} else {
		botSession.ChannelMessageSendComplex(channelId, msg)
	}
}

func getStateFromAction(action events.Action) string {
	switch action {
	case events.ActionStart:
		return "running"
	// The restart event is propagate after the start so the state is "running"
	case events.ActionRestart:
		return "running"
	case events.ActionCreate:
		return "created"
	case events.ActionRemove:
		return "removing"
	case events.ActionPause:
		return "paused"
	case events.ActionStop:
		return "exited"
	case events.ActionDie:
		return "dead"
	case events.ActionKill:
		return "dead"
	default:
		log.Fatalf("Unknown event action: %v", action)
		return ""
	}
}
