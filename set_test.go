package immutable_test

import (
	"fmt"
	"testing"

	"github.com/fishy/go-immutable"
)

func TestSetBuilder(t *testing.T) {
	s := immutable.SetLiteral("a", "b", "c")
	size := s.Len()
	if size != 3 {
		t.Errorf("Len() expected 3, got %d", size)
	}

	var item immutable.Hashable = "b"
	var ok bool
	ok = s.Contains(item)
	if !ok {
		t.Errorf("%v should be in the set", item)
	}
	item = "z"
	ok = s.Contains(item)
	if ok {
		t.Errorf("%v should not be in the set", item)
	}

	saw := immutable.NewSetBuilder()
	if err := s.Range(func(x immutable.Hashable) error {
		if saw.Contains(x) {
			t.Errorf("Already iterated %v", x)
		}
		saw.Add(x)
		return nil
	}); err != nil {
		t.Errorf("Range() should return nil, got: %v", err)
	}
	size = saw.Len()
	if size != 3 {
		t.Errorf("Should iterated 3 items, got %d", size)
	}
}

func TestSetString(t *testing.T) {
	target1 := "[a b]"
	target2 := "[b a]"
	s := immutable.SetLiteral("a", "b")
	setStr := fmt.Sprintf("%v", s)
	if setStr != target1 && setStr != target2 {
		t.Errorf(
			"Set.String() expected either %q or %q, got %q",
			target1,
			target2,
			setStr,
		)
	}
}
