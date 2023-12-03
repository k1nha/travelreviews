package handlers

import "fmt"

type PlaceRequest struct {
	Name   string `json:"name"`
	Street string `json:"street"`
	City   string `json:"city"`
}

func (p *PlaceRequest) Validate() error {
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
