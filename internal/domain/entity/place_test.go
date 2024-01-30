package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlace(t *testing.T) {
	place := NewPlace("example name", "restaurant")

	assert.NotNil(t, place)
	assert.Equal(t, "example name", place.Name)
	assert.Equal(t, "restaurant", place.Type)
}

func TestUpdatePlace(t *testing.T) {
	place := NewPlace("example name", "restaurant")
	assert.NotNil(t, place)
	place.Touch()
	assert.NotEqual(t, place.UpdateAt, place.CreatedAt)
}
