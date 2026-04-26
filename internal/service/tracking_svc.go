// internal/service/tracking.go
package service

import "github.com/MamangPermen/tracking-service/internal/repository"

type TrackingService struct {
	repo repository.TrackingRepository
}

func NewTrackingService(repo repository.TrackingRepository) *TrackingService {
	return &TrackingService{repo: repo}
}

// Main logic
func (s *TrackingService) CheckStatus(id string) (string, error) {
	return "", nil
}