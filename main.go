package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	commandPrefix string
	botID         string
)

func main() {
	/// Create a new discord session
	discord, err := discordgo.New("Bot NjI2Mjk4OTg2MjIzMTA4MDk3.XYsS3A.Lgx9vaYYGzOck-Nx0b8wwd2YTkk")
	errCheck("error creating discord session", err)
	/// Get bot account info
	user, err := discord.User("@me")
	errCheck("error retrieving account", err)

	botID = user.ID
	/// Create a Handler for the discord session
	discord.AddHandler(commandHandler)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		/// Set Bot discord status
		err = discord.UpdateStatus(0, "Dobby has no master, Dobby is a free bot, and Dobby has come to save Discord.")
		errCheck("Error attempting to set my status", err)
		/// Get a list of all servers (guilds) bot is connected to
		servers := discord.State.Guilds
		fmt.Printf("Dobby has begun service on %d servers", len(servers))
	})

	/// Try to open session
	err = discord.Open()
	errCheck("Error opening connection to Discord", err)
	defer discord.Close()

	/// Set up prefix that goes before bot commands
	commandPrefix = "!"

	/// Create a channel that takes an empty struct that waits for input.
	/// Hacky way of making main func sit and wiat forever, not using CPU
	<-make(chan struct{})
}

// errCheck
// Helpher function that allows us to check for errors and log reason
func errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
		panic(err)
	}
}

// commandHandler
// Starting framework of command handler
func commandHandler(discord *discordgo.Session, msg *discordgo.MessageCreate) {
	/// Check to make sure the message recieved did not come from another bot
	/// or from Dobby itself
	user := msg.Author
	if user.ID == botID || user.Bot {
		/// Do nothing because the bot is talking
		return
	}

	//content := msg.Content

	fmt.Printf("Message: %+v || From: %s\n", msg.Message, msg.Author)
}
