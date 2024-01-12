package readonly

import "errors"

type readOnly[T any] struct {
	value T
}

// Get returns the value stored in the readOnly struct.
func (r readOnly[T]) Get() T {
	return r.value
}

//  NewReadOnly is a function that creates a new instance of the readOnly struct with the given immutableValue.
func NewReadOnly[T any](immutableValue T) readOnly[T] {
	return readOnly[T]{value: immutableValue}
}

type final[T any] struct {
	assigned bool
	value    T
}

// NewFinalWithValue is a function that creates a new instance of the final struct with the given value.
// It is recommended to use `NewReadOnly` instead.
func NewFinalWithValue[T any](value T) final[T] {
	return final[T]{
		assigned: true,
		value:    value,
	}
}

// NewFinal is a function that creates a new instance of the final struct with the given type.
func NewFinal[T any]() final[T] {
	return final[T]{}
}

// Get returns the value stored in the final struct.
// It returns an error if the value has not been assigned.
func (f final[T]) Get() (T, error) {
	var err error
	if !f.assigned {
		err = errors.New("value is not assigned")
	}
	return f.value, err
}

// MustGet returns the value stored in the final struct.
// It panics with the message "value is not assigned" if the value has not been assigned.
// The method returns the assigned value.
func (f final[T]) MustGet() T {
	if !f.assigned {
		panic("value is not assigned")
	}
	return f.value
}

// Set assigns a new value to the final struct.
func (f *final[T]) Set(value T) error {
	if f.assigned {
		return errors.New("a final value can not be reassigned")
	}
	f.value = value
	f.assigned = true
	return nil
}

// MustSet is a method that assigns a new value to the final struct.
// It panics with the message "a final value can not be reassigned" if the value has already been assigned.
// The method does not return any value.
func (f *final[T]) MustSet(value T) {
	if f.assigned {
		panic("a final value can not be reassigned")
	}
	f.value = value
	f.assigned = true
}
