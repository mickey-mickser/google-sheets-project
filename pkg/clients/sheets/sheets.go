package sheets

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"os"
)

type ctxSheets struct {
	ctx context.Context
}
type Sheets interface {
	CliSheets(serviceKeyPath string) (*sheets.Service, error)
}

func NewSheets(ctx context.Context) Sheets {
	return &ctxSheets{ctx: ctx}
}
func (s *ctxSheets) CliSheets(serviceKeyPath string) (*sheets.Service, error) {
	credPath := serviceKeyPath
	if _, err := os.Stat(credPath); err != nil {
		credPath = "." + credPath
	}
	client, err := sheets.NewService(s.ctx, option.WithCredentialsFile(credPath))
	if err != nil {
		return nil, fmt.Errorf("failed to create Google Sheets client: %w", err)

	}

	return client, nil
}
