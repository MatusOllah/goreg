package goreg

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"iter"
	"regexp"
	"sync"
)

// StandardRegistry is a standard registry. It uses a map under the hood.
type StandardRegistry[T any] struct {
	objs     map[string]T
	stringRe *regexp.Regexp
	mu       sync.RWMutex
}

// NewStandardRegistry creates a new [StandardRegistry].
func NewStandardRegistry[T any]() *StandardRegistry[T] {
	return &StandardRegistry[T]{objs: make(map[string]T), stringRe: regexp.MustCompile(`\{.*?\}`)}
}

// Register registers an object under the ID.
func (r *StandardRegistry[T]) Register(id string, obj T) {
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

// Reset wipes the registry.
func (r *StandardRegistry[T]) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.objs = make(map[string]T)
}

// Iter returns an iterator over key-value pairs. See the [iter] package documentation for more details.
//
// Note that you should NOT call any other methods in the for loop. It will cause it to lock.
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

// String returns a string representation of the registry.
func (r *StandardRegistry[T]) String() string {
	return r.stringRe.FindString(fmt.Sprintf("%#v", r.objs))
}

// MarshalJSON implements the [encoding/json.Marshaler] interface.
func (r *StandardRegistry[T]) MarshalJSON() ([]byte, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return json.Marshal(r.objs)
}

// UnmarshalJSON implements the [encoding/json.Unmarshaler] interface.
func (r *StandardRegistry[T]) UnmarshalJSON(data []byte) error {
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
