package entity

import (
	"github.com/k1nha/travelreviews/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
}

func NewUser(name, password, email string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Password: string(hash),
		Email:    email,
	}, nil
}

func (u *User) ValidatePassword(pwrd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwrd))
	return err != nil
}
