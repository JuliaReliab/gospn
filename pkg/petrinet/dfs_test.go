package petrinet

import (
	"fmt"
	"testing"
)

func (m1 *Mark) equals(m2 *Mark) bool {
	if m1.n != m2.n {
		return false
	}
	s1 := m1.toSlice()
	s2 := m2.toSlice()
	for i := 0; i < m1.n; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestMarkStack(t *testing.T) {
	stack := NewMarkStack()
	m1 := newMark([]MarkInt{1, 2, 3})
	m2 := newMark([]MarkInt{1, 0, 1})
	// fmt.Println(m1)
	// fmt.Println(m2)
	// fmt.Println(&stack)
	stack.push(m1)
	stack.push(m2)
	// fmt.Println(&stack)
	m3 := stack.pop()
	// fmt.Println(m3)
	if m3 != m2 {
		t.Errorf("fail")
	}
}

func TestMarkSet(t *testing.T) {
	set := NewMarkSet()
	m1 := newMark([]MarkInt{1, 2, 3})
	m2 := newMark([]MarkInt{1, 0, 1})
	// fmt.Println(m1)
	// fmt.Println(m2)
	// fmt.Println(&set)
	set.add(m1)
	// fmt.Println(&set)
	if set.exist(m1) != true {
		t.Errorf("fail")
	}
	if set.exist(m2) != false {
		t.Errorf("fail")
	}
}

func TestMarkMap(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 10)
	p2 := net.NewPlace("p2", 10)
	p3 := net.NewPlace("p3", 10)
	t1 := net.NewExpTrans("t1", 0, true, 1.0)
	t2 := net.NewExpTrans("t2", 0, true, 1.0)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.NewInArc(p3, t2, 1)
	net.NewOutArc(t2, p1, 1)
	net.Finalize()

	d := new(dfs)
	mm := NewMarkGenerator(len(net.placelist))
	m0 := mm.genMark([]MarkInt{4, 4, 6})
	s1 := fmt.Sprintf("%p", m0)
	m1 := mm.genMark([]MarkInt{5, 5, 5})
	d.markGenerator = mm
	dest, _ := d.createNextMarking(net, m1, t1)
	s2 := fmt.Sprintf("%p", dest)
	if s1 != s2 {
		t.Errorf("fail %s %s\n", s1, s2)
	}
}

func TestGenVecMap1(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 10)
	p2 := net.NewPlace("p2", 10)
	p3 := net.NewPlace("p3", 10)
	gen1 := NewDistribution("exponential", 1.0)
	gen2 := NewDistribution("exponential", 1.0)
	t1 := net.NewGenTrans("t1", 0, true, gen1, GenTransPolicyPRI)
	t2 := net.NewGenTrans("t2", 0, true, gen2, GenTransPolicyPRI)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.NewInhibitArc(p3, t2, 1)
	net.NewOutArc(t2, p1, 1)
	net.Finalize()

	d := new(dfs)
	gm := NewGenVecGenerator(len(net.genlist))
	g0 := gm.genGenVec([]TransStatus{ENABLE, PREEMPTION})
	m0 := newMark([]MarkInt{1, 1, 1})
	s1 := fmt.Sprintf("%p", g0)
	d.genVecGenerator = gm
	gv := d.createGenVec(net, m0)
	s2 := fmt.Sprintf("%p", gv)
	if s1 != s2 {
		fmt.Println(g0, gv)
		t.Errorf("fail %s %s\n", s1, s2)
	}
}

func TestGenVecMap2(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 10)
	p2 := net.NewPlace("p2", 10)
	p3 := net.NewPlace("p3", 10)
	t1 := net.NewExpTrans("t1", 0, true, 1.0)
	t2 := net.NewExpTrans("t2", 0, true, 1.0)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.NewInhibitArc(p3, t2, 1)
	net.NewOutArc(t2, p1, 1)
	net.Finalize()

	d := new(dfs)
	gm := NewGenVecGenerator(len(net.genlist))
	g0 := gm.genGenVec([]TransStatus{})
	d.genVecGenerator = gm
	m0 := newMark([]MarkInt{1, 1, 1})
	s1 := fmt.Sprintf("%p", g0)
	gv := d.createGenVec(net, m0)
	fmt.Println(gv)
	s2 := fmt.Sprintf("%p", gv)
	if s1 != s2 {
		t.Errorf("fail %s %s\n", s1, s2)
	}
}

