package goreg_test

import (
	"testing"

	"github.com/MatusOllah/goreg"
)

func TestClone(t *testing.T) {
	reg := goreg.NewStandardRegistry[int]()
	reg.Register("kajsmentke", 42)
	reg.Register("kozmeker", 69)

	cloned := goreg.Clone(reg)

	if cloned.Len() != reg.Len() {
		t.Errorf("expected length %d, got %d", reg.Len(), cloned.Len())
	}
}

func TestCollect(t *testing.T) {
	reg := goreg.NewStandardRegistry[int]()
	reg.Register("kajsmentke", 42)
	reg.Register("kozmeker", 69)

	collected := goreg.Collect(reg)
	if len(collected) != 2 {
		t.Errorf("expected map length 2, got %d", len(collected))
	}
}

func TestCopy(t *testing.T) {
	dst := goreg.NewStandardRegistry[int]()
	dst.Register("kajsmentke", 42)

	src := goreg.NewStandardRegistry[int]()
	src.Register("kozmeker", 69)

	goreg.Copy(dst, src)

	if dst.Len() != 2 {
		t.Errorf("expected dst length 2, got %d", dst.Len())
	}
}

func TestUnregisterFunc(t *testing.T) {
	reg := goreg.NewStandardRegistry[int]()
	reg.Register("kajsmentke", 42)
	reg.Register("kozmeker", 69)

	goreg.UnregisterFunc(reg, func(k string, v int) bool {
		return k == "kajsmentke"
	})

	if reg.Len() != 1 {
		t.Errorf("expected length 1, got %d", reg.Len())
	}
}

func TestEqual(t *testing.T) {
	reg1 := goreg.NewStandardRegistry[int]()
	reg2 := goreg.NewStandardRegistry[int]()
	reg1.Register("kajsmentke", 42)
	reg2.Register("kajsmentke", 42)

	if !goreg.Equal(reg1, reg2) {
		t.Error("expected equal, but not equal")
	}

	reg2.Unregister("kajsmentke")
	if goreg.Equal(reg1, reg2) {
		t.Error("expected not equal, but actually equal")
	}

	reg2.Register("kozmeker", 69)
	if goreg.Equal(reg1, reg2) {
		t.Error("expected not equal, but actually equal")
	}
}

func TestEqualFunc(t *testing.T) {
	reg1 := goreg.NewStandardRegistry[int]()
	reg2 := goreg.NewStandardRegistry[float64]()
	reg1.Register("a", 3)
	reg2.Register("a", 3.0)

	eq := func(a int, b float64) bool {
		return float64(a) == b
	}

	if !goreg.EqualFunc(reg1, reg2, eq) {
		t.Error("expected equal, but not equal")
	}

	reg2.Unregister("a")
	if goreg.EqualFunc(reg1, reg2, eq) {
		t.Error("expected not equal, but actually equal")
	}

	reg2.Register("b", 4.5)
	if goreg.EqualFunc(reg1, reg2, eq) {
		t.Error("expected not equal, but actually equal")
	}
}
