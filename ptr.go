package ptr

// New creates a pointer to a value
func New[T any](value T) *T {
	return &value
}
