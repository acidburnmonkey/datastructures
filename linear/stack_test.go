package linear

import (
	"testing"
)

func TestPush(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if len(s.Container) != 3 {
		t.Errorf("expected 2 items, got %d", len(s.Container))
	}
}
