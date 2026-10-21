package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SimonKimDev/frodo-bot/internal/bot"
	"github.com/SimonKimDev/frodo-bot/internal/config"
	"github.com/bwmarrin/discordgo"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("exiting with error: %v", err)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("config failed: %w", err)
	}

	discord, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		return fmt.Errorf("bot failed to init: %w", err)
	}

	b := bot.New(cfg)
	b.Register(discord)

	// open session
	err = discord.Open()
	if err != nil {
		return fmt.Errorf("failed to open discord session: %w", err)
	}

	// executes after run returns
	defer discord.Close()

	fmt.Println("bot is running")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	<-c

	return nil
}
