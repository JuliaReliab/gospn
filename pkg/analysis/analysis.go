// Package analysis turns what the engine computed into a format-neutral
// result.Result. It used to live inline in cmd/gospn.go, where nothing could reach it:
// cmd is package main, so no test could run a subcommand and look at what it produced.
package analysis

import (
	"fmt"
	"strings"

	"github.com/okamumu/gospn/pkg/petrinet"
	"github.com/okamumu/gospn/pkg/result"
)

// MarkResult collects the matrices and vectors of a marking graph.
//
// Go map iteration order is random, so every group of elements is sorted by name
// before the next one starts: the same net has to produce the same file twice.
func MarkResult(mg *petrinet.MarkingGraph) *result.Result {
	res := &result.Result{}
	expmat, immmat, genmat := mg.TransMatrix()
	grouplabel := mg.GroupLabels()
	grouptranslabel := mg.TransLabels()
	for _, mats := range []map[petrinet.GroupTrans]*petrinet.CSC{expmat, immmat, genmat} {
		first := res.Len()
		for tr, m := range mats {
			label := fmt.Sprintf("%s%s%s", grouplabel[tr.GetSrc()], grouplabel[tr.GetDest()], grouptranslabel[tr])
			dim, nnz, rowind, colptr, val := m.Get()
			res.AddSparse(label, dim, nnz, rowind, colptr, val)
		}
		res.SortFrom(first)
	}
	first := res.Len()
	for g, v := range mg.InitVector() {
		res.AddDense(fmt.Sprintf("init%s", grouplabel[g]), v)
	}
	res.SortFrom(first)
	first = res.Len()
	for rewardlabel, rv := range mg.RewardVector() {
		for g, v := range rv {
			res.AddDense(fmt.Sprintf("%s%s", rewardlabel, grouplabel[g]), v)
		}
	}
	res.SortFrom(first)

	// Which marking each row is. Without this the file carries matrices and nothing
	// that says what a row means: the markings went to a separate text file, and only
	// when -s was given. `place` is the column order, shared by every mark<G>.
	//
	// A reward called "mark" would collide with these names; Result.Validate refuses
	// the file rather than letting a writer silently keep the last of the two.
	res.AddText("place", strings.Join(mg.Net().PlaceLabels(), "\n"))
	first = res.Len()
	for g, marks := range mg.StateMarkings() {
		nplaces := len(mg.Net().PlaceLabels())
		values := make([]int32, len(marks)*nplaces)
		for k, m := range marks {
			for i, n := range m {
				values[i*len(marks)+k] = int32(n)
			}
		}
		res.AddDenseMatrix(fmt.Sprintf("mark%s", grouplabel[g]), len(marks), nplaces, values)
	}
	res.SortFrom(first)
	return res
}

// SimRun is everything one `gospn sim` run produced, in the order RunAll returns it.
type SimRun struct {
	Irwd        map[string][]float64
	Crwd        map[string][]float64
	Lastrwd     map[string][]float64
	ElapsedTime []float64
	Count       []int32
	Config      petrinet.PNSimConfig
	Seed        int64
	Clamped     int
}

// SimResult collects the rewards of a simulation, followed by what the run needs to be
// repeated. The seed and the horizon are what identify a run, and the version matters
// because 0.16.0 changed which random stream a given seed produces -- but the version
// is provenance, added by the caller.
func SimResult(run SimRun) *result.Result {
	res := &result.Result{}
	for _, rwd := range []struct {
		suffix string
		values map[string][]float64
	}{{"_irwd", run.Irwd}, {"_crwd", run.Crwd}, {"_lastrwd", run.Lastrwd}} {
		first := res.Len()
		for rlabel, v := range rwd.values {
			res.AddDense(rlabel+rwd.suffix, v)
		}
		res.SortFrom(first)
	}
	res.AddDense("elapsedtime", run.ElapsedTime)
	res.AddDense("count", run.Count)
	return res
}

// SimRunInfo adds the scalars that describe the run. It is separate from SimResult
// because provenance goes between the two, and the element order is part of the file.
func SimRunInfo(res *result.Result, run SimRun) {
	res.AddDense("seed", []int64{run.Seed})
	res.AddDense("simulations", []int32{int32(run.Config.NumOfSimulation)})
	res.AddDense("endingtime", []float64{run.Config.EndingTime})
	res.AddDense("firings", []int32{run.Config.NumOfFiring})
	res.AddDense("parallel", []int32{int32(run.Config.Parallel)})

	// Clamping is reported on stderr, but the fact that a run was not exact has to
	// travel with the data as well -- whoever reads the file later did not see it.
	res.AddDense("clamped", []int32{int32(run.Clamped)})
}

// PathStep is one firing of a simulated path. petrinet's own event type is unexported,
// so the caller converts; the accessors on it are Time, Mark and TransLabel.
type PathStep struct {
	Time  float64
	Mark  []petrinet.MarkInt
	Trans string // empty for the initial marking, which nothing fired to reach
}

// PathResult collects one simulated path: the marking after every firing, when it
// happened, and what fired. `gospn test` printed this to stdout and nothing else,
// which made it the one subcommand whose output could not be read back by anything.
//
// Marking column i is place PlaceLabels()[i]; the states are stored column-major, so
// column i of the matrix is that place's token count over the whole path.
func PathResult(net *petrinet.Net, path []PathStep) *result.Result {
	labels := net.PlaceLabels()
	res := &result.Result{}
	res.AddText("place", strings.Join(labels, "\n"))

	times := make([]float64, len(path))
	trans := make([]string, len(path))
	states := make([]int32, len(path)*len(labels))
	for k, e := range path {
		times[k] = e.Time
		trans[k] = e.Trans
		for i, n := range e.Mark {
			states[i*len(path)+k] = int32(n)
		}
	}
	res.AddDense("path_time", times)
	res.AddText("path_trans", strings.Join(trans, "\n"))
	res.AddDenseMatrix("path_state", len(path), len(labels), states)
	return res
}
