package goreg

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"iter"
	"sync"
)

// StandardRegistry is a standard registry. It uses a map under the hood.
type StandardRegistry[T any] struct {
	objs map[string]T
	mu   sync.RWMutex
}

// NewStandardRegistry creates a new [StandardRegistry].
func NewStandardRegistry[T any]() *StandardRegistry[T] {
	return &StandardRegistry[T]{objs: make(map[string]T)}
}

// Register registers an object under the ID.
func (r *StandardRegistry[T]) Register(id string, obj T) {
	if any(obj) == nil {
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.objs[id] = obj
}

// Unregister unregisters an object under the ID.
func (r *StandardRegistry[T]) Unregister(id string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.objs, id)
}

// Get returns the object under the ID.
func (r *StandardRegistry[T]) Get(id string) (obj T, ok bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	obj, ok = r.objs[id]
	return
}

// Len returns the number of items in the registry.
func (r *StandardRegistry[T]) Len() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.objs)
}

// Iter returns an iterator over ID-object (key-value) pairs. See the [iter] package documentation for more details.
func (r *StandardRegistry[T]) Iter() iter.Seq2[string, T] {
	return func(yield func(string, T) bool) {
		r.mu.Lock()
		defer r.mu.Unlock()

		for id, obj := range r.objs {
			if !yield(id, obj) {
				return
			}
		}
	}
}

// MarshalJSON implements the [encoding/json.Marshaler] interface.
func (r *StandardRegistry[T]) MarshalJSON() ([]byte, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return json.Marshal(r.objs)
}

// UnmarshalJSON implements the [encoding/json.Unmarshaler] interface.
func (r *StandardRegistry[T]) UnmarhsalJSON(data []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return json.Unmarshal(data, &r.objs)
}

// GobEncode implements the [encoding/gob.GobEncoder] interface.
func (r *StandardRegistry[T]) GobEncode() ([]byte, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var bf bytes.Buffer
	if err := gob.NewEncoder(&bf).Encode(r.objs); err != nil {
		return nil, err
	}

	return bf.Bytes(), nil
}

// GobDecode implements the [encoding/gob.GobDecoder] interface.
func (r *StandardRegistry[T]) GobDecode(data []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return gob.NewDecoder(bytes.NewReader(data)).Decode(&r.objs)
}
