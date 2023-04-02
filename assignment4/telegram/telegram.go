package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
)

func SendPhotoFromURL(bot *tgbotapi.BotAPI, chatID int64, photoURL string) error {
	// Fetch the photo from the URL
	resp, err := http.Get(photoURL)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	// Create a new PhotoConfig with the photo bytes
	photoConfig := tgbotapi.NewPhotoUpload(chatID, tgbotapi.FileReader{
		Name:   "photo.jpg",
		Reader: resp.Body,
		Size:   resp.ContentLength,
	})

	// Send the photo to the user
	_, err = bot.Send(photoConfig)
	return err
}

// replyToMessage sends a message to a Telegram user in response to a previous message
func replyToMessage(bot *tgbotapi.BotAPI, chatID int64, messageID int, text string) error {
	// Create a new MessageConfig with the chat ID, message ID, and text
	messageConfig := tgbotapi.NewMessage(chatID, text)
	messageConfig.ReplyToMessageID = messageID

	// Send the message to the user
	_, err := bot.Send(messageConfig)
	return err
}

func Hancom(bot *tgbotapi.BotAPI, update tgbotapi.Update, url string) {
	command := update.Message.Command()
	switch command {
	case "start":
		// Reply to the user with a welcome message
		err := replyToMessage(bot, update.Message.Chat.ID, update.Message.MessageID, "Welcome to my bot! To get started, type /image to receive a random photo.")
		if err != nil {
			log.Println(err)
		}
	case "image":
		// Get a random photo URL

		// Send the photo to the user
		err := SendPhotoFromURL(bot, update.Message.Chat.ID, url)
		if err != nil {
			log.Println(err)
		}
	default:
		// Reply to the user with a message indicating that the command is not supported
		err := replyToMessage(bot, update.Message.Chat.ID, update.Message.MessageID, "Sorry, I don't support that command.")
		if err != nil {
			log.Println(err)
		}
	}
}
