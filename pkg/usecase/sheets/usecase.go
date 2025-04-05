package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"os"
)

//go:generate mockgen -source=usecase.go -destination=mocks/mock.go

type SheetUseCase interface {
	BatchUpdate(ctx context.Context, sheetId string, jsonBody []byte) (*sheets.BatchUpdateSpreadsheetResponse, error)
	Get(ctx context.Context, sheetId, param string) ([][]interface{}, error)
	Create(ctx context.Context, jsonBody []byte) (*sheets.Spreadsheet, error)
	Delete(ctx context.Context, sheetId, param string, jsonBody []byte) (*sheets.ClearValuesResponse, error)
}
type sheetUseCase struct {
	cli *sheets.Service
	log *logrus.Logger
}

func NewSheetUse(cli *sheets.Service, log *logrus.Logger) SheetUseCase {
	return &sheetUseCase{
		cli: cli,
		log: log,
	}
}
func (s *sheetUseCase) BatchUpdate(ctx context.Context, sheetId string, jsonBody []byte) (*sheets.BatchUpdateSpreadsheetResponse, error) {
	var batchUpdateRequest sheets.BatchUpdateSpreadsheetRequest
	err := json.Unmarshal(jsonBody, &batchUpdateRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal batchUpdateRequest body: %w", err)
	}

	resp, err := s.cli.Spreadsheets.BatchUpdate(sheetId, &batchUpdateRequest).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("error while updating data in Google Sheets: %w", err)
	}
	s.logRequestTypes(batchUpdateRequest.Requests)
	return resp, nil
}

func (s *sheetUseCase) Get(ctx context.Context, sheetId, param string) ([][]interface{}, error) {
	if param == "" {
		param = "A1:Z1000"
	}

	resp, err := s.cli.Spreadsheets.Values.Get(sheetId, param).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("error while getting data from Google Sheets: %w", err)
	}
	s.log.Printf("Get request - Sheet ID: %s, Range: %s", sheetId, param)
	return resp.Values, nil
}

func (s *sheetUseCase) Delete(ctx context.Context, sheetId, param string, jsonBody []byte) (*sheets.ClearValuesResponse, error) {
	if param == "" {
		param = "A1:Z1000"
	}
	var clearValuesRequest sheets.ClearValuesRequest
	err := json.Unmarshal(jsonBody, &clearValuesRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal clearValuesRequest body: %w", err)
	}

	resp, err := s.cli.Spreadsheets.Values.Clear(sheetId, param, &clearValuesRequest).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("error clearing values in Google Sheets %s: %w", sheetId, err)
	}
	s.log.Printf("Table cleared: %s\n", sheetId)
	return resp, nil
}

func (s *sheetUseCase) Create(ctx context.Context, jsonBody []byte) (*sheets.Spreadsheet, error) {
	var createRequest sheets.Spreadsheet

	err := json.Unmarshal(jsonBody, &createRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal createRequest body: %w", err)
	}
	resp, err := s.cli.Spreadsheets.Create(&createRequest).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("error creating table in Google Sheets: %w", err)
	}
	err = grantSheetAccess(ctx, resp.SpreadsheetId)
	if err != nil {
		return nil, fmt.Errorf("failed to set permissions: %w", err)
	}
	s.log.Printf("Table created: %s\n", resp.SpreadsheetUrl)
	return resp, nil
}

func grantSheetAccess(ctx context.Context, sheetID string) error {
	var credPath string

	if _, err := os.Stat("configs/service-account-key.json"); err == nil {
		credPath = "configs/service-account-key.json"
	} else {
		credPath = "/configs/service-account-key.json"
	}

	srv, err := drive.NewService(ctx, option.WithCredentialsFile(credPath))
	if err != nil {
		return fmt.Errorf("failed to create Google Drive client: %w", err)
	}

	perm := &drive.Permission{
		Type:         "user",
		Role:         "writer",
		EmailAddress: "msunavong@gmail.com",
	}

	_, err = srv.Permissions.Create(sheetID, perm).Do()
	if err != nil {
		return fmt.Errorf("failed to set permission: %w", err)
	}
	return nil
}
