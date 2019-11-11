package commands

import (
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// ExecuteCommand parses and executes the command from the calling code
func ExecuteCommand(s *discordgo.Session, m *discordgo.Message, t0 time.Time) {
	/// strip prefix from command
	msg := strings.Split(strings.TrimSpace(m.Content), "!")[1]
	if len(msg) > 2 {
		msg = strings.Split(strings.Split(m.Content, " ")[0], "!")[1]
	}

	/// handle commands
	switch msg {
	//case "help":
	//HandleHelpCommand(s, m, t0)
	case "info":
		HandleInfoCommand(s, m, t0)
	case "ping":
		HandlePingCommand(s, m)
	default:
		HandleUnknownCommand(s, m, msg)
	}
}

// HandleUnknownCommand handles commands not recognized in ExecuteCommand
func HandleUnknownCommand(s *discordgo.Session, m *discordgo.Message, msg string) {
	c, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		println("Unable to open User Channel: ", err)
		return
	}
	s.ChannelMessageSend(c.ID, "The command `"+msg+"` is not recognized.")
}
