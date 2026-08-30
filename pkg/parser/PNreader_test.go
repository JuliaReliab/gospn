package parser

import (
	"os"
	"testing"

	"github.com/okamumu/gospn/pkg/petrinet"
)

// These tests used to read nine nets and print them. That every bundled net parses and
// evaluates is now covered end to end by TestEveryExampleEvaluates in test/, so what is
// left here is what only these entry points can be asked: that the two of them agree,
// and that a missing file is an error rather than an empty net.

var readerNets = []string{
	"../../example/spnp_example1.spn",
	"../../example/spnp_example2.spn",
	"../../example/spnp_example3.spn",
	"../../example/spnp_example4.spn",
	"../../example/spnp_example5.spn",
	"../../example/spnp_example6.spn",
	"../../example/iaas_cloud.spn",
	"../../example/raid6.spn",
	"../../example/raid10.spn",
}

// PNreadFromFile and PNreadFromText are separate copies of the same five steps, so they
// can drift apart. Reading a file both ways has to give the same net.
func TestPNreadFromFileMatchesFromText(t *testing.T) {
	for _, file := range readerNets {
		file := file
		t.Run(file, func(t *testing.T) {
			net, imark, err := PNreadFromFile(file)
			if err != nil {
				t.Fatal(err)
			}
			text, err := os.ReadFile(file)
			if err != nil {
				t.Fatal(err)
			}
			net2, imark2 := PNreadFromText(string(text))

			if !equalStrings(net.PlaceLabels(), net2.PlaceLabels()) {
				t.Errorf("places differ: %v vs %v", net.PlaceLabels(), net2.PlaceLabels())
			}
			if !equalMarks(imark, imark2) {
				t.Errorf("initial markings differ: %v vs %v", imark, imark2)
			}
			if len(net.PlaceLabels()) == 0 {
				t.Error("the net has no places")
			}
		})
	}
}

func TestPNreadFromFileMissing(t *testing.T) {
	net, imark, err := PNreadFromFile("../../example/there-is-no-such-net.spn")
	if err == nil {
		t.Fatal("a missing file was read without an error")
	}
	if net != nil {
		t.Error("a net was returned for a missing file")
	}
	if len(imark) != 0 {
		t.Errorf("the initial marking is %v, want empty", imark)
	}
}

// The initial marking is indexed by place, in the order Finalize sorted them into.
func TestPNreadInitialMarking(t *testing.T) {
	net, imark, err := PNreadFromFile("../../example/raid6.spn")
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]petrinet.MarkInt{"Pn": 6, "Po": 1, "Pdf": 0, "Pr": 0, "Pc": 0}
	labels := net.PlaceLabels()
	if len(labels) != len(want) {
		t.Fatalf("%d places, want %d", len(labels), len(want))
	}
	for i, label := range labels {
		if imark[i] != want[label] {
			t.Errorf("%s starts with %d tokens, want %d", label, imark[i], want[label])
		}
	}
}

func equalMarks(a, b []petrinet.MarkInt) bool {
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
