// Package result is the in-memory form of what an analysis produced, sitting between
// the analysis engine and the output writers.
//
// Before it existed, cmd/gospn.go built MATLAB elements directly, so the file format
// was decided in the same place the numbers were collected and adding a second format
// meant duplicating that assembly. A Result is format-neutral: the subcommands fill one
// in, and a Writer turns it into .mat, .npz or .json.
package result

import (
	"fmt"
	"sort"
)

// Kind says how an element's payload is laid out. It is not a data type: the element
// type of Values is carried by Values itself.
type Kind int

const (
	// KindDense is a full array; Values holds every entry.
	KindDense Kind = iota
	// KindSparse is CSC with a 0-origin, exactly as MarkingGraph.CSC reports it.
	KindSparse
	// KindText is a string; Values is nil and Text holds it.
	KindText
)

// Element is one named value in a result file. The name is what the reader will see:
// MATLAB variable, npz key, JSON object key.
type Element struct {
	Name string
	Kind Kind

	// Dims is the shape. gospn writes vectors with a single dimension, which is what
	// the MATLAB writer has always been handed, so this is not forced to length 2.
	Dims []int32

	// Sparse only. Ir is rowind, Jc is colptr, both 0-origin.
	NNZ int
	Ir  []int
	Jc  []int

	// Dense only: []float64, []int32 or []int64.
	Values interface{}

	// Text only.
	Text string
}

// Result is an ordered list of elements. Order is preserved because it is what the
// writers emit, and a stable order is what makes a JSON golden diffable.
type Result struct {
	Elements []Element
}

// Len is the number of elements added so far. Taking it before adding a group and
// passing it to SortFrom is how a subcommand makes a map's worth of output ordered.
func (r *Result) Len() int { return len(r.Elements) }

// SortFrom sorts elements from index i onwards by name. Go map iteration order is
// random, so without this the same net produces a differently ordered file on every
// run -- harmless for MATLAB, fatal for a text golden.
func (r *Result) SortFrom(i int) {
	tail := r.Elements[i:]
	sort.Slice(tail, func(a, b int) bool { return tail[a].Name < tail[b].Name })
}

// AddDense appends a vector. values must be []float64, []int32 or []int64.
func (r *Result) AddDense(name string, values interface{}) {
	n, err := valuesLen(values)
	if err != nil {
		panic(fmt.Sprintf("result: %s: %v", name, err))
	}
	r.Elements = append(r.Elements, Element{
		Name: name, Kind: KindDense, Dims: []int32{int32(n)}, Values: values,
	})
}

// AddSparse appends a CSC matrix. dims is {rows, cols}.
func (r *Result) AddSparse(name string, dims []int32, nnz int, ir, jc []int, pr []float64) {
	r.Elements = append(r.Elements, Element{
		Name: name, Kind: KindSparse, Dims: dims, NNZ: nnz, Ir: ir, Jc: jc, Values: pr,
	})
}

// AddText appends a string element.
func (r *Result) AddText(name, text string) {
	r.Elements = append(r.Elements, Element{
		Name: name, Kind: KindText, Dims: []int32{1, int32(len([]rune(text)))}, Text: text,
	})
}

// Validate reports a duplicate name. MATLAB's `load` silently keeps the last variable
// of a given name, which is how `gospn sim` lost its instantaneous reward vector for
// several releases without any test noticing; npz and JSON lose it just as quietly.
func (r *Result) Validate() error {
	seen := make(map[string]bool, len(r.Elements))
	for _, e := range r.Elements {
		if seen[e.Name] {
			return fmt.Errorf("result: duplicate element name %q", e.Name)
		}
		seen[e.Name] = true
	}
	return nil
}

func valuesLen(values interface{}) (int, error) {
	switch v := values.(type) {
	case []float64:
		return len(v), nil
	case []int32:
		return len(v), nil
	case []int64:
		return len(v), nil
	default:
		return 0, fmt.Errorf("unsupported value type %T (want []float64, []int32 or []int64)", values)
	}
}
