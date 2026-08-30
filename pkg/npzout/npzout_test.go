package npzout

import (
	"archive/zip"
	"bytes"
	"encoding/binary"
	"io"
	"sort"
	"strings"
	"testing"

	"github.com/okamumu/gospn/pkg/result"
)

func write(t *testing.T, r *result.Result) *zip.Reader {
	t.Helper()
	var buf bytes.Buffer
	if err := Write(&buf, r); err != nil {
		t.Fatal(err)
	}
	z, err := zip.NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	if err != nil {
		t.Fatal(err)
	}
	return z
}

// readMember returns the .npy header dict and the raw data of one archive member.
func readMember(t *testing.T, z *zip.Reader, name string) (string, []byte) {
	t.Helper()
	for _, f := range z.File {
		if f.Name != name+".npy" {
			continue
		}
		rc, err := f.Open()
		if err != nil {
			t.Fatal(err)
		}
		defer rc.Close()
		b, err := io.ReadAll(rc)
		if err != nil {
			t.Fatal(err)
		}
		if string(b[:8]) != "\x93NUMPY\x01\x00" {
			t.Fatalf("%s: bad magic %q", name, b[:8])
		}
		n := int(binary.LittleEndian.Uint16(b[8:10]))
		// The data has to start on a 64-byte boundary or NumPy's mmap path misreads it.
		if (10+n)%64 != 0 {
			t.Errorf("%s: preamble is %d bytes, not a multiple of 64", name, 10+n)
		}
		if b[10+n-1] != '\n' {
			t.Errorf("%s: header does not end with a newline", name)
		}
		return string(b[10 : 10+n]), b[10+n:]
	}
	t.Fatalf("member %q not found in %v", name, names(z))
	return "", nil
}

func names(z *zip.Reader) []string {
	var out []string
	for _, f := range z.File {
		out = append(out, f.Name)
	}
	sort.Strings(out)
	return out
}

func TestSparseIsSplitIntoCSCMembers(t *testing.T) {
	r := &result.Result{}
	r.AddSparse("A", []int32{2, 3}, 2, []int{0, 1}, []int{0, 1, 2, 2}, []float64{1.5, 2.5})
	z := write(t, r)

	hdr, data := readMember(t, z, "A.data")
	if !strings.Contains(hdr, "'descr': '<f8'") || !strings.Contains(hdr, "'shape': (2,)") {
		t.Errorf("A.data header = %q", hdr)
	}
	var got [2]float64
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &got); err != nil {
		t.Fatal(err)
	}
	if got != [2]float64{1.5, 2.5} {
		t.Errorf("A.data = %v", got)
	}

	hdr, _ = readMember(t, z, "A.indices")
	if !strings.Contains(hdr, "'descr': '<i8'") {
		t.Errorf("A.indices header = %q", hdr)
	}
	readMember(t, z, "A.indptr")

	_, data = readMember(t, z, "A.shape")
	var shape [2]int64
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &shape); err != nil {
		t.Fatal(err)
	}
	if shape != [2]int64{2, 3} {
		t.Errorf("A.shape = %v, want [2 3]", shape)
	}
}

func TestDenseAndTextMembers(t *testing.T) {
	r := &result.Result{}
	r.AddDense("v", []float64{1, 2, 3})
	r.AddDense("count", []int32{7})
	r.AddText("gospn_version", "1.2.3")
	z := write(t, r)

	hdr, _ := readMember(t, z, "v")
	if !strings.Contains(hdr, "'shape': (3,)") {
		t.Errorf("v header = %q", hdr)
	}
	if hdr, _ = readMember(t, z, "count"); !strings.Contains(hdr, "'descr': '<i4'") {
		t.Errorf("count header = %q", hdr)
	}

	// A string is a uint8 array of UTF-8 bytes. NumPy's '<U' would be nicer to read in
	// Python but makes NPZ.jl reject the entire archive.
	hdr, data := readMember(t, z, "gospn_version")
	if !strings.Contains(hdr, "'descr': '|u1'") || !strings.Contains(hdr, "'shape': (5,)") {
		t.Errorf("gospn_version header = %q", hdr)
	}
	if string(data) != "1.2.3" {
		t.Errorf("gospn_version = %q", data)
	}
}

// A 1-by-n row vector should load as a 1-D array; indexing a (1, n) with [i] in Python
// gives a row, not a number.
func TestRowVectorLoadsAsOneDimensional(t *testing.T) {
	r := &result.Result{}
	r.Elements = append(r.Elements, result.Element{
		Name: "row", Kind: result.KindDense, Dims: []int32{1, 3}, Values: []float64{1, 2, 3},
	})
	hdr, _ := readMember(t, write(t, r), "row")
	if !strings.Contains(hdr, "'shape': (3,)") {
		t.Errorf("row header = %q", hdr)
	}
}

func TestWriteRejectsDuplicateNames(t *testing.T) {
	r := &result.Result{}
	r.AddDense("x", []float64{1})
	r.AddDense("x", []float64{2})
	if err := Write(&bytes.Buffer{}, r); err == nil {
		t.Fatal("duplicate name accepted")
	}
}

// A 2-D dense element is stored column-major, so the header has to say so: NumPy and
// Julia both read the buffer in the order fortran_order names, and getting it wrong
// transposes the matrix silently.
func TestDenseMatrixIsFortranOrdered(t *testing.T) {
	r := &result.Result{}
	r.AddDenseMatrix("path_state", 3, 2, []int32{0, 1, 2, 3, 4, 5})
	r.AddDense("path_time", []float64{0, 1, 2})
	z := write(t, r)

	hdr, data := readMember(t, z, "path_state")
	if !strings.Contains(hdr, "'fortran_order': True") {
		t.Errorf("path_state: header is %q", hdr)
	}
	if !strings.Contains(hdr, "'shape': (3,2,)") {
		t.Errorf("path_state: shape missing from %q", hdr)
	}
	if len(data) != 6*4 {
		t.Errorf("path_state: %d bytes, want %d", len(data), 6*4)
	}

	// A vector is not column-major, and saying it is would be a lie a reader can see.
	hdr, _ = readMember(t, z, "path_time")
	if !strings.Contains(hdr, "'fortran_order': False") {
		t.Errorf("path_time: header is %q", hdr)
	}
}
