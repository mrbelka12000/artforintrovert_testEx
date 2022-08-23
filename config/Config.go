package config

import (
	"fmt"
	"os"
	"sync"

	"go.uber.org/zap"
)

type config struct {
	Api struct {
		Port string
	}
	MongoDB struct {
		MongoUrl   string
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
		cfg = buildConfig()
		if cfg == nil {
			return
		}
	})
	return cfg
}

func buildConfig() *config {
	conf := &config{}

	conf.Api.Port = os.Getenv("PORT")
	conf.MongoDB.Database = os.Getenv("MONGO_DATABASE")
	conf.MongoDB.Collection = os.Getenv("MONGO_COLLECTION")
	conf.MongoDB.MongoUrl = os.Getenv("MONGO_URI")
	conf.MongoDB.Host = os.Getenv("MONGO_HOST")
	conf.MongoDB.Port = os.Getenv("MONGO_PORT")

	if conf.Api.Port == "" {
		zap.S().Error("no port in env")
		return nil
	}
	if conf.MongoDB.Database == "" {
		zap.S().Error("no mongo database in env")
		return nil
	}
	if conf.MongoDB.Collection == "" {
		zap.S().Error("no mongo collection in env")
		return nil
	}
	if conf.MongoDB.MongoUrl == "" {
		conf.MongoDB.MongoUrl = getConnectionToMongoString(conf)
		if conf.MongoDB.MongoUrl == "" {
			zap.S().Error("no mongo connection data in env")
			return nil
		}
	}

	return conf
}

func getConnectionToMongoString(conf *config) (connStr string) {
	connStr = fmt.Sprintf("mongodb://%v:%v", conf.MongoDB.Host, conf.MongoDB.Port)
	return
}
