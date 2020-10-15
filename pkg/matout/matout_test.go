package matout

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"runtime"
	"testing"
	"time"
)

func TestHeaderString1(t *testing.T) {
	header := MATLABHeader{
		level:           "5.0",
		platform:        runtime.GOOS,
		created:         time.Now(),
		version:         0x0100,
		endianIndicator: "IM",
	}
	buf := bytes.NewBuffer(make([]byte, 0, 128))
	header.ToBytes(NewMATLABBuffer(buf, binary.LittleEndian))
	fmt.Println(buf.String())
	fmt.Println(buf.Bytes())
	fmt.Println(buf.Len())
}

func TestHeaderString2(t *testing.T) {
	header := MATLABHeader{
		level:           "5.0",
		platform:        runtime.GOOS,
		created:         time.Now(),
		version:         0x0100,
		endianIndicator: "MI",
	}
	buf := bytes.NewBuffer(make([]byte, 0, 128))
	header.ToBytes(NewMATLABBuffer(buf, binary.BigEndian))
	fmt.Println(buf.String())
	fmt.Println(buf.Bytes())
	fmt.Println(buf.Len())
}

func TestStringData1(t *testing.T) {
	data := CreateMATLABDataString("test string")
	buf := bytes.NewBuffer(make([]byte, 0, 128))
	data.ToBytes(NewMATLABBuffer(buf, binary.BigEndian))
	fmt.Println(buf.String())
	fmt.Println(buf.Bytes())
	fmt.Println(buf.Len())
}

func TestStringData2(t *testing.T) {
	data := CreateMATLABDataString("a")
	buf := bytes.NewBuffer(make([]byte, 0, 128))
	data.ToBytes(NewMATLABBuffer(buf, binary.BigEndian))
	fmt.Println(buf.String())
	fmt.Println(buf.Bytes())
	fmt.Println(buf.Len())
}

func TestArrayData1(t *testing.T) {
	data := CreateMATLABArray([]int32{1, 2, 3, 4, 5})
	buf := bytes.NewBuffer(make([]byte, 0, 128))
	data.ToBytes(NewMATLABBuffer(buf, binary.BigEndian))
	fmt.Println(buf.Bytes())
	fmt.Println(buf.Len())
}

func TestArrayData2(t *testing.T) {
	d1 := []float64{1.0, 2, 3, 4, 5}
	data := CreateMATLABArray(d1)
	buf := bytes.NewBuffer(make([]byte, 0, 128))
	data.ToBytes(NewMATLABBuffer(buf, binary.BigEndian))
	fmt.Println(buf.Bytes())
	fmt.Println(buf.Len())
}

func TestArrayData3(t *testing.T) {
	d1 := []int64{1, 2, 3, 4, 5}
	data := CreateMATLABArray(d1)
	buf := bytes.NewBuffer(make([]byte, 0, 128))
	data.ToBytes(NewMATLABBuffer(buf, binary.BigEndian))
	fmt.Println(buf.Bytes())
	fmt.Println(buf.Len())
}

func TestArrayData4(t *testing.T) {
	d1 := uint64(100)
	data := CreateMATLABArray(d1)
	buf := bytes.NewBuffer(make([]byte, 0, 128))
	data.ToBytes(NewMATLABBuffer(buf, binary.BigEndian))
	fmt.Println(buf.Bytes())
	fmt.Println(buf.Len())
}

func TestArrayData5(t *testing.T) {
	d1 := []int32{}
	data := CreateMATLABArray(d1)
	buf := bytes.NewBuffer(make([]byte, 0, 128))
	data.ToBytes(NewMATLABBuffer(buf, binary.BigEndian))
	fmt.Println(buf.Bytes())
	fmt.Println(buf.Len())
}

func TestMatrixData1(t *testing.T) {
	d1 := []int32{1, 2, 3, 4, 5}
	data := CreateMATLABMatrix(len(d1), "test", d1)
	buf := bytes.NewBuffer(make([]byte, 0, 128))
	data.ToBytes(NewMATLABBuffer(buf, binary.BigEndian))
	fmt.Println(buf.Bytes())
	fmt.Println(buf.Len())
}

