package petrinet

import (
	"reflect"
)

type markInt uint

func toarray(x []markInt) reflect.Value {
	n := len(x)
	rt := reflect.ArrayOf(n, reflect.TypeOf(x[0]))
	rx := reflect.New(rt).Elem()
	for i, v := range x {
		rx.Index(i).Set(reflect.ValueOf(v))
	}
	return rx
}

func toslice(x *Mark) []markInt {
	y := make([]markInt, x.n, x.n)
	for i,_ := range y {
		y[i] = x.mark.Index(i).Interface().(markInt)
	}
	return y
}

type Mark struct {
	n int
	mark reflect.Value
}

func NewMark(mark []markInt) *Mark {
	return &Mark{
		n: len(mark),
		mark: toarray(mark),
	}
}

func (m *Mark) Get(i int) markInt {
	return m.mark.Index(i).Interface().(markInt)
}

type MarkMap struct {
	amap reflect.Value
}

func NewMarkMap(n int) *MarkMap {
	rkey := reflect.ArrayOf(n, reflect.TypeOf(markInt(0)))
	rval := reflect.TypeOf(NewMark(make([]markInt, n)))
	return &MarkMap{
		amap: reflect.MakeMap(reflect.MapOf(rkey, rval)),
	}
}

func (m *MarkMap) Set(key *Mark, val *Mark) {
	m.amap.SetMapIndex(key.mark, reflect.ValueOf(val))
}

func (m *MarkMap) Get(key *Mark) *Mark {
	result := m.amap.MapIndex(key.mark)
	if result.IsValid() {
		return result.Interface().(*Mark)
	} else {
		return nil
	}
}
