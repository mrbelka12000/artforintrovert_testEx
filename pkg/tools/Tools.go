package tools

import (
	"fmt"
	"os"
	
	"github.com/mrbelka12000/artforintrovert_testEx/config"
)

func GetConnectionString() (connStr string) {
	cfg := config.GetConf()

	connStr = os.Getenv("MONGO_URI")

	if connStr == "" {
		connStr = fmt.Sprintf("mongodb://%v:%v", cfg.MongoDB.Host, cfg.MongoDB.Port)
	}

	return
}
