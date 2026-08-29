package petrinet

import (
	"errors"
	"strings"
	"testing"
)

// A firing that stays inside every place's capacity must return a true nil error, not a
// typed nil pointer wrapped in an error interface.
func TestDoFiringNoClamp(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 10)
	p2 := net.NewPlace("p2", 10)
	t1 := net.NewExpTrans("t1", 0, true, 1.0)
	net.NewInArc(p1, t1, 1)
	net.NewOutArc(t1, p2, 1)
	net.Finalize()

	mark, err := t1.DoFiring(net, []MarkInt{1, 0})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if mark[0] != 0 || mark[1] != 1 {
		t.Errorf("expected [0 1], got %v", mark)
	}
}

// Going past a place's capacity is clamped, and reported as a *ClampError naming the
// place and the bound it ran into.
func TestDoFiringClampsAtMax(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 10)
	p2 := net.NewPlace("p2", 3)
	t1 := net.NewExpTrans("t1", 0, true, 1.0)
	net.NewInArc(p1, t1, 1)
	net.NewOutArc(t1, p2, 1)
	net.Finalize()

	mark, err := t1.DoFiring(net, []MarkInt{1, 3})
	if mark[1] != 3 {
		t.Errorf("expected p2 clamped at its max 3, got %d", mark[1])
	}
	var ce *ClampError
	if !errors.As(err, &ce) {
		t.Fatalf("expected a *ClampError, got %v", err)
	}
	if len(ce.Events) != 1 {
		t.Fatalf("expected 1 clamp event, got %d", len(ce.Events))
	}
	want := ClampEvent{Trans: "t1", Place: "p2", Bound: ClampMax, Limit: 3}
	if ce.Events[0] != want {
		t.Errorf("expected %v, got %v", want, ce.Events[0])
	}
}

// The same at the lower bound.
func TestDoFiringClampsAtZero(t *testing.T) {
	net := NewNet()
	p1 := net.NewPlace("p1", 10)
	t1 := net.NewExpTrans("t1", 0, true, 1.0)
	net.NewInArc(p1, t1, 2)
	net.Finalize()

	mark, err := t1.DoFiring(net, []MarkInt{1})
	if mark[0] != 0 {
		t.Errorf("expected p1 clamped at 0, got %d", mark[0])
	}
	var ce *ClampError
	if !errors.As(err, &ce) {
		t.Fatalf("expected a *ClampError, got %v", err)
	}
	want := ClampEvent{Trans: "t1", Place: "p1", Bound: ClampMin, Limit: 0}
	if len(ce.Events) != 1 || ce.Events[0] != want {
		t.Errorf("expected [%v], got %v", want, ce.Events)
	}
}

// The recorder collapses repeats of the same (transition, place, bound) and counts them.
func TestClampRecorderCounts(t *testing.T) {
	var r clampRecorder
	a := ClampEvent{Trans: "t1", Place: "p1", Bound: ClampMax, Limit: 3}
	b := ClampEvent{Trans: "t2", Place: "p1", Bound: ClampMax, Limit: 3}
	r.record(nil)
	r.record(errors.New("something else"))
	r.record(&ClampError{Events: []ClampEvent{a}})
	r.record(&ClampError{Events: []ClampEvent{a, b}})

	got := r.events()
	if len(got) != 2 {
		t.Fatalf("expected 2 summaries, got %d (%v)", len(got), got)
	}
	if got[0].ClampEvent != a || got[0].Count != 2 {
		t.Errorf("expected %v counted twice, got %v", a, got[0])
	}
	if got[1].ClampEvent != b || got[1].Count != 1 {
		t.Errorf("expected %v counted once, got %v", b, got[1])
	}
}

// Building a marking graph over a net that clamps must surface it, so a caller is not
// left with a generator matrix whose rows silently do not sum to zero.
func TestMarkingGraphReportsClamping(t *testing.T) {
	build := func() (*Net, []MarkInt) {
		net := NewNet()
		p1 := net.NewPlace("p1", 10)
		p2 := net.NewPlace("p2", 1) // too small to take what t1 produces
		t1 := net.NewExpTrans("t1", 0, true, 1.0)
		net.NewInArc(p1, t1, 1)
		net.NewOutArc(t1, p2, 1)
		net.Finalize()
		return net, []MarkInt{2, 1}
	}

	net, m0 := build()
	events := mustGraph(CreateMarkingGraphWithDFS(net, m0)).ClampEvents()
	if len(events) != 1 {
		t.Fatalf("expected 1 clamp summary from DFS, got %v", events)
	}
	if events[0].Place != "p2" || events[0].Bound != ClampMax || events[0].Limit != 1 {
		t.Errorf("unexpected summary %v", events[0])
	}

	net, m0 = build()
	if events := mustGraph(CreateMarkingGraphWithDFSTangible(net, m0)).ClampEvents(); len(events) != 1 {
		t.Fatalf("expected 1 clamp summary from the tangible DFS, got %v", events)
	}

	// And a net that never clamps reports nothing.
	net = NewNet()
	p1 := net.NewPlace("p1", 10)
	p2 := net.NewPlace("p2", 10)
	t1 := net.NewExpTrans("t1", 0, true, 1.0)
	net.NewInArc(p1, t1, 1)
	net.NewOutArc(t1, p2, 1)
	net.Finalize()
	if events := mustGraph(CreateMarkingGraphWithDFS(net, []MarkInt{2, 0})).ClampEvents(); len(events) != 0 {
		t.Errorf("expected no clamping, got %v", events)
	}
}

func TestFormatClampEvents(t *testing.T) {
	if s := FormatClampEvents(nil); s != "" {
		t.Errorf("expected an empty report for no events, got %q", s)
	}
	s := FormatClampEvents([]ClampSummary{{
		ClampEvent: ClampEvent{Trans: "t1", Place: "p2", Bound: ClampMax, Limit: 255},
		Count:      7,
	}})
	for _, want := range []string{"t1", "p2", "max 255", "7 time(s)", "not exact"} {
		if !strings.Contains(s, want) {
			t.Errorf("expected the report to mention %q, got %q", want, s)
		}
	}
}

// mustGraph is for tests whose nets are small enough that the state limit cannot
// apply; a limit error there is a bug in the test, not a result to handle.
func mustGraph(mg *MarkingGraph, err error) *MarkingGraph {
	if err != nil {
		panic(err)
	}
	return mg
}
