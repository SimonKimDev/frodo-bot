package bot

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

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

func shouldReply(content string) bool {
	return strings.Contains(strings.ToLower(content), "secret")
}
