package bot

import (
	"bytes"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Playing struct {
	Cmd         string
	Description string
}

func NewPlayingCmd() Cmd {
	return &Playing{
		Cmd:         "!playing",
		Description: "<game_name> - show users who play the specified game",
	}
}

func (p Playing) Do(ctx Context, sess *discordgo.Session, msg *discordgo.MessageCreate) error {
	if ctx.Args == "" {
		return ErrEmptyArgs
	}

	var respMsg bytes.Buffer
	// Find among the users who are present in the game
	for _, presence := range sess.State.Guilds[0].Presences {
		if presence.Game != nil {
			if strings.EqualFold(presence.Game.Name, ctx.Args) {
				// Preload user info
				user, err := sess.User(presence.User.ID)
				if err != nil {
					return err
				}

				respMsg.WriteString(user.Username)
				respMsg.WriteString("\n")
			}
		}
	}
	if len(respMsg.Bytes()) == 0 {
		respMsg.WriteString(fmt.Sprintf("No one is currently playing in %s", ctx.Args))
	}

	_, err := sess.ChannelMessageSend(msg.ChannelID, respMsg.String())
	if err != nil {
		return err
	}

	return nil
}

func (p Playing) Name() string {
	return p.Cmd
}

func (p Playing) Help() string {
	return p.Description
}
