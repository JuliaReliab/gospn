package petrinet

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func (e event) String(net *Net) string {
	str := make([]string, 0)
	for i, n := range e.mark {
		if n > 0 {
			str = append(str, fmt.Sprintf("%s:%d", net.placelist[i].label, n))
		}
	}
	return fmt.Sprintf("%.4f {%s}", e.time, strings.Join(str, ","))
}

func TestSim1(t *testing.T) {
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
	sim := PNSimulation{
		net:        net,
		endingtime: 100,
	}
	s := rand.NewSource(1)
	result := sim.runSimulation(m0, rand.New(s))
	for i, x := range result {
		fmt.Println(i, x.String(net))
	}
}

func TestSim2(t *testing.T) {
	net, m0 := buildRaid6()
	sim := PNSimulation{
		net:        net,
		endingtime: 1000.0,
	}
	s := rand.NewSource(1)
	result := sim.runSimulation(m0, rand.New(s))
	for i, x := range result {
		fmt.Println(i, x.String(net))
	}
}
