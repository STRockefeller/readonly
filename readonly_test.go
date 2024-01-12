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

func TestFinalGet(t *testing.T) {
	const value = "test"
	f := NewFinalWithValue(value)

	got, err := f.Get()
	assert.NoError(t, err)
	assert.Equal(t, value, got)
}

func TestFinalGet_BadCase(t *testing.T) {
	f := NewFinal[string]()

	_, err := f.Get()
	assert.Error(t, err)
}

func TestFinalMustGet(t *testing.T) {
	const value = "test"
	f := NewFinalWithValue(value)

	got := f.MustGet()
	assert.Equal(t, value, got)
}

func TestFinalMustGet_BadCase(t *testing.T) {
	f := NewFinal[string]()

	assert.Panics(t, func() { f.MustGet() })
}

func TestFinalSet(t *testing.T) {
	const value = "test"
	f := NewFinal[string]()

	err := f.Set(value)
	assert.NoError(t, err)

	got, err := f.Get()
	assert.NoError(t, err)
	assert.Equal(t, value, got)
}

func TestFinalSet_BadCase(t *testing.T) {
	const value = "test"
	f := NewFinal[string]()

	assert.NoError(t, f.Set(value))
	assert.Error(t, f.Set(value))
}

func TestFinalMustSet(t *testing.T) {
	const value = "test"
	f := NewFinal[string]()

	assert.NotPanics(t, func() { f.MustSet(value) })

	got, err := f.Get()
	assert.NoError(t, err)
	assert.Equal(t, value, got)
}

func TestFinalMustSet_BadCase(t *testing.T) {
	const value = "test"
	f := NewFinal[string]()

	assert.NotPanics(t, func() { f.MustSet(value) })
	assert.Panics(t, func() { f.MustSet(value) })
}
