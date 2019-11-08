package bot

import (
	"fmt"

	"strings"

	"github.com/WikiWikiWasp/Dobby/config"
	"github.com/bwmarrin/discordgo"
)

var (
	BotID string
)

// Start starts the bot up
func Start() {
	fmt.Println("Starting up Dobby...")
	conf := config.New()
	// BotToken, exists := os.LookupEnv("DCB_TOKEN")
	if conf.BotToken == "" {
		fmt.Printf("Error: Dobby Token not found...\n")
		return
	}

	/// Create a new discord session
	goBot, err := discordgo.New("Bot " + conf.BotToken)
	errCheck("error creating discord session", err)

	/// Get bot account info
	fmt.Printf("Getting Dobby account info using token...\n")
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
		fmt.Printf("\nDobby has begun service on %d servers\n", len(servers))
	})

	/// Try to open session
	err = goBot.Open()
	errCheck("Error opening connection to Discord", err)
}

// errCheck
// Helpher function that allows us to check for errors and log reason
func errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v\n", msg, err)
	}
}

// messageHandler
// Starting framework of message handler
func messageHandler(sess *discordgo.Session, msg *discordgo.MessageCreate) {
	/// need to use conf.BotPrefix instead of hardcodeing "!"
	if strings.HasPrefix(msg.Content, "!") {
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
