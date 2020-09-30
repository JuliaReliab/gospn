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
