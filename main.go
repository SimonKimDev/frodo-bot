package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var KeepYourSecrets string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load env file %s", err)
	}

	botToken := os.Getenv("DISCORD_TOKEN")
	if len(botToken) == 0 {
		log.Fatal("bot token is empty")
	}
	KeepYourSecrets = os.Getenv("KEEP_YOUR_SECRETS")
	if len(KeepYourSecrets) == 0 {
		log.Fatal("gif link is empty")
	}

	discord, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatalf("failed to init bot with token %s", err)
	}

	discord.AddHandler(Ready)
	discord.AddHandler(MessageWatcher)

	discord.Identify.Intents = discordgo.IntentsMessageContent | discordgo.IntentsGuildMessages

	// open session
	err = discord.Open()
	if err != nil {
		log.Fatalf("failed to open discord session, %s", err)
	}

	// this runs after main
	defer discord.Close()

	fmt.Println("Bot is running")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func Ready(session *discordgo.Session, ready *discordgo.Ready) {
	user := ready.User
	log.Printf("Logged in as %s", user.Username)
}

func MessageWatcher(session *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == session.State.User.ID {
		return
	}
	if m.Author.Bot == true {
		return
	}

	userMessage := strings.ToLower(m.Content)
	log.Printf("%s was sent by %s", userMessage, m.Author.Username)

	if strings.Contains(userMessage, "secret") {
		_, err := session.ChannelMessageSendReply(m.ChannelID, KeepYourSecrets, m.Reference())
		if err != nil {
			log.Fatalf("Failed to send messages with %s", err)
		}
	}
}
