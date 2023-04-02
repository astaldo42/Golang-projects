package main

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	tg "normal-bot/telegram"
	"normal-bot/unsplash"
)

func main() {
	// Create a new bot with your API token
	bot, err := tgbotapi.NewBotAPI(access2)
	if err != nil {
		log.Fatal(err)
	}

	// Set up updates channel
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}
	//take the photo

	//photo = unsplash.CutAfterMax(photo)
	// Process incoming updates
	for update := range updates {
		if update.Message == nil {
			continue
		}
		photo, err := unsplash.RandomPhoto()
		if err != nil {
			log.Fatal(err)
		}
		if update.Message.IsCommand() {
			// Handle the command
			tg.Hancom(bot, update, photo)
		}
	}
}
