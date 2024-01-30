package entity

import (
	"testing"

	"github.com/k1nha/travelreviews/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewReview(t *testing.T) {
	userId := entity.NewID()
	placeId := entity.NewID()

	review := NewReview(placeId, userId, 5, 3.0, "example description")

	assert.NotNil(t, review)
	assert.NotNil(t, review.ID)
	assert.Equal(t, "example description", review.Description)
	assert.Equal(t, 5, review.Stars)
	assert.Equal(t, userId, review.UserId)

	err := review.Validate()
	assert.Nil(t, err)
}

func TestReviewCanValidate(t *testing.T) {
	userId := entity.NewID()
	placeId := entity.NewID()

	review := NewReview(placeId, userId, 5, 3.0, "example description")
	assert.NotNil(t, review)

	err := review.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "avaliação precisa ser entre 0 a 5 estrelas")
}
