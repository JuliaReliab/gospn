package test

import (
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/okamumu/gospn/pkg/parser"
	"github.com/okamumu/gospn/pkg/petrinet"
)

// example/iaas_cloud.spn is parameterised by n, the number of servers in each pool,
// and its state space grows 5-7x for every step of n:
//
//	n=3     25,652 states     n=5    910,731 states
//	n=4    176,487 states     n=6   ~4-5 million (projected)
//
// so it is the net that says whether the search fits in a given machine. n is overridden
// the same way `gospn mark -post` does it -- a later assignment wins -- which needs no
// second copy of the file.
var memNets = []struct {
	name string
	n    int
}{
	{"iaas_n3", 3},
	{"iaas_n4", 4},
	{"iaas_n5", 5},
}

func iaasNet(tb testing.TB, n int) (*petrinet.Net, []petrinet.MarkInt) {
	tb.Helper()
	defs, err := os.ReadFile("../example/iaas_cloud.spn")
	if err != nil {
		tb.Fatal(err)
	}
	net, imark := parser.PNreadFromText(fmt.Sprintf("%s\nn = %d\n", defs, n))
	return net, imark
}

// BenchmarkMarkingGraphMemory reports the heap the marking graph *retains*, which is
// what decides whether a net can be enumerated at all. Go's own -benchmem reports
// allocation volume, and the two differ by more than a factor of two here: most of what
// the search allocates is garbage by the time it finishes, and none of that is the
// reason a large net fails.
func BenchmarkMarkingGraphMemory(b *testing.B) {
	for _, mn := range memNets {
		b.Run(mn.name, func(b *testing.B) {
			net, imark := iaasNet(b, mn.n)
			var states, retained, held uint64
			for i := 0; i < b.N; i++ {
				var before, after runtime.MemStats
				runtime.GC()
				runtime.ReadMemStats(&before)
				mg, err := petrinet.CreateMarkingGraphWithDFSOpts(net, imark,
					petrinet.SearchOptions{MaxStates: 0})
				if err != nil {
					b.Fatal(err)
				}
				// Twice: once before collecting, which includes the garbage the search
				// left behind and is what the machine actually had to hold, and once
				// after, which is what the graph keeps. A net fails on the first number
				// and is analysed out of the second, and they differ by about 2x.
				var peak runtime.MemStats
				runtime.ReadMemStats(&peak)
				runtime.GC()
				runtime.ReadMemStats(&after)
				retained = after.HeapAlloc - before.HeapAlloc
				held = peak.HeapAlloc - before.HeapAlloc
				states = 0
				for _, ms := range mg.StateMarkings() {
					states += uint64(len(ms))
				}
				runtime.KeepAlive(mg)
			}
			b.ReportMetric(float64(states), "states")
			b.ReportMetric(float64(retained)/float64(states), "B/state")
			b.ReportMetric(float64(retained)/1e6, "MB")
			b.ReportMetric(float64(held)/float64(states), "B/state-peak")
		})
	}
}
