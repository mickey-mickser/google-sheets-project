package config

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
)

type Config interface {
	Log() *logrus.Logger
}

type config struct {
	logger
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
