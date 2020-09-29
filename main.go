package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/nc30/hayashimo-go/presenter/http/middleware"
	"github.com/nc30/hayashimo-go/presenter/http/router"
)

func main() {
	f := fiber.New()

	middleware.SetMddlewares(f)
	router.SetRoute(f)

	f.Listen(":3001")
}
