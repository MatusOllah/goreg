package goreg

import "maps"

// Collect collects key-value pairs from the registry into a new map and returns it.
func Collect[T any](reg Registry[T]) map[string]T {
	return maps.Collect(reg.Iter())
}

// Copy copies all objects in src adding them to dst.
// When a ID in src is already present in dst,
// the value in dst will be overwritten by the value associated
// with the ID in src.
func Copy[T any](dst, src Registry[T]) {
	for id, obj := range src.Iter() {
		dst.Register(id, obj)
	}
}

// Equal reports whether two registries contain the same objects.
// Objects are compared using ==.
func Equal[T comparable](reg1, reg2 Registry[T]) bool {
	if reg1.Len() != reg2.Len() {
		return false
	}

	for id, obj1 := range reg1.Iter() {
		if obj2, ok := reg2.Get(id); !ok || obj1 != obj2 {
			return false
		}
	}

	// Do the same for the other registry
	for id, obj1 := range reg2.Iter() {
		if obj2, ok := reg1.Get(id); !ok || obj1 != obj2 {
			return false
		}
	}
	return true
}

// EqualFunc is like Equal, but compares values using eq.
// Objects are still compared with ==.
func EqualFunc[T1, T2 any](reg1 Registry[T1], reg2 Registry[T2], eq func(T1, T2) bool) bool {
	if reg1.Len() != reg2.Len() {
		return false
	}
	for id, obj1 := range reg1.Iter() {
		if obj2, ok := reg2.Get(id); !ok || !eq(obj1, obj2) {
			return false
		}
	}
	return true
}
