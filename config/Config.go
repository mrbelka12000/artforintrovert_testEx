package config

import (
	"os"
	"sync"

	"go.uber.org/zap"

	"github.com/tkanos/gonfig"
)

const cfgPath = "config/config.json"

type config struct {
	Api struct {
		Port string
	}
	MongoDB struct {
		Database   string
		Collection string
		Host       string
		Port       string
	}
}

var (
	cfg  *config
	once sync.Once
)

func GetConf() *config {
	once.Do(func() {
		cfg = parseConf()
		if cfg == nil {
			return
		}
	})
	return cfg
}

func parseConf() *config {
	path := os.Getenv("cfgPath")
	if path == "" {
		path = cfgPath
	}
	cfg = &config{}

	err := gonfig.GetConf(path, cfg)
	if err != nil {
		zap.S().Debugf("parse config error: %v ", err.Error())
		return nil
	}
	return cfg
}
