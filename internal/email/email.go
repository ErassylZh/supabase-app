package email

import (
	"errors"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func NewEmail(title string, body []byte, from string, to ...string) (e *email.Email, err error) {

	return &email.Email{
		Subject: title,
		HTML:    body,
		To:      to,
		From:    from,
	}, nil
}

type Auth struct {
	username, password string
}

func NewAuth(username, password string) smtp.Auth {
	return &Auth{username, password}
}

func (a *Auth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *Auth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown from server")
		}
	}
	return nil, nil
}
