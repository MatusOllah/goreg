package goreg_test

import (
	"testing"

	"github.com/MatusOllah/goreg"
)

func TestStandardRegistry_Register(t *testing.T) {
	reg := goreg.NewStandardRegistry[int]()

	expected := 42

	reg.Register("horalky", expected)

	obj, ok := reg.Get("horalky")
	if !ok {
		t.Errorf("expected obj=%v but not found", obj)
	}
	if obj != expected {
		t.Errorf("expected obj=%d, got obj=%d", expected, obj)
	}
}

func TestStandardRegistry_Unregister(t *testing.T) {
	reg := goreg.NewStandardRegistry[int]()

	reg.Register("horalky", 42)
	reg.Unregister("horalky")

	if _, ok := reg.Get("horalky"); ok {
		t.Errorf("expected to not get id=horalky but actually did")
	}
}

func TestStandardRegistry_Len(t *testing.T) {
	reg := goreg.NewStandardRegistry[int]()

	m := map[string]int{
		"jeden": 1,
		"dva":   2,
		"tri":   3,
		"styri": 4,
	}

	for k, v := range m {
		reg.Register(k, v)
	}

	expected := len(m)
	regLen := reg.Len()

	if regLen != expected {
		t.Errorf("expected length %d, got %d", expected, regLen)
	}
}

func TestStandardRegistry_Iter(t *testing.T) {
	reg := goreg.NewStandardRegistry[int]()

	m := map[string]int{
		"jeden": 1,
		"dva":   2,
		"tri":   3,
		"styri": 4,
	}

	for k, v := range m {
		reg.Register(k, v)
	}

	i := 0
	for id, obj := range reg.Iter() {
		expected, ok := m[id]
		if !ok {
			t.Errorf("expected id=%s, obj=%v but not found", id, obj)
		}
		if obj != expected {
			t.Errorf("expected id=%s obj=%d, got obj=%d", id, expected, obj)
		}
		i++
	}

	if i != len(m) {
		t.Errorf("expected length %d, got %d", len(m), i)
	}
}
