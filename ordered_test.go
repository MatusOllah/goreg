package goreg_test

import (
	"testing"

	"github.com/MatusOllah/goreg"
)

func TestOrderedRegistry_RegisterAndGet(t *testing.T) {
	reg := goreg.NewOrderedRegistry[int]()
	reg.Register("kajsmentke", 42)
	reg.Register("kozmeker", 69)

	if val, ok := reg.Get("kajsmentke"); !ok || val != 42 {
		t.Errorf("expected 42, got %v", val)
	}
	if val, ok := reg.Get("kozmeker"); !ok || val != 69 {
		t.Errorf("expected 69, got %v", val)
	}
}

func TestOrderedRegistry_GetInvalid(t *testing.T) {
	reg := goreg.NewOrderedRegistry[int]()

	if _, ok := reg.Get("invalid"); ok {
		t.Error("expected key to be not found")
	}
}

func TestOrderedRegistry_Unregister(t *testing.T) {
	reg := goreg.NewOrderedRegistry[int]()
	reg.Register("kajsmentke", 42)
	reg.Register("kozmeker", 69)

	reg.Unregister("kajsmentke")

	if _, ok := reg.Get("kajsmentke"); ok {
		t.Error("expected key kajsmentke to be not found")
	}

	if reg.Len() != 1 {
		t.Errorf("expected length 1, got %d", reg.Len())
	}
}

func TestOrderedRegistry_Len(t *testing.T) {
	reg := goreg.NewOrderedRegistry[int]()

	if reg.Len() != 0 {
		t.Errorf("Expected length 0, got %d", reg.Len())
	}

	reg.Register("kajsmentke", 42)
	if reg.Len() != 1 {
		t.Errorf("Expected length 1, got %d", reg.Len())
	}

	reg.Register("kozmeker", 69)
	if reg.Len() != 2 {
		t.Errorf("Expected length 2, got %d", reg.Len())
	}

	reg.Unregister("kajsmentke")
	if reg.Len() != 1 {
		t.Errorf("expected length 1, got %d", reg.Len())
	}
}

func TestOrderedRegistry_Iter(t *testing.T) {
	reg := goreg.NewOrderedRegistry[int]()
	reg.Register("kajsmentke", 42)
	reg.Register("kozmeker", 69)

	values := map[string]int{}
	reg.Iter()(func(_ string, _ int) bool {
		return false
	})
	reg.Iter()(func(k string, v int) bool {
		values[k] = v
		return true
	})

	if values["kajsmentke"] != 42 {
		t.Errorf("expected 42, got %d", values["kajsmentke"])
	}
	if values["kozmeker"] != 69 {
		t.Errorf("expected 69, got %d", values["kozmeker"])
	}

	if len(values) != reg.Len() {
		t.Errorf("expected length %d, got %d", reg.Len(), len(values))
	}
}

func TestOrderedRegistry_JSONCodec(t *testing.T) {
	reg := goreg.NewOrderedRegistry[int]()
	reg.Register("kajsmentke", 42)
	reg.Register("kozmeker", 69)

	data, err := reg.MarshalJSON()
	if err != nil {
		t.Errorf("failed to marshal JSON: %v", err)
	}

	newReg := goreg.NewOrderedRegistry[int]()
	if err := newReg.UnmarhsalJSON(data); err != nil {
		t.Errorf("failed to unmarshal JSON: %v", err)
	}

	if val, ok := reg.Get("kajsmentke"); !ok || val != 42 {
		t.Errorf("expected 42, got %v", val)
	}
	if val, ok := reg.Get("kozmeker"); !ok || val != 69 {
		t.Errorf("expected 69, got %v", val)
	}
}

func TestOrderedRegistry_GobCodec(t *testing.T) {
	reg := goreg.NewOrderedRegistry[int]()
	reg.Register("kajsmentke", 42)
	reg.Register("kozmeker", 69)

	data, err := reg.GobEncode()
	if err != nil {
		t.Errorf("failed to encode gob: %v", err)
	}

	newReg := goreg.NewOrderedRegistry[int]()
	if err := newReg.GobDecode(data); err != nil {
		t.Errorf("failed to decode gob: %v", err)
	}

	if val, ok := reg.Get("kajsmentke"); !ok || val != 42 {
		t.Errorf("expected 42, got %v", val)
	}
	if val, ok := reg.Get("kozmeker"); !ok || val != 69 {
		t.Errorf("expected 69, got %v", val)
	}
}
