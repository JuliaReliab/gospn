package petrinet

import (
	"fmt"
	"testing"
)

func TestMarkingSet(t *testing.T) {
	fmt.Println(toarray([]markInt{1, 2, 3}))
	m := NewMarkMap(3)
	m1 := NewMark([]markInt{1, 2, 3})
	m2 := NewMark([]markInt{1, 2, 3})
	m.Set(m1, m1)
	m.Set(m2, m2)
	if m.Get(m1) != m.Get(m2) {
		t.Errorf("fail")
	}
	if m1.Get(0) != 1 {
		t.Errorf("fail2")
	}
}
