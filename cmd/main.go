package main

import (
	"context"
	"log"

	"github.com/mrbelka12000/artforintrovert_testEx/internal/app"
)

func main() {
	ctx := context.Background()
	app.Run(ctx)
	err := ctx.Err()
	if err != nil {
		log.Printf("context error: %v \n", err)
	}
}
