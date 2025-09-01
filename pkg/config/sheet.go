package config

import (
	"github.com/pkg/errors"
)

type SheetStruct struct {
	SheetID string `json:"sheet_id"`
}
type sheet struct {
	sheetParams `json:"sheet"`
}
type sheetParams struct {
	SheetID string `json:"sheet_id"`
}

func (s *sheet) check() error {
	if s.SheetID == "" {
		return errors.New("SheetID is empty")
	}

	return nil
}
func (s *sheet) validate() error {
	return errors.Wrap(s.check(), "failed to validate sheet")
}

func (s *sheet) sheetID() string {

	return s.sheetParams.SheetID
}

func (s *sheet) Sheet() *SheetStruct {

	return &SheetStruct{
		SheetID: s.sheetID(),
	}
}
