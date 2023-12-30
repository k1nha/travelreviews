package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlace(t *testing.T) {
	place := NewPlace("example name", "example street", "example city")

	assert.NotNil(t, place)
	assert.Equal(t, "example name", place.Name)
	assert.Equal(t, "example street", place.Street)
}

func TestUpdatePlace(t *testing.T) {
	place := NewPlace("example name", "example street", "example city")
	assert.NotNil(t, place)
	place.Touch()
	assert.NotEqual(t, place.UpdateAt, place.CreatedAt)
}
