package service

import (
	"errors"
	"github.com/MamangPermen/tracking-service/internal/model"
	"github.com/MamangPermen/tracking-service/internal/repository"
)

type LogEventService struct {
	repo repository.LogEventRepository
}

func NewLogEventService(repo repository.LogEventRepository) *LogEventService {
	return &LogEventService{repo: repo}
}

// ProcessLog bertugas memvalidasi dan menyimpan log baru.
func (s *LogEventService) ProcessLog(log model.TrackingLog) error {
	// Dummy error buat testing
	return errors.New("intentional error for testing purposes")
}