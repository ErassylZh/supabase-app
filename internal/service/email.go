package service

import (
	"bytes"
	"context"
	"fmt"
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

	msg := email.NewEmail()
	msg.From = message.From
	msg.To = message.To
	msg.Subject = message.Subject
	msg.Text = []byte(message.Body)

	// Добавляем вложения из списка
	for _, attachment := range attachments {
		msg.Attachments = append(msg.Attachments, attachment)
	}

	fmt.Println(message.FileData)
	// Добавляем файл в сообщение, если он указан в виде []byte
	if len(message.FileData) > 0 {
		fileReader := bytes.NewReader(message.FileData)
		_, err := msg.Attach(fileReader, message.FileName, "application/octet-stream")
		if err != nil {
			return err
		}
	}
	fmt.Println(msg.Attachments)

	auth := mail.NewAuth(e.username, e.password)
	return msg.Send(e.addr, auth)
}
