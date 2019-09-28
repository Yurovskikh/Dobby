package config

import (
	"os"
)

type ConfigStruct struct {
	BotToken  string
	BotPrefix string
}

// New returns a new Config struct
func New() *ConfigStruct {
	return &ConfigStruct{
		BotToken:  getEnv("DCB_TOKEN", ""),
		BotPrefix: getEnv("BOT_PREFIX", "!"),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
