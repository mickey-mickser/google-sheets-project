package config

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
	"log"
	"sync"
)

type tgBotToken struct {
	tgTokenParams `json:"tg_bot_token"`
	o             sync.Once
}

type tgTokenParams struct {
	TgToken string `json:"tg_token"`
}

func (p *tgBotToken) validate() error {
	return errors.Wrap(p.check(), "failed to validate telegram token")
}

func (p *tgBotToken) check() error {
	if p.TgToken == "" {
		return errors.New("TgToken is empty")
	}

	return nil
}

func (p *tgBotToken) TelegramTokenCli() *tgbotapi.BotAPI {
	botCli, err := tgbotapi.NewBotAPI(p.TgToken)
	if err != nil {
		log.Panic(err)
	}
	return botCli
}
