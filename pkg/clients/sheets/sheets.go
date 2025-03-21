package sheets

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
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
	client, err := sheets.NewService(s.ctx, option.WithCredentialsFile("configs/service-account-key.json"))
	if err != nil {
		log.Fatalf("Failed to create Google Sheets client: %v", err)
	}

	return client
}
