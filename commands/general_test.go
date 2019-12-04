package commands

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestHandleHelpCommand(t *testing.T) {
	type args struct {
		s *discordgo.Session
		m *discordgo.Message
	}
	commands := map[string]string{
		"help": "Help Command",
		"ping": "Ping Command",
		"info": "Info Command",
	}
	tests := []struct {
		name    string
		cmd     string
		args    args
		wantErr bool
	}{
		{
			name:    "command found",
			cmd:     "help",
			wantErr: false,
		},
		{
			name:    "command not found",
			cmd:     "nope",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, ok := commands[tt.cmd]; !ok && !tt.wantErr {
				t.Log("Error: Could not find command")
				t.Fail()
			}
			com := "!" + tt.cmd
			if com
			HandleHelpCommand(tt.args.s, tt.args.m)
		})
	}
}
