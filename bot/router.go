package bot

func NewRouter() map[string]Cmd {
	playing := NewPlayingCmd()
	ping := NewPingCmd()
	return map[string]Cmd{
		ping.Name():    ping,
		playing.Name(): playing,
	}
}
