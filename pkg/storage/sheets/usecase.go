package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/sheets/v4"
)

type SheetUseCase interface {
	BatchUpdate(ctx context.Context, sheetId string, jsonBody []byte) (*sheets.BatchUpdateSpreadsheetResponse, error)
}
type sheetUseCase struct {
	cli *sheets.Service
}

func NewSheetUse(cli *sheets.Service) SheetUseCase {
	return &sheetUseCase{cli: cli}
}
func (s *sheetUseCase) BatchUpdate(ctx context.Context, sheetId string, jsonBody []byte) (*sheets.BatchUpdateSpreadsheetResponse, error) {
	var batchUpdateRequest sheets.BatchUpdateSpreadsheetRequest
	err := json.Unmarshal(jsonBody, &batchUpdateRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal request body: %w", err)
	}

	resp, err := s.cli.Spreadsheets.BatchUpdate(sheetId, &batchUpdateRequest).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("error while updating data in Google Sheets: %w", err)
	}
	logRequestTypes(batchUpdateRequest.Requests)
	return resp, nil
}
