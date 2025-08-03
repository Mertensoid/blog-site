package main

import (
	"blog-site/config"
	"blog-site/internal/home"
	"blog-site/package/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	loggerConf := config.NewLogConfig()

	logger := logger.NewLogger(loggerConf)
	logger.Info().Msg(dbConf.Url)

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: logger,
	}))
	app.Static("/public", "./public")

	home.NewHandler(app, logger)

	app.Listen(":5001")
}
