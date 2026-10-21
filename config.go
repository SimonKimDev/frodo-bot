package main

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token  string
	GIFURL string
}

func loadConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf(".env file is missing, using environment variable")
	}

	botToken := os.Getenv("DISCORD_TOKEN")
	if len(botToken) == 0 {
		return Config{}, errors.New("bot token is empty")
	}

	giflink := os.Getenv("KEEP_YOUR_SECRETS")
	if len(giflink) == 0 {
		return Config{}, errors.New("GIFURL is empty")
	}

	return Config{
		Token:  botToken,
		GIFURL: giflink,
	}, nil
}
