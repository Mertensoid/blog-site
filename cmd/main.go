package main

import (
	"blog-site/config"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	loggerConf := config.NewLogConfig()
}
