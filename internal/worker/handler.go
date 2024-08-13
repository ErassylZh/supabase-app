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
	airTableSync := NewAirTableSync(
		h.repositories.Airtable,
		h.repositories.Product,
	)
	if err = s.Every(uint64(300)).Seconds().Do(airTableSync.Run); err != nil {
		log.Println("worker failed", "err", err.Error())
	}

	<-s.Start()
	return
}
