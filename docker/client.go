package docker

import (
	"context"
	"fmt"
	"kaki-tech/kt-bot/models"
	"log"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
	dkr "github.com/docker/docker/client"
	"github.com/samber/lo"
)

type DockerClient interface {
	// List all visible containers
	GetContainers() []*models.KtContainer

	// Get container details with his Id
	GetContainer(id string) (*models.KtContainer, error)

	// Start the container with name
	StartContainer(name string) error

	// Stop the container with name
	StopContainer(name string) error

	// Restart the container with name
	RestartContainer(name string) error

	// Update the container image
	UpdateContainer(name string) error

	// TODO : Wrap in a custom API and expose only KtContainer
	// cf. https://pkg.go.dev/github.com/docker/docker/client#Client.Events
	Events(ctx context.Context, options types.EventsOptions) (<-chan events.Message, <-chan error)
}

type KtClient struct {
	*dkr.Client
}

var Filter = filters.NewArgs(filters.KeyValuePair{
	Key:   "label",
	Value: "visibility=bot-discord",
})

func (c *KtClient) GetContainers() []*models.KtContainer {
	containers, err := c.ContainerList(context.Background(), container.ListOptions{
		All:     true,
		Filters: Filter,
	})

	if err != nil {
		log.Fatalf("Error while listing containers: %v", err)
	}

	result := lo.Map(containers, func(c types.Container, index int) *models.KtContainer {
		return &models.KtContainer{
			ID:    c.ID,
			Name:  strings.Replace(c.Names[0], "/", "", 1),
			Image: c.Image,
			State: c.State,
		}
	})

	return result
}

func (c *KtClient) StartContainer(name string) error {
	return c.ExecuteWithOneContainer(name, func(ctr *models.KtContainer) error {
		return c.ContainerStart(context.Background(), ctr.ID, container.StartOptions{})
	})
}

func (c *KtClient) StopContainer(name string) error {
	return c.ExecuteWithOneContainer(name, func(ctr *models.KtContainer) error {
		return c.ContainerStop(context.Background(), ctr.ID, container.StopOptions{})
	})
}

func (c *KtClient) RestartContainer(name string) error {
	return c.ExecuteWithOneContainer(name, func(ctr *models.KtContainer) error {
		return c.ContainerRestart(context.Background(), ctr.ID, container.StopOptions{})
	})
}

func (c *KtClient) UpdateContainer(name string) error {
	return c.ExecuteWithOneContainer(name, func(ctr *models.KtContainer) error {
		// Image must be specified without tag
		imageToUpdate := strings.Split(ctr.Image, ":")[0]
		return updateWithWatchtower(imageToUpdate)
	})
}

func (c *KtClient) GetContainer(id string) (*models.KtContainer, error) {
	// TODO : duplicate from GetContainers, need refactoring

	filter := Filter.Clone()
	filter.Add("id", id)

	containers, err := c.ContainerList(context.Background(), container.ListOptions{
		All:     true,
		Filters: filter,
	})

	if err != nil {
		log.Fatalf("Error while listing containers: %v", err)
	}

	result := lo.Map(containers, func(c types.Container, index int) *models.KtContainer {
		return &models.KtContainer{
			ID:    c.ID,
			Name:  strings.Replace(c.Names[0], "/", "", 1),
			Image: c.Image,
			State: c.State,
		}
	})

	if len(result) > 0 {
		return result[0], nil
	} else {
		return nil, fmt.Errorf("container with id: %v not found", id)
	}
}

func (c *KtClient) ExecuteWithOneContainer(name string, handle func(ctr *models.KtContainer) error) error {
	containers := c.GetContainers()

	foundContainer, found := lo.Find(containers, func(c *models.KtContainer) bool { return c.Name == name })

	if found {
		return handle(foundContainer)
	} else {
		return fmt.Errorf("unable to find container with name: %v", name)
	}
}

func GetClient() DockerClient {
	ktClient, err := dkr.NewClientWithOpts(dkr.FromEnv)

	if err != nil {
		log.Panic(err)
	}

	configureWatchtower()

	return &KtClient{Client: ktClient}
}
