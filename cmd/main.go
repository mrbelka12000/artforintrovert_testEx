package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mrbelka12000/artforintrovert_testEx/config"
	"github.com/mrbelka12000/artforintrovert_testEx/internal/handler"
	"github.com/mrbelka12000/artforintrovert_testEx/internal/repository"
	"github.com/mrbelka12000/artforintrovert_testEx/internal/router"
	"github.com/mrbelka12000/artforintrovert_testEx/internal/server"
	"github.com/mrbelka12000/artforintrovert_testEx/internal/service"
	"github.com/mrbelka12000/artforintrovert_testEx/pkg/logger"
	"go.uber.org/zap"
)

func main() {

	mainCtx, mainCancel := context.WithCancel(context.Background())
	done := make(chan os.Signal, 1)
	wait := make(chan struct{})
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	_, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("failed to prepare logger: %v", err)
	}

	client, err := repository.GetMongoDBClient()
	if err != nil {
		zap.S().Debugf("failed to get connection: %v", err)
		return
	}

	cfg := config.GetConf()
	repo := repository.NewRepository(client)
	srv := service.NewService(repo)
	handler := handler.NewHandler(srv)
	mux := router.SetUpMux(handler)
	server := server.NewServer(mux)
	//srv.Insert()

	go repository.Updater(client, mainCtx, wait)
	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.S().Errorf("failed to stop server: %v", err)
			return
		}
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Api.Port
	}

	zap.S().Info("Server started on port :" + port)
	<-done

	ctx, _ := context.WithTimeout(mainCtx, 5*time.Second)
	zap.S().Info("Server stopped")

	defer func() {
		err = client.Disconnect(ctx)
		if err != nil {
			zap.S().Warnf("failed to close client: %v", err)
		} else {
			zap.S().Info("client successfully closed")
		}
		zap.S().Sync()
		close(done)
		close(wait)
	}()

	err = server.Shutdown(ctx)
	if err != nil {
		zap.S().Errorf("failed to shutdown server: %v", err)
		return
	}

	mainCancel()

	<-wait
}
