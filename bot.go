package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	gifURL string
}

func newBot(cfg Config) *Bot {
	return &Bot{gifURL: cfg.GIFURL}
}

func (b *Bot) onReady(s *discordgo.Session, r *discordgo.Ready) {
	user := r.User
	log.Printf("logged in as %v", user.Username)
}

func (b *Bot) onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if !shouldReply(m.Content) {
		return
	}

	if _, err := s.ChannelMessageSendReply(m.ChannelID, b.gifURL, m.Reference()); err != nil {
		log.Printf("failed to send messages with %v", err)
	}
}
