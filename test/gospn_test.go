package test

import (
	"testing"

	"github.com/okamumu/gospn/pkg/parser"
	"github.com/okamumu/gospn/pkg/petrinet"
)

// The nets used as benchmark inputs, roughly in order of marking-graph size.
var benchNets = []struct {
	name string
	file string
}{
	{"P1", "../example/spnp_example1.spn"},
	{"P2", "../example/spnp_example2.spn"},
	{"P3", "../example/spnp_example3.spn"},
	{"P4", "../example/spnp_example4.spn"},
	{"P5", "../example/spnp_example5.spn"},
	{"P6", "../example/spnp_example6.spn"},
	{"P7", "../example/raid6.spn"},
	{"P8", "../example/raid10.spn"},
}

// The net is parsed once, outside the timed loop: what is being measured is the
// marking-graph construction, not the parser. A read error fails the benchmark rather
// than being swallowed -- a benchmark that silently measures nothing is worse than none.
func BenchmarkGoSPN(b *testing.B) {
	for _, bn := range benchNets {
		b.Run(bn.name, func(b *testing.B) {
			net, imark, err := parser.PNreadFromFile(bn.file)
			if err != nil {
				b.Fatalf("cannot read %s: %v", bn.file, err)
			}
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				mg := petrinet.CreateMarkingGraphWithDFS(net, imark)
				mg.TransMatrix()
				mg.GroupLabels()
				mg.TransLabels()
				mg.InitVector()
			}
		})
	}
}

// TestMarkingGraphSizes pins the marking graph produced for every benchmark net, so the
// optimisation work can prove it did not change what gospn computes.
func TestMarkingGraphSizes(t *testing.T) {
	for _, bn := range benchNets {
		t.Run(bn.name, func(t *testing.T) {
			net, imark, err := parser.PNreadFromFile(bn.file)
			if err != nil {
				t.Fatalf("cannot read %s: %v", bn.file, err)
			}
			mg := petrinet.CreateMarkingGraphWithDFS(net, imark)
			t.Logf("%s groups=%d translabels=%d", bn.name, len(mg.GroupLabels()), len(mg.TransLabels()))
		})
	}
}
