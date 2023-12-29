package authusecase

import (
	"fmt"

	"github.com/k1nha/travelreviews/internal/domain"
	"github.com/k1nha/travelreviews/internal/domain/adapter"
)

type TokenInput struct {
	Email    string
	Username string
}

type TokenOutput struct {
	Token string
}

type GenerateToken struct {
	UserRepository domain.UserRepository
	JwtAdapter     adapter.JwtAdapter
}

func (g *GenerateToken) Execute(i TokenInput) (*TokenOutput, error) {
	user, err := g.UserRepository.FindByEmail(i.Email)

	fmt.Println(user)
	if err != nil {
		return nil, err
	}

	tkn, err := g.JwtAdapter.CreateToken(i.Email, i.Username)

	if err != nil {
		return nil, err
	}

	return &TokenOutput{
		Token: tkn,
	}, nil
}
