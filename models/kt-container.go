package models

type KtContainer struct {
	// ID of the container
	ID string
	// Name of the container without "/"
	Name string
	// Name of the image of the container. Containing the tag
	Image string
	// State: created|restarting|running|removing|paused|exited|dead
	State string
}
