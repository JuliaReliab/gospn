package test

import (
	"testing"

	"github.com/okamumu/gospn/pkg/parser"
	"github.com/okamumu/gospn/pkg/petrinet"
)

// k8s.spn is a 53-place, 64-transition GSPN with places holding up to 1000 tokens --
// far too large to enumerate, so simulation is the only way it is analysed. That
// makes it the realistic case for the RunAll path.
func BenchmarkSimK8s(b *testing.B) {
	net, imark, err := parser.PNreadFromFile("../example/k8s.spn")
	if err != nil {
		b.Fatalf("cannot read: %v", err)
	}
	cfg := petrinet.PNSimConfig{
		EndingTime:      1000.0,
		NumOfSimulation: 10,
		Rewards: []string{"perf_Iq", "perf_Ipp", "throughput", "jobs",
			"drop_probability", "hpa_D1pt", "ca_Ne1"},
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sim := petrinet.NewPNSimulation(net, cfg)
		sim.RunAll(imark, 12345)
	}
}
