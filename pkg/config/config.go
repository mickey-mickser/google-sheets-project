package config

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
)

type Config interface {
	Log() *logrus.Logger
	Permissions() *PermissionsStruct
}

type config struct {
	logger
	permissions
}

func NewConfig(cfgPath string) (Config, error) {
	file, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			logrus.Error(err)
		}
	}()
	cfg := config{}
	if err = json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}
