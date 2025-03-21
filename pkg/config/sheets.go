package config

import (
	"github.com/pkg/errors"
	"sync"
)

type googleSheet struct {
	sheet `json:"google_sheets"`
	o     sync.Once
}

type sheet struct {
	SheetID string `json:"spread_sheet_id"`
}

func (p *googleSheet) validate() error {
	return errors.Wrap(p.check(), "failed to validate google sheet")
}

func (p *googleSheet) check() error {
	if p.SheetID == "" {
		return errors.New("sheet id is empty")
	}

	return nil
}

func (p *googleSheet) GoogleSheetID() string {
	return p.SheetID
}
