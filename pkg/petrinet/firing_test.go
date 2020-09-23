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
	t1 := net.NewTrans("t1", 0, true)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.Indexing()
	mark := NewMark([]markInt{1,1,1})
	fmt.Println(IsEnabled(net, t1, mark))
	if IsEnabled(net, t1, mark) != ENABLE {
	    t.Errorf("fail")
	}
}

func TestEnable2(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 10)
	p2 := net.NewPlace("p2", 10)
	p3 := net.NewPlace("p3", 10)
	t1 := net.NewTrans("t1", 0, true)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.Indexing()
	mark := NewMark([]markInt{1,1,0})
	fmt.Println(IsEnabled(net, t1, mark))
	if IsEnabled(net, t1, mark) != ENABLE {
	    t.Errorf("fail")
	}
	fmt.Println(toslice(DoFiring(net, t1, mark)))
}

func TestEnable3(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 10)
	p2 := net.NewPlace("p2", 10)
	p3 := net.NewPlace("p3", 10)
	t1 := net.NewTrans("t1", 0, true)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.Indexing()
	mark := NewMark([]markInt{1,0,0})
	fmt.Println(IsEnabled(net, t1, mark))
	if IsEnabled(net, t1, mark) != DISABLE {
	    t.Errorf("fail")
	}
}

func TestEnable4(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 10)
	p2 := net.NewPlace("p2", 10)
	p3 := net.NewPlace("p3", 10)
	t1 := net.NewTrans("t1", 0, true)
	a1 := net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.Indexing()
	net.infunc[a1] = func(m *Mark) markInt { return 10 }
	mark := NewMark([]markInt{10,1,0})
	fmt.Println(IsEnabled(net, t1, mark))
	if IsEnabled(net, t1, mark) != ENABLE {
	    t.Errorf("fail")
	}
	fmt.Println(toslice(DoFiring(net, t1, mark)))
}
