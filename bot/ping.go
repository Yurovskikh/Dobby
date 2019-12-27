package bot

import "github.com/bwmarrin/discordgo"

type Ping struct {
	Cmd         string
	Description string
}

func NewPingCmd() Cmd {
	return &Ping{
		Cmd:         "!ping",
		Description: "Dobby has no master. Dobby is a free elf!",
	}
}

func (p Ping) Do(ctx Context, sess *discordgo.Session, msg *discordgo.MessageCreate) error {
	/// when a user sends "ping" the bot replies with "pong"
	_, err := sess.ChannelMessageSend(msg.ChannelID, "pong!")
	return err
}

func (p Ping) Name() string {
	return p.Cmd
}

func (p Ping) Help() string {
	return p.Description
}
