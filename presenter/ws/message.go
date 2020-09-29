package ws

import ()

type Payload interface {
}

type State struct {
	Type         string       `json:"type"`
	Kintai       int          `json:"kintai"`
	Notifcations Notifcations `json:"notifcations"`
	Message      string       `json:"message"`
}

type Notifcations struct {
	Github bool `json:"github"`
	Gitlab bool `json:"gitlab"`
	Slack  bool `json:"slack"`
	Email  bool `json:"email"`
}
