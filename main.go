package main

import (
	"os"
	"log"
	"time"
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"

	"github.com/nc30/hayashimo-go/presenter/http/middleware"
	"github.com/nc30/hayashimo-go/domain/service"
)

type State struct {
	Type string `json:"type"`
	Kintai       int `json:"kintai"`
	Notifcations Notifcations `json:"notifcations"`
	Message string `json:"message"`
}

type Notifcations struct{
	Github bool `json:"github"`
	Gitlab bool `json:"gitlab"`
	Slack  bool `json:"slack"`
	Email  bool `json:"email"`
}

func main() {

	f := fiber.New()

	middleware.SetMddlewares(f)

	f.Get("/status", func(c *fiber.Ctx) error {
		return c.JSON(State{Kintai: 0})
	})

	f.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		if c.Params("id") != os.Getenv("WEBSOCKET_KEY"){
			c.Close()
			return
		}

		c.SetReadDeadline(time.Now().Add((60 + 5) * time.Second))
		c.SetPongHandler(func(string) error { c.SetReadDeadline(time.Now().Add((60 + 5) * time.Second)); return nil })

		ctx := context.Background()

		go func() {
			ct, cancel := context.WithCancel(ctx)
			defer cancel()
			defer recover()
			ticker := time.NewTicker(60 * time.Second)
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					go func(){
						defer recover()
						c.WriteMessage(websocket.PingMessage, nil)
					}()
				case <-ct.Done():
					return
				}
			}
		}()

		ct, cancel := context.WithCancel(ctx)
		defer cancel()
		ticker := time.NewTicker(60 * time.Second)
		defer ticker.Stop()
		defer recover()

		for {
			select {
			case <-ticker.C:
				l, err := service.GetNotifcationLength(ct, os.Getenv("GITHUB_KEY"))

				s := &State{Type: "notifcations"}

				if err == nil{
					s.Notifcations.Github = l > 0
				}

				js, err := json.Marshal(s)
				if err == nil{
					go func(){
						defer func(){
							err := recover()
							if err != nil{
								log.Println(err)
							}
						}()
						c.WriteMessage(websocket.TextMessage, js)
						log.Println("sended.")
					}()
				}else{
					log.Println(err)
				}
			case <-ct.Done():
				return
			}
		}

		<-ctx.Done()
	}))

	f.Listen(":3001")

}
