package matout

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type MATLABDataType uint32

const (
	MATLABDataTypeUnknown MATLABDataType = iota
	miINT8
	miUINT8
	miINT16
	miUINT16
	miINT32
	miUINT32
	miSINGLE
	_
	miDOUBLE
	_
	_
	miINT64
	miUINT64
	miMATRIX
	miCOMPRESSED
	miUTF8
	miUTF16
	miUTF32
)

type MATLABArrayType uint32

const (
	MATLABArrayTypeUnknown MATLABArrayType = iota
	mxCELL_CLASS
	mxSTRUCT_CLASS
	mxOBJECT_CLASS
	mxCHAR_CLASS
	mxSPARSE_CLASS
	mxDOUBLE_CLASS
	mxSINGLE_CLASS
	mxINT8_CLASS
	mxUINT8_CLASS
	mxINT16_CLASS
	mxUINT16_CLASS
	mxINT32_CLASS
	mxUINT32_CLASS
	mxINT64_CLASS
	mxUINT64_CLASS
)

// The interface for Writer which is used in MATLABBuffer
type MATLABWriter interface {
	io.Writer
	WriteByte(x byte) error
	WriteString(x string) (int, error)
}

// The interface for data which can be written to the buffer. This is used in MATLABMatFile to represent data elements.
type MATLABByteInterface interface {
	byteLen() int                            // the method to get the size of bytes
	ToBytes(buf *MATLABBuffer) *MATLABBuffer // the method to write the buffer
}

// The structure of Buffer stream. `buf` is a pointer of bytes.Buffer. `endian` is to indicate LittleEndian or BigEndian.
// Note: BigEndian does not work until now.
type MATLABBuffer struct {
	buf    MATLABWriter
	endian binary.ByteOrder
}

// The function to create MATLABBuffer. `size` is the capacity of buffer.
func NewMATLABBuffer(buf MATLABWriter, endian binary.ByteOrder) *MATLABBuffer {
	return &MATLABBuffer{
		buf:    buf,
		endian: endian,
	}
}

// The method to write a string to the buffer
func (b *MATLABBuffer) WriteString(s string) *MATLABBuffer {
	b.buf.WriteString(s)
	return b
}

// The method to write data to the buffer. The data can be intXX, uintXX,
// floatXX and their slices, where XX is the number of bits.
func (b *MATLABBuffer) Write(data interface{}) *MATLABBuffer {
	binary.Write(b.buf, b.endian, data)
	return b
}

// The method to write a byte data.
func (b *MATLABBuffer) WriteByte(data byte) *MATLABBuffer {
	b.buf.WriteByte(data)
	return b
}

// The method to write zero to the buffer. numBytes is the number of bytes.
func (b *MATLABBuffer) Padding(numBytes int) *MATLABBuffer {
	for i := 0; i < numBytes; i++ {
		b.buf.WriteByte(0x00)
	}
	return b
}

// The structure of Matfile, which consists of header and elements. The elements are a list of data elements
// such as matrix and sparse matrix.
type MATLABMatFile struct {
	header   *MATLABHeader
	elements []MATLABByteInterface
	endian   binary.ByteOrder
}

// The function to provide the indicator of endianess in MatFile.
func getEndian(littleEndian bool) (string, binary.ByteOrder) {
	if littleEndian {
		return "IM", binary.LittleEndian
	} else {
		return "MI", binary.BigEndian
	}
}

// The function to create MatFile. `endian` is
func CreateMATLABMatFile(littleEndian bool) *MATLABMatFile {
	endianIndicator, endian := getEndian(littleEndian)
	header := &MATLABHeader{
		level:           "5.0",
		platform:        runtime.GOOS,
		created:         time.Now(),
		version:         0x0100,
		endianIndicator: endianIndicator,
	}
	return &MATLABMatFile{
		header:   header,
		elements: make([]MATLABByteInterface, 0),
		endian:   endian,
	}
}

// The method to add data element.
func (m *MATLABMatFile) AddElement(e MATLABByteInterface) {
	m.elements = append(m.elements, e)
}

// The method to write the buffer
func (m *MATLABMatFile) ToBytes(buf *MATLABBuffer) *MATLABBuffer {
	m.header.ToBytes(buf)
	for _, e := range m.elements {
		e.ToBytes(buf)
	}
	return buf
}

