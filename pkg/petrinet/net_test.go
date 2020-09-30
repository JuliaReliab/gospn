package petrinet

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPetri01(t *testing.T) {
	p1 := newPlace("p1", 0, 10)
	t1 := newImmTrans("t1", 0, 0, true, 1.0)
	a1 := newInArc(p1, t1, 1)
	// fmt.Println(p1)
	if a1.src != p1 {
		t.Errorf("fail")
	}
	if a1.dest != t1.getTrans() {
		t.Errorf("fail")
	}
}

func TestPNDo1(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 100)
	p2 := net.NewPlace("p2", 100)
	p3 := net.NewPlace("p3", 100)
	t1 := net.NewExpTrans("t1", 0, true, 1)
	gen1 := NewDistribution("exponential", []float64{1.0})
	t2 := net.NewGenTrans("t2", 0, true, gen1, GenTransPolicyPRD)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.NewInArc(p3, t2, 1)
	net.NewOutArc(t2, p1, 1)
	net.NewOutArc(t2, p2, 1)
	net.Finalize()

	writer := bytes.NewBuffer(make([]byte, 0, 256))
	net.ToPNDot(writer)
	fmt.Println(writer.String())
}

func TestPNDo2(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 100)
	p2 := net.NewPlace("p2", 100)
	p3 := net.NewPlace("p3", 100)
	net.NewPlace("p4", 100)
	t1 := net.NewExpTrans("t1", 0, true, 1)
	gen1 := NewDistribution("exponential", []float64{1.0})
	t2 := net.NewGenTrans("t2", 0, true, gen1, GenTransPolicyPRD)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.NewInArc(p3, t2, 1)
	net.NewOutArc(t2, p1, 1)
	net.NewOutArc(t2, p2, 1)
	net.Finalize()

	writer := bytes.NewBuffer(make([]byte, 0, 256))
	net.ToPNDot(writer)
	fmt.Println(writer.String())
}
