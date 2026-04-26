package handler

import (
	"net/http"
)

// TrackingAPIHandler bertanggung jawab melayani endpoint REST/gRPC untuk sisi klien (web/app).
// Harus stateless agar mudah di-scale up secara horizontal di Kubernetes.
type TrackingAPIHandler struct {
	// Ketergantungan ke layer service disuntikkan di sini
}

// GetHistory menangani HTTP GET request untuk memuat riwayat pelacakan.
// Operasi ini sangat read-heavy.
func (h *TrackingAPIHandler) GetHistory(w http.ResponseWriter, r *http.Request) {
	// Ekstraksi parameter ID Resi dari URL.
	// Pemanggilan layer service untuk mengambil data agregasi dari NoSQL.
	// Konversi struktur data model.TrackingHistory ke format JSON untuk respons klien.
}