// MATLABHeader is a struct for the header of a matfile. The level is a version of matfile. In many cases, the level
// is "5.0". The platform is the name of OS. By using "runtime" package, it can be given by runtime.GOOS. The created
// the date when it creates, which is usually `time.Now()` as "time" package. The version is to be set as 0x0100.
// The endiaanIndicator is a string to indicate the endianness. If the endianness is the little endian, it becomes "IM".
// If it is the big endian, the indicator is "MI".
type MATLABHeader struct {
	level           string
	platform        string
	created         time.Time
	version         uint16
	endianIndicator string
}

// get the number of bytes.
func (h *MATLABHeader) byteLen() int {
	return 128
}

// get a byte buffer for the header
func (h *MATLABHeader) ToBytes(buf *MATLABBuffer) *MATLABBuffer {
	// header string. The header string should be 116 bytes.
	slen := 116
	s := fmt.Sprintf("MATLAB %s MAT-file, Platform: %s, Created on: %s", h.level, h.platform, h.created.Format(time.ANSIC))
	if len(s) > slen {
		buf.WriteString(s[:slen])
	} else {
		buf.WriteString(s)
		buf.Padding(slen - len(s)) // padding 0x00 up to slen bytes
	}
	buf.Write(uint32(0x0)) // subsys data offset
	buf.Write(uint32(0x0)) // subsys data offset
	buf.Write(h.version)
	buf.WriteString(h.endianIndicator)
	return buf
}

// MATLABDataElementHeader is a struct for the header of data element. dataType represents the format of data such as uint8,
// int16, etc. numOfDataByte is a length of data size. The size of header should be 8 bytes in usual. But if the data length
// is less than 4 bytes, the size of header should be 4 by compressing datatype and numofdatabyte to uint16. At that time, the
// flag `smallsize` becomes true.
type MATLABDataElementHeader struct {
	dataType      MATLABDataType
	numOfDataByte uint32
	smallsize     bool
}

func (h *MATLABDataElementHeader) byteLen() int {
	if h.smallsize {
		return 4
	} else {
		return 8
	}
}

func (h *MATLABDataElementHeader) ToBytes(buf *MATLABBuffer) *MATLABBuffer {
	if h.smallsize {
		buf.Write(uint16(h.dataType))
		buf.Write(uint16(h.numOfDataByte))
	} else {
		buf.Write(h.dataType)
		buf.Write(h.numOfDataByte)
	}
	return buf
}

// The struct to represent zero padding. This is added to the last of each data element so that the end of data element
// becomes the multiple of 8.
type MATLABPadding struct {
	size int
}

func (b *MATLABPadding) byteLen() int {
	return b.size
}

func (mb *MATLABPadding) ToBytes(b *MATLABBuffer) *MATLABBuffer {
	b.Padding(mb.size)
	return b
}

func createMATLABPadding(totalBytes int, segmentByte int) *MATLABPadding {
	size := (segmentByte - totalBytes%segmentByte) % segmentByte
	return &MATLABPadding{
		size: size,
	}
}

// The function is to create data header and padding. The format of each data element is
// - header
// - data
// - padding
func createMATLABDataElementHeader(dataType MATLABDataType, dataByte int) (*MATLABDataElementHeader, *MATLABPadding) {
	header := &MATLABDataElementHeader{
		dataType:      dataType,
		numOfDataByte: uint32(dataByte),
		smallsize:     dataByte <= 4,
	}
	padding := createMATLABPadding(header.byteLen()+dataByte, 8)
	return header, padding
}

// The structure for a string in MATLAB.
type MATLABString struct {
	header  *MATLABDataElementHeader
	data    string
	padding *MATLABPadding
}

// The function to create MATLAB string
func CreateMATLABDataString(data string) *MATLABString {
	header, padding := createMATLABDataElementHeader(miINT8, len(data))
	return &MATLABString{
		header:  header,
		data:    data,
		padding: padding,
	}
}

func (d *MATLABString) byteLen() int {
	return d.header.byteLen() + int(d.header.numOfDataByte) + d.padding.byteLen()
}

