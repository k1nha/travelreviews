package placehandler

import "fmt"

type CreatePlaceRequest struct {
	Name   string `json:"name"`
	Street string `json:"street"`
	City   string `json:"city"`
}

func (p *CreatePlaceRequest) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if p.Street == "" {
		return fmt.Errorf("street cannot be empty")
	}

	if p.City == "" {
		return fmt.Errorf("city cannot be empty")
	}

	return nil
}
