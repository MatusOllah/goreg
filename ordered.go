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
	Key   string `json:"id"`
	Value T      `json:"value"`
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
	r.objs = append(r.objs, kvPair[T]{Key: id, Value: obj})
}

func (r *OrderedRegistry[T]) findIndex(id string) (i int, ok bool) {
	return slices.BinarySearchFunc(r.objs, kvPair[T]{Key: id}, func(a, b kvPair[T]) int {
		return strings.Compare(a.Key, b.Key)
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

	return r.objs[i].Value, ok
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
			if !yield(obj.Key, obj.Value) {
				return
			}
		}
	}
}

// MarshalJSON implements the [encoding/json.Marshaler] interface.
func (r *OrderedRegistry[T]) MarshalJSON() ([]byte, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return json.Marshal(r.objs)
}

// UnmarshalJSON implements the [encoding/json.Unmarshaler] interface.
func (r *OrderedRegistry[T]) UnmarhsalJSON(data []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return json.Unmarshal(data, &r.objs)
}

// GobEncode implements the [encoding/gob.GobEncoder] interface.
func (r *OrderedRegistry[T]) GobEncode() ([]byte, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var bf bytes.Buffer
	if err := gob.NewEncoder(&bf).Encode(r.objs); err != nil {
		return nil, err
	}

	return bf.Bytes(), nil
}

// GobDecode implements the [encoding/gob.GobDecoder] interface.
func (r *OrderedRegistry[T]) GobDecode(data []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return gob.NewDecoder(bytes.NewReader(data)).Decode(&r.objs)
}
