package entity

import (
	"testing"

	"github.com/k1nha/travelreviews/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestReview(t *testing.T) {
	pId := entity.NewID()

	review := NewReview(pId, 5, "example description")

	assert.NotNil(t, review)

}
