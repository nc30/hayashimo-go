package model

import (
	"github.com/gofiber/websocket/v2"
)

type User struct {
	githubKey string
	Channels  map[*websocket.Conn]struct{}
}

func (u *User) Send(p *Payload) error {
	js, err := json.Marshal(p)
	if err != nil {
		return err
	}

	for c, _ := range u.Channels {
		go func() {
			defer recover()
			c.WriteMessage(websocket.TextMessage, js)
		}()
	}

	return nil
}

func (u *User) Fetch() {
	// github
	g := new(chan bool)
	go func() {
		l, err := service.GetNotifcationLength(ct, u.githubKey)
		if err == nil {
			s.Notifcations.Github = l > 0
		}
	}()

	s := &State{
		Type:   "notifcations",
		Github: <-g,
	}

	u.Send(s)
}
