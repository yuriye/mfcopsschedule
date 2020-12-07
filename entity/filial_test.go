package entity_test

import (
	"github.com/stretchr/testify/assert"
	"mfcopsschedule/entity"
	"testing"
)

func TestNewFilial(t *testing.T) {
	f, err := entity.NewFilial("Петропавловск-Камчатский", nil, nil)
	assert.Nil(t, err)
	assert.NotNil(t, f.ID)
	assert.Equal(t, f.Name, "Петропавловск-Камчатский")
	assert.Nil(t, f.Offices)
	assert.Nil(t, f.Employees)
}
