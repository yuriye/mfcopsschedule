package entity_test

import (
	"github.com/stretchr/testify/assert"
	"mfcopsschedule/entity"
	"testing"
)

func TestNewEmployee(t *testing.T) {
	u, err := entity.NewEmployee("Марина", "", "Иванова")
	assert.Nil(t, err)
	assert.Equal(t, u.FirstName, "Марина")
	assert.NotNil(t, u.ID)
}
