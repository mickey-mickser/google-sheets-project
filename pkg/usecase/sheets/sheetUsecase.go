package sheetUsecase

import (
	"context"
	"fmt"
	"github.com/mickey-mickser/google-sheets-project/pkg/config"
	"github.com/mickey-mickser/google-sheets-project/pkg/models"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/sheets/v4"
)

// PermissionSetter is an interface for setting access rights to a table.
// All logic for setting rights is moved here so that it can be mocked up in tests.
//
//go:generate mockgen -source=sheetUsecase.go -destination=../../mocks/mock.go
type PermissionSetter interface {
	SetPermission(ctx context.Context, sheetID string) error
}

// realPermissionSetter is a real implementation of PermissionSetter,
// which sets permissions via Google API with parameters from the config.
// Everything is strictly tied to specific fields, otherwise it is not real.
type realPermissionSetter struct {
	Role         string
	Type         string
	CredPath     string
	EmailAddress string
}

// SheetUseCase — interface of the main use case for working with Google Sheets.
// Here are the main methods for working with tables — updating, reading, creating, deleting, and setting permissions.
// The SetPermissionSetter method is needed to replace the behavior of setting permissions (for tests, or if you suddenly want to change the implementation).
type SheetUseCase interface {
	BatchUpdate(ctx context.Context, sheetId string, updates []models.UpdateRequest) (*sheets.BatchUpdateValuesResponse, error)
	Get(ctx context.Context, sheetId, param string) ([][]interface{}, error)
	Create(ctx context.Context, body models.CreateRequest) (*sheets.Spreadsheet, error)
	Delete(ctx context.Context, sheetId, param string) (*sheets.ClearValuesResponse, error)
	SetPermissionSetter(ps PermissionSetter) // SetPermissionSetter allows replacing the PermissionSetter. Mainly intended for tests and mocking.
	SharePermission(ctx context.Context, body models.ShareRequest) (*drive.Permission, error)
}

// sheetUseCase is the actual implementation of SheetUseCase.
// Stores the Google Sheets service, logger, permissions configs, and the current PermissionSetter.
type sheetUseCase struct {
	cli        *sheets.Service
	driveSvc   *drive.Service
	log        *logrus.Logger
	permission *config.PermissionsStruct
	permSetter PermissionSetter
	sheetID    string
}

// NewSheetUse initializes a new SheetUseCase
func NewSheetUse(cli *sheets.Service, driveSvc *drive.Service, log *logrus.Logger, permission *config.PermissionsStruct, sheetID string) SheetUseCase {
	return &sheetUseCase{
		cli:        cli,
		driveSvc:   driveSvc,
		log:        log,
		permission: permission,
		permSetter: &realPermissionSetter{
			CredPath:     permission.ServiceKeyPath,
			EmailAddress: permission.EmailAddress,
			Role:         permission.PermissionRole,
			Type:         permission.PermissionType,
		},
		sheetID: sheetID,
	}
}

// BatchUpdate applies multiple update requests to a Google Sheet
func (s *sheetUseCase) BatchUpdate(ctx context.Context, sheetId string, updates []models.UpdateRequest) (*sheets.BatchUpdateValuesResponse, error) {
	if updates == nil {
		return nil, fmt.Errorf("batchUpdate usecase: invalid parameter(updates: %s)", updates)
	}

	valueRanges, err := s.buildValueRanges(updates)
	if err != nil {
		return nil, err
	}

	req := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "RAW",
		Data:             valueRanges,
	}
	s.logRequestTypes(updates)

	resp, err := s.cli.Spreadsheets.Values.
		BatchUpdate(sheetId, req).
		Context(ctx).
		Do()
	if err != nil {
		return nil, fmt.Errorf("error while updating values in Google Sheets: %w", err)
	}

	return resp, nil
}

// Get retrieves a range of values from a Google Sheet
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

// Delete clears a range of values in a Google Sheet
func (s *sheetUseCase) Delete(ctx context.Context, sheetId, param string) (*sheets.ClearValuesResponse, error) {
	clearReq := &sheets.ClearValuesRequest{}

	resp, err := s.cli.Spreadsheets.Values.Clear(sheetId, param, clearReq).Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("error clearing values in Google Sheets %s: %w", sheetId, err)
	}
	s.log.Printf("Table cleared: %s\n", sheetId)
	return resp, nil
}

// Create makes a new spreadsheet and grants access
func (s *sheetUseCase) Create(ctx context.Context, body models.CreateRequest) (*sheets.Spreadsheet, error) {
	if body.Properties.Title == "" {
		return nil, fmt.Errorf("create usecase: invalid parameter(properties: %s)", body.Properties.Title)
	}
	createRequest := sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: body.Properties.Title,
		},
	}

	resp, err := s.cli.Spreadsheets.Create(&createRequest).Context(ctx).Do()
	if err != nil {
		if gErr, ok := err.(*googleapi.Error); ok {
			logrus.Errorf("Sheets API error: code=%d, body=%s", gErr.Code, gErr.Body)
		}
		return nil, fmt.Errorf("create usecase: error creating table in Google Sheets: %w", err)
	}
	err = s.grantSheetAccess(ctx, resp.SpreadsheetId)
	if err != nil {
		return nil, fmt.Errorf("failed to set permissions: %w", err)
	}
	s.log.Printf("Table created: %s\n", resp.SpreadsheetUrl)
	return resp, nil
}

// SharePermission grants reader access to a sheet for the given email.
func (s *sheetUseCase) SharePermission(ctx context.Context, body models.ShareRequest) (*drive.Permission, error) {
	if body.Email == "" {
		return nil, fmt.Errorf("email is required")
	}

	perm := &drive.Permission{
		Type:         "user",
		Role:         "reader",
		EmailAddress: body.Email,
	}

	call := s.driveSvc.Permissions.
		Create(s.sheetID, perm).
		SupportsAllDrives(true).
		Fields("id").
		Context(ctx)

	if body.SendEmail {
		call = call.SendNotificationEmail(true)
		if body.EmailMessage != "" {
			call = call.EmailMessage(body.EmailMessage)
		}
	} else {
		call = call.SendNotificationEmail(false)
	}

	p, err := call.Do()
	if err != nil {
		if gerr, ok := err.(*googleapi.Error); ok && gerr.Code == 409 {
			return &drive.Permission{Id: ""}, nil
		}
		return nil, err
	}
	return p, nil
}