func TestVisitIMM(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 10)
	p2 := net.NewPlace("p2", 10)
	p3 := net.NewPlace("p3", 10)
	t1 := net.NewImmTrans("t1", 0, true, 1)
	t2 := net.NewImmTrans("t2", 0, true, 1)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.NewInArc(p3, t2, 1)
	net.NewOutArc(t2, p1, 1)
	net.NewOutArc(t2, p2, 1)
	net.Finalize()

	dfs := new(dfs)
	dfs.markGenerator = NewMarkGenerator(len(net.placelist))
	dfs.novisited = NewMarkStack()
	dfs.markToGenvec = make(map[*Mark]*GenVec)
	dfs.markToGroupType = make(map[*Mark]GroupType)
	dfs.links = make([]Link, 0)

	m0 := dfs.markGenerator.genMark([]MarkInt{1, 1, 0})
	dfs.visitImmMark(net, m0)
	// fmt.Println(dfs.novisited)
	m1 := dfs.markGenerator.genMark([]MarkInt{0, 0, 1})
	if m2 := dfs.novisited.pop(); !m1.equals(m2) {
		fmt.Println(m1.toSlice())
		fmt.Println(m2.toSlice())
		t.Errorf("fail")
	}
}

func TestVisitIMM2(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 10)
	p2 := net.NewPlace("p2", 10)
	p3 := net.NewPlace("p3", 10)
	t1 := net.NewImmTrans("t1", 0, true, 1)
	t2 := net.NewImmTrans("t2", 0, true, 1)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.NewInArc(p3, t2, 1)
	net.NewOutArc(t2, p1, 1)
	net.NewOutArc(t2, p2, 1)
	net.Finalize()

	dfs := new(dfs)
	dfs.markGenerator = NewMarkGenerator(len(net.placelist))
	dfs.novisited = NewMarkStack()
	dfs.markToGenvec = make(map[*Mark]*GenVec)
	dfs.markToGroupType = make(map[*Mark]GroupType)
	dfs.links = make([]Link, 0)

	m0 := dfs.markGenerator.genMark([]MarkInt{1, 1, 1})
	dfs.visitImmMark(net, m0)
	// fmt.Println(dfs.novisited)
	// fmt.Println(dfs.links)
	m1 := dfs.markGenerator.genMark([]MarkInt{2, 2, 0})
	if m2 := dfs.novisited.pop(); !m1.equals(m2) {
		fmt.Println(m1)
		fmt.Println(m2)
		t.Errorf("fail")
	}
	m3 := dfs.markGenerator.genMark([]MarkInt{0, 0, 2})
	if m4 := dfs.novisited.pop(); !m3.equals(m4) {
		fmt.Println(m3)
		fmt.Println(m4)
		t.Errorf("fail")
	}
}

func TestVisitGEN1(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 10)
	p2 := net.NewPlace("p2", 10)
	p3 := net.NewPlace("p3", 10)
	t1 := net.NewExpTrans("t1", 0, true, 1)
	gen1 := NewDistribution("exponential", 1.0)
	t2 := net.NewGenTrans("t2", 0, true, gen1, GenTransPolicyPRD)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.NewInArc(p3, t2, 1)
	net.NewOutArc(t2, p1, 1)
	net.NewOutArc(t2, p2, 1)
	net.Finalize()

	dfs := new(dfs)
	dfs.markGenerator = NewMarkGenerator(len(net.placelist))
	dfs.novisited = NewMarkStack()
	dfs.markToGenvec = make(map[*Mark]*GenVec)
	dfs.markToGroupType = make(map[*Mark]GroupType)
	dfs.links = make([]Link, 0)

	m0 := dfs.markGenerator.genMark([]MarkInt{1, 1, 1})
	dfs.visitGenMark(net, m0)
	// fmt.Println(dfs.novisited)
	// fmt.Println(dfs.links)
	m1 := dfs.markGenerator.genMark([]MarkInt{0, 0, 2})
	if m2 := dfs.novisited.pop(); !m1.equals(m2) {
		fmt.Println(m1)
		fmt.Println(m2)
		t.Errorf("fail")
	}
	m3 := dfs.markGenerator.genMark([]MarkInt{2, 2, 0})
	if m4 := dfs.novisited.pop(); !m3.equals(m4) {
		fmt.Println(m3)
		fmt.Println(m4)
		t.Errorf("fail")
	}
}

