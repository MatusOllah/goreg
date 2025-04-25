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

// A GetIndexRegistry is a registry with a GetIndex method.
type GetIndexRegistry[T any] interface {
	Registry[T]

	// GetIndex returns the object under the index.
	GetIndex(i int) (obj T, ok bool)
}

// A MustGetRegistry is a registry with a MustGet method.
type MustGetRegistry[T any] interface {
	Registry[T]

	// MustGet returns the object under the ID and logs error if not found.
	MustGet(id string) T
}

// A MustGetIndexRegistry is a registry with a MustGetIndex method.
type MustGetIndexRegistry[T any] interface {
	Registry[T]

	// MustGetIndex returns the object under the index and logs error if not found.
	MustGetIndex(i int) T
}
