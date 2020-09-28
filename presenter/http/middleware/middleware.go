package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetMddlewares(f *fiber.App) {
	f.Use(NewCors())
	f.Use(NewLogger())
	f.Use(NewCompressor())
}

func NewCors() fiber.Handler {
	return cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	})
}

func NewLogger() fiber.Handler {
	return logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Tokyo",
		Output:     os.Stdout,
	})
}

func NewCompressor() fiber.Handler {
	return compress.New()
}
