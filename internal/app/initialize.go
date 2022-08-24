// Package app generates and builds all parts of the application.
package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mrbelka12000/artforintrovert_testEx/config"
	"github.com/mrbelka12000/artforintrovert_testEx/internal/handler"
	"github.com/mrbelka12000/artforintrovert_testEx/internal/service"
	"github.com/mrbelka12000/artforintrovert_testEx/internal/service/repository"
	"github.com/mrbelka12000/artforintrovert_testEx/pkg/cache"
	"github.com/mrbelka12000/artforintrovert_testEx/pkg/logger"
	"github.com/mrbelka12000/artforintrovert_testEx/pkg/mongodb"
	"github.com/mrbelka12000/artforintrovert_testEx/pkg/server"
)

const waitLimitForGS = 5 * time.Second

func Run(ctx context.Context) {
	l, err := logger.NewLogger()
	if err != nil {
		log.Printf("failed to create logger: %v \n", err)
		return
	}

	cfg, err := config.GetConf()
	if err != nil {
		l.Debugf("failed to get config: %v", err)
		return
	}

	client, err := mongodb.GetMongoDBClient(ctx)
	if err != nil {
		l.Debugf("failed to get connection: %v", err)
		return
	}

	runCtx, runCancel := context.WithCancel(ctx)
	done := make(chan os.Signal, 1)
	wait := make(chan struct{})
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	repo := repository.NewRepo(client, l)
	srv := service.NewService(repo, l)
	hndl := handler.NewHandler(srv, l)
	mux := handler.SetUpMux(hndl)
	httpServer := server.NewServer(mux)

	//to manually fill in the database
	srv.InsertProduct()

	go cache.Updater(client, runCtx, wait)

	l.Infof("Server started on port: %v", cfg.Api.Port)
	select {
	case s := <-done:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("http server notify error: %w", err).Error())
	}

	l.Info("Server stopped")

	runCancel()

	stopCtx, _ := context.WithTimeout(ctx, waitLimitForGS)

	err = client.Disconnect(stopCtx)
	if err != nil {
		l.Errorf("failed to close client: %v", err)
	}

	err = httpServer.Shutdown()
	if err != nil {
		l.Errorf("failed to shutdown server: %v", err)
	}

	err = l.Sync()
	if err != nil {
		log.Printf("logger sync error: %v \n", err)
	}

	<-wait
	close(done)
	close(wait)
}