func TestFullMatrixData1(t *testing.T) {
	matfile := CreateMATLABMatFile(true)
	d1 := []int32{1, 2, 3, 4, 5}
	data := CreateMATLABMatrix(len(d1), "test", d1)
	matfile.AddElement(data)
	buf := bytes.NewBuffer(make([]byte, 0, 128))
	matfile.ToBytes(NewMATLABBuffer(buf, binary.BigEndian))
	fmt.Println(buf.Bytes())
	fmt.Println(buf.Len())
}

func TestFullMatrixData2(t *testing.T) {
	matfile := CreateMATLABMatFile(true)
	d1 := []int32{1, 2, 3, 4, 5}
	data := CreateMATLABMatrix(len(d1), "test1", d1)
	matfile.AddElement(data)
	d2 := []float64{1, 2, 3, 4, 5}
	data2 := CreateMATLABMatrix(len(d2), "test2", d2)
	matfile.AddElement(data2)
	buf := bytes.NewBuffer(make([]byte, 0, 128))
	matfile.ToBytes(NewMATLABBuffer(buf, binary.BigEndian))
	fmt.Println(buf.Bytes())
	fmt.Println(buf.Len())
}

func TestBufio(t *testing.T) {
	matfile := CreateMATLABMatFile(true)
	d1 := []int64{1, 2, 3, 4, 5}
	data := CreateMATLABMatrix(len(d1), "test1", d1)
	matfile.AddElement(data)
	d2 := []float64{1, 2, 3, 4, 5}
	data2 := CreateMATLABMatrix(len(d2), "test2", d2)
	matfile.AddElement(data2)

	file, err := os.Create("tmp.mat")
	if err != nil {
		t.Errorf("error")
	}
	defer file.Close()

	buf := bufio.NewWriter(file)
	matfile.ToBytes(NewMATLABBuffer(buf, binary.LittleEndian)) // binary.BigEndian causes error when it reads
	buf.Flush()

	// # in julia
	// using MAT
	// file = matopen("tmp.mat")
	// read(file, "test1")
	// read(file, "test2")
	// close(file)
}

func TestBufio2(t *testing.T) {
	matfile := CreateMATLABMatFile(true)
	d1 := []int64{1, 2, 3, 4, 5, 10}
	data := CreateMATLABMatrix([]int32{3, 2}, "test1", d1)
	matfile.AddElement(data)
	d2 := []float32{1, 2, 3, 4, 5, 8}
	data2 := CreateMATLABMatrix([]int32{2, 3}, "test2", d2)
	matfile.AddElement(data2)

	file, err := os.Create("tmp.mat")
	if err != nil {
		t.Errorf("error")
	}
	defer file.Close()

	buf := bufio.NewWriter(file)
	matfile.ToBytes(NewMATLABBuffer(buf, binary.LittleEndian)) // binary.BigEndian causes error when it reads
	buf.Flush()

	// # in julia
	// using MAT
	// file = matopen("tmp.mat")
	// read(file, "test1")
	// read(file, "test2")
	// close(file)
}

func TestBufio3(t *testing.T) {
	matfile := CreateMATLABMatFile(true)
	rowind := []int{0, 1, 2, 0, 2, 1}
	colptr := []int{0, 3, 5, 6}
	val := []float64{3, 4, 5, 1, 1, 5}
	data := CreateMATLABSparseMatrix([]int32{3, 3}, "sp1", len(val), rowind, colptr, val)
	matfile.AddElement(data)

	file, err := os.Create("tmp.mat")
	if err != nil {
		t.Errorf("error")
	}
	defer file.Close()

	buf := bufio.NewWriter(file)
	matfile.ToBytes(NewMATLABBuffer(buf, binary.LittleEndian)) // binary.BigEndian causes error when it reads
	buf.Flush()

	// # in julia
	// using MAT
	// file = matopen("tmp.mat")
	// read(file, "sp1")
	// close(file)
}
