package sheets

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"os"
)

type ctxSheets struct {
	ctx context.Context
}
type Sheets interface {
	CliSheets() *sheets.Service
}

func NewSheets(ctx context.Context) Sheets {
	return &ctxSheets{ctx: ctx}
}
func (s *ctxSheets) CliSheets() *sheets.Service {
	var credPath string
	if _, err := os.Stat("configs/service-account-key.json"); err == nil {
		credPath = "configs/service-account-key.json"
	} else {
		credPath = "/configs/service-account-key.json"
	}
	client, err := sheets.NewService(s.ctx, option.WithCredentialsFile(credPath))
	if err != nil {
		log.Fatalf("Failed to create Google Sheets client: %v", err)
	}

	return client
}
