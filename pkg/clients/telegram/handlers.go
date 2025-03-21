package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const commandStart = "start"

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	_, err := b.bot.Send(msg)
	if err != nil {
		b.log.Errorf("Failed to send message: %v", err)
	}
}
func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Error: /help")

	switch message.Command() {
	case commandStart:
		msg.Text = "Text: /start"
		_, err := b.bot.Send(msg)
		return err

	default:
		_, err := b.bot.Send(msg)

		return err

	}
}
