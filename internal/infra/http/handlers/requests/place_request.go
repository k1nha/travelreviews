package requests

import "fmt"

type CreatePlaceRequest struct {
	Name   string `json:"name"`
	Typeof string `json:"typeof"`
}

func (p *CreatePlaceRequest) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if p.Typeof == "" {
		return fmt.Errorf("typeof cannot be empty")
	}

	return nil
}
