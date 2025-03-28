package goreg

import (
	"fmt"
	"iter"
)

// Registry represents a generic registry.
type Registry[T any] interface {
	// Register registers an object under the ID.
	Register(id string, obj T)

	// Unregister unregisters an object under the ID.
	Unregister(id string)

	// Get returns the object under the ID.
	Get(id string) (obj T, ok bool)

	// Len returns the number of items in the registry.
	Len() int

	// Reset wipes the registry.
	Reset()

	// Iter returns an iterator over key-value pairs. See the [iter] package documentation for more details.
	Iter() iter.Seq2[string, T]

	fmt.Stringer
}
