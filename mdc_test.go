package over

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMdcAdapter_Set(t *testing.T) {
	ResetGlobalMdcAdapter()
	MDC().Set("key1", "value1")
	MDC().Set("key2", "value2")

	assert.Equal(t, 2, MDC().Count())
}

func TestMdcAdapter_SetMap(t *testing.T) {
	ResetGlobalMdcAdapter()
	var items = map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	MDC().SetMap(items)

	assert.Equal(t, 3, MDC().Count())
}

func TestMdcAdapter_SetIfAbsent(t *testing.T) {
	ResetGlobalMdcAdapter()
	MDC().SetIfAbsent("key1", "value1")

	ok := MDC().SetIfAbsent("key1", "value2")
	assert.False(t, ok)
}

func TestMdcAdapter_Get(t *testing.T) {
	t.Run("get missing key", func(t *testing.T) {
		ResetGlobalMdcAdapter()

		val, ok := MDC().Get("key1")
		assert.False(t, ok)
		assert.Nil(t, val)
	})

	t.Run("get", func(t *testing.T) {
		ResetGlobalMdcAdapter()

		MDC().Set("key1", "value1")
		val, ok := MDC().Get("key1")
		assert.True(t, ok)
		assert.NotNil(t, val)
		assert.Equal(t, "value1", val)
	})
}

func TestMdcAdapter_GetString(t *testing.T) {
	t.Run("getString missing key", func(t *testing.T) {
		ResetGlobalMdcAdapter()

		val := MDC().GetString("key1")
		assert.Equal(t, "", val)
	})

	t.Run("getString", func(t *testing.T) {
		ResetGlobalMdcAdapter()

		MDC().Set("key1", "value1")
		val := MDC().GetString("key1")
		assert.Equal(t, "value1", val)
	})
}

func TestMdcAdapter_Count(t *testing.T) {
	ResetGlobalMdcAdapter()
	count := MDC().Count()
	assert.Equal(t, 0, count)

	var items = map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	MDC().SetMap(items)
	count = MDC().Count()
	assert.Equal(t, 3, count)
}

func TestMdcAdapter_Has(t *testing.T) {
	ResetGlobalMdcAdapter()
	has := MDC().Has("key1")
	assert.False(t, has)
}

func TestMdcAdapter_Remove(t *testing.T) {
	ResetGlobalMdcAdapter()
	MDC().Set("key1", "value1")
	assert.Equal(t, 1, MDC().Count())

	MDC().Remove("key1")
	val, ok := MDC().Get("key1")
	assert.Equal(t, 0, MDC().Count())
	assert.False(t, ok)
	assert.Nil(t, val)
}

func TestMdcAdapter_IsEmpty(t *testing.T) {
	ResetGlobalMdcAdapter()
	assert.True(t, MDC().IsEmpty())
}

func TestMdcAdapter_Keys(t *testing.T) {
	ResetGlobalMdcAdapter()
	MDC().Set("key1", "value1")
	assert.Equal(t, MDC().getUniqueKey("key1"), MDC().Keys()[0])
}