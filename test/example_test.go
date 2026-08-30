package test

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/okamumu/gospn/pkg/parser"
	"github.com/okamumu/gospn/pkg/petrinet"
)

// TestEveryExampleEvaluates parses every net in example/ and evaluates its guards,
// rates, weights and rewards once.
//
// This is the third time a bundled example has shipped with an expression that cannot be
// evaluated: `ret_val` commented out in spnp_example6.spn (0.11.1), the same net again in
// 0.13.x, and cold_vm_reju.spn naming Tvreset.rate and Thfai.rate, which are typos for
// Tvrestart.rate and Thfail.rate. Each one was found by a person running the tool, not by
// a test, because nothing ran the bundled nets. This does.
func TestEveryExampleEvaluates(t *testing.T) {
	files, err := filepath.Glob("../example/*.spn")
	if err != nil {
		t.Fatal(err)
	}
	if len(files) == 0 {
		t.Fatal("no nets found in ../example")
	}
	for _, file := range files {
		file := file
		t.Run(filepath.Base(file), func(t *testing.T) {
			net, imark, err := parser.PNreadFromFile(file)
			if err != nil {
				t.Fatal(err)
			}
			if err := net.CheckExpressions(imark); err != nil {
				t.Error(err)
			}
		})
	}
}

// TestEveryEnumerableExampleBuildsMatrices goes further for the nets small enough to
// enumerate: it builds the marking graph and the transition matrices, which is where an
// expression that only fails for some markings would show up. A net over the limit is
// skipped rather than failed -- example/k8s.spn cannot be enumerated at any size.
func TestEveryEnumerableExampleBuildsMatrices(t *testing.T) {
	files, err := filepath.Glob("../example/*.spn")
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		file := file
		t.Run(filepath.Base(file), func(t *testing.T) {
			net, imark, err := parser.PNreadFromFile(file)
			if err != nil {
				t.Fatal(err)
			}
			mg, err := petrinet.CreateMarkingGraphWithDFSOpts(net, imark,
				petrinet.SearchOptions{MaxStates: 200_000})
			if err != nil {
				var limit *petrinet.StateLimitError
				if errors.As(err, &limit) {
					t.Skipf("%d states is over the cap this test uses", limit.Found)
				}
				t.Fatal(err)
			}
			exp, imm, gen := mg.TransMatrix()
			if len(exp)+len(imm)+len(gen) == 0 {
				t.Errorf("%s: no transition matrices", file)
			}
			mg.RewardVector()
			mg.InitVector()
		})
	}
}
