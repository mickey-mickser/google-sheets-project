// @title Google Sheets API
// @version 1.0
// @description This is a sample API for Google Sheets interaction
// @termsOfService http://example.com/terms/

// @contact.name Your Name
// @contact.url http://www.example.com
// @contact.email example@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
// @schemes http
package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/mickey-mickser/google-sheets-project/cmd/docs"
	"github.com/mickey-mickser/google-sheets-project/pkg/clients/sheets"
	"github.com/mickey-mickser/google-sheets-project/pkg/config"
	"github.com/mickey-mickser/google-sheets-project/pkg/server"
	"github.com/mickey-mickser/google-sheets-project/pkg/server/handler"
	"github.com/sirupsen/logrus"
	_ "google.golang.org/api/sheets/v4"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Setup context with graceful shutdown
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
	//if _, err := os.Stat(".env"); err == nil {
	//	if err := godotenv.Load(); err != nil {
	//		log.Printf("warning: failed to load .env: %v", err)
	//	}
	//} else {
	//	log.Println("note: .env file not found, skipping godotenv")
	//}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("cannot load .env: %v", err)
	}
	cfg, err := config.NewConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		logrus.Fatalf("cannot load config: %v", err)
	}
	//if err := godotenv.Load(); err != nil {
	//	logrus.Fatalf("error loading env variables: %s", err.Error())
	//}
	//configPath := os.Getenv("CONFIG_PATH") // "/configs/config.json"
	//cfg, err := config.NewConfig(configPath)
	//if err != nil {
	//	log.Fatalf("cannot load config %s: %v", configPath, err)
	//}

	//configPath := os.Getenv("CONFIG_PATH") // "/configs/config.json"
	//if _, err := os.Stat(configPath); err != nil {
	//	configPath = "." + configPath
	//}
	//cfg, err := config.NewConfig(configPath)
	//if err != nil {
	//	panic(err)
	//}

	log := cfg.Log()
	permission := cfg.Permissions()
	ctx, cancel := ctxWithSig()
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
			cancel()
		}
	}()

	// Initialize services
	sheetCli := sheets.NewClients(log)

	if err := sheetCli.CliSheetsADC(ctx); err != nil {
		logrus.Fatalf("failed to init Sheets client: %v", err)
	}

	sheetSvc := sheetCli.ServiceSheets()
	driveSvc := sheetCli.ServiceDrive()

	handlers := handler.NewHandler(log, sheetSvc, driveSvc, permission, cfg.Sheet().SheetID)

	//Block main() until all background goroutines (like the server) complete
	wg := new(sync.WaitGroup)

	// Run server with graceful shutdown
	server.NewServer(log, os.Getenv("PORT"), handlers.InitRoutes()).Run(ctx, wg)

	wg.Wait()
}
