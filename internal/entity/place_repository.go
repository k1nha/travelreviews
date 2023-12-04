package entity

type PlaceInterface interface {
	Save(place *Place) error
	GetById(id string) (*Place, error)
}
