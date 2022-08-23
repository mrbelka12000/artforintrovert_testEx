package main

import (
	"log"

	"github.com/mrbelka12000/artforintrovert_testEx/internal/app"
	"github.com/mrbelka12000/artforintrovert_testEx/pkg/logger"
)

func main() {
	_, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("failed to prepare logger: %v", err)
	}

	app.Run()
}
