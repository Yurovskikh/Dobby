package bot

import (
	"fmt"

	"strings"

	config "github.com/WikiWikiWasp/Dobby/config"
	"github.com/bwmarrin/discordgo"
)

var (
	BotID string
	goBat *discordgo.Session
)

// Start starts the bot up
func Start() {
	fmt.Println("Starting up Dobby...")
	/// Create a new discord session
	goBot, err := discordgo.New("Bot " + config.Token)
	//discord, err := discordgo.New("Bot t6-cdNMAaQnIPY39UrHl1-wUJqIDjMLg")
	errCheck("error creating discord session", err)
	fmt.Println("Getting Dobby account info...")
	/// Get bot account info
	user, err := goBot.User("@me")
	errCheck("error retrieving account", err)

	BotID = user.ID
	/// Create a Handler for the discord session
	goBot.AddHandler(messageHandler)
	goBot.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		/// Set Bot discord status
		err = discord.UpdateStatus(0, "Dobby is...free")
		errCheck("Error attempting to set my status", err)
		/// Get a list of all servers (guilds) bot is connected to
		servers := discord.State.Guilds
		fmt.Printf("Dobby has begun service on %d servers\n", len(servers))
	})

	/// Try to open session
	err = goBot.Open()
	errCheck("Error opening connection to Discord", err)
	//defer goBot.Close()
}

// errCheck
// Helpher function that allows us to check for errors and log reason
func errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v\n", msg, err)
		//panic(err)
	}
}

// messageHandler
// Starting framework of message handler
func messageHandler(sess *discordgo.Session, msg *discordgo.MessageCreate) {
	if strings.HasPrefix(msg.Content, config.BotPrefix) {
		/// Check to make sure the message recieved did not come from another bot
		/// or from Dobby itself
		user := msg.Author
		if user.ID == BotID || user.Bot {
			/// Do nothing because the bot is talking
			return
		}

		//fmt.Printf("Message: %+v || From: %s\n", msg.Content, msg.Author)
		if msg.Content == "!ping" {
			fmt.Printf("Ping received from: %s\n", msg.Author)
			/// when a user sends "ping" the bot replies with "pong"
			_, _ = sess.ChannelMessageSend(msg.ChannelID, "pong!")
		}
	}
}
