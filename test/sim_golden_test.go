package test

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/okamumu/gospn/pkg/parser"
	"github.com/okamumu/gospn/pkg/petrinet"
)

// Nothing pinned the simulation output before this: the golden test covers the
// marking graph, and RunAll is a separate path that evaluates the same guard, rate
// and reward expressions over a sample path rather than over a state space. With a
// fixed seed the whole run is deterministic, so it can be pinned exactly.
//
// Values are rendered at 12 significant digits, as in TestMarkingGraphGolden.
//
// These hashes changed in 0.16.0: replications now take a random stream derived from
// their index rather than sharing one sequential stream, so a given seed produces
// different (statistically equivalent) numbers than it did up to 0.15.1.
func simFingerprint(t *testing.T, file string, rewards []string) string {
	return simFingerprintP(t, file, rewards, 0)
}

func simFingerprintP(t *testing.T, file string, rewards []string, parallel int) string {
	net, imark, err := parser.PNreadFromFile(file)
	if err != nil {
		t.Fatalf("cannot read %s: %v", file, err)
	}
	sim := petrinet.NewPNSimulation(net, petrinet.PNSimConfig{
		EndingTime:      10000.0,
		NumOfSimulation: 50,
		Rewards:         rewards,
		Parallel:        parallel,
	})
	irwd, crwd, lastrwd, elapsed, count := sim.RunAll(imark, 12345)

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
	"P4": "7878a16422dfcc4476de77b7ff4532ae9b33b9b913e900181f09536b3f65d58a",
	"P5": "438b97b95d85f64e050b86e28696439b2a5efc5f3e3024f2fc9285b516d1403a",
	"P7": "239815164a6cf9e4c95970568f718a998fcbbc7f7b189eb9d8181852b5738960",
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

// Replications run on several workers at once, so the result must not depend on how
// many. Replication k always takes the random stream derived from k, never from the
// worker that happened to pick it up; if that were ever reversed, the numbers would
// quietly start depending on the machine and on scheduling.
func TestSimIsIndependentOfWorkerCount(t *testing.T) {
	for _, sn := range simNets {
		t.Run(sn.name, func(t *testing.T) {
			want := simFingerprintP(t, sn.file, sn.rewards, 1)
			for _, p := range []int{2, 3, 4, 8, 16} {
				if got := simFingerprintP(t, sn.file, sn.rewards, p); got != want {
					t.Errorf("Parallel=%d gives a different result from Parallel=1", p)
				}
			}
		})
	}
}
