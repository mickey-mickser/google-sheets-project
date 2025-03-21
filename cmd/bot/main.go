package main

import (
	"context"
	"github.com/mickey-mickser/telegram-project/pkg/clients/sheets"
	"github.com/mickey-mickser/telegram-project/pkg/clients/telegram"
	"github.com/mickey-mickser/telegram-project/pkg/config"
	"github.com/mickey-mickser/telegram-project/pkg/http"
	"github.com/mickey-mickser/telegram-project/pkg/storage/sheets"
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
	cfg, err := config.NewConfig("./configs/config.json")
	if err != nil {
		panic(err)
	}

	log := cfg.Log()
	botCli := cfg.TelegramTokenCli()
	//dbGorm := cfg.DB()
	//sqlDB, err := dbGorm.DB()
	//if err != nil {
	//	panic(err)
	//}
	//sheetId := cfg.GoogleSheetID()

	ctx, cancel := ctxWithSig()
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
			cancel()
		}
	}()

	cli := sheets.NewSheets(ctx).CliSheets()

	sheetUse := usecase.NewSheetUse(cli)

	botCli.Debug = true

	wg := new(sync.WaitGroup)
	http.NewHttp(log, sheetUse).Run(ctx, wg)
	if err := telegram.NewBot(botCli, log).Start(ctx); err != nil {
		log.Panic(err)
	}

	wg.Wait()

}
