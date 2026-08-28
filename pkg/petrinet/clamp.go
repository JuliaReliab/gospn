package petrinet

import (
	"errors"
	"fmt"
	"strings"
)

// A firing whose destination marking runs past a place's capacity (or below zero) is
// clamped rather than rejected. That is a deliberate choice, but it makes the marking
// graph inexact: the transition's real destination is lost, and the generator matrix
// ends up with rows that do not sum to zero.
//
// DoFiring reports this through its error return. The types below let the analysis
// collect those reports so a caller can see what happened, instead of every call site
// having to discard the error.
//
// This is easy to run into because the default place capacity is 255 (see
// parser.PNcompile): a model that legitimately holds more tokens in a place, but does
// not spell out `max`, is truncated.

// ClampBound tells which bound a firing ran into.
type ClampBound int

const (
	ClampMin ClampBound = iota + 1 // the place would have gone below zero
	ClampMax                       // the place would have gone above its capacity
)

func (b ClampBound) String() string {
	switch b {
	case ClampMin:
		return "min"
	case ClampMax:
		return "max"
	default:
		return "unknown"
	}
}

// ClampEvent identifies one place that a transition clamped.
type ClampEvent struct {
	Trans string
	Place string
	Bound ClampBound
	Limit MarkInt
}

func (e ClampEvent) String() string {
	return fmt.Sprintf("transition %s clamped place %s at %s %d", e.Trans, e.Place, e.Bound, e.Limit)
}

// ClampSummary is a ClampEvent together with how often it occurred during an analysis.
type ClampSummary struct {
	ClampEvent
	Count int
}

func (s ClampSummary) String() string {
	return fmt.Sprintf("%s (%d time(s))", s.ClampEvent, s.Count)
}

// ClampError is the error returned by DoFiring when the destination marking had to be
// clamped. It carries every place that was clamped by that single firing.
type ClampError struct {
	Events []ClampEvent
}

func (e *ClampError) Error() string {
	msgs := make([]string, 0, len(e.Events))
	for _, ev := range e.Events {
		msgs = append(msgs, ev.String())
	}
	return strings.Join(msgs, "; ")
}

func (e *ClampError) add(ev ClampEvent) {
	e.Events = append(e.Events, ev)
}

// clampRecorder accumulates the clamping reported by DoFiring over a whole analysis,
// collapsing repeats of the same (transition, place, bound). It is not safe for
// concurrent use, like the rest of the search state it lives in.
type clampRecorder struct {
	index map[ClampEvent]int
	list  []ClampSummary
}

// record takes the error returned by DoFiring. Anything that is not a *ClampError,
// including nil, is ignored.
func (r *clampRecorder) record(err error) {
	var ce *ClampError
	if !errors.As(err, &ce) {
		return
	}
	for _, ev := range ce.Events {
		if i, ok := r.index[ev]; ok {
			r.list[i].Count++
			continue
		}
		if r.index == nil {
			r.index = make(map[ClampEvent]int)
		}
		r.index[ev] = len(r.list)
		r.list = append(r.list, ClampSummary{ClampEvent: ev, Count: 1})
	}
}

// events returns the clamping seen so far, in the order it was first seen.
func (r *clampRecorder) events() []ClampSummary {
	if len(r.list) == 0 {
		return nil
	}
	out := make([]ClampSummary, len(r.list))
	copy(out, r.list)
	return out
}

// FormatClampEvents renders a report for a caller that wants to show the user what was
// clamped. It returns "" when nothing was clamped.
func FormatClampEvents(events []ClampSummary) string {
	if len(events) == 0 {
		return ""
	}
	var b strings.Builder
	fmt.Fprintf(&b, "%d place(s) were clamped during firing; the result is not exact.\n", len(events))
	fmt.Fprintf(&b, "Give the place an explicit `max` if it is meant to hold more tokens "+
		"(the default capacity is 255).\n")
	for _, e := range events {
		fmt.Fprintf(&b, "  %s\n", e)
	}
	return b.String()
}
