package test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
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
//
// wide is here for a different reason: anything whose cost is *per place* -- the intern
// key, the marking itself -- is invisible in a comparison that only varies the state
// count. example/k8s.spn has 53 places against iaas_cloud's 17 but cannot be enumerated
// at any size, and a capped search returns an error rather than a graph, so the wide
// case is a synthetic net instead: k independent two-place subsystems, 2k places and
// 2^k markings, with k chosen to keep the run short.
var memNets = []struct {
	name       string
	n          int // iaas_cloud's server count; 0 selects the synthetic wide net
	subsystems int
}{
	{"iaas_n3", 3, 0},
	{"iaas_n4", 4, 0},
	{"iaas_n5", 5, 0},
	{"wide_32places", 0, 16},
}

func memNet(tb testing.TB, n, subsystems int) (*petrinet.Net, []petrinet.MarkInt) {
	tb.Helper()
	if n == 0 {
		return wideNet(subsystems)
	}
	return iaasNet(tb, n)
}

// wideNet is k copies of `a_i <-> b_i`, one token each: 2k places, 2^k markings, and
// nothing shared between the subsystems.
func wideNet(k int) (*petrinet.Net, []petrinet.MarkInt) {
	net := petrinet.NewNet()
	for i := 0; i < k; i++ {
		a := net.NewPlace(fmt.Sprintf("a%d", i), 1)
		b := net.NewPlace(fmt.Sprintf("b%d", i), 1)
		ab := net.NewExpTrans(fmt.Sprintf("t%d", i), 0, true, 1.0)
		ba := net.NewExpTrans(fmt.Sprintf("u%d", i), 0, true, 1.0)
		net.NewInArc(a, ab, 1)
		net.NewOutArc(ab, b, 1)
		net.NewInArc(b, ba, 1)
		net.NewOutArc(ba, a, 1)
	}
	net.Finalize()
	imark := make(map[string]petrinet.MarkInt, k)
	for i := 0; i < k; i++ {
		imark[fmt.Sprintf("a%d", i)] = 1
	}
	return net, net.MakeMark(imark)
}

func iaasNet(tb testing.TB, n int) (*petrinet.Net, []petrinet.MarkInt) {
	tb.Helper()
	defs, err := os.ReadFile("../example/iaas_cloud.spn")
	if err != nil {
		tb.Fatal(err)
	}
	net, imark, err := parser.PNreadFromText(fmt.Sprintf("%s\nn = %d\n", defs, n))
	if err != nil {
		tb.Fatal(err)
	}
	return net, imark
}

// writeHeapProfile dumps the heap while the graph is still alive, which is the only way
// to see what it retains: `go test -memprofile` writes after the benchmark returns, when
// the graph is already garbage and the profile is empty. It writes only when
// GOSPN_HEAPPROF names a directory.
func writeHeapProfile(tb testing.TB, name string, mg *petrinet.MarkingGraph) {
	tb.Helper()
	dir := os.Getenv("GOSPN_HEAPPROF")
	if dir == "" {
		return
	}
	f, err := os.Create(filepath.Join(dir, name+".heap"))
	if err != nil {
		tb.Fatal(err)
	}
	defer f.Close()
	runtime.GC()
	if err := pprof.WriteHeapProfile(f); err != nil {
		tb.Fatal(err)
	}
	runtime.KeepAlive(mg)
	tb.Logf("heap profile: %s", f.Name())
}

// BenchmarkMarkingGraphMemory reports the heap the marking graph *retains*, which is
// what decides whether a net can be enumerated at all. Go's own -benchmem reports
// allocation volume, and the two differ by more than a factor of two here: most of what
// the search allocates is garbage by the time it finishes, and none of that is the
// reason a large net fails.
func BenchmarkMarkingGraphMemory(b *testing.B) {
	for _, mn := range memNets {
		b.Run(mn.name, func(b *testing.B) {
			net, imark := memNet(b, mn.n, mn.subsystems)
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
				// Two numbers, because they answer different questions and differ by
				// 2-3x: what the graph *keeps*, and the high-water mark of heap the
				// process had to obtain, which is what decides whether a net can be
				// enumerated at all.
				//
				// The high-water mark is HeapSys, which only grows. Sampling HeapAlloc
				// straight after the search instead makes the number depend on where
				// the collector last ran: it varied from 742 to 1353 B/state between
				// runs of the same build, which is more than the effects being measured.
				runtime.GC()
				runtime.ReadMemStats(&after)
				retained = after.HeapAlloc - before.HeapAlloc
				held = after.HeapSys - before.HeapSys
				states = 0
				for _, ms := range mg.StateMarkings() {
					states += uint64(len(ms))
				}
				writeHeapProfile(b, mn.name, mg)
				runtime.KeepAlive(mg)
			}
			b.ReportMetric(float64(len(net.PlaceLabels())), "places")
			b.ReportMetric(float64(states), "states")
			b.ReportMetric(float64(retained)/float64(states), "B/state")
			b.ReportMetric(float64(retained)/1e6, "MB")
			// HeapSys is process-wide and only grows, so a net that runs after a larger
			// one sees no growth at all and reports zero. Measure one net per process:
			//
			//	go test ./test -run XXX -bench MarkingGraphMemory/iaas_n5 -benchtime 1x
			if held == 0 {
				b.Log("the peak is 0: an earlier net in this process already grew the heap. " +
					"Run this net on its own to measure it.")
			}
			b.ReportMetric(float64(held)/float64(states), "B/state-peak")
		})
	}
}
