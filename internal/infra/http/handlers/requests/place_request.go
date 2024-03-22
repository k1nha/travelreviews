package requests

import "fmt"

type CreatePlaceRequest struct {
	Name    string `json:"name"`
	Typeof  string `json:"typeof"`
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
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
