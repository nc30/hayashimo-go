package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"

	"github.com/nc30/hayashimo-go/presenter/http/handler"
	"github.com/nc30/hayashimo-go/presenter/ws"
)

func SetRoute(f *fiber.App) {
	f.Get("/status", handler.StatusHandler)

	f.Get("/ws/:id", websocket.New(ws.WebsocketHandler))
}
