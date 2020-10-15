package petrinet

import (
	// "fmt"
	"testing"
)

func TestGenVecSet(t *testing.T) {
	// v := TransStatusSlice{ENABLE, ENABLE, ENABLE}
	// fmt.Println(v.toarray())
	m := NewGenVecGenerator(3)
	m1 := m.genGenVec([]TransStatus{ENABLE, DISABLE, PREEMPTION})
	m2 := m.genGenVec([]TransStatus{ENABLE, DISABLE, PREEMPTION})
	if m1 != m2 {
		t.Errorf("fail")
	}
	if m1.toSlice()[0] != ENABLE {
		t.Errorf("fail2")
	}
}
