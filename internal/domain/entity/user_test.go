package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John", "12345", "johndoe", "john@doe.com")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "John", user.Name)
	assert.Equal(t, "john@doe.com", user.Email)

}

func TestNewUserWithValidEmail(t *testing.T) {
	user, err := NewUser("John", "12345", "johndoe", "john@doe.com")

	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.True(t, user.ValidateEmail("john@doe.com"))
	assert.False(t, user.ValidateEmail("john@.com"))
}

func TestNewUserWithHashPassword(t *testing.T) {
	user, err := NewUser("John", "12345", "johndoe", "john@doe.com")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, user.ValidatePassword("12345"))
	assert.False(t, user.ValidatePassword("123"))
	assert.NotEqual(t, "12345", user.Password)
}
