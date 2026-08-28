package petrinet

import (
	"fmt"
	"testing"
)

func TestEnable1(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 10)
	p2 := net.NewPlace("p2", 10)
	p3 := net.NewPlace("p3", 10)
	t1 := net.NewExpTrans("t1", 0, true, 1.0)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.Finalize()
	mark := newMark([]MarkInt{1, 1, 1})
	// fmt.Println(net)
	// fmt.Println(t1.IsEnabled(net, mark))
	if t1.IsEnabled(net, mark.toSlice()) != ENABLE {
		t.Errorf("fail")
	}
}

func TestEnable2(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 10)
	p2 := net.NewPlace("p2", 10)
	p3 := net.NewPlace("p3", 10)
	t1 := net.NewImmTrans("t1", 0, true, 1.0)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.Finalize()
	mark := newMark([]MarkInt{1, 1, 0})
	// fmt.Println(t1.IsEnabled(net, mark))
	if t1.IsEnabled(net, mark.toSlice()) != ENABLE {
		t.Errorf("fail")
	}
	// fmt.Println(t1.doFiring(net, mark).toslice())
}

func TestEnable3(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 10)
	p2 := net.NewPlace("p2", 10)
	p3 := net.NewPlace("p3", 10)
	t1 := net.NewImmTrans("t1", 0, true, 1.0)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.Finalize()
	mark := newMark([]MarkInt{1, 0, 0})
	// fmt.Println(t1.IsEnabled(net, mark))
	if t1.IsEnabled(net, mark.toSlice()) != DISABLE {
		t.Errorf("fail")
	}
}

func TestEnabe4(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 10)
	p2 := net.NewPlace("p2", 10)
	p3 := net.NewPlace("p3", 10)
	t1 := net.NewImmTrans("t1", 0, true, 1.0)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	bb := true
	net.SetGuard(t1, "t1", func([]MarkInt) bool {
		fmt.Println("guard")
		return bb
	})
	net.Finalize()
	mark := newMark([]MarkInt{1, 1, 1})
	bb = false
	if t1.IsEnabled(net, mark.toSlice()) != DISABLE {
		t.Errorf("fail")
	}
}

// A firing that would push a place past its capacity is clamped and reported through
// the returned error. The callers in the marking-graph search used to discard that
// error, so the destination marking was silently wrong and the resulting generator
// matrix had rows that did not sum to zero. Keep the contract pinned here.
func TestDoFiringClampsAtMax(t *testing.T) {
	t.Setenv("GOSPN_NO_WARN", "1")
	net := NewNet()
	p1 := net.NewPlace("p1", 10)
	p2 := net.NewPlace("p2", 3) // capacity below what the firing would produce
	t1 := net.NewExpTrans("t1", 0, true, 1.0)
	net.NewInArc(p1, t1, 1)
	net.NewOutArc(t1, p2, 1)
	net.Finalize()

	mark, err := t1.DoFiring(net, []MarkInt{1, 3})
	if err == nil {
		t.Errorf("expected an error when a firing exceeds the place capacity")
	}
	if mark[1] != 3 {
		t.Errorf("expected p2 to be clamped at its max 3, got %d", mark[1])
	}
}

// The same for the lower bound.
func TestDoFiringClampsAtZero(t *testing.T) {
	t.Setenv("GOSPN_NO_WARN", "1")
	net := NewNet()
	p1 := net.NewPlace("p1", 10)
	t1 := net.NewExpTrans("t1", 0, true, 1.0)
	net.NewInArc(p1, t1, 2)
	net.Finalize()

	mark, err := t1.DoFiring(net, []MarkInt{1})
	if err == nil {
		t.Errorf("expected an error when a firing takes a place below zero")
	}
	if mark[0] != 0 {
		t.Errorf("expected p1 to be clamped at 0, got %d", mark[0])
	}
}
