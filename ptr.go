package ptr

// New creates a pointer to a value
func New[T any](value T) *T {
	return &value
}

// Value returns the underlying value of a pointer
func Value[T any](ptr *T) T {
	if ptr == nil {
		var ret T
		return ret
	}
	return *ptr
}
