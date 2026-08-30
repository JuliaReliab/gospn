package test

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/okamumu/gospn/pkg/parser"
	"github.com/okamumu/gospn/pkg/petrinet"
)

// `gospn mark -t` runs a second, separate search (dfstangible.go) that vanishes
// immediate markings as it goes. It changes the state space -- 9 states become 5 on
// raid6.spn, 910,731 become 568,584 on iaas_cloud.spn with n=5 -- and until now nothing
// pinned its result: five of its functions were never executed by any test.
//
// What must hold is that vanishing a marking does not change which *tangible* markings
// exist. The two searches disagree on the vanishing ones by construction, which is the
// point of the flag.

// tangibleMarkings is the markings of the groups that are not vanishing, as strings so
// they can be compared as a set. GroupLabels names a vanishing group I<k>.
func tangibleMarkings(t *testing.T, mg *petrinet.MarkingGraph) []string {
	t.Helper()
	labels := mg.GroupLabels()
	var out []string
	for g, marks := range mg.StateMarkings() {
		if strings.HasPrefix(labels[g], "I") {
			continue
		}
		for _, m := range marks {
			out = append(out, fmt.Sprint(m))
		}
	}
	sort.Strings(out)
	return out
}

func TestTangibleSearchKeepsTheTangibleMarkings(t *testing.T) {
	for _, bn := range benchNets {
		bn := bn
		t.Run(bn.name, func(t *testing.T) {
			net, imark, err := parser.PNreadFromFile(bn.file)
			if err != nil {
				t.Fatal(err)
			}
			plain, err := petrinet.CreateMarkingGraphWithDFS(net, imark)
			if err != nil {
				t.Fatal(err)
			}
			net2, imark2, err := parser.PNreadFromFile(bn.file)
			if err != nil {
				t.Fatal(err)
			}
			tangible, err := petrinet.CreateMarkingGraphWithDFSTangible(net2, imark2)
			if err != nil {
				t.Fatal(err)
			}

			want, got := tangibleMarkings(t, plain), tangibleMarkings(t, tangible)
			if !equalStringSlices(want, got) {
				t.Errorf("the tangible markings differ.\n  plain: %v\n     -t: %v", want, got)
			}

			// And -t never has more states overall: vanishing markings are removed,
			// never added.
			if countStates(tangible) > countStates(plain) {
				t.Errorf("-t has %d states, plain has %d", countStates(tangible), countStates(plain))
			}
		})
	}
}

func countStates(mg *petrinet.MarkingGraph) int {
	n := 0
	for _, ms := range mg.StateMarkings() {
		n += len(ms)
	}
	return n
}

func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
