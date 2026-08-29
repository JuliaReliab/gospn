// Package npzout writes a result as a NumPy .npz archive.
//
// An .npz is a zip of .npy members, so it needs nothing outside the standard library --
// which matters because gospn is built without cgo and cross-compiled in CI, ruling out
// HDF5 (and therefore MATLAB v7.3). On the reading side it is one np.load() call, and
// Julia reads it with NPZ.jl.
//
// A sparse matrix named M becomes four members:
//
//	M.data     float64, the nonzeros
//	M.indices  int64, row indices (CSC, 0-origin)
//	M.indptr   int64, column pointers
//	M.shape    int64, {rows, cols}
//
// so that in Python:
//
//	z = np.load("out.npz")
//	A = scipy.sparse.csc_matrix((z["M.data"], z["M.indices"], z["M.indptr"]), shape=tuple(z["M.shape"]))
//
// Dense vectors and strings are single members under their own name.
package npzout

import (
	"archive/zip"
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/okamumu/gospn/pkg/result"
)

// Write writes r as an .npz archive.
func Write(w io.Writer, r *result.Result) error {
	if err := r.Validate(); err != nil {
		return err
	}
	z := zip.NewWriter(w)
	for _, e := range r.Elements {
		var err error
		switch e.Kind {
		case result.KindSparse:
			rows, cols := int64(0), int64(0)
			if len(e.Dims) > 0 {
				rows = int64(e.Dims[0])
			}
			if len(e.Dims) > 1 {
				cols = int64(e.Dims[1])
			}
			if err = writeArray(z, e.Name+".data", []int{len(e.Ir)}, e.Values); err == nil {
				if err = writeArray(z, e.Name+".indices", []int{len(e.Ir)}, toInt64(e.Ir)); err == nil {
					if err = writeArray(z, e.Name+".indptr", []int{len(e.Jc)}, toInt64(e.Jc)); err == nil {
						err = writeArray(z, e.Name+".shape", []int{2}, []int64{rows, cols})
					}
				}
			}
		case result.KindDense:
			err = writeArray(z, e.Name, dimsToShape(e.Dims), e.Values)
		case result.KindText:
			err = writeString(z, e.Name, e.Text)
		}
		if err != nil {
			return fmt.Errorf("%s: %v", e.Name, err)
		}
	}
	return z.Close()
}

// A member is stored as <name>.npy because that is the suffix np.load strips to make
// the key; a member without it is not loadable as an array.
func member(z *zip.Writer, name string) (io.Writer, error) {
	return z.CreateHeader(&zip.FileHeader{Name: name + ".npy", Method: zip.Deflate})
}

func writeArray(z *zip.Writer, name string, shape []int, values interface{}) error {
	descr, err := descrOf(values)
	if err != nil {
		return err
	}
	w, err := member(z, name)
	if err != nil {
		return err
	}
	if err := writeHeader(w, descr, shape); err != nil {
		return err
	}
	return binary.Write(w, binary.LittleEndian, values)
}

// writeString stores a string as a uint8 array of its UTF-8 bytes:
//
//	Python: z["net"].tobytes().decode()
//	Julia:  String(z["net"])
//
// NumPy's own '<U' dtype would read back as a plain str in Python and is the obvious
// choice there, but NPZ.jl cannot parse it and fails the *whole* archive with
// "unsupported type U3", not just that one member. A format that only one of the two
// intended readers can open is not worth the nicer Python call.
func writeString(z *zip.Writer, name, s string) error {
	w, err := member(z, name)
	if err != nil {
		return err
	}
	data := []byte(s)
	if err := writeHeader(w, "|u1", []int{len(data)}); err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

// writeHeader writes the .npy v1.0 preamble: magic, version, a uint16 header length,
// and the dict, space-padded so that the whole preamble is a multiple of 64 bytes and
// the data that follows is aligned.
func writeHeader(w io.Writer, descr string, shape []int) error {
	var sb strings.Builder
	for _, n := range shape {
		fmt.Fprintf(&sb, "%d,", n)
	}
	dict := fmt.Sprintf("{'descr': '%s', 'fortran_order': False, 'shape': (%s), }", descr, sb.String())
	const preamble = 6 + 2 + 2 // magic + version + header length
	pad := 64 - (preamble+len(dict)+1)%64
	if pad == 64 {
		pad = 0
	}
	dict += strings.Repeat(" ", pad) + "\n"

	if _, err := io.WriteString(w, "\x93NUMPY\x01\x00"); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, uint16(len(dict))); err != nil {
		return err
	}
	_, err := io.WriteString(w, dict)
	return err
}

func descrOf(values interface{}) (string, error) {
	switch values.(type) {
	case []float64:
		return "<f8", nil
	case []int32:
		return "<i4", nil
	case []int64:
		return "<i8", nil
	default:
		return "", fmt.Errorf("unsupported value type %T", values)
	}
}

// dimsToShape drops a trailing length-1 dimension only when it is the leading one of a
// row vector, so that a vector gospn wrote as 1-by-n loads as a 1-D array rather than
// a 2-D one -- the shape a caller in Python or Julia expects to index with [i].
func dimsToShape(dims []int32) []int {
	if len(dims) == 2 && dims[0] == 1 {
		return []int{int(dims[1])}
	}
	shape := make([]int, len(dims))
	for i, d := range dims {
		shape[i] = int(d)
	}
	return shape
}

func toInt64(x []int) []int64 {
	y := make([]int64, len(x))
	for i, v := range x {
		y[i] = int64(v)
	}
	return y
}
