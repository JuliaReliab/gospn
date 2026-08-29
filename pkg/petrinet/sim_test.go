package petrinet

import (
	"fmt"
	"math/rand"
	"testing"
)

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
	config := PNSimConfig{
		EndingTime:  100,
		NumOfFiring: 0,
	}
	sim := NewPNSimulation(net, config)
	s := rand.NewSource(1)
	result, nn, tt := sim.RunSimulation(m0, rand.New(s))
	fmt.Println(nn, tt)
	for i, x := range result {
		fmt.Println(i, x.String(net))
	}
}

func TestSim2(t *testing.T) {
	net, m0 := buildRaid6()
	config := PNSimConfig{
		EndingTime:  0,
		NumOfFiring: 10,
	}
	sim := NewPNSimulation(net, config)
	s := rand.NewSource(1)
	result, nn, tt := sim.RunSimulation(m0, rand.New(s))
	fmt.Println(nn, tt)
	for i, x := range result {
		fmt.Println(i, x.String(net))
	}
}

func TestSim3(t *testing.T) {
	net, m0 := buildRaid6()
	config := PNSimConfig{
		EndingTime:      0,
		NumOfFiring:     10,
		NumOfSimulation: 100,
		Rewards:         []string{"avail"},
	}
	sim := NewPNSimulation(net, config)
	irwd, crwd, lastrwd, _, _ := sim.RunAll(m0, 1)
	fmt.Println(irwd)
	fmt.Println(crwd)
	fmt.Println(lastrwd)
}

func TestJSON1(t *testing.T) {
	json := `
	{
		"time": 1.0,
		"firings": 5,
		"simulations": 10,
		"rewards": ["avail", "unavail"]
	}`
	result, err := ReadConfigFromJson([]byte(json))
	fmt.Println(result)
	fmt.Println(err)
}

func TestSim4(t *testing.T) {
	net, m0 := buildRaid6()
	json := `
	{
		"time": 0,
		"firings": 10,
		"simulations": 100,
		"rewards": ["avail", "unavail"]
	}`
	if config, err := ReadConfigFromJson([]byte(json)); err == nil {
		sim := NewPNSimulation(net, config)
		irwd, crwd, lastrwd, _, _ := sim.RunAll(m0, 1)
		fmt.Println(irwd)
		fmt.Println(crwd)
		fmt.Println(lastrwd)
	}
}

// Each worker records clamping into its own recorder, since a clampRecorder is not
// safe to share; the counts are merged when the replications finish. A merge that
// dropped or double-counted events would show up as a total that depends on how many
// workers ran.
func TestClampCountsSurviveTheMerge(t *testing.T) {
	build := func() (*Net, []MarkInt) {
		net := NewNet()
		p1 := net.NewPlace("p1", 10)
		p2 := net.NewPlace("p2", 1) // capacity 1: firing into it clamps
		tr := net.NewExpTrans("t", 0, true, 1.0)
		net.NewInArc(p1, tr, 1)
		net.NewOutArc(tr, p2, 1)
		net.Finalize()
		return net, []MarkInt{5, 0}
	}

	var want []ClampSummary
	for _, parallel := range []int{1, 2, 4, 8} {
		net, m0 := build()
		sim := NewPNSimulation(net, PNSimConfig{
			NumOfFiring: 5, NumOfSimulation: 40, Parallel: parallel,
		})
		sim.RunAll(m0, 7)
		got := sim.ClampEvents()
		if len(got) != 1 {
			t.Fatalf("Parallel=%d: expected one clamp summary, got %v", parallel, got)
		}
		if want == nil {
			want = got
			if want[0].Count == 0 {
				t.Fatalf("expected the net to clamp at all, got %v", want)
			}
			continue
		}
		if got[0] != want[0] {
			t.Errorf("Parallel=%d: clamp summary %v, want %v", parallel, got[0], want[0])
		}
	}
}
