// Package jsonout writes a result as JSON.
//
// It is the format that can be read without a library and diffed in a review: a .mat
// or .npz says nothing to `git diff`, which is why a variable-name collision in the
// simulation output survived several releases unnoticed.
package jsonout

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/okamumu/gospn/pkg/result"
)

// FormatVersion is bumped when the shape of the document changes incompatibly.
const FormatVersion = 1

type document struct {
	Format   string    `json:"format"`
	Version  int       `json:"version"`
	Elements []element `json:"elements"`
}

type element struct {
	Name   string      `json:"name"`
	Kind   string      `json:"kind"`
	Dims   []int32     `json:"dims,omitempty"`
	DType  string      `json:"dtype,omitempty"`
	NNZ    int         `json:"nnz,omitempty"`
	RowInd []int       `json:"rowind,omitempty"`
	ColPtr []int       `json:"colptr,omitempty"`
	Values interface{} `json:"values,omitempty"`
	Text   string      `json:"text,omitempty"`
}

// Write writes r as an indented JSON document.
func Write(w io.Writer, r *result.Result) error {
	if err := r.Validate(); err != nil {
		return err
	}
	doc := document{Format: "gospn-result", Version: FormatVersion}
	for _, e := range r.Elements {
		switch e.Kind {
		case result.KindSparse:
			dt, err := dtype(e.Values)
			if err != nil {
				return fmt.Errorf("%s: %v", e.Name, err)
			}
			doc.Elements = append(doc.Elements, element{
				Name: e.Name, Kind: "sparse", Dims: e.Dims, DType: dt,
				NNZ: e.NNZ, RowInd: e.Ir, ColPtr: e.Jc, Values: e.Values,
			})
		case result.KindDense:
			dt, err := dtype(e.Values)
			if err != nil {
				return fmt.Errorf("%s: %v", e.Name, err)
			}
			doc.Elements = append(doc.Elements, element{
				Name: e.Name, Kind: "dense", Dims: e.Dims, DType: dt, Values: e.Values,
			})
		case result.KindText:
			doc.Elements = append(doc.Elements, element{
				Name: e.Name, Kind: "text", Text: e.Text,
			})
		}
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(doc)
}

func dtype(values interface{}) (string, error) {
	switch values.(type) {
	case []float64:
		return "float64", nil
	case []int32:
		return "int32", nil
	case []int64:
		return "int64", nil
	default:
		return "", fmt.Errorf("unsupported value type %T", values)
	}
}
