package usecase

import (
	"context"
	"fmt"
	"google.golang.org/api/sheets/v4"
)

type SheetUseCase interface {
	Get(ctx context.Context, table, param string) ([][]interface{}, error)
}
type sheetUseCase struct {
	cli *sheets.Service
}

func NewSheetUse(cli *sheets.Service) SheetUseCase {
	return &sheetUseCase{cli: cli}
}
func (s *sheetUseCase) Get(ctx context.Context, table, param string) ([][]interface{}, error) {
	if param == "" {
		param = "A1:Z1000"
	}

	resp, err := s.cli.Spreadsheets.Values.Get(table, param).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("error while getting data from Google Sheets: %w", err)
	}

	return resp.Values, nil
}
