package matout

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/okamumu/gospn/pkg/result"
)

// The adapter must produce exactly what the hand-assembled matfile did; that is what
// makes the format-neutral Result safe to route the subcommands through.
func TestWriteResultMatchesHandBuiltMatFile(t *testing.T) {
	r := &result.Result{}
	r.AddSparse("A", []int32{2, 2}, 2, []int{0, 1}, []int{0, 1, 2}, []float64{1.5, 2.5})
	r.AddDense("initA", []float64{1, 0})
	r.AddDense("count", []int32{7})
	r.AddText("gospn_version", "1.2.3")

	var got bytes.Buffer
	if err := WriteResult(&got, r); err != nil {
		t.Fatal(err)
	}

	matfile := CreateMATLABMatFile(true)
	matfile.AddElement(CreateMATLABSparseMatrix([]int32{2, 2}, "A", 2, []int{0, 1}, []int{0, 1, 2}, []float64{1.5, 2.5}))
	matfile.AddElement(CreateMATLABMatrix([]int32{2}, "initA", []float64{1, 0}))
	matfile.AddElement(CreateMATLABMatrix([]int32{1}, "count", []int32{7}))
	matfile.AddElement(CreateMATLABCharMatrix("gospn_version", "1.2.3"))
	var want bytes.Buffer
	w := bufio.NewWriter(&want)
	matfile.ToBytes(NewMATLABBuffer(w, binary.LittleEndian))
	w.Flush()

	// The header carries a timestamp, so compare everything after it.
	const headerLen = 128
	if !bytes.Equal(got.Bytes()[headerLen:], want.Bytes()[headerLen:]) {
		t.Errorf("adapter output differs from the hand-built matfile (%d vs %d bytes)", got.Len(), want.Len())
	}
}

func TestWriteResultRejectsDuplicateNames(t *testing.T) {
	r := &result.Result{}
	r.AddDense("x", []float64{1})
	r.AddDense("x", []float64{2})
	if err := WriteResult(&bytes.Buffer{}, r); err == nil {
		t.Fatal("duplicate name accepted")
	}
}
