package worker

import (
	"github.com/jasonlvhit/gocron"
	"log"
	"work-project/internal/config"
	"work-project/internal/repository"
	"work-project/internal/service"
)

type Handler struct {
	cfg          *config.Config
	services     *service.Services
	repositories *repository.Repositories
}

func NewHandlerWorker(
	cfg *config.Config,
	services *service.Services,
	repositories *repository.Repositories,
) *Handler {
	return &Handler{
		cfg:          cfg,
		services:     services,
		repositories: repositories,
	}
}

func (h *Handler) Init() {
	var (
		err error
		s   = gocron.NewScheduler()
	)
	pushNotificationReader := NewPushNotificationCheckerHandler(
		h.repositories.FirebaseMessaging,
		h.repositories.PushNotification,
	)
	airTableSync := NewAirTableSync(
		h.repositories.Airtable,
		h.repositories.Product,
		h.repositories.Post,
		h.repositories.StorageClient,
		h.repositories.Image,
	)
	if err = s.Every(uint64(h.cfg.Integration.PushNotificationReadPeriod)).Seconds().Do(pushNotificationReader.Run); err != nil {
		log.Println("worker failed", "err", err.Error())
	}
	if err = s.Every(uint64(h.cfg.Integration.AirtableSyncPeriod)).Seconds().Do(airTableSync.Run); err != nil {
		log.Println("worker failed", "err", err.Error())
	}

	<-s.Start()
	return
}
