package tokenhandler

import (
	"fmt"

	"github.com/k1nha/travelreviews/util"
)

type TokenRequest struct {
	Email    string
	Username string
}

func (t *TokenRequest) Validate() error {
	if t.Email == "" || t.Username == "" {
		return fmt.Errorf("email or/and username cannot be empty")
	}

	if !util.ValidEmail(t.Email) {
		return fmt.Errorf("email invalid")

	}

	return nil
}
