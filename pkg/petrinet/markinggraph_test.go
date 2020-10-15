package petrinet

import (
	"bytes"
	"fmt"
	"testing"
)

func TestMarkDot1(t *testing.T) {
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

	m0 := []MarkInt{10, 1, 1}
	mg := CreateMarkingGraphWithDFS(net, m0)

	writer := bytes.NewBuffer(make([]byte, 0, 256))
	mg.ToMarkDot(writer)
	fmt.Println(writer.String())
}

func TestMarkGraph1(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 100)
	p2 := net.NewPlace("p2", 100)
	p3 := net.NewPlace("p3", 100)
	t1 := net.NewExpTrans("t1", 0, true, 1)
	t2 := net.NewExpTrans("t2", 0, true, 1)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.NewInArc(p3, t2, 1)
	net.NewOutArc(t2, p1, 1)
	net.NewOutArc(t2, p2, 1)
	net.Finalize()

	m0 := []MarkInt{10, 1, 1}
	mg := CreateMarkingGraphWithDFS(net, m0)

	fmt.Println(len(mg.groups))
	fmt.Println(mg.groupTransToLink)
	e, i, g := mg.TransMatrix()
	label1 := mg.GroupLabels()
	label2 := mg.TransLabels()
	for tr, m := range e {
		fmt.Printf("%s%s%s ", label1[tr.src], label1[tr.dest], label2[tr])
		fmt.Println(m)
	}
	for tr, m := range i {
		fmt.Printf("%s%s%s ", label1[tr.src], label1[tr.dest], label2[tr])
		fmt.Println(m)
	}
	for tr, m := range g {
		fmt.Printf("%s%s%s ", label1[tr.src], label1[tr.dest], label2[tr])
		fmt.Println(m)
	}
}

func TestMarkGraph2(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 100)
	p2 := net.NewPlace("p2", 100)
	p3 := net.NewPlace("p3", 100)
	t1 := net.NewExpTrans("t1", 0, true, 1)
	t2 := net.NewImmTrans("t2", 0, true, 1)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.NewInArc(p3, t2, 1)
	net.NewOutArc(t2, p1, 1)
	net.NewOutArc(t2, p2, 1)
	net.Finalize()

	m0 := []MarkInt{10, 1, 1}
	mg := CreateMarkingGraphWithDFS(net, m0)

	fmt.Println(len(mg.groups))
	fmt.Println(mg.groupTransToLink)
	e, i, g := mg.TransMatrix()
	label1 := mg.GroupLabels()
	label2 := mg.TransLabels()
	for tr, m := range e {
		fmt.Printf("%s%s%s ", label1[tr.src], label1[tr.dest], label2[tr])
		fmt.Println(m)
	}
	for tr, m := range i {
		fmt.Printf("%s%s%s ", label1[tr.src], label1[tr.dest], label2[tr])
		fmt.Println(m)
	}
	for tr, m := range g {
		fmt.Printf("%s%s%s ", label1[tr.src], label1[tr.dest], label2[tr])
		fmt.Println(m)
	}
}

func TestMarkGraph3(t *testing.T) {
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

	m0 := []MarkInt{10, 1, 1}
	mg := CreateMarkingGraphWithDFS(net, m0)

	fmt.Println(len(mg.groups))
	fmt.Println(mg.groupTransToLink)
	e, i, g := mg.TransMatrix()
	label1 := mg.GroupLabels()
	label2 := mg.TransLabels()
	for tr, m := range e {
		fmt.Printf("%s%s%s ", label1[tr.src], label1[tr.dest], label2[tr])
		fmt.Println(m)
	}
	for tr, m := range i {
		fmt.Printf("%s%s%s ", label1[tr.src], label1[tr.dest], label2[tr])
		fmt.Println(m)
	}
	for tr, m := range g {
		fmt.Printf("%s%s%s ", label1[tr.src], label1[tr.dest], label2[tr])
		fmt.Println(m)
	}

	writer := bytes.NewBuffer(make([]byte, 0, 256))
	mg.ToGroupMarkDot(writer)
	fmt.Println(writer.String())

	iv := mg.InitVector()
	for g, v := range iv {
		fmt.Printf("%s ", label1[g])
		fmt.Println(v)
	}
}
