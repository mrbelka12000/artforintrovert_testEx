package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/mrbelka12000/artforintrovert_testEx/config"
	"github.com/mrbelka12000/artforintrovert_testEx/internal/handler"
	"github.com/mrbelka12000/artforintrovert_testEx/internal/service"
	"github.com/mrbelka12000/artforintrovert_testEx/internal/service/repository"
	"github.com/mrbelka12000/artforintrovert_testEx/pkg/mongodb"
	"github.com/mrbelka12000/artforintrovert_testEx/pkg/server"
)

const waitLimitForGS = 5 * time.Second

func Run(ctx context.Context) {
	cfg := config.GetConf()
	if cfg == nil {
		zap.S().Debug("failed to prepare config")
		return
	}
	fmt.Printf("%+v \n", cfg)

	client, err := mongodb.GetMongoDBClient(ctx)
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

	httpServer := server.NewServer(mux)
	srv.Product.Insert()

	go mongodb.Updater(client, runCtx, wait)

	zap.S().Infof("Server started on port: %v", cfg.Api.Port)

	select {
	case s := <-done:
		zap.S().Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		zap.S().Error(fmt.Errorf("http server notify error: %w", err))
	}

	zap.S().Info("Server stopped")

	runCancel()

	stopCtx, _ := context.WithTimeout(ctx, waitLimitForGS)

	err = client.Disconnect(stopCtx)
	if err != nil {
		zap.S().Warnf("failed to close client: %v", err)
	}

	err = httpServer.Shutdown()
	if err != nil {
		zap.S().Errorf("failed to shutdown server: %v", err)
	}

	<-wait
	close(done)
	close(wait)
}
