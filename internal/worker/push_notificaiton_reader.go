package worker

import (
	"context"
	"firebase.google.com/go/messaging"
	"google.golang.org/appengine/log"
	"work-project/internal/repository"
)

type PushNotificationCheckerHandler struct {
	firebaseClient   repository.FirebaseMessaging
	pushNotification repository.PushNotification
}

func NewPushNotificationCheckerHandler(firebaseClient repository.FirebaseMessaging, pushNotification repository.PushNotification) *PushNotificationCheckerHandler {
	return &PushNotificationCheckerHandler{firebaseClient: firebaseClient, pushNotification: pushNotification}
}

func (h *PushNotificationCheckerHandler) Run() (err error) {
	ctx := context.Background()

	pushNotificationModelsNotSended, err := h.pushNotification.GetAllNotSended(ctx)
	if err != nil {
		return err
	}
	for _, pushNotificationModel := range pushNotificationModelsNotSended {
		message := messaging.Message{
			Notification: &messaging.Notification{
				Title: pushNotificationModel.Title,
				Body:  pushNotificationModel.Text,
			},
		}
		if pushNotificationModel.Topic != nil {
			message.Topic = *pushNotificationModel.Topic
		} else {
			message.Topic = "all"
		}
		if pushNotificationModel.Condition != nil {
			message.Topic = *pushNotificationModel.Condition
		}
		if pushNotificationModel.Token != nil {
			message.Topic = *pushNotificationModel.Token
		}

		err := h.firebaseClient.SendMessage(ctx, message)
		if err != nil {
			log.Errorf(ctx, "error while send firebase send message worker", "err", err)
		}

		pushNotificationModel.IsSended = true
		err = h.pushNotification.Update(ctx, pushNotificationModel)
		if err != nil {
			log.Errorf(ctx, "error while update push notification model sended", "err", err)
		}
	}

	return
}
