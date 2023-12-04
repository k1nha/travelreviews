package entity

import (
	"testing"

	"github.com/google/uuid"
)

func TestReview(t *testing.T) {
	id := uuid.New()
	pId := uuid.New()

	review := NewReview(id, pId, 5, "example description")

	if review.ID != id {
		t.Errorf("review.Id = %v, want %v", review.ID, id)
	}
}
