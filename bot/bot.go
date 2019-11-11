package bot

import (
	"fmt"
	"time"

	"strings"

	"github.com/WikiWikiWasp/Dobby/commands"
	"github.com/WikiWikiWasp/Dobby/config"
	"github.com/bwmarrin/discordgo"
)

var (
	BotID string
	t0    time.Time
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
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	/// Check to make sure the message recieved did not come from another bot
	/// or from Dobby itself
	if s == nil || m == nil || m.Author.ID == BotID || m.Author.Bot {
		return
	}

	/// need to use conf.BotPrefix instead of hardcodeing "!"
	if strings.HasPrefix(m.Content, "!") && strings.Count(m.Content, "!") < 2 {

		commands.ExecuteCommand(s, m.Message, t0)
		return
	}
}
