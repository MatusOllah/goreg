package goreg

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"iter"
	"slices"
	"strings"
	"sync"
)

type kvPair[T any] struct {
	key   string
	value T
}

// OrderedRegistry is a ordered registry. It uses a slice under the hood.
type OrderedRegistry[T any] struct {
	objs []kvPair[T]
	mu   sync.RWMutex
}

// NewOrderedRegistry creates a new [OrderedRegistry].
func NewOrderedRegistry[T any]() *OrderedRegistry[T] {
	return &OrderedRegistry[T]{objs: []kvPair[T]{}}
}

// Register registers an object under the ID.
func (r *OrderedRegistry[T]) Register(id string, obj T) {
	if any(obj) == nil {
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.objs = append(r.objs, kvPair[T]{key: id, value: obj})
}

func (r *OrderedRegistry[T]) findIndex(id string) (i int, ok bool) {
	return slices.BinarySearchFunc(r.objs, kvPair[T]{key: id}, func(a, b kvPair[T]) int {
		return strings.Compare(a.key, b.key)
	})
}

// Unregister unregisters an object under the ID.
func (r *OrderedRegistry[T]) Unregister(id string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	i, ok := r.findIndex(id)
	if !ok {
		return
	}

	r.objs = slices.Delete(r.objs, i, i+1)
}

// Get returns the object under the ID.
func (r *OrderedRegistry[T]) Get(id string) (obj T, ok bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	i, ok := r.findIndex(id)
	if !ok {
		return
	}

	return r.objs[i].value, ok
}

// Len returns the number of items in the registry.
func (r *OrderedRegistry[T]) Len() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.objs)
}

// Iter returns an iterator over ID-object (key-value) pairs. See the [iter] package documentation for more details.
func (r *OrderedRegistry[T]) Iter() iter.Seq2[string, T] {
	return func(yield func(string, T) bool) {
		r.mu.Lock()
		defer r.mu.Unlock()

		for _, obj := range r.objs {
			if !yield(obj.key, obj.value) {
				return
			}
		}
	}
}

// MarshalJSON implements the [encoding/json.Marshaler] interface.
func (r *OrderedRegistry[T]) MarshalJSON() ([]byte, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return json.Marshal(Collect(r))
}

// UnmarshalJSON implements the [encoding/json.Unmarshaler] interface.
func (r *OrderedRegistry[T]) UnmarhsalJSON(data []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	m := make(map[string]T)
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	for k, v := range m {
		r.Register(k, v)
	}

	return nil
}

// GobEncode implements the [encoding/gob.GobEncoder] interface.
func (r *OrderedRegistry[T]) GobEncode() ([]byte, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var bf bytes.Buffer
	if err := gob.NewEncoder(&bf).Encode(Collect(r)); err != nil {
		return nil, err
	}

	return bf.Bytes(), nil
}

// GobDecode implements the [encoding/gob.GobDecoder] interface.
func (r *OrderedRegistry[T]) GobDecode(data []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	m := make(map[string]T)
	if err := gob.NewDecoder(bytes.NewReader(data)).Decode(&m); err != nil {
		return err
	}

	for k, v := range m {
		r.Register(k, v)
	}

	return nil
}
