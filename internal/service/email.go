package service

import (
	"context"
	mail "work-project/internal/email"
	"work-project/internal/schema"

	"github.com/jordan-wright/email"
)

type EmailSender interface {
	Send(ctx context.Context, message schema.Message) error
}

type EmailSenderService struct {
	username string
	password string
	addr     string
}

func NewEmailSenderService(username string, password string, addr string) *EmailSenderService {
	return &EmailSenderService{
		username: username,
		password: password,
		addr:     addr,
	}
}

func (e *EmailSenderService) Send(ctx context.Context, message schema.Message) error {
	attachments, ok := message.Attachments.([]*email.Attachment)
	if !ok {
		attachments = nil
	}
	message.From = e.username

	msg := &email.Email{
		ReplyTo:     nil,
		From:        message.From,
		To:          message.To,
		Bcc:         nil,
		Cc:          nil,
		Subject:     message.Subject,
		Text:        []byte(message.Body),
		HTML:        nil,
		Sender:      message.From,
		Headers:     nil,
		Attachments: attachments,
		ReadReceipt: nil,
	}
	auth := mail.NewAuth(e.username, e.password)

	return msg.Send(e.addr, auth)
}
