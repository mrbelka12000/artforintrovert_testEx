package config

import (
	"context"
	"sync"

	"github.com/sethvargo/go-envconfig"
	"go.uber.org/zap"
)

type config struct {
	Api struct {
		Port string `env:"PORT,default=8080"`
	}
	MongoDB struct {
		MongoUrl   string `env:"MONGO_URI,required"`
		Database   string `env:"MONGO_DATABASE,required"`
		Collection string `env:"MONGO_COLLECTION,required"`
	}
}

var (
	cfg  *config
	once sync.Once
)

func GetConf() *config {
	once.Do(func() {
		cfg = buildConfig()
		if cfg == nil {
			return
		}
	})
	return cfg
}

func buildConfig() *config {
	conf := config{}
	err := envconfig.Process(context.Background(), &conf)
	if err != nil {
		zap.S().Errorf("failed to build config: %v", err)
		return nil
	}

	return &conf
}
