package main

import (
	"context"
	"fmt"
	"log"

	bot "github.com/WikiWikiWasp/Dobby/bot"
	"github.com/google/go-github/github"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	/// starts up bot
	bot.Start()

	/// Construct a new github client to access github services
	client := github.NewClient(nil)

	/// get list of releases for the Dobby repo
	opt := &github.ListOptions{}
	repos, _, err := client.Repositories.ListReleases(
		context.Background(),
		"WikiWikiWasp",
		"Dobby",
		opt)
	if err != nil {
		panic(err)
	}

	/// print logo and the current release tag

	fmt.Printf(`
	$$$$$$$\            $$\       $$\                 
	$$  __$$\           $$ |      $$ |                
	$$ |  $$ | $$$$$$\  $$$$$$$\  $$$$$$$\  $$\   $$\ 
	$$ |  $$ |$$  __$$\ $$  __$$\ $$  __$$\ $$ |  $$ |
	$$ |  $$ |$$ /  $$ |$$ |  $$ |$$ |  $$ |$$ |  $$ |
	$$ |  $$ |$$ |  $$ |$$ |  $$ |$$ |  $$ |$$ |  $$ |
	$$$$$$$  |\$$$$$$  |$$$$$$$  |$$$$$$$  |\$$$$$$$ |
	\_______/  \______/ \_______/ \_______/  \____$$ |`)
	fmt.Printf("\n\t\t\t\t%+v", *repos[0].TagName)
	fmt.Printf(`     $$\   $$ |
	                                        \$$$$$$  |
	                                         \______/ `)

	fmt.Printf("\n\n")
	/// Create a channel that takes an empty struct that waits for input.
	/// Hacky way of making main func sit and wiat forever, not using CPU
	<-make(chan struct{})

	return
}
