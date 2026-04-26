package service

import (
	"github.com/MamangPermen/tracking-service/internal/model"
	"github.com/MamangPermen/tracking-service/internal/repository"
)

type TrackingService struct {
	repo repository.TrackingRepository
}

func NewTrackingService(repo repository.TrackingRepository) *TrackingService {
	return &TrackingService{repo: repo}
}

// GetHistory bertugas mengambil seluruh histori paket berdasarkan ID Resi.
func (s *TrackingService) GetHistory(resiID string) (*model.TrackingHistory, error) {
	// Sengaja balikin array kosong biar unit test failed
	return &model.TrackingHistory{
		ResiID:  resiID,
		History: []model.TrackingLog{}, 
	}, nil
}