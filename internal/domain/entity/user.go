package entity

import (
	"regexp"
	"time"

	"github.com/k1nha/travelreviews/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(name, password, username, email string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        entity.NewID(),
		Name:      name,
		Username:  username,
		Password:  string(hash),
		Email:     email,
		CreatedAt: time.Now(),
	}, nil
}

func (u *User) ValidatePassword(pwrd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwrd))
	return err == nil
}

func (u *User) ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

type UserRepository interface {
	Save(user *User) error
	FindByEmail(email string) (*User, error)
}
