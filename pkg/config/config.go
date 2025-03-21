package config

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
)

type Config interface {
	Log() *logrus.Logger
	DB() *gorm.DB
	TelegramTokenCli() *tgbotapi.BotAPI
	GoogleSheetID() string
}

type config struct {
	db
	logger
	tgBotToken
	googleSheet
}

func NewConfig(cfgPath string) (Config, error) {
	file, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	cfg := config{}
	if err = json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}
