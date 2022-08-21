package config

import (
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
	cfg = &config{}

	err := gonfig.GetConf(cfgPath, cfg)
	if err != nil {
		zap.S().Debugf("parse config error: %v ", err.Error())
		return nil
	}
	return cfg
}
