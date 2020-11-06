package petrinet

import (
	"strings"
)

type GenVec struct {
	n        int
	vecSlice []TransStatus
	label    string
}

func newGenVec(v []TransStatus) *GenVec {
	result := make([]string, len(v))
	for i, x := range v {
		result[i] = x.String()
	}
	label := "{" + strings.Join(result, ",") + "}"
	return &GenVec{
		n:        len(v),
		vecSlice: v,
		label:    label,
	}
}

func (g *GenVec) toSlice() []TransStatus {
	return g.vecSlice
}

func (g *GenVec) IsAnyEnabled() bool {
	for i := 0; i < g.n; i++ {
		if g.vecSlice[i] == ENABLE {
			return true
		}
	}
	return false
}

func (g *GenVec) String() string {
	return g.label
}

// GenVecGenerator

type GenVecGeneratorInterface interface {
	genGenVec([]TransStatus) *GenVec
}

type GenVecGenerator struct {
	key  []byte
	data map[string]*GenVec
}

func NewGenVecGenerator(n int) *GenVecGenerator {
	return &GenVecGenerator{
		key:  make([]byte, 0, 1*n), // estimate 1 character for one gentrans
		data: make(map[string]*GenVec),
	}
}

func (g *GenVecGenerator) genGenVec(vec []TransStatus) *GenVec {
	g.key = g.key[:0]
	for _, x := range vec {
		g.key = append(g.key, x.String()...)
		g.key = append(g.key, ',')
	}
	key := string(g.key)
	if gv, ok := g.data[key]; ok {
		return gv
	} else {
		newgv := newGenVec(vec)
		g.data[key] = newgv
		return newgv
	}
}
