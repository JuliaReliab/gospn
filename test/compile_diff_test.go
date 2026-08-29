package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/okamumu/gospn/pkg/parser"
	"github.com/okamumu/gospn/pkg/petrinet"
)

// The compiled closures in pkg/parser/compile.go are a second implementation of the
// expression semantics, alongside the AST interpreter they fall back to. This is the
// test that keeps the two from drifting: with CheckCompiled set, every compiled
// closure also runs the interpreter and panics on the first disagreement, so simply
// building the marking graph and taking the matrices and rewards checks every
// compiled expression at every reachable marking.
//
// It has already earned its place: it identified a compiled `#place` reading the
// wrong index -- Net.Finalize assigns the indices after the closures are built --
// which the golden test could only show as a runaway state space.
func TestCompiledAgreesWithInterpreter(t *testing.T) {
	parser.CheckCompiled = true
	defer func() { parser.CheckCompiled = false }()
	for _, bn := range benchNets {
		t.Run(bn.name, func(t *testing.T) {
			net, imark, err := parser.PNreadFromFile(bn.file)
			if err != nil {
				t.Fatalf("cannot read %s: %v", bn.file, err)
			}
			mg, _ := petrinet.CreateMarkingGraphWithDFS(net, imark)
			mg.TransMatrix() // exercises the rate closures

			// RewardVector exercises the reward closures, but it can also fail for a
			// reason that has nothing to do with compilation: spnp_example6.spn names
			// an undefined variable (ret_val), which evaluates to its own name as a
			// string, and the interpreter panics on it with or without this package.
			// Only a disagreement between the two implementations is this test's
			// business, so anything else is reported and not failed.
			func() {
				defer func() {
					if r := recover(); r != nil {
						if msg := fmt.Sprint(r); strings.Contains(msg, "disagrees with the interpreter") {
							t.Fatalf("%v", r)
						} else {
							t.Logf("RewardVector fails for this net independently of compilation: %v", r)
						}
					}
				}()
				mg.RewardVector()
			}()
		})
	}
}

// TestCompileCoverage reports how much of each net compiles. A drop here is not a
// correctness problem -- anything that does not compile is interpreted -- but it is
// how a silent loss of the optimisation would show up.
func TestCompileCoverage(t *testing.T) {
	for _, bn := range benchNets {
		t.Run(bn.name, func(t *testing.T) {
			parser.CompileStatsReset()
			if _, _, err := parser.PNreadFromFile(bn.file); err != nil {
				t.Fatalf("cannot read %s: %v", bn.file, err)
			}
			s := parser.CompileStats
			if s.RateFallback+s.GuardFallback+s.MultiFallback+s.RewardFallback > 0 {
				t.Logf("not everything compiled: rate %d/%d guard %d/%d multi %d/%d reward %d/%d",
					s.RateCompiled, s.RateCompiled+s.RateFallback,
					s.GuardCompiled, s.GuardCompiled+s.GuardFallback,
					s.MultiCompiled, s.MultiCompiled+s.MultiFallback,
					s.RewardCompiled, s.RewardCompiled+s.RewardFallback)
			}
		})
	}
}
