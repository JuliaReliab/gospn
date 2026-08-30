package petrinet

import (
	"fmt"
	"runtime"
)

// A reachability search has no natural stopping point: a net that is unbounded, or
// merely large, simply keeps producing markings until the process is killed. That is
// what `gospn mark` did on a net meant for simulation -- it printed "Create marking..."
// and was OOM-killed with no output, no state count and no clue how far it had got.
//
// So the search takes a limit. It is expressed in states rather than in bytes on
// purpose: a memory-triggered abort would make the same input behave differently on
// different machines, and reproducibility is worth more here than precision. The error
// reports the memory actually used, so a limit can be chosen for a given machine.

// DefaultMaxStates is the limit the search uses when none is given.
//
// The cost per state depends on the net -- on iaas_cloud.spn (17 places, 3.4 links per
// state) it is about 400 bytes retained and about 900 bytes at peak, so this limit is
// roughly 1 GB of peak heap there. A net with more places or a higher branching factor
// costs more. BenchmarkMarkingGraphMemory in test/ reports both numbers.
//
// Raise it with `gospn mark -maxstates`, or pass 0 for no limit. Enumerating
// iaas_cloud.spn with n = 6 (3.8 million states, about 3.5 GB) needs that.
const DefaultMaxStates = 1_000_000

// DefaultProgressEvery is how many newly discovered states pass between progress
// reports when a progress function is set.
const DefaultProgressEvery = 100_000

// SearchOptions configures a reachability search.
type SearchOptions struct {
	// MaxStates stops the search once this many distinct markings have been found.
	// Zero means no limit.
	MaxStates int

	// Progress, if set, is called with the number of states found so far, every
	// ProgressEvery states. A long search is otherwise completely silent.
	Progress func(states int)

	// ProgressEvery is the number of states between Progress calls. Zero selects
	// DefaultProgressEvery.
	ProgressEvery int
}

// DefaultSearchOptions is what the plain Create... entry points use.
func DefaultSearchOptions() SearchOptions {
	return SearchOptions{MaxStates: DefaultMaxStates}
}

func (o SearchOptions) progressEvery() int {
	if o.ProgressEvery > 0 {
		return o.ProgressEvery
	}
	return DefaultProgressEvery
}

// StateLimitError reports that the search stopped at its state limit. The marking
// graph it would have produced is incomplete, so no graph is returned with it: a
// truncated graph yields transition matrices that look valid and are not.
type StateLimitError struct {
	Found     int    // states discovered when the search stopped
	Limit     int    // the limit it hit
	HeapBytes uint64 // heap in use at that point, for choosing a limit
}

func (e *StateLimitError) Error() string {
	return fmt.Sprintf("the reachability search stopped after %d states, at the limit of %d "+
		"(%.1f MB of heap in use, about %.0f bytes per state).\n"+
		"Raise the limit with -maxstates, or pass -maxstates 0 to remove it.\n"+
		"If the net is unbounded, or is large enough that its state space cannot be "+
		"enumerated, analyse it with `gospn sim` instead of `gospn mark`.",
		e.Found, e.Limit, float64(e.HeapBytes)/(1<<20), float64(e.HeapBytes)/float64(e.Found))
}

// stateCounter tracks the search against its limit and drives progress reporting.
type stateCounter struct {
	opts       SearchOptions
	lastReport int
}

func newStateCounter(opts SearchOptions) *stateCounter {
	return &stateCounter{opts: opts}
}

// check is called after each newly discovered marking. It returns a *StateLimitError
// once the limit is reached, and nil otherwise.
func (c *stateCounter) check(found int) error {
	if c.opts.Progress != nil && found-c.lastReport >= c.opts.progressEvery() {
		c.lastReport = found
		c.opts.Progress(found)
	}
	if c.opts.MaxStates > 0 && found >= c.opts.MaxStates {
		return &StateLimitError{Found: found, Limit: c.opts.MaxStates, HeapBytes: heapInUse()}
	}
	return nil
}

// heapInUse reports the heap currently in use, for the limit error's benefit only.
func heapInUse() uint64 {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	return ms.HeapInuse
}
