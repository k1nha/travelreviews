package user_usecase

import "github.com/k1nha/travelreviews/internal/domain/entity"

type input struct {
	name     string
	password string
	username string
	email    string
}

type output struct {
	user *entity.User
}

type CreateUser struct {
	UserRepository entity.UserRepository
}

func (c *CreateUser) Execute(v input) (*output, error) {
	user, err := entity.NewUser(v.name, v.password, v.username, v.email)
	if err != nil {
		return nil, err
	}

	err = c.UserRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return &output{
		user,
	}, nil
}
