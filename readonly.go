package readonly

import (
	"sync"
)

type readOnly[T any] struct {
	value T
}

// Get returns the value stored in the readOnly struct.
func (r readOnly[T]) Get() T {
	return r.value
}

// NewReadOnly creates a new instance of the readOnly struct with the given immutableValue.
func NewReadOnly[T any](immutableValue T) readOnly[T] {
	return readOnly[T]{value: immutableValue}
}

type final[T any] struct {
	once  sync.Once
	value T
}

// NewFinal creates a new instance of the final struct with the given type.
func NewFinal[T any]() final[T] {
	return final[T]{}
}

// Get returns the value stored in the final struct.
func (f *final[T]) Get() T {
	return f.value
}

// Set assigns a new value to the final struct.
// It only works the first time it's called.
func (f *final[T]) Set(value T) {
	f.once.Do(
		func() {
			f.value = value
		},
	)
}
