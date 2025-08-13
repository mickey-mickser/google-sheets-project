// package sheets
//
// import (
//
//	"context"
//	"fmt"
//	"golang.org/x/oauth2/google"
//	"google.golang.org/api/option"
//	"google.golang.org/api/sheets/v4"
//	"os"
//
// )
//
//	type ctxSheets struct {
//		ctx context.Context
//	}
//
//	type Sheets interface {
//		CliSheets(serviceKeyPath string) (*sheets.Service, error)
//	}
//
//	func NewSheets(ctx context.Context) Sheets {
//		return &ctxSheets{ctx: ctx}
//	}
//
//	func (s *ctxSheets) CliSheets(serviceKeyPath string) (*sheets.Service, error) {
//		credPath := serviceKeyPath
//		if _, err := os.Stat(credPath); err != nil {
//			credPath = "." + credPath
//		}
//		client, err := sheets.NewService(s.ctx, option.WithCredentialsFile(credPath))
//		if err != nil {
//			return nil, fmt.Errorf("failed to create Google Sheets client: %w", err)
//
//		}
//
//		return client, nil
//	}
//
// // Client — оболочка над Sheets API
//
//	type Client struct {
//		srv *sheets.Service
//	}
//
// // NewSheet возвращает новый пустой клиент
//
//	func NewSheet() *Client {
//		return &Client{}
//	}
//
// // CliSheetsADC инициализирует sheets.Service через
// // Application Default Credentials (Workload Identity)
//
//	func (c *Client) CliSheetsADC(ctx context.Context) error {
//		// Запрашиваем область чтения таблиц
//		scope := sheets.SpreadsheetsReadonlyScope
//
//		// Найдём креды: в проде это metadata Cloud Run
//		creds, err := google.FindDefaultCredentials(ctx, scope)
//		if err != nil {
//			return fmt.Errorf("find default credentials: %w", err)
//		}
//
//		// Создадим сам сервис
//		srv, err := sheets.NewService(ctx, option.WithCredentials(creds))
//		if err != nil {
//			return fmt.Errorf("new sheets service: %w", err)
//		}
//
//		c.srv = srv
//		return nil
//	}
//
// // FetchRange получает диапазон из Google Sheets
//
//	func (c *Client) FetchRange(ctx context.Context, spreadsheetID, readRange string) (*sheets.ValueRange, error) {
//		return c.srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Context(ctx).Do()
//	}
package sheets

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	sheetsapi "google.golang.org/api/sheets/v4"
)

// Client — оболочка над Sheets API и Drive API
type Client struct {
	srvSheets *sheetsapi.Service
	srvDrive  *drive.Service
	log       *logrus.Logger
}

// NewSheet returns a new client with a logger
func NewSheet(log *logrus.Logger) *Client {
	return &Client{log: log}
}

// CliSheetsADC initializes Sheets and Drive services via Metadata Server
// with an OAuth2 token containing the required scope, and grants SA access to the table.
func (c *Client) CliSheetsADC(ctx context.Context, spreadsheetID string) error {
	// OAuth2 scopes для Sheets и Drive
	scopes := []string{
		sheetsapi.SpreadsheetsScope,
		drive.DriveScope,
	}

	// 1) Запрашиваем токен из Metadata Server
	tok, err := fetchToken(ctx, scopes)
	if err != nil {
		return fmt.Errorf("fetch metadata token: %w", err)
	}

	// 2) Создаём TokenSource на основе полученного токена
	ts := oauth2.StaticTokenSource(tok)

	// 3) Инициализируем Sheets клиент
	c.srvSheets, err = sheetsapi.NewService(ctx,
		option.WithTokenSource(ts),
	)
	if err != nil {
		return fmt.Errorf("new sheets service: %w", err)
	}

	// 4) Инициализируем Drive клиент
	c.srvDrive, err = drive.NewService(ctx,
		option.WithTokenSource(ts),
	)
	if err != nil {
		return fmt.Errorf("new drive service: %w", err)
	}

	// 5) Программно выдаём SA роль writer на указанный Spreadsheet
	if err := c.GrantPermission(ctx, spreadsheetID); err != nil {
		return fmt.Errorf("grant permission: %w", err)
	}

	c.log.Infof("Initialized ADC and granted access on spreadsheet %s", spreadsheetID)
	return nil
}

