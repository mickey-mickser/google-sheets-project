package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	log logrus.Ext1FieldLogger
}

func NewBot(bot *tgbotapi.BotAPI, log logrus.Ext1FieldLogger) *Bot {
	return &Bot{
		bot: bot,
		log: log,
	}
}

func (b *Bot) Start(ctx context.Context) error {

	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}

	b.handleUpdates(ctx, updates)
	return nil
}

func (b *Bot) handleUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel) {
	for {
		select {
		case update := <-updates:
			if update.Message == nil {
				continue
			}
			if update.Message.IsCommand() {
				if err := b.handleCommand(update.Message); err != nil {
					b.log.Errorf("Failed to handle command: %v", err)
				}
				continue
			}

			b.handleMessage(update.Message)
		case <-ctx.Done():
			b.log.Info("Bot stopped.")
			return
		}
	}
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return nil, err
	}
	return updates, nil
}
