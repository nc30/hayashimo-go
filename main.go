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

		var (
			mt  int
			msg []byte
			err error
		)

		ctx := context.Background()

		go func(){
			ct, cancel := context.WithCancel(ctx)
			defer recover()
			defer cancel()

			ticker := time.NewTicker(60 * time.Second)
			for {
				select {
				case <-ticker.C:
					l, _ := service.GetNotifcationLength(ct, os.Getenv("GITHUB_KEY"))
					log.Println(l)

					s := &State{Type: "notifcations"}

					if err == nil{
						s.Notifcations.Github = l > 0
					}

					js, err := json.Marshal(s)
					if err == nil{
						go func(){
							defer recover()
							c.WriteMessage(websocket.TextMessage, js)
						}()
					}else{
						log.Println(err)
					}
				case <-ct.Done():
					return
				}
			}
		}()

		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				break
			}
		}

	}))

	f.Listen(":3001")

}
