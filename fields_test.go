package over

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetGlobalFields(t *testing.T) {
	fields := []string{"field1", "field2"}
	SetGlobalFields(fields)
	assert.Equal(t, fields, _globalFields)
}

func TestGetGlobalFields(t *testing.T) {
	fields := []string{"field4", "field5"}
	SetGlobalFields(fields)
	assert.Equal(t, fields, _globalFields)

	actual := GetGlobalFields()
	assert.Equal(t, _globalFields, actual)
}

func TestAddGlobalFields(t *testing.T) {
	AddGlobalFields("field3")
	assert.Contains(t, _globalFields, "field3")
}

func TestClearGlobalFields(t *testing.T) {
	fields := []string{"field4", "field5"}
	SetGlobalFields(fields)
	assert.Equal(t, fields, _globalFields)

	ClearGlobalFields()
	assert.Empty(t, _globalFields)
}