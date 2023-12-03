package handlers

import (
	"fmt"
)

type ReviewRequest struct {
	PlaceId     string `json:"place_id"`
	Description string `json:"description"`
	Stars       int    `json:"stars"`
}

func (c *ReviewRequest) Validate() error {
	if c.PlaceId == "" {
		return fmt.Errorf("PlaceId cannot be empty")
	}

	if c.Description == "" {
		return fmt.Errorf("description cannot be empty")
	}

	if c.Stars < 0 && c.Stars > 5 {
		return fmt.Errorf("stars error validation")
	}

	return nil
}
