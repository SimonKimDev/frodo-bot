package bot

import (
	"log"

	"github.com/SimonKimDev/frodo-bot/internal/config"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	gifURL string
}

func New(cfg config.Config) *Bot {
	return &Bot{gifURL: cfg.GIFURL}
}

func (b *Bot) Register(s *discordgo.Session) {
	s.AddHandler(b.onReady)
	s.AddHandler(b.onMessage)
	s.Identify.Intents = discordgo.IntentsMessageContent | discordgo.IntentsGuildMessages
}

func (b *Bot) onReady(s *discordgo.Session, r *discordgo.Ready) {
	user := r.User
	log.Printf("logged in as %v", user.Username)
}
