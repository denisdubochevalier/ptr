package ptr_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/denisdubochevalier/ptr"
)

// TestNew uses a table-driven approach to verify the behavior of the ptr.New function.
func TestNew(t *testing.T) {
	t.Parallel()

	// Define a struct for a custom test case
	type customStruct struct {
		ID   int
		Name string
	}

	// testCases defines the inputs and names for each test scenario.
	testCases := []struct {
		name  string
		input any
	}{
		{
			name:  "with integer",
			input: 123,
		},
		{
			name:  "with string",
			input: "hello world",
		},
		{
			name:  "with boolean",
			input: true,
		},
		{
			name: "with struct",
			input: customStruct{
				ID:   1,
				Name: "Test",
			},
		},
		{
			name:  "with zero value integer",
			input: 0,
		},
		{
			name:  "with empty string",
			input: "",
		},
	}

	// Iterate over each test case
	for _, tc := range testCases {
		// Capture range variable to prevent concurrency issues
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			// Call the generic helper to perform the checks.
			// This avoids the need for a type switch.
			checkResult(is, tc.input)
		})
	}
}

// checkResult is a generic helper function that verifies the behavior of ptr.New.
func checkResult[T any](is *require.Assertions, value T) {
	result := ptr.New(value)

	// 1. Check that the result is not nil
	is.NotNil(result)

	// 2. Check that the dereferenced value matches the input
	is.Equal(value, *result)

	// 3. Check that the pointer is to a new memory location.
	// ptr.New receives a copy of `value`, so the address of the pointer
	// it returns will be different from the address of the `value` variable here.
	is.NotSame(&value, result)
}
