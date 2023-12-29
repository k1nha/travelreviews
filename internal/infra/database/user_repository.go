package database

import (
	"database/sql"
	"errors"

	"github.com/k1nha/travelreviews/internal/domain/entity"
)

type UserRepository struct {
	Db *sql.DB
}

func (u *UserRepository) Save(user *entity.User) error {
	_, err := u.Db.Exec("INSERT INTO users(id, email, username, password, created_at)  VALUES (?, ?, ?, ?, ?)", user.ID, user.Email, user.Username, user.Password, user.CreatedAt)

	if err != nil {
		return errors.New("error on insert user on table users")
	}

	return nil
}

func (u *UserRepository) FindByEmail(email string) (*entity.User, error) {
	row := u.Db.QueryRow("SELECT * FROM users WHERE email = ?", email)

	if err != nil {
		return nil, errors.New("error on exec query to find by email in users")
	}

	var user entity.User

	err = row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt)

	return &user, nil
}