func (d *MATLABString) ToBytes(buf *MATLABBuffer) *MATLABBuffer {
	d.header.ToBytes(buf)
	buf.WriteString(d.data)
	d.padding.ToBytes(buf)
	return buf
}

// The structure for array in MATLAB, which consists of header and data. []intXX, []uintXX,
// []floatXX are available for data.
type MATLABArray struct {
	header  *MATLABDataElementHeader
	data    interface{}
	padding *MATLABPadding
}

// The function to create MATLAB array
func CreateMATLABArray(data interface{}) *MATLABArray {
	datatype, bytelen := getdatatype(data)
	header, padding := createMATLABDataElementHeader(datatype, bytelen)
	return &MATLABArray{
		header:  header,
		data:    data,
		padding: padding,
	}
}

func (d *MATLABArray) byteLen() int {
	return d.header.byteLen() + int(d.header.numOfDataByte) + d.padding.byteLen()
}

func (d *MATLABArray) ToBytes(buf *MATLABBuffer) *MATLABBuffer {
	d.header.ToBytes(buf)
	buf.Write(d.data)
	d.padding.ToBytes(buf)
	return buf
}

// The function provides datatype and bytes.
func getdatatype(data interface{}) (MATLABDataType, int) {
	switch data.(type) {
	case int32:
		return miINT32, 4
	case uint32:
		return miUINT32, 4
	case int64:
		return miINT64, 8
	case uint64:
		return miUINT64, 8
	case float32:
		return miSINGLE, 4
	case float64:
		return miDOUBLE, 8
	case []int32:
		return miINT32, 4 * len(data.([]int32))
	case []uint32:
		return miUINT32, 4 * len(data.([]uint32))
	case []int64:
		return miINT64, 8 * len(data.([]int64))
	case []uint64:
		return miUINT64, 8 * len(data.([]uint64))
	case []float32:
		return miSINGLE, 4 * len(data.([]float32))
	case []float64:
		return miDOUBLE, 8 * len(data.([]float64))
	default:
		log.Panic("The type should include the siez of bits like []int32, []int64.")
		return MATLABDataTypeUnknown, 0
	}
}

// bits for ArrayFlags.
// complexBit indicates whether the array containts complex numbers or not.
// globalBit indicates whether the array is defined as the global variable or not.
// logicalBit indicates whether the array containts logical values or not?
const (
	MATLABcomplexBit uint32 = 0x00000700
	MATLABglobalBit         = 0x00000400
	MATLABlogicalBit        = 0x00002000
)

func CreateMATLABArrayFlags(x uint32, t MATLABArrayType, complex bool, global bool, logical bool) *MATLABArray {
	flags := []uint32{uint32(t), x}
	if complex {
		flags[0] |= MATLABcomplexBit
	}
	if global {
		flags[0] |= MATLABglobalBit
	}
	if logical {
		flags[0] |= MATLABlogicalBit
	}
	return CreateMATLABArray(flags)
}

// The structure for MATLAB matrix which consists of header, arrayFlags,
// dimensionArray, arrayName and realValue. `complexValue` is needed to be
// implemented in future.
type MATLABMatrix struct {
	header          *MATLABDataElementHeader
	arrayFlags      *MATLABArray
	dimensionsArray *MATLABArray
	arrayName       *MATLABString
	realValue       *MATLABArray
	padding         *MATLABPadding
}

func (d *MATLABMatrix) byteLen() int {
	return d.header.byteLen() + int(d.header.numOfDataByte) + d.padding.byteLen()
}

func (d *MATLABMatrix) ToBytes(buf *MATLABBuffer) *MATLABBuffer {
	d.header.ToBytes(buf)
	d.arrayFlags.ToBytes(buf)
	d.dimensionsArray.ToBytes(buf)
	d.arrayName.ToBytes(buf)
	d.realValue.ToBytes(buf)
	d.padding.ToBytes(buf)
	return buf
}

func CreateMATLABMatrix(dims interface{}, name string, pr interface{}) *MATLABMatrix {
	matrixtype := getmatrixtype(pr)
	arrayFlags := CreateMATLABArrayFlags(0, matrixtype, false, false, false)
	dimensionsArray := CreateMATLABArray(toInt32(dims))
	arrayName := CreateMATLABDataString(name)
	realValue := CreateMATLABArray(pr)
	header, padding := createMATLABDataElementHeader(miMATRIX,
		arrayFlags.byteLen()+dimensionsArray.byteLen()+arrayName.byteLen()+realValue.byteLen())
	return &MATLABMatrix{
		header:          header,
		arrayFlags:      arrayFlags,
		dimensionsArray: dimensionsArray,
		arrayName:       arrayName,
		realValue:       realValue,
		padding:         padding,
	}
}

