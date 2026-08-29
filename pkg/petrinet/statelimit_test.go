package petrinet

import (
	"errors"
	"testing"
)

// A net with one place of capacity 1000 and a transition that fills it has 1001
// reachable markings, which is enough to trip a small limit.
func limitNet() (*Net, []MarkInt) {
	net := NewNet()
	p := net.NewPlace("p", 1000)
	t := net.NewExpTrans("t", 0, true, 1.0)
	net.NewOutArc(t, p, 1)
	net.Finalize()
	return net, []MarkInt{0}
}

func TestSearchStopsAtTheLimit(t *testing.T) {
	net, m0 := limitNet()
	mg, err := CreateMarkingGraphWithDFSOpts(net, m0, SearchOptions{MaxStates: 100})
	if err == nil {
		t.Fatalf("expected the search to stop at the limit, got a graph of %d states", nstates(mg))
	}
	if mg != nil {
		t.Error("no graph should be returned with a limit error: a truncated graph " +
			"produces transition matrices that look valid and are not")
	}
	var sle *StateLimitError
	if !errors.As(err, &sle) {
		t.Fatalf("expected a *StateLimitError, got %T: %v", err, err)
	}
	if sle.Limit != 100 {
		t.Errorf("Limit = %d, want 100", sle.Limit)
	}
	if sle.Found < 100 {
		t.Errorf("Found = %d, want at least the limit", sle.Found)
	}
}

func TestNoLimitWhenZero(t *testing.T) {
	net, m0 := limitNet()
	mg, err := CreateMarkingGraphWithDFSOpts(net, m0, SearchOptions{MaxStates: 0})
	if err != nil {
		t.Fatalf("MaxStates 0 means no limit, got %v", err)
	}
	if n := nstates(mg); n != 1001 {
		t.Errorf("got %d states, want 1001", n)
	}
}

func TestLimitDoesNotFireBelowIt(t *testing.T) {
	net, m0 := limitNet()
	if _, err := CreateMarkingGraphWithDFSOpts(net, m0, SearchOptions{MaxStates: 5000}); err != nil {
		t.Errorf("a limit above the state count must not fire: %v", err)
	}
}

func TestProgressIsReported(t *testing.T) {
	net, m0 := limitNet()
	var calls []int
	_, err := CreateMarkingGraphWithDFSOpts(net, m0, SearchOptions{
		Progress:      func(n int) { calls = append(calls, n) },
		ProgressEvery: 100,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(calls) < 5 {
		t.Errorf("expected progress every 100 of ~1001 states, got %v", calls)
	}
	for i := 1; i < len(calls); i++ {
		if calls[i] <= calls[i-1] {
			t.Errorf("progress must increase, got %v", calls)
			break
		}
	}
}

// nstates is a test helper; the marking graph does not export a state count.
func nstates(mg *MarkingGraph) int {
	n := 0
	for _, marks := range mg.groupToMark {
		n += len(marks)
	}
	return n
}
