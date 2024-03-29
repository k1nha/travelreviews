package entity

import (
	"time"

	"github.com/k1nha/travelreviews/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            entity.ID `json:"id"`
	Name          string    `json:"name"`
	Password      string    `json:"password"`
	Email         string    `json:"email"`
	
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func NewUser(name, password, username, email string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:            entity.NewID(),
		Name:          name,
		Password:      string(hash),
		Email:         email,
		UpdatedAt:     time.Now(),
		CreatedAt:     time.Now(),
	}, nil
}

func (u *User) ValidatePassword(pwrd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwrd))
	return err == nil
}

func (u *User) Touch() {
	u.UpdatedAt = time.Now()
}

type UserRepository interface {
	Save(user *User) error
	FindByEmail(email string) (*User, error)
	FindById(id string) (*User, error)
}
