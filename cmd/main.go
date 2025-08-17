package main

import (
	"blog-site/config"
	"blog-site/internal/home"
	"blog-site/internal/register"
	"blog-site/package/bcrypt"
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
	cryptograf := bcrypt.NewCrypto(logger)

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: logger,
	}))
	app.Static("/public", "./public")
	dbpool := database.CreateDbPool(dbConf, logger)
	defer dbpool.Close()

	repository := register.NewUsersRepository(dbpool, logger, cryptograf)

	home.NewHandler(app, logger, repository, cryptograf)
	register.NewHandler(app, logger, repository, cryptograf)

	app.Listen(":5001")
}
