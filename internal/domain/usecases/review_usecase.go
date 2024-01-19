package usecase

import (
	"github.com/k1nha/travelreviews/internal/domain/entity"
	pkgentity "github.com/k1nha/travelreviews/pkg/entity"
)

type ReviewUseCases struct {
	ReviewRepository entity.ReviewRepository
}

func NewReviewUseCases(reviewRepo entity.ReviewRepository) *ReviewUseCases {
	return &ReviewUseCases{
		ReviewRepository: reviewRepo,
	}
}

func (r *ReviewUseCases) Create(placeId, userId, description string, stars int) error {
	placeID, err := pkgentity.ParseID(placeId)
	if err != nil {
		return err
	}

	userID, err := pkgentity.ParseID(userId)
	if err != nil {
		return err
	}

	review := entity.NewReview(placeID, userID, stars, description)
	err = r.ReviewRepository.Save(review)
	if err != nil {
		return err
	}

	return nil
}
