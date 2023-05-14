// Package slices offers various utilities for handling and manipulating
// slice data structures in Go.
package slices

// Flatten return a slice after flattening the input slice by one level.
func Flatten[T any](a [][]T) (b []T) {
	b = []T{}
	for _, e := range a {
		b = append(b, e...)
	}

	return
}
