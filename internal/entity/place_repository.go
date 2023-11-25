package entity

type PlaceInterface interface {
	Save(place *Place) error
}
