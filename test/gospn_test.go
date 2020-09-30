package test

import (
	"../pkg/parser"
	"../pkg/petrinet"
	"testing"
)

func BenchmarkGoSPNP1(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	if net, imark, err := parser.PNreadFromFile("./data/spnp_example1.spn"); err == nil {
		b.ResetTimer()
		mg := petrinet.CreateMarkingGraphWithDFS(net, imark)
		mg.TransMatrix()
		mg.GroupLabels()
		mg.TransLabels()
		mg.InitVector()
	}
	// }
}

func BenchmarkGoSPNP2(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	if net, imark, err := parser.PNreadFromFile("./data/spnp_example2.spn"); err == nil {
		b.ResetTimer()
		mg := petrinet.CreateMarkingGraphWithDFS(net, imark)
		mg.TransMatrix()
		mg.GroupLabels()
		mg.TransLabels()
		mg.InitVector()
	}
	// }
}

func BenchmarkGoSPNP3(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	if net, imark, err := parser.PNreadFromFile("./data/spnp_example3.spn"); err == nil {
		b.ResetTimer()
		mg := petrinet.CreateMarkingGraphWithDFS(net, imark)
		mg.TransMatrix()
		mg.GroupLabels()
		mg.TransLabels()
		mg.InitVector()
	}
	// }
}

func BenchmarkGoSPNP4(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	if net, imark, err := parser.PNreadFromFile("./data/spnp_example4.spn"); err == nil {
		b.ResetTimer()
		mg := petrinet.CreateMarkingGraphWithDFS(net, imark)
		mg.TransMatrix()
		mg.GroupLabels()
		mg.TransLabels()
		mg.InitVector()
	}
	// }
}

func BenchmarkGoSPNP5(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	if net, imark, err := parser.PNreadFromFile("./data/spnp_example5.spn"); err == nil {
		b.ResetTimer()
		mg := petrinet.CreateMarkingGraphWithDFS(net, imark)
		mg.TransMatrix()
		mg.GroupLabels()
		mg.TransLabels()
		mg.InitVector()
	}
	// }
}

func BenchmarkGoSPNP6(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	if net, imark, err := parser.PNreadFromFile("./data/spnp_example6.spn"); err == nil {
		b.ResetTimer()
		mg := petrinet.CreateMarkingGraphWithDFS(net, imark)
		mg.TransMatrix()
		mg.GroupLabels()
		mg.TransLabels()
		mg.InitVector()
	}
	// }
}

func BenchmarkGoSPNP7(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	if net, imark, err := parser.PNreadFromFile("./data/raid6.spn"); err == nil {
		b.ResetTimer()
		mg := petrinet.CreateMarkingGraphWithDFS(net, imark)
		mg.TransMatrix()
		mg.GroupLabels()
		mg.TransLabels()
		mg.InitVector()
	}
	// }
}

func BenchmarkGoSPNP8(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	if net, imark, err := parser.PNreadFromFile("./data/raid10.spn"); err == nil {
		b.ResetTimer()
		mg := petrinet.CreateMarkingGraphWithDFS(net, imark)
		mg.TransMatrix()
		mg.GroupLabels()
		mg.TransLabels()
		mg.InitVector()
	}
	// }
}