// GrantPermission grants the service account the writer role on the spreadsheetID file.
func (c *Client) GrantPermission(ctx context.Context, spreadsheetID string) error {
	perm := &drive.Permission{
		Type:         "user",
		Role:         "writer",
		EmailAddress: "my-cloudrun-service-account@propane-ripsaw-465110-v5.iam.gserviceaccount.com",
	}

	res, err := c.srvDrive.Permissions.Create(spreadsheetID, perm).
		SupportsAllDrives(true).
		Do()
	if err != nil {
		if gerr, ok := err.(*googleapi.Error); ok {
			c.log.Errorf("Drive API error: code=%d, body=%s", gerr.Code, gerr.Body)
		}
		return err
	}

	c.log.Infof("Granted %s access to %s on %s", perm.Role, perm.EmailAddress, spreadsheetID)
	c.log.Infof("Response: %v", res)
	return nil
}

// ServiceSheets возвращает внутренний *sheets.Service для прямого доступа
func (c *Client) ServiceSheets() *sheetsapi.Service {
	return c.srvSheets
}
func (c *Client) ServiceDrive() *drive.Service {
	return c.srvDrive
}

// FetchRange получает диапазон из Google Sheets
func (c *Client) FetchRange(ctx context.Context, spreadsheetID, readRange string) (*sheetsapi.ValueRange, error) {
	return c.srvSheets.Spreadsheets.Values.Get(spreadsheetID, readRange).
		Context(ctx).
		Do()
}

//// CliSheetsADC инициализирует sheets и drive через Application Default Credentials (ADC)
//// и выдаёт сервис‑аккаунту доступ к указанной таблице.
//func (c *Client) CliSheetsADC(ctx context.Context, spreadsheetID string) error {
//	// Запрашиваем учётные данные с нужными OAuth2 scopes
//	scopes := []string{
//		sheetsapi.SpreadsheetsScope,
//		sheetsapi.DriveScope,
//	}
//
//	ts, err := google.DefaultTokenSource(ctx, scopes...)
//	if err != nil {
//		return fmt.Errorf("DefaultTokenSource: %w", err)
//	}
//	creds, err := google.FindDefaultCredentials(ctx, scopes...)
//	if err != nil {
//		return fmt.Errorf("find default credentials: %w", err)
//	}
//
//	// Инициализируем Sheets service
//	srvSheets, err := sheetsapi.NewService(ctx,
//		option.WithCredentials(creds),
//		option.WithScopes(sheetsapi.SpreadsheetsScope),
//	)
//	if err != nil {
//		return fmt.Errorf("new sheets service: %w", err)
//	}
//	c.srvSheets = srvSheets
//
//	// Инициализируем Drive service
//	srvDrive, err := drive.NewService(ctx,
//		option.WithCredentials(creds),
//		option.WithScopes(scopes...),
//	)
//	if err != nil {
//		return fmt.Errorf("new drive service: %w", err)
//	}
//	c.srvDrive = srvDrive
//
//	// Выдаём сервис‑аккаунту роль writer на указанную таблицу
//	if err := c.GrantPermission(ctx, spreadsheetID); err != nil {
//		return fmt.Errorf("grant permission: %w", err)
//	}
//
//	c.log.Infof("ADC initialized and access granted for AS on spreadsheet %s", spreadsheetID)
//	return nil
//}

//// GrantPermission выдаёт сервис‑аккаунту (email) роль writer на файл spreadsheetID.
//func (c *Client) GrantPermission(ctx context.Context, spreadsheetID string) error {
//	perm := &drive.Permission{
//		Type:         "user",   // можно "domain"/"group"/"anyone"
//		Role:         "writer", // можно "reader"/"commenter"
//		EmailAddress: "my-cloudrun-service-account@propane-ripsaw-465110-v5.iam.gserviceaccount.com",
//	}
//
//	res, err := c.srvDrive.Permissions.Create(spreadsheetID, perm).
//		SupportsAllDrives(true).
//		Do()
//	if err != nil {
//		if gerr, ok := err.(*googleapi.Error); ok {
//			c.log.Errorf("Drive API error: code=%d, body=%s", gerr.Code, gerr.Body)
//		}
//		return err
//	}
//
//	c.log.Infof("Granted %s access to %s on %s", perm.Role, perm.EmailAddress, spreadsheetID)
//	c.log.Infof("Response: %v", res)
//
//	return nil
//}
//
//// ServiceSheets возвращает внутренний *sheets.Service для работы напрямую
//func (c *Client) ServiceSheets() *sheetsapi.Service {
//	return c.srvSheets
//}
//
//// FetchRange получает диапазон из Google Sheets
//func (c *Client) FetchRange(ctx context.Context, spreadsheetID, readRange string) (*sheetsapi.ValueRange, error) {
//	return c.srvSheets.Spreadsheets.Values.Get(spreadsheetID, readRange).
//		Context(ctx).
//		Do()
//}
