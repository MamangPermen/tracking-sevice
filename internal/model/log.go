package model

import "time"

// TrackingLog merepresentasikan data log pemindaian tunggal dari titik transit.
// Sesuai untuk skema tabel log di database NoSQL (ScyllaDB/Cassandra).
type TrackingLog struct {
	ResiID       string    `json:"resi_id"`
	LocationCode string    `json:"location_code"`
	ActivityCode string    `json:"activity_code"`
	PhotoURL     string    `json:"photo_url,omitempty"`
	Timestamp    time.Time `json:"timestamp"`
}

// TrackingHistory merepresentasikan struktur balikan untuk API pelanggan.
// Berisi kumpulan log yang sudah diagregasi berdasarkan ID Resi.
type TrackingHistory struct {
	ResiID  string        `json:"resi_id"`
	History []TrackingLog `json:"history"`
}