package main

import (
	"context"
	"github.com/mickey-mickser/google-sheets-project/pkg/clients/sheets"
	"github.com/mickey-mickser/google-sheets-project/pkg/config"
	"github.com/mickey-mickser/google-sheets-project/pkg/http"
	"github.com/mickey-mickser/google-sheets-project/pkg/http/handler"
	"github.com/mickey-mickser/google-sheets-project/pkg/usecase/sheets"
	_ "google.golang.org/api/sheets/v4"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func ctxWithSig() (context.Context, func()) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	go func() {
		select {
		case <-ch:
			cancel()
		}

	}()

	return ctx, cancel
}
func main() {

	var configPath string
	if _, err := os.Stat("/configs/config.json"); err == nil {
		configPath = "/configs/config.json"
	} else {
		configPath = "./configs/config.json"
	}

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		panic(err)
	}

	log := cfg.Log()

	ctx, cancel := ctxWithSig()
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
			cancel()
		}
	}()

	cli := sheets.NewSheets(ctx).CliSheets()

	sheetUse := usecase.NewSheetUse(cli, log)

	wg := new(sync.WaitGroup)

	handlers := handler.NewHandler(log, sheetUse)
	http.NewHttp(log, handlers.InitRoutes()).Run(ctx, wg)

	wg.Wait()
}