// The function to change []int and []uint to []int32
func toInt32(x interface{}) []int32 {
	switch x.(type) {
	case int32:
		return []int32{x.(int32)}
	case []int32:
		return x.([]int32)
	case int:
		return []int32{int32(x.(int))}
	case []int:
		y := x.([]int)
		result := make([]int32, len(y), len(y))
		for i, v := range y {
			result[i] = int32(v)
		}
		return result
	case int64:
		return []int32{int32(x.(int64))}
	case []int64:
		y := x.([]int64)
		result := make([]int32, len(y), len(y))
		for i, v := range y {
			result[i] = int32(v)
		}
		return result
	default:
		log.Panic("toInt32 should be int, []int, int32, []int32, int64 or []int64")
		return []int32{}
	}
}

// The function provides matrixtype
func getmatrixtype(data interface{}) MATLABArrayType {
	switch data.(type) {
	case []int32:
		return mxINT32_CLASS
	case []uint32:
		return mxUINT32_CLASS
	case []int64:
		return mxINT64_CLASS
	case []uint64:
		return mxUINT64_CLASS
	case []float32:
		return mxSINGLE_CLASS
	case []float64:
		return mxDOUBLE_CLASS
	default:
		log.Panic("matrix type should be int32, uint32, int64, uint64, float32 or float64")
		return MATLABArrayTypeUnknown
	}
}

// The structure for MATLAB sparse matrix. The format is CSC (compressed sparse column) with 0 origin.
// In the context of CSC, rowIndex indicates rowind. columnIndex indicates colptr. Although the origin of
// array index in Matlab is 1, the matfile should be the zero origin.
type MATLABSparseMatrix struct {
	header          *MATLABDataElementHeader
	arrayFlags      *MATLABArray
	dimensionsArray *MATLABArray
	arrayName       *MATLABString
	rowIndex        *MATLABArray
	columnIndex     *MATLABArray
	realValue       *MATLABArray
	padding         *MATLABPadding
}

func (d *MATLABSparseMatrix) byteLen() int {
	return d.header.byteLen() + int(d.header.numOfDataByte) + d.padding.byteLen()
}

func (d *MATLABSparseMatrix) ToBytes(buf *MATLABBuffer) *MATLABBuffer {
	d.header.ToBytes(buf)
	d.arrayFlags.ToBytes(buf)
	d.dimensionsArray.ToBytes(buf)
	d.arrayName.ToBytes(buf)
	d.rowIndex.ToBytes(buf)
	d.columnIndex.ToBytes(buf)
	d.realValue.ToBytes(buf)
	d.padding.ToBytes(buf)
	return buf
}

// The function to write
func CreateMATLABSparseMatrix(dims interface{}, name string, nnz int,
	ir interface{}, jc interface{}, pr interface{}) *MATLABSparseMatrix {
	arrayFlags := CreateMATLABArrayFlags(uint32(nnz), mxSPARSE_CLASS, false, false, false)
	dimensionsArray := CreateMATLABArray(toInt32(dims))
	arrayName := CreateMATLABDataString(name)
	rowIndex := CreateMATLABArray(toInt32(ir))
	columnIndex := CreateMATLABArray(toInt32(jc))
	realValue := CreateMATLABArray(pr)
	header, padding := createMATLABDataElementHeader(miMATRIX,
		arrayFlags.byteLen()+
			dimensionsArray.byteLen()+
			arrayName.byteLen()+
			rowIndex.byteLen()+
			columnIndex.byteLen()+
			realValue.byteLen())
	return &MATLABSparseMatrix{
		header:          header,
		arrayFlags:      arrayFlags,
		dimensionsArray: dimensionsArray,
		arrayName:       arrayName,
		rowIndex:        rowIndex,
		columnIndex:     columnIndex,
		realValue:       realValue,
		padding:         padding,
	}
}
