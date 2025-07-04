// Package ptr provides a utility function for creating pointers from literal values.
//
// In Go, it is not possible to take the address of a literal or a constant directly.
// For example, the following code is invalid:
//
//	// Invalid code
//	timeout := &30
//
// To work around this, one would typically need to declare an intermediate variable:
//
//	// Standard Go workaround
//	tmp := 30
//	timeout := &tmp
//
// This package offers a generic helper function, `New`, to streamline this process,
// making code more concise, especially when initializing structs with pointer fields.
package ptr

// New creates a pointer to a value of any type.
//
// It is a generic function that accepts a value `v` and returns a pointer
// to it (*T). This is particularly useful for creating pointers from literals
// (e.g., ptr.New(100), ptr.New("hello"), ptr.New(true)) when initializing
// struct fields that are pointers. Using a pointer field allows distinguishing
// between a zero value (like 0, "", or false) and a value that was not
// specified (nil).
func New[T any](value T) *T {
	return &value
}
