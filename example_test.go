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
	err := f.Set(42)
	if err != nil {
		// Handle the error
	}

	// Get the value from the final instance
	val, err := f.Get()
	if err != nil {
		// Handle the error
	}

	// Print the value
	fmt.Println(val)

	// Output: 42
}
