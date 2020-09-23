package petrinet

import (
	"fmt"
	"testing"
)

func TestPetriNode(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 10)
	t1 := net.NewTrans("p1", 0, true)
	fmt.Printf("p1 %p\n", p1)
	fmt.Printf("t1 %p\n", t1)
	net.NewInArc(p1, t1, 1)
	net.NewOutArc(t1, p1, 1)
	if net.places[0] != p1 {
		t.Errorf("failure: p1 is not assigned")
	}
	// if a1 != p1.outarcs[0] {
	//     t.Errorf("a1 fails p1:%p t1:%p p1.node%p", p1, t1, &p1.node)
	// }
}
