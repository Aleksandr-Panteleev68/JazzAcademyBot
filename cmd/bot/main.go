package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"JazzAcademy/internal/config"
	"JazzAcademy/internal/db"
	"JazzAcademy/internal/utils"

	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	logger := utils.NewLogger(cfg.LogLevel)
	defer logger.Sync()

	pool, err := db.NewPool(cfg.DBConnString)
	if err != nil {
		logger.Fatal("Failed to connect to DB", zap.Error(err))
	}
	defer pool.Close()

	// TODO: примерение миграций?

	// TODO: инициализировать Googel клиент, ТГ бот, сервисы, хэндлеры

	ctx, cancel := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		cancel()
	}()

	logger.Info("Bot skeleton ready. Waiting for signal to shutdown.")
	<-ctx.Done()
	logger.Info("Shutting down")
}
