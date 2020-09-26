package main

import (
	"github.com/gofiber/fiber/v2"
)

type State struct {
	Kintai       int `json:"kintai"`
	Notifcations struct {
		Github bool `json:"github"`
		Gitlab bool `json:"gitlab"`
		Slack  bool `json:"slack"`
		Email  bool `json:"email"`
	} `json:"notifcations"`
	Messages []struct {
		Type    string `form:"json:"type"`
		Message string `form:"json:"message"`
	} `json:"messages"`
}

func main() {

	f := fiber.New()

	f.Get("/:name", func(c *fiber.Ctx) error {
		return c.JSON(State{Kintai: 0})
	})

	f.Listen(":3001")

}
