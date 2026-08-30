// Package analysis turns what the engine computed into a format-neutral
// result.Result. It used to live inline in cmd/gospn.go, where nothing could reach it:
// cmd is package main, so no test could run a subcommand and look at what it produced.
package analysis

import (
	"fmt"

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
