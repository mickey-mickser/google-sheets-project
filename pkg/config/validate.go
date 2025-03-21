package config

type validate func() error

func (cfg *config) validate() error {
	validations := []validate{
		cfg.db.validate,
		cfg.logger.validate,
		cfg.tgBotToken.validate,
		cfg.googleSheet.validate,
	}

	for _, validate := range validations {
		if err := validate(); err != nil {
			return err
		}
	}

	return nil
}
