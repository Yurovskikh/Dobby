package commands

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// HandleHelpCommand returns a list of available commands and actions
func HandleHelpCommand(s *discordgo.Session, m *discordgo.Message) {
	commands := map[string]string{
		"help": "Help Command",
		"ping": "Ping Command",
		"info": "Info Command",
	}

	/// strip '!' from command
	msg := strings.Split(strings.TrimSpace(m.Content), "!")[1]
	/// split command
	helpCommand := strings.Fields(msg)

	/// handle help details
	if len(helpCommand) > 2 { /// too many arguments
		HandleUnknownCommand(s, m, m.Content)
	} else if len(helpCommand) == 2 && helpCommand[0] == "help" { /// help for a specific command
		/// look for command in available command map above
		_, found := commands[helpCommand[1]]
		if found {
			s.ChannelMessageSend(m.ChannelID, "These are the help details for the `"+helpCommand[1]+"` command.")
			fmt.Printf("Command Discord Message:\n%s", m.Content)
		} else {
			HandleUnknownCommand(s, m, m.Content)
		}
	} else if len(helpCommand) == 1 && helpCommand[0] == "help" { /// general help on commands
		/// output all the available commands and their descriptions
		/// formatting of the command list
		fmsg := fmt.Sprintf(
			"\n%s\n%s\n",
			"Dobby Commands",
			strings.Repeat("-", len("Dobby Commands")))

		/// grab commands and descriptions from map and concatenated
		for com, desc := range commands {
			fmsg = fmsg + "%-10s%-20s\n"
			fmsg = fmt.Sprintf(
				fmsg,
				com,
				desc)
		}
		fmsg = "```txt" + fmsg + "```"
		s.ChannelMessageSend(m.ChannelID, fmsg)
		// output discord message content
		fmt.Printf("Command Discord Message:\n%+v", *m)
		jdump, _ := json.MarshalIndent(m, "", "  ")
		fmt.Printf("Command Discord Message Pretty:\n%+v", string(jdump))
	}
}

// HandleInfoCommand returns the info of the current channel
func HandleInfoCommand(s *discordgo.Session, m *discordgo.Message, t0 time.Time) {
	t1 := time.Now()
	dchan, err := s.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("[ERROR] Issue finding channel by ID: ", err)
		return
	}

	dchanName := dchan.Name
	msg := "```txt\n%s\n%s\n%-16s%-20s\n%-16s%-20s\n%-16s%-20s```"
	msg = fmt.Sprintf(
		msg,
		"Dobby Information",
		strings.Repeat("-", len("Dobby Information")),
		"ChannelID",
		m.ChannelID,
		"Channel Name",
		dchanName,
		"Uptime",
		(t1.Sub(t0).String()))
	s.ChannelMessageSend(m.ChannelID, msg)
}

// HandlePingCommand returns the message "Pong" if bot is running
func HandlePingCommand(s *discordgo.Session, m *discordgo.Message) {
	fmt.Printf("Ping received from: %s\n", m.Author)
	_, _ = s.ChannelMessageSend(m.ChannelID, "pong!")
}
