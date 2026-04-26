// internal/repository/tracking.go
package repository

// struct buat nampung data resi
type Resi struct {
	ID     string
	Status string
}

// interface untuk repository tracking
type TrackingRepository interface {
	GetResiStatus(id string) (*Resi, error)
}