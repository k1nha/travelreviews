package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewReview(t *testing.T) {
	user, _ := NewUser("John", "12345", "johndoe", "john@doe.com")
	p := NewPlace("example name", "example street", "example city")

	review := NewReview(p.ID, user.ID, 5, "example description")

	assert.NotNil(t, review)
	assert.Equal(t, "example description", review.Description)
	assert.Equal(t, 5, review.Stars)
	assert.Equal(t, user.ID, review.UserId)
}
