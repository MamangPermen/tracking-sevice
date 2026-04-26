//go:build functional

package service

import "testing"

// Menguji integrasi penulisan log secara langsung ke database.
func TestProcessLog_DB_Failed(t *testing.T) {
	t.Errorf("Functional test for DB log insertion intentionally failed")
}