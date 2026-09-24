package sheets

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	sheetsapi "google.golang.org/api/sheets/v4"
)

type Client struct {
	srvSheets *sheetsapi.Service
	srvDrive  *drive.Service
	log       *logrus.Logger
}

// NewSheet returns a new client with a logger
func NewClients(log *logrus.Logger) *Client {
	return &Client{log: log}
}

// CliSheetsADC initializes Sheets and Drive services via Metadata Server
// with an OAuth2 token containing the required scope, and grants SA access to the table.
func (c *Client) CliSheetsADC(ctx context.Context) error {
	creds, err := google.FindDefaultCredentials(ctx,
		"https://www.googleapis.com/auth/cloud-platform",

		sheetsapi.SpreadsheetsScope,
		drive.DriveScope,
	)
	if err != nil {
		return fmt.Errorf("find default credentials: %w", err)
	}

	s, err := sheetsapi.NewService(ctx, option.WithCredentials(creds))
	if err != nil {
		return fmt.Errorf("new sheets service: %w", err)
	}
	d, err := drive.NewService(ctx, option.WithCredentials(creds))
	if err != nil {
		return fmt.Errorf("new drive service: %w", err)
	}
	tok, _ := creds.TokenSource.Token()
	about, _ := d.About.Get().Fields("user(emailAddress,displayName)").Do()
	logrus.Infof("token exp=%s type=%s", tok.Expiry, tok.TokenType)
	logrus.Infof("drive user=%s (%s)", about.User.EmailAddress, about.User.DisplayName)
	c.srvSheets = s
	c.srvDrive = d
	c.log.Infof("Initialized Sheets/Drive via ADC")
	return nil
}

// GrantPermission grants the service account the writer role on the spreadsheetID file.
func (c *Client) GrantPermission(ctx context.Context, spreadsheetID string) error {
	perm := &drive.Permission{
		Type:         "user",
		Role:         "writer",
		EmailAddress: "msunavong@gmail.com",
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

// ServiceSheets returns the internal *sheets.Service for direct access
func (c *Client) ServiceSheets() *sheetsapi.Service {
	return c.srvSheets
}
func (c *Client) ServiceDrive() *drive.Service {
	return c.srvDrive
}

// FetchRange gets a range from Google Sheets
func (c *Client) FetchRange(ctx context.Context, spreadsheetID, readRange string) (*sheetsapi.ValueRange, error) {
	return c.srvSheets.Spreadsheets.Values.Get(spreadsheetID, readRange).
		Context(ctx).
		Do()
}
