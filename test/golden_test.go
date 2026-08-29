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

// dumpMarkingGraph renders everything that feeds the MATLAB output -- the per-group state
// labels and the three sparse transition matrices -- as deterministic text. Groups and
// transitions are keyed by their label and sorted, so the dump does not depend on map
// iteration order or on where the allocator happened to put a marking. The DOT export
// cannot be used for this: it prints raw pointers.
// floats renders a float slice at 12 significant digits. TransMatrix() accumulates the
// generator diagonal while ranging over a map, so the summation order -- and with it the
// last ULP of each diagonal entry -- varies between runs. Full precision would therefore
// make this dump non-reproducible for reasons unrelated to any change being tested.
func floats(xs []float64) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%.12g", x)
	}
	return "[" + strings.Join(parts, " ") + "]"
}

func dumpMarkingGraph(mg *petrinet.MarkingGraph) string {
	var sb strings.Builder

	grouplabel := mg.GroupLabels()
	states := mg.StateLabels()
	initv := mg.InitVector()

	type groupRow struct {
		label  string
		states []string
		init   []float64
	}
	rows := make([]groupRow, 0, len(grouplabel))
	for g, label := range grouplabel {
		rows = append(rows, groupRow{label: label, states: states[g], init: initv[g]})
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].label < rows[j].label })
	for _, r := range rows {
		fmt.Fprintf(&sb, "group %s n=%d\n", r.label, len(r.states))
		for i, s := range r.states {
			fmt.Fprintf(&sb, "  state[%d] %s\n", i, s)
		}
		fmt.Fprintf(&sb, "  init %s\n", floats(r.init))
	}

	translabel := mg.TransLabels()
	expmat, immmat, genmat := mg.TransMatrix()
	for _, m := range []struct {
		kind string
		mat  map[petrinet.GroupTrans]*petrinet.CSC
	}{{"EXP", expmat}, {"IMM", immmat}, {"GEN", genmat}} {
		type matRow struct {
			label string
			csc   *petrinet.CSC
		}
		mrows := make([]matRow, 0, len(m.mat))
		for gt, csc := range m.mat {
			// The same composite name cmd/gospn.go gives the MATLAB variable. TransLabels
			// alone is not unique -- every non-GEN group transition is labelled "E" or "I".
			name := grouplabel[gt.GetSrc()] + grouplabel[gt.GetDest()] + translabel[gt]
			mrows = append(mrows, matRow{label: name, csc: csc})
		}
		sort.Slice(mrows, func(i, j int) bool { return mrows[i].label < mrows[j].label })
		for _, r := range mrows {
			dim, nnz, rowind, colptr, value := r.csc.Get()
			fmt.Fprintf(&sb, "%s %s dim=%v nnz=%d\n  rowind=%v\n  colptr=%v\n  value=%s\n",
				m.kind, r.label, dim, nnz, rowind, colptr, floats(value))
		}
	}
	return sb.String()
}

// The hash of dumpMarkingGraph for each benchmark net. Any optimisation of the marking
// graph construction must leave these untouched; a changed hash means gospn now computes
// something different, not merely faster.
var goldenHash = map[string]string{
	"P1": "2adfcd93bef713bbf870994f745aaee5df9c4a386aca78b9d4ba67935f4751b7",
	"P2": "59a252bcc2a463a6b773d8fda0fbe7c6e0d9ae318fdd418ed0ccbcaddb8e84cd",
	"P3": "f2dcb25a917f68d8af119d128ea497eded771cc686e48aecb25d52758a54cb82",
	"P4": "a40b79a7cc08921c982c5f1662d86b2cd14cfa81e17671f6582af9c665aebaf7",
	"P5": "35f73f54472f99505013678e388e418ee9611f8999128c88d999d0eab9c1a2bb",
	"P6": "158dc5c6bcd8c50a4c5ba2cfb876c7ef1cb7b27070707363ab69e2e18d351ad4",
	"P7": "bdfab5e04db482be4e4a592d3c6ef8a95cc40764d82d9b04017e1c1d4c0ea85d",
	"P8": "0c5410cc1bdf704335dd7cac6de2224354ad6d7f9fd05f37415b4a788d81d204",
}

func TestMarkingGraphGolden(t *testing.T) {
	for _, bn := range benchNets {
		t.Run(bn.name, func(t *testing.T) {
			net, imark, err := parser.PNreadFromFile(bn.file)
			if err != nil {
				t.Fatalf("cannot read %s: %v", bn.file, err)
			}
			dump := dumpMarkingGraph(mustGraph(petrinet.CreateMarkingGraphWithDFS(net, imark)))
			sum := sha256.Sum256([]byte(dump))
			got := hex.EncodeToString(sum[:])
			want, ok := goldenHash[bn.name]
			if !ok {
				t.Logf("GOLDEN %q: %q,", bn.name, got)
				return
			}
			if got != want {
				t.Errorf("marking graph changed for %s:\n got %s\nwant %s", bn.name, got, want)
			}
		})
	}
}

// The dump must not depend on map iteration order or allocation addresses.
func TestMarkingGraphGoldenIsDeterministic(t *testing.T) {
	for _, bn := range benchNets {
		t.Run(bn.name, func(t *testing.T) {
			var first string
			for i := 0; i < 5; i++ {
				net, imark, err := parser.PNreadFromFile(bn.file)
				if err != nil {
					t.Fatalf("cannot read %s: %v", bn.file, err)
				}
				dump := dumpMarkingGraph(mustGraph(petrinet.CreateMarkingGraphWithDFS(net, imark)))
				if i == 0 {
					first = dump
				} else if dump != first {
					t.Fatalf("dump is not deterministic across runs (iteration %d)", i)
				}
			}
		})
	}
}

// mustGraph is for the bundled nets, all of which are far below the state limit.
func mustGraph(mg *petrinet.MarkingGraph, err error) *petrinet.MarkingGraph {
	if err != nil {
		panic(err)
	}
	return mg
}
