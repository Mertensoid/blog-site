package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Url string
}

type LogConfig struct {
	Level  int
	Format string
}

func Init() {
	if err := godotenv.Load("/Users/admin/Documents/Обучение/GOLANG/PurpleSchool/blog-site/.env"); err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(".env loaded")
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Url: getString("DATABASE_URL", ""),
	}
}

func NewLogConfig() *LogConfig {
	return &LogConfig{
		Level:  getInt("LOG_LEVEL", 0),
		Format: getString("LOG_FORMAT", "json"),
	}
}

func getString(s, defaultValue string) string {
	res := os.Getenv(s)
	if res == "" {
		return defaultValue
	} else {
		return res
	}
}

func getInt(s string, defaultValue int) int {
	res, err := strconv.Atoi(os.Getenv(s))
	if err != nil {
		return defaultValue
	}
	return res
}

func getBool(s string, defaultValue bool) bool {
	res, err := strconv.ParseBool(os.Getenv(s))
	if err != nil {
		return defaultValue
	}
	return res
}
