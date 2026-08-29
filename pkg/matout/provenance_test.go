package matout

import (
	"bytes"
	"encoding/binary"
	"strings"
	"testing"
)

// A char matrix must land in the file as a named mxCHAR_CLASS array; the check that
// matters for its purpose is that the name and the text are both recoverable.
func TestCharMatrixRoundsIntoTheFile(t *testing.T) {
	f := CreateMATLABMatFile(true)
	f.AddElement(CreateMATLABCharMatrix("gospn_version", "0.17.0"))
	f.AddElement(CreateMATLABMatrix(3, "vec", []float64{1, 2, 3}))
	var buf bytes.Buffer
	f.ToBytes(NewMATLABBuffer(&buf, binary.LittleEndian))
	b := buf.Bytes()
	if !bytes.Contains(b, []byte("gospn_version")) {
		t.Error("the variable name is not in the file")
	}
	// UTF-16LE of "0.17.0"
	want := make([]byte, 0, 12)
	for _, r := range "0.17.0" {
		want = append(want, byte(r), 0)
	}
	if !bytes.Contains(b, want) {
		t.Error("the string value is not in the file as UTF-16")
	}
	if !bytes.Contains(b, []byte("vec")) {
		t.Error("a numeric matrix alongside it went missing")
	}
	if strings.Count(string(b), "gospn_version") != 1 {
		t.Error("the name should appear once")
	}
}

// The reward vectors used to collide: lastrwd was written under the same name as
// irwd, so loading the file gave back lastrwd for both and the instantaneous reward
// could not be recovered at all. A file must not name the same variable twice.
func TestDuplicateNamesAreVisible(t *testing.T) {
	f := CreateMATLABMatFile(true)
	f.AddElement(CreateMATLABMatrix(2, "avail_irwd", []float64{1, 2}))
	f.AddElement(CreateMATLABMatrix(2, "avail_crwd", []float64{3, 4}))
	f.AddElement(CreateMATLABMatrix(2, "avail_lastrwd", []float64{5, 6}))
	var buf bytes.Buffer
	f.ToBytes(NewMATLABBuffer(&buf, binary.LittleEndian))
	s := string(buf.Bytes())
	for _, name := range []string{"avail_irwd", "avail_crwd", "avail_lastrwd"} {
		if n := strings.Count(s, name); n != 1 {
			t.Errorf("%s appears %d times, want 1", name, n)
		}
	}
}
