package config

import (
	"context"
	"fmt"
	"sync"

	"github.com/sethvargo/go-envconfig"
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

func GetConf() (conf *config, err error) {
	once.Do(func() {
		conf, err = buildConfig()
		if conf == nil {
			return
		}
		cfg = conf
	})

	return cfg, err
}

func buildConfig() (*config, error) {
	conf := config{}
	err := envconfig.Process(context.Background(), &conf)
	if err != nil {
		return nil, fmt.Errorf("failed to build config: %w", err)
	}

	return &conf, nil
}
