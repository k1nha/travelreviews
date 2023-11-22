package entity

type ReviewRepository interface {
	Save(review *Review) error
	GetAll() []Review
}
