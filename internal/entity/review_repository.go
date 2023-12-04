package entity

type ReviewInterface interface {
	Save(review *Review) error
	//GetAll() []Review
}
