package test

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/okamumu/gospn/pkg/mt"
	"github.com/okamumu/gospn/pkg/parser"
	"github.com/okamumu/gospn/pkg/petrinet"
)

// Nothing pinned the simulation output before this: the golden test covers the
// marking graph, and RunAll is a separate path that evaluates the same guard, rate
// and reward expressions over a sample path rather than over a state space. With a
// fixed seed the whole run is deterministic, so it can be pinned exactly.
//
// Values are rendered at 12 significant digits, as in TestMarkingGraphGolden.
func simFingerprint(t *testing.T, file string, rewards []string) string {
	net, imark, err := parser.PNreadFromFile(file)
	if err != nil {
		t.Fatalf("cannot read %s: %v", file, err)
	}
	sim := petrinet.NewPNSimulation(net, petrinet.PNSimConfig{
		EndingTime:      10000.0,
		NumOfSimulation: 50,
		Rewards:         rewards,
	})
	rng := mt.NewMT64()
	rng.Seed(12345)
	irwd, crwd, lastrwd, elapsed, count := sim.RunAll(imark, rng)

	var sb strings.Builder
	dump := func(name string, m map[string][]float64) {
		keys := make([]string, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Fprintf(&sb, "%s[%s]", name, k)
			for _, v := range m[k] {
				fmt.Fprintf(&sb, " %.12g", v)
			}
			sb.WriteByte('\n')
		}
	}
	dump("irwd", irwd)
	dump("crwd", crwd)
	dump("lastrwd", lastrwd)
	fmt.Fprintf(&sb, "elapsed")
	for _, v := range elapsed {
		fmt.Fprintf(&sb, " %.12g", v)
	}
	fmt.Fprintf(&sb, "\ncount %v\n", count)
	return sb.String()
}

var simGolden = map[string]string{
	"P4": "96d420c83f1bb1ae2d3b3b152e423ada826865a7f553440f31197ec634601df4",
	"P5": "369a32dff0b5846942d661746ca29261b52bd0f5ae5057b722b1cbfccf47deae",
	"P7": "c0b0f71ff14dea331abdb21fd6c5b329c0caaca46e2730c5720fe900a3e40c2f",
}

func TestSimGolden(t *testing.T) {
	for _, sn := range simNets {
		t.Run(sn.name, func(t *testing.T) {
			sum := sha256.Sum256([]byte(simFingerprint(t, sn.file, sn.rewards)))
			got := hex.EncodeToString(sum[:])
			want, ok := simGolden[sn.name]
			if !ok {
				t.Logf("GOLDEN %q: %q,", sn.name, got)
				return
			}
			if got != want {
				t.Errorf("simulation output changed for %s:\n got %s\nwant %s", sn.name, got, want)
			}
		})
	}
}

// The same seed must give the same run.
func TestSimIsDeterministic(t *testing.T) {
	for _, sn := range simNets {
		t.Run(sn.name, func(t *testing.T) {
			first := simFingerprint(t, sn.file, sn.rewards)
			for i := 0; i < 3; i++ {
				if simFingerprint(t, sn.file, sn.rewards) != first {
					t.Fatalf("run %d differs from the first", i+1)
				}
			}
		})
	}
}
