// internal/service/tracking_test.go
package service

import (
	"testing"

	"github.com/MamangPermen/tracking-service/internal/repository"
	"github.com/MamangPermen/tracking-service/internal/repository/mock"
	"github.com/golang/mock/gomock"
)

func TestCheckStatus_Failed(t *testing.T) {
	// Setup Gomock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockTrackingRepository(ctrl)

	// expect DB pura-pura balikin status Delivered buat resi ini
	dummyResi := &repository.Resi{ID: "TEST-12345", Status: "Delivered"}

	// kirim request ke repo, expect balikin dummyResi
	mockRepo.EXPECT().GetResiStatus("TEST-12345").Return(dummyResi, nil).Times(1)

	// Inject mock repo ke service
	svc := NewTrackingService(mockRepo)

	// Eksekusi function-nya
	result, err := svc.CheckStatus("TEST-12345")

	if err != nil {
		t.Fatalf("Unexpected error, but got: %v", err)
	}

	// tes 
	if result != "Delivered" {
		t.Errorf("Test failed! Expected 'Delivered', got: '%v'", result)
	}
}