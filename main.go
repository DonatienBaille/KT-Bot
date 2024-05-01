package main

import (
	"kaki-tech/kt-bot/config"
	"kaki-tech/kt-bot/discord"
	"log"
	"os"
	"os/signal"
)

func main() {
	config.ConfigureEnvironmentVariables()
	go discord.StartBot()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Exiting application...")
	discord.StopBot()
	log.Println("Goodbye!")
}
