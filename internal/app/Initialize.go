package app

import (
	"context"
	"errors"
	"fmt"
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
	"github.com/mrbelka12000/artforintrovert_testEx/internal/server"
	"github.com/mrbelka12000/artforintrovert_testEx/pkg/service"
)

const waitLimitForGS = 5 * time.Second

func Run(ctx context.Context) {
	cfg := config.GetConf()
	if cfg == nil {
		zap.S().Debug("failed to prepare config")
		return
	}
	fmt.Printf("%+v \n", cfg)

	client, err := db.GetMongoDBClient(ctx)
	if err != nil {
		zap.S().Debugf("failed to get connection: %v", err)
		return
	}

	runCtx, runCancel := context.WithCancel(ctx)
	done := make(chan os.Signal, 1)
	wait := make(chan struct{})
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	repo := repository.NewRepository(client)
	srv := service.NewService(repo)
	hndl := handler.NewHandler(srv)
	mux := handler.SetUpMux(hndl)
	server := server.NewServer(mux)
	srv.Product.Insert()

	go db.Updater(client, runCtx, wait)
	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.S().Errorf("failed to stop server: %v", err)
			return
		}
	}()

	zap.S().Infof("Server started on port: %v", cfg.Api.Port)
	<-done

	stopCtx, _ := context.WithTimeout(ctx, waitLimitForGS)
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
	runCancel()

	err = server.Shutdown(stopCtx)
	if err != nil {
		zap.S().Errorf("failed to shutdown server: %v", err)
		return
	}

	<-wait
}
