package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	place := NewPlace("Lugar", "Rua dos Lugares", "Luggar")

	assert.NotNil(t, place)
	assert.Equal(t, place.Name, "Lugas")
}
