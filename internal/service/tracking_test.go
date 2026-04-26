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

	// Skenario: Kita expect DB pura-pura balikin status Delivered buat resi ini
	dummyResi := &repository.Resi{ID: "BUBU-12345", Status: "Delivered"}

	// Ngasih tau si mock: "Kalo ada yang manggil GetResiStatus pake BUBU-12345, kasih dummyResi ya"
	mockRepo.EXPECT().GetResiStatus("BUBU-12345").Return(dummyResi, nil).Times(1)

	// Inject mock repo ke service
	svc := NewTrackingService(mockRepo)

	// Eksekusi function-nya
	result, err := svc.CheckStatus("BUBU-12345")

	if err != nil {
		t.Fatalf("Ga expect error, tapi dapet: %v", err)
	}

	// Ini PASTI FAILED karena di service aslinya tadi cuma return string kosong ""
	if result != "Delivered" {
		t.Errorf("Test failed! Expected 'Delivered', malah dapet: '%v'", result)
	}
}