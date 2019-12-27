package bot

import "strings"

type Context struct {
	Cmd  string
	Args string
}

func NewFromContent(s string) Context {
	split := strings.Split(s, " ")
	cmd := split[0]

	var args string
	if len(split) > 1 {
		args = strings.TrimSpace(strings.TrimPrefix(s, cmd))
	}

	return Context{
		Cmd:  cmd,
		Args: args,
	}
}
