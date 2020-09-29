package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/nc30/hayashimo-go/presenter/ws"
)

func StatusHandler(c *fiber.Ctx) error {
	return c.JSON(ws.State{Kintai: 0})
}
