package bot

import (
	"log"

	"github.com/SimonKimDev/frodo-bot/internal/config"
	"github.com/SimonKimDev/frodo-bot/internal/wound"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	gifURL string
	wounds *wound.Tracker
}

func New(cfg config.Config) *Bot {
	return &Bot{
		wounds: wound.New(),
		gifURL: cfg.GIFURL,
	}
}

// Register adds the bot's event handlers and intents to the current Session.
// Should be called before calling session.Open
func (b *Bot) Register(s *discordgo.Session) {
	s.AddHandler(b.onReady)
	s.AddHandler(b.onMessage)
	s.AddHandler(b.onInteraction)
	s.Identify.Intents = discordgo.IntentsMessageContent | discordgo.IntentsGuildMessages
}

func (b *Bot) onReady(s *discordgo.Session, r *discordgo.Ready) {
	user := r.User
	log.Printf("logged in as %v", user.Username)
	if _, err := s.ApplicationCommandCreate(user.ID, "", morgulBladeCommand()); err != nil {
		log.Printf("failed to create command %v", err)
	}
}
