package service

import (
	"context"
	"firebase.google.com/go/messaging"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type PushNotification interface {
	Send(ctx context.Context, text, header string, token, topic, condition *string) error
	GetAllByUser(ctx context.Context, deviceToken string) ([]model.PushNotification, error)
}

type PushNotificationService struct {
	pushNotification  repository.PushNotification
	firebaseMessaging repository.FirebaseMessaging
}

func NewPushNotificationService(pushNotification repository.PushNotification, firebaseMessaging repository.FirebaseMessaging) *PushNotificationService {
	return &PushNotificationService{pushNotification: pushNotification, firebaseMessaging: firebaseMessaging}
}

func (s *PushNotificationService) Send(ctx context.Context, text, header string, token, topic, condition *string) error {
	message := messaging.Message{
		Topic: "all",
		Notification: &messaging.Notification{
			Title: header,
			Body:  text,
		},
	}
	pushNotificationModel := model.PushNotification{
		Text:      text,
		Title:     header,
		Token:     token,
		Topic:     topic,
		Condition: condition,
	}
	if token != nil {
		message.Token = *token
	}
	if topic != nil {
		message.Topic = *topic
	}
	if condition != nil {
		message.Condition = *condition
	}
	err := s.firebaseMessaging.SendMessage(ctx, message)
	if err != nil {
		return err
	}
	return s.pushNotification.Create(ctx, pushNotificationModel)
}

func (s *PushNotificationService) GetAllByUser(ctx context.Context, deviceToken string) ([]model.PushNotification, error) {
	var token *string
	if len(deviceToken) > 0 {
		token = &deviceToken
	}
	return s.pushNotification.GetByToken(ctx, token)
}
