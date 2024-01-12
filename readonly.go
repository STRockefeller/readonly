package readonly

import (
	"sync"
)

type ReadOnly[T any] struct {
	value T
}

// Get returns the value stored in the readOnly struct.
func (r ReadOnly[T]) Get() T {
	return r.value
}

// NewReadOnly creates a new instance of the readOnly struct with the given immutableValue.
func NewReadOnly[T any](immutableValue T) ReadOnly[T] {
	return ReadOnly[T]{value: immutableValue}
}

type Final[T any] struct {
	once  sync.Once
	value T
}

// Get returns the value stored in the final struct.
func (f *Final[T]) Get() T {
	return f.value
}

// Set assigns a new value to the final struct.
// It only works the first time it's called.
func (f *Final[T]) Set(value T) {
	f.once.Do(
		func() {
			f.value = value
		},
	)
}
