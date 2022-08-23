package main

import (
	"context"
	"log"

	"github.com/mrbelka12000/artforintrovert_testEx/internal/app"
	"github.com/mrbelka12000/artforintrovert_testEx/pkg/logger"
)

func main() {
	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("failed to prepare logger: %v", err)
	}
	defer logger.Sync()

	ctx := context.Background()
	app.Run(ctx)
	err = ctx.Err()
	if err != nil {
		logger.Errorf("context error: %v", err)
	}
}
