package service

import (
	"context"
	"firebase.google.com/go/messaging"
	"work-project/internal/model"
	"work-project/internal/repository"
	"work-project/internal/repository/integration"
)

type PushNotification interface {
	Send(ctx context.Context, text, header string, token, topic, condition *string) error
}

type PushNotificationService struct {
	pushNotification  repository.PushNotification
	firebaseMessaging integration.FirebaseMessaging
}

func NewPushNotificationService(pushNotification repository.PushNotification, firebaseMessaging integration.FirebaseMessaging) *PushNotificationService {
	return &PushNotificationService{pushNotification: pushNotification, firebaseMessaging: firebaseMessaging}
}

func (s *PushNotificationService) Send(ctx context.Context, text, header string, token, topic, condition *string) error {
	message := messaging.Message{
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
