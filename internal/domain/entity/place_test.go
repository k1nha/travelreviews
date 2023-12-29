package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlace(t *testing.T) {
	p := NewPlace("example name", "example street", "example city")

	assert.NotNil(t, p)

}
