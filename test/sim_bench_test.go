package test

import (
	"testing"

	"github.com/okamumu/gospn/pkg/mt"
	"github.com/okamumu/gospn/pkg/parser"
	"github.com/okamumu/gospn/pkg/petrinet"
)

// The marking-graph benchmarks only exercise the guard and rate closures through the
// reachability search. Simulation runs the same closures on a different path -- one
// sample path at a time, with the update and multiplicity functions in play too --
// so it is measured separately.
var simNets = []struct {
	name    string
	file    string
	rewards []string
}{
	{"P4", "../example/spnp_example4.spn", []string{"reliab", "reward_rate"}},
	{"P5", "../example/spnp_example5.spn", []string{"enallrwd", "reliab"}},
	{"P7", "../example/raid6.spn", []string{"avail", "unavail"}},
}

func BenchmarkSim(b *testing.B) {
	for _, sn := range simNets {
		b.Run(sn.name, func(b *testing.B) {
			net, imark, err := parser.PNreadFromFile(sn.file)
			if err != nil {
				b.Fatalf("cannot read %s: %v", sn.file, err)
			}
			cfg := petrinet.PNSimConfig{
				EndingTime:      10000.0,
				NumOfFiring:     0,
				NumOfSimulation: 200,
				Rewards:         sn.rewards,
			}
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				sim := petrinet.NewPNSimulation(net, cfg)
				rng := mt.NewMT64()
				rng.Seed(12345)
				sim.RunAll(imark, rng)
			}
		})
	}
}
