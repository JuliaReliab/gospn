package petrinet

import (
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
}

type MarkGenerator struct {
	key  []byte
	data map[string]*Mark
}

func NewMarkGenerator(n int) *MarkGenerator {
	return &MarkGenerator{
		key:  make([]byte, 0, 5*n), // estimate 5 characters for one place
		data: make(map[string]*Mark),
	}
}

func (g *MarkGenerator) genMark(m []MarkInt) *Mark {
	g.key = g.key[:0]
	for _, x := range m {
		g.key = append(g.key, strconv.Itoa(int(x))...)
		g.key = append(g.key, ',')
	}
	key := string(g.key)
	if mark, ok := g.data[key]; ok {
		return mark
	} else {
		newmark := newMark(m)
		g.data[key] = newmark
		return newmark
	}
}
