package entity_test

import (
	"github.com/stretchr/testify/assert"
	"mfcopsschedule/entity"
	"testing"
)

func TestNewUser(t *testing.T) {
	u, err := entity.NewUser("sjobs@apple.com", "new_password", "Steve", "", "Jobs")
	assert.Nil(t, err)
	assert.Equal(t, u.FirstName, "Steve")
	assert.NotNil(t, u.ID)
	assert.NotEqual(t, u.Password, "new_password")
}
