package worker

import (
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
	//var (
	//	err error
	//	s   = gocron.NewScheduler()
	//)
	//if err = s.Every(uint64(h.cfg.Integration.AntiPlagiarismCheckerPeriod)).Seconds().Do(antiPlagiarismResultHandler.Run); err != nil {
	//	log.Println("worker failed", "err", err.Error())
	//}
	//if err = s.Every(uint64(h.cfg.Integration.GrammarCheckerPeriod)).Seconds().Do(grammarCheckerHandler.Run); err != nil {
	//	log.Println("worker failed", "err", err.Error())
	//}
	//
	//<-s.Start()
	return
}
