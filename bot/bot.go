package bot

import (
	"fmt"
	"github.com/WikiWikiWasp/Dobby/config"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"strings"
)

const cmdPrefix = "!"

type Cmd interface {
	Do(ctx Context, sess *discordgo.Session, msg *discordgo.MessageCreate) error
	Name() string
	Help() string
}

type Bot struct {
	cfg    *config.ConfigStruct
	id     string
	router map[string]Cmd
	log    *logrus.Logger
}

func New(cfg *config.ConfigStruct, log *logrus.Logger) *Bot {
	return &Bot{
		cfg:    cfg,
		router: NewRouter(),
		log:    log,
	}
}

// Start starts the bot up
func (b Bot) Start() {
	fmt.Println("Starting up Dobby...")
	/// Create a new discord session
	session, err := discordgo.New(fmt.Sprintf("Bot %s", b.cfg.BotToken))
	if err != nil {
		b.log.WithError(err).Fatal("unable to creating discord session")
	}

	/// Get bot account info
	b.log.Info("Getting Dobby account info using token...")
	user, err := session.User("@me")
	if err != nil {
		b.log.WithError(err).Fatal("unable to retrieving account")
	}

	b.id = user.ID

	/// Create a Handler for the discord session
	session.AddHandler(b.messageHandler)
	session.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		/// Set Bot discord status
		if err = discord.UpdateStatus(0, "Dobby is...free"); err != nil {
			b.log.WithError(err).Error("unable to attempting to set my status")
			return
		}

		/// Get a list of all servers (guilds) bot is connected to
		servers := discord.State.Guilds
		fmt.Printf("Dobby has begun service on %d servers\n", len(servers))
	})

	/// Try to open session
	if err = session.Open(); err != nil {
		b.log.WithError(err).Fatal("unable to opening connection to Discord")
	}
}

// messageHandler
// Starting framework of message handler
func (b Bot) messageHandler(sess *discordgo.Session, msg *discordgo.MessageCreate) {
	/// need to use conf.BotPrefix instead of hardcodeing "!"
	if !strings.HasPrefix(msg.Content, cmdPrefix) {
		return
	}

	/// Check to make sure the message recieved did not come from another bot
	/// or from Dobby itself
	user := msg.Author
	if user.ID == b.id || user.Bot {
		/// Do nothing because the bot is talking
		return
	}

	ctx := NewFromContent(msg.Content)
	cmd, ok := b.router[ctx.Cmd]
	if !ok {
		// todo cmd = b.Router["!help"]
		return
	}

	if err := cmd.Do(ctx, sess, msg); err != nil {
		if err == ErrEmptyArgs {
			_, err := sess.ChannelMessageSend(msg.ChannelID, cmd.Help())
			if err != nil {
				b.log.WithField("cmd", cmd.Name()).WithError(err).Error("unable to send message")
			}
		} else {
			b.log.WithField("cmd", cmd.Name()).WithError(err).Error("unable to do cmd")
		}
	}
}
