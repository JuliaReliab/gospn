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
	s := rand.NewSource(1)
	irwd, crwd, lastrwd, _, _ := sim.RunAll(m0, rand.New(s))
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
		s := rand.NewSource(1)
		irwd, crwd, lastrwd, _, _ := sim.RunAll(m0, rand.New(s))
		fmt.Println(irwd)
		fmt.Println(crwd)
		fmt.Println(lastrwd)
	}
}
