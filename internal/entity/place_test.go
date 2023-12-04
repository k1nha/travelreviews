package entity

import (
	"testing"

	"github.com/google/uuid"
)

func TestPlace(t *testing.T) {
	id := uuid.New()

	p := NewPlace(id, "example name", "example street", "example city")

	if p.ID != id {
		t.Errorf("place.Id = %v, want %v", p.ID, id)
	}
}
