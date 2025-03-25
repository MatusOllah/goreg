package goreg

import "maps"

// Collect collects ID-object (key-value) pairs from the registry into a new map and returns it.
func Collect[T any](reg Registry[T]) map[string]T {
	return maps.Collect(reg.Iter())
}
