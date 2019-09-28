package main

import (
	"log"

	bot "github.com/WikiWikiWasp/Dobby/bot"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	bot.Start()

	/// Create a channel that takes an empty struct that waits for input.
	/// Hacky way of making main func sit and wiat forever, not using CPU
	<-make(chan struct{})

	return
}
