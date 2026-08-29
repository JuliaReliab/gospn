package petrinet

import (
	"encoding/binary"
	"strconv"
	"strings"
)

type MarkInt int

type Mark struct {
	n         int
	markSlice []MarkInt
}

func (x *Mark) toSlice() []MarkInt {
	return x.markSlice
}

func newMark(mark []MarkInt) *Mark {
	return &Mark{
		n:         len(mark),
		markSlice: mark,
	}
}

func (m *Mark) String() string {
	result := make([]string, m.n)
	for i := 0; i < m.n; i++ {
		result[i] = strconv.Itoa(int(m.markSlice[i]))
	}
	return "[" + strings.Join(result, ",") + "]"
}

// MarkGenerator

type MarkGeneratorInterface interface {
	genMark([]MarkInt) *Mark
	size() int
}

// MarkGenerator interns markings: the same token vector always yields the same *Mark, so
// the rest of the search can compare and key markings by pointer.
//
// The key is the raw little-endian bytes of the vector rather than its decimal rendering.
// Formatting each count cost a strconv call per place per edge, and the lookup is written
// as the literal g.data[string(g.key)] form because that is the only shape the compiler
// recognises well enough to skip allocating the string; binding it to a variable first
// made every lookup -- hit or miss -- allocate.
type MarkGenerator struct {
	key  []byte
	data map[string]*Mark
}

func NewMarkGenerator(n int) *MarkGenerator {
	return &MarkGenerator{
		key:  make([]byte, markKeyWidth*n),
		data: make(map[string]*Mark),
	}
}

// The number of bytes one place's token count occupies in a lookup key.
const markKeyWidth = 8

// size is the number of distinct markings interned so far. The state limit watches
// this rather than the number of markings already expanded: the intern table is what
// holds the memory, and it runs ahead of the expansion by the branching factor.
func (g *MarkGenerator) size() int { return len(g.data) }

func (g *MarkGenerator) genMark(m []MarkInt) *Mark {
	if n := markKeyWidth * len(m); len(g.key) != n {
		g.key = make([]byte, n)
	}
	for i, x := range m {
		binary.LittleEndian.PutUint64(g.key[markKeyWidth*i:], uint64(x))
	}
	if mark, ok := g.data[string(g.key)]; ok {
		return mark
	}
	// Only on a miss is the marking retained, and it is copied first: callers reuse the
	// slice they passed in for the next firing.
	own := make([]MarkInt, len(m))
	copy(own, m)
	newmark := newMark(own)
	g.data[string(g.key)] = newmark
	return newmark
}
