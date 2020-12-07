package entity_test

import (
	"github.com/stretchr/testify/assert"
	"mfcopsschedule/entity"
	"testing"
)

func TestNewOffice(t *testing.T) {
	o, err := entity.NewOffice("Савченко 23", nil)
	assert.Nil(t, err)
	assert.Equal(t, o.Name, "Савченко 23")
	assert.NotNil(t, o.ID)
	assert.Nil(t, o.DefaultEmployeers)
}
