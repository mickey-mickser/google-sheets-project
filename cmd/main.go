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
	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
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
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	configPath := os.Getenv("CONFIG_PATH")
	if _, err := os.Stat(configPath); err != nil {
		configPath = "." + configPath
	}
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		panic(err)
	}

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
	cli, err := sheets.NewSheets(ctx).CliSheets(permission.ServiceKeyPath)
	if err != nil {
		panic(err)
	}
	handlers := handler.NewHandler(log, cli, permission)

	//Block main() until all background goroutines (like the server) complete
	wg := new(sync.WaitGroup)

	// Run server with graceful shutdown
	server.NewServer(log, os.Getenv("PORT"), handlers.InitRoutes()).Run(ctx, wg)

	wg.Wait()
}
