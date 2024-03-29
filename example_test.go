package readonly

import (
	"fmt"
)

// ExampleReadOnly demonstrates the usage of the readOnly struct.
func ExampleReadOnly() {
	// Create a new readOnly instance with an immutable string value
	ro := NewReadOnly("Hello, World!")

	// Get the value from the readOnly instance
	val := ro.Get()

	// Print the value
	fmt.Println(val)

	// Output: Hello, World!
}

// ExampleFinal demonstrates the usage of the final struct.
func ExampleFinal() {
	// Create a new final instance with an integer type
	var f Final[int]

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
