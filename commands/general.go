package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// HandleInfoCommand returns a list of available commands and actions
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
		(t1.Sub(t0).String())
	s.ChannelMessageSend(m.ChannelID, msg)
}

// HandlePingCommand returns the message "Pong" if bot is running
func HandlePingCommand(s *discordgo.Session, m *discordgo.Message) {
	fmt.Printf("Ping received from: %s\n", m.Author)
	_, _ = s.ChannelMessageSend(m.ChannelID, "pong!")
}
