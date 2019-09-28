package main

import (
	bot "github.com/WikiWikiWasp/Dobby/bot"
	config "github.com/WikiWikiWasp/Dobby/config"
)

func main() {

	err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	bot.Start()

	/// Create a channel that takes an empty struct that waits for input.
	/// Hacky way of making main func sit and wiat forever, not using CPU
	<-make(chan struct{})

	return
}