func TestCreateMarking1(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 100)
	p2 := net.NewPlace("p2", 100)
	p3 := net.NewPlace("p3", 100)
	t1 := net.NewExpTrans("t1", 0, true, 1)
	gen1 := NewDistribution("exponential", 1.0)
	t2 := net.NewGenTrans("t2", 0, true, gen1, GenTransPolicyPRD)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.NewInArc(p3, t2, 1)
	net.NewOutArc(t2, p1, 1)
	net.NewOutArc(t2, p2, 1)
	net.Finalize()

	dfs := new(dfs)
	dfs.markGenerator = NewMarkGenerator(len(net.placelist))
	dfs.genVecGenerator = NewGenVecGenerator(len(net.genlist))
	dfs.visited = NewMarkSet()
	dfs.novisited = NewMarkStack()
	dfs.markToGenvec = make(map[*Mark]*GenVec)
	dfs.markToGroupType = make(map[*Mark]GroupType)
	dfs.links = make([]Link, 0)

	m0 := dfs.markGenerator.genMark([]MarkInt{10, 1, 1})
	// dfs.markGenerator.Set(m0, m0)
	dfs.novisited.push(m0)
	dfs.createMarking(net)
	fmt.Println(dfs.visited)
	fmt.Println(dfs.markToGenvec)
	fmt.Println(dfs.links)
}

func TestCreateMarking2(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 100)
	p2 := net.NewPlace("p2", 100)
	p3 := net.NewPlace("p3", 100)
	t1 := net.NewImmTrans("t1", 0, true, 1)
	gen1 := NewDistribution("exponential", 1.0)
	t2 := net.NewGenTrans("t2", 0, true, gen1, GenTransPolicyPRD)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.NewInArc(p3, t2, 1)
	net.NewOutArc(t2, p1, 1)
	net.NewOutArc(t2, p2, 1)
	net.Finalize()

	dfs := new(dfs)
	dfs.markGenerator = NewMarkGenerator(len(net.placelist))
	dfs.genVecGenerator = NewGenVecGenerator(len(net.genlist))
	dfs.visited = NewMarkSet()
	dfs.novisited = NewMarkStack()
	dfs.markToGenvec = make(map[*Mark]*GenVec)
	dfs.markToGroupType = make(map[*Mark]GroupType)
	dfs.links = make([]Link, 0)

	m0 := dfs.markGenerator.genMark([]MarkInt{10, 1, 1})
	// dfs.markGenerator.Set(m0, m0)
	dfs.novisited.push(m0)
	dfs.createMarking(net)
	fmt.Println(dfs.visited)
	fmt.Println(dfs.markToGenvec)
	fmt.Println(dfs.links)
}

func TestCreateMarking3(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 100)
	p2 := net.NewPlace("p2", 100)
	p3 := net.NewPlace("p3", 100)
	t1 := net.NewExpTrans("t1", 0, true, 1)
	gen1 := NewDistribution("exponential", 1.0)
	t2 := net.NewGenTrans("t2", 0, true, gen1, GenTransPolicyPRD)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.NewInArc(p3, t2, 1)
	net.NewOutArc(t2, p1, 1)
	net.NewOutArc(t2, p2, 1)
	net.Finalize()

	m0 := []MarkInt{10, 1, 1}
	mg := CreateMarkingGraph(net, m0, new(dfs))
	fmt.Println(mg)
}

// markGenerator   *MarkMap
// genVecGenerator *GenVecMap
// novisited      markStack
// visited        markSet
// markToGenvec   map[*Mark]*GenVec     // map from Mark to GenVec
// genVecType     map[*GenVec]TransType // the genvec type
// links          []Link                // links
