package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Fabio", "fabio@email.com", "senha")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Fabio", user.Name)
	assert.Equal(t, "fabio@email.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Fabio", "fabio@email.com", "senha")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("senha"))
	assert.False(t, user.ValidatePassword("senha_errada"))
	assert.NotEqual(t, "senha", user.Password)
}
