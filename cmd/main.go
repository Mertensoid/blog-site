package main

import (
	"blog-site/config"
	"blog-site/package/logger"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	loggerConf := config.NewLogConfig()

	logger := logger.NewLogger(loggerConf)
}
