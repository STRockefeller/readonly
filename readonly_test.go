package readonly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadOnlyGet(t *testing.T) {
	const value = "test"
	r := NewReadOnly(value)

	got := r.Get()
	assert.Equal(t, value, got)
}

func TestFinalSetAndGet(t *testing.T) {
	const value = "test"
	var f Final[string]
	f.Set(value)

	got := f.Get()
	assert.Equal(t, value, got)
}

func TestFinalSetTwice(t *testing.T) {
	const value = "test"
	const value2 = "test2"
	var f Final[string]

	f.Set(value)
	f.Set(value2)

	got := f.Get()
	assert.Equal(t, value, got)
}
