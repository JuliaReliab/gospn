package petrinet

import (
	"bytes"
	"strings"
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

// pnDotNet is p1,p2 -> t1 -> p3 -> t2 -> p1,p2, optionally with a place nothing
// connects to -- ToPNDot draws each connected component as its own subgraph, so an
// isolated place is a second cluster.
func pnDotNet(isolated bool) *Net {
	net := NewNet()
	p1 := net.NewPlace("p1", 100)
	p2 := net.NewPlace("p2", 100)
	p3 := net.NewPlace("p3", 100)
	if isolated {
		net.NewPlace("p4", 100)
	}
	t1 := net.NewExpTrans("t1", 0, true, 1)
	t2 := net.NewGenTrans("t2", 0, true, NewDistribution("exponential", 1.0), GenTransPolicyPRD)
	net.NewInArc(p1, t1, 1)
	net.NewInArc(p2, t1, 1)
	net.NewOutArc(t1, p3, 1)
	net.NewInArc(p3, t2, 1)
	net.NewOutArc(t2, p1, 1)
	net.NewOutArc(t2, p2, 1)
	net.Finalize()
	return net
}

func pnDot(net *Net) string {
	var buf bytes.Buffer
	net.ToPNDot(&buf)
	return buf.String()
}

// ToPNDot used to name its nodes by their addresses, so drawing the same net twice gave
// two different files. These tests printed the result and checked none of it.
func TestPNDotIsDeterministicAndNamesNodesByLabel(t *testing.T) {
	for _, isolated := range []bool{false, true} {
		isolated := isolated
		name := "connected"
		if isolated {
			name = "with an isolated place"
		}
		t.Run(name, func(t *testing.T) {
			got := pnDot(pnDotNet(isolated))
			if again := pnDot(pnDotNet(isolated)); got != again {
				t.Error("two drawings of the same net differ")
			}
			if strings.Contains(got, "0x") {
				t.Errorf("a node is named by its address:\n%s", got)
			}
			for _, want := range []string{`"p_p1" [shape=circle`, `"t_t1" [shape=box`, `"p_p1"->"t_t1"`} {
				if !strings.Contains(got, want) {
					t.Errorf("missing %s in:\n%s", want, got)
				}
			}
			// An isolated place is drawn, in a cluster of its own.
			if strings.Contains(got, `"p_p4"`) != isolated {
				t.Errorf("p4 present = %v, want %v", strings.Contains(got, `"p_p4"`), isolated)
			}
			if want := 2; isolated && strings.Count(got, "subgraph cluster") != want {
				t.Errorf("%d clusters, want %d", strings.Count(got, "subgraph cluster"), want)
			}
		})
	}
}
