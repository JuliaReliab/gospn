package petrinet

import (
	// "fmt"
	"testing"
)

func TestMarkingSet(t *testing.T) {
	// v := []MarkInt{1, 2, 3}
	// fmt.Println(v.toarray())
	m := NewMarkGenerator(3)
	m1 := m.genMark([]MarkInt{1, 2, 3})
	m2 := m.genMark([]MarkInt{1, 2, 3})
	if m1 != m2 {
		t.Errorf("fail")
	}
	if m1.toSlice()[0] != 1 {
		t.Errorf("fail2")
	}
}
