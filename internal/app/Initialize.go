package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/mrbelka12000/artforintrovert_testEx/config"
	"github.com/mrbelka12000/artforintrovert_testEx/db"
	"github.com/mrbelka12000/artforintrovert_testEx/internal/handler"
	"github.com/mrbelka12000/artforintrovert_testEx/internal/repository"
	"github.com/mrbelka12000/artforintrovert_testEx/internal/router"
	"github.com/mrbelka12000/artforintrovert_testEx/internal/server"
	"github.com/mrbelka12000/artforintrovert_testEx/pkg/service"
)

func Run() {
	cfg := config.GetConf()
	if cfg == nil {
		zap.S().Debug("failed to prepare config")
		return
	}

	client, err := db.GetMongoDBClient()
	if err != nil {
		zap.S().Debugf("failed to get connection: %v", err)
		return
	}

	runCtx, runCancel := context.WithCancel(context.Background())
	done := make(chan os.Signal, 1)
	wait := make(chan struct{})
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	repo := repository.NewRepository(client)
	srv := service.NewService(repo)
	handler := handler.NewHandler(srv)
	mux := router.SetUpMux(handler)
	server := server.NewServer(mux)
	srv.Insert()

	go db.Updater(client, runCtx, wait)
	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.S().Errorf("failed to stop server: %v", err)
			return
		}
	}()

	zap.S().Info("Server started on port :" + cfg.Api.Port)
	<-done

	stopCtx, _ := context.WithTimeout(runCtx, 5*time.Second)
	zap.S().Info("Server stopped")

	defer func() {
		err = client.Disconnect(stopCtx)
		if err != nil {
			zap.S().Warnf("failed to close client: %v", err)
		} else {
			zap.S().Info("client successfully closed")
		}
		close(done)
		close(wait)
	}()

	err = server.Shutdown(stopCtx)
	if err != nil {
		zap.S().Errorf("failed to shutdown server: %v", err)
	}

	runCancel()
	<-wait
}
