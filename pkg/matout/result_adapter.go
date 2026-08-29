package matout

import (
	"bufio"
	"encoding/binary"
	"io"

	"github.com/okamumu/gospn/pkg/result"
)

// WriteResult writes a result as a MATLAB v5 .mat file. It is the adapter that keeps
// the subcommands from assembling MATLAB elements by hand; the element constructors
// below it are unchanged and still the way to build a .mat directly.
func WriteResult(w io.Writer, r *result.Result) error {
	if err := r.Validate(); err != nil {
		return err
	}
	matfile := CreateMATLABMatFile(true)
	for _, e := range r.Elements {
		switch e.Kind {
		case result.KindSparse:
			matfile.AddElement(CreateMATLABSparseMatrix(e.Dims, e.Name, e.NNZ, e.Ir, e.Jc, e.Values))
		case result.KindDense:
			matfile.AddElement(CreateMATLABMatrix(e.Dims, e.Name, e.Values))
		case result.KindText:
			matfile.AddElement(CreateMATLABCharMatrix(e.Name, e.Text))
		}
	}
	buf := bufio.NewWriter(w)
	matfile.ToBytes(NewMATLABBuffer(buf, binary.LittleEndian))
	return buf.Flush()
}
