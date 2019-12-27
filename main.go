package main

import (
	"github.com/WikiWikiWasp/Dobby/config"
	"github.com/sirupsen/logrus"
	"log"

	bot "github.com/WikiWikiWasp/Dobby/bot"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.WarnLevel)

	cfg := config.New()
	// BotToken, exists := os.LookupEnv("DCB_TOKEN")
	if cfg.BotToken == "" {
		logger.Fatal("Error: Dobby Token not found...")
		return
	}

	bot.New(cfg, logger).Start()

	/// Create a channel that takes an empty struct that waits for input.
	/// Hacky way of making main func sit and wiat forever, not using CPU
	<-make(chan struct{})

	return
}
