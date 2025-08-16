package main

import (
	"blog-site/config"
	"blog-site/internal/home"
	"blog-site/internal/register"
	"blog-site/package/database"
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
	dbpool := database.CreateDbPool(dbConf, logger)
	defer dbpool.Close()

	repository := register.NewUsersRepository(dbpool, logger)

	home.NewHandler(app, logger)
	register.NewHandler(app, logger, repository)

	app.Listen(":5001")
}
