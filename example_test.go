package readonly

import (
	"fmt"
)

// ExamplereadOnly demonstrates the usage of the readOnly struct.
func ExamplereadOnly() {
	// Create a new readOnly instance with an immutable string value
	ro := NewReadOnly("Hello, World!")

	// Get the value from the readOnly instance
	val := ro.Get()

	// Print the value
	fmt.Println(val)

	// Output: Hello, World!
}

// Examplefinal demonstrates the usage of the final struct.
func Examplefinal() {
	// Create a new final instance with an integer type
	f := NewFinal[int]()

	// Set the value of the final instance
	f.Set(42)

	// The Set method only works the first time it's called.
	f.Set(24)

	// Get the value from the final instance
	val := f.Get()

	// Print the value
	fmt.Println(val)

	// Output: 42
}
