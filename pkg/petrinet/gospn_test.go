package petrinet

// TODO: this should be implemented to ensure the marking result.
// When parser is imported, a cycle is happen.

import (
	// 	"bufio"
	"bytes"
	"fmt"
	// 	"os"
	"testing"
)

func buildRaid6() (*Net, []MarkInt) {
	/*
		  Example: RAID6
		  F. Machida, R. Xia and K.S. Trivedi,
		  Performability modeling for RAID storage systems by Markov regenerative process,
		  IEEE Transactions on Dependable and Secure Computing

				// HDD model

				place Pn (init = 6)
				place Pdf
				exp Tdfail (guard = gfail, rate = Tdfail_rate)
				gen Trebuild (guard = gfail, dist = Trebuild_dist)
				imm Tinit (guard = ginit)
				arc Pn to Tdfail
				arc Tdfail to Pdf
				arc Pdf to Trebuild
				arc Pdf to Tinit
				arc Trebuild to Pn
				arc Tinit to Pn

				// Reconstruction model

				place Po (init = 1)
				place Pr
				place Pc
				imm Tstart (guard = gstart)
				gen Trecon (dist = Trecon_dist)
				imm Tend (guard = gend)
				arc Po to Tstart
				arc Tstart to Pr
				arc Pr to Trecon
				arc Trecon to Pc
				arc Pc to Tend
				arc Tend to Po

				// rate and gurads

				Tdfail_rate = #Pn * lambda
				gfail = #Po == 1
				gstart = #Pdf > 2
				ginit = #Pc == 1
				gend = #Pdf == 0

				// params

				Trebuild_dist = det(MTTR1)
				Trecon_dist = det(MTTR2)

				MTTF = 1.0e+6 // [hours]
				lambda = 1/MTTF
				MTTR1 = 2 // [hours]
				MTTR2 = 24 // [hours]
				reward avail ifelse(#Po == 1, 1, 0)
				reward unavail ifelse(#Po == 1, 0, 1)
	*/
	net := NewNet()

	// parameters
	MTTF := 1.0e+6
	lambda := 1 / MTTF
	MTTR1 := 2.0
	MTTR2 := 24.0
	Trebuild_dist := NewDistribution("constant", MTTR1)
	Trecon_dist := NewDistribution("constant", MTTR2)

	Pn := net.NewPlace("Pn", 10)
	Pdf := net.NewPlace("Pdf", 10)
	Tdfail := net.NewExpTrans("Tdfail", 0, true, 1.0)
	Trebuild := net.NewGenTrans("Trebuild", 0, true, Trebuild_dist, GenTransPolicyPRD)
	Tinit := net.NewImmTrans("Tinit", 0, true, 1.0)
	net.NewInArc(Pn, Tdfail, 1)
	net.NewOutArc(Tdfail, Pdf, 1)
	net.NewInArc(Pdf, Trebuild, 1)
	net.NewInArc(Pdf, Tinit, 1)
	net.NewOutArc(Trebuild, Pn, 1)
	net.NewOutArc(Tinit, Pn, 1)

	Po := net.NewPlace("Po", 10)
	Pr := net.NewPlace("Pr", 10)
	Pc := net.NewPlace("Pc", 10)
	Tstart := net.NewImmTrans("Tstart", 0, true, 1.0)
	Trecon := net.NewGenTrans("Trecon", 0, true, Trecon_dist, GenTransPolicyPRD)
	Tend := net.NewImmTrans("Tend", 0, true, 1.0)
	net.NewInArc(Po, Tstart, 1)
	net.NewOutArc(Tstart, Pr, 1)
	net.NewInArc(Pr, Trecon, 1)
	net.NewOutArc(Trecon, Pc, 1)
	net.NewInArc(Pc, Tend, 1)
	net.NewOutArc(Tend, Po, 1)

	// rate and guards
	Tdfail_rate := func(m []MarkInt) float64 {
		return float64(m[Pn.index]) * lambda
	}
	gfail := func(m []MarkInt) bool {
		return m[Po.index] == 1
	}
	gstart := func(m []MarkInt) bool {
		return m[Pdf.index] > 2
	}
	ginit := func(m []MarkInt) bool {
		return m[Pc.index] == 1
	}
	gend := func(m []MarkInt) bool {
		return m[Pdf.index] == 0
	}
	net.SetWeightRate(Tdfail, Tdfail_rate)
	net.SetGuard(Tdfail, "#Po==1", gfail)
	net.SetGuard(Trebuild, "#Po==1", gfail)
	net.SetGuard(Tinit, "#Pc==1", ginit)
	net.SetGuard(Tstart, "#Pdf>2", gstart)
	net.SetGuard(Tend, "#Pdf==0", gend)

	net.SetReward("avail", func(m []MarkInt) float64 {
		if m[Po.index] == 1 {
			return 1.0
		} else {
			return 0.0
		}
	})
	net.SetReward("unavail", func(m []MarkInt) float64 {
		if m[Po.index] == 1 {
			return 0.0
		} else {
			return 1.0
		}
	})

	net.Finalize()

	m0 := net.MakeMark(map[string]MarkInt{"Pn": 6, "Po": 1})

	return net, m0
}

// func (mg *MarkingGraph) GetMarks() []*Mark {
// 	return mg.marks
// }

// func (mg *MarkingGraph) GetSizeGroup(g GroupType) int {
// 	sum := 0
// 	for markgroup, mset := range mg.groupToMark {
// 		if markgroup.gtype == g {
// 			sum += len(mset)
// 		}
// 	}
// 	return sum
// }

// func (mg *MarkingGraph) GetNNZSizeGroup(g GroupType) int {
// 	sum := 0
// 	for grouptr, lset := range mg.groupTransToLink {
// 		if grouptr.src.gtype == g {
// 			sum += len(lset)
// 		}
// 	}
// 	return sum
// }

// func TestGoSPNP1(t *testing.T) {
// 	if net, imark, err := parser.PNreadFromFile("./data/spnp_example1.spn"); err == nil {
// 		fmt.Println("Crate marking...")
// 		mg := CreateMarkingGraphWithDFS(net, imark)
// 		if !(mg.GetSizeGroup(IMMGroup) == 0 &&
// 			mg.GetSizeGroup(GENGroup) == 30 &&
// 			mg.GetSizeGroup(ABSGroup) == 0) {
// 			t.Errorf("Error")
// 			panic("")
// 		}
// 		if !(mg.GetNNZSizeGroup(IMMGroup) == 0 &&
// 			mg.GetNNZSizeGroup(GENGroup) == 88 &&
// 			mg.GetNNZSizeGroup(ABSGroup) == 0) {
// 			t.Errorf("Error")
// 		}
// 	}
// }

// func TestGoSPNP2(t *testing.T) {
// 	if net, imark, err := parser.PNreadFromFile("./data/spnp_example2.spn"); err == nil {
// 		fmt.Println("Crate marking...")
// 		mg := CreateMarkingGraphWithDFS(net, imark)
// 		if !(mg.GetSizeGroup(IMMGroup) == 4 &&
// 			mg.GetSizeGroup(GENGroup) == 6 &&
// 			mg.GetSizeGroup(ABSGroup) == 1) {
// 			t.Errorf("Error")
// 		}
// 		if !(mg.GetNNZSizeGroup(IMMGroup) == 7 &&
// 			mg.GetNNZSizeGroup(GENGroup) == 8 &&
// 			mg.GetNNZSizeGroup(ABSGroup) == 0) {
// 			t.Errorf("Error")
// 		}
// 	}
// }

// func TestGoSPNP3(t *testing.T) {
// 	if net, imark, err := parser.PNreadFromFile("./data/spnp_example3.spn"); err == nil {
// 		fmt.Println("Crate marking...")
// 		mg := CreateMarkingGraphWithDFS(net, imark)
// 		if !(mg.GetSizeGroup(IMMGroup) == 0 &&
// 			mg.GetSizeGroup(GENGroup) == 51 &&
// 			mg.GetSizeGroup(ABSGroup) == 0) {
// 			t.Errorf("Error")
// 		}
// 		if !(mg.GetNNZSizeGroup(IMMGroup) == 0 &&
// 			mg.GetNNZSizeGroup(GENGroup) == 100 &&
// 			mg.GetNNZSizeGroup(ABSGroup) == 0) {
// 			t.Errorf("Error")
// 		}
// 		fmt.Println(mg.TransMatrix())
// 	}
// }

// func TestGoSPNP4(t *testing.T) {
// 	if net, imark, err := parser.PNreadFromFile("./data/spnp_example4.spn"); err == nil {
// 		fmt.Println("Crate marking...")
// 		mg := CreateMarkingGraphWithDFS(net, imark)
// 		if !(mg.GetSizeGroup(IMMGroup) == 63 &&
// 			mg.GetSizeGroup(GENGroup) == 49 &&
// 			mg.GetSizeGroup(ABSGroup) == 1) {
// 			t.Errorf("Error")
// 		}
// 		if !(mg.GetNNZSizeGroup(IMMGroup) == 63 &&
// 			mg.GetNNZSizeGroup(GENGroup) == 147 &&
// 			mg.GetNNZSizeGroup(ABSGroup) == 0) {
// 			t.Errorf("Error")
// 		}
// 		file, err := os.Create("markdot_example4.dot")
// 		if err != nil {
// 			panic("file open error")
// 		}
// 		defer file.Close()
// 		writer := bufio.NewWriter(file)
// 		mg.ToMarkDot(writer)
// 		writer.Flush()
// 	}
// }

// func TestGoSPNP5(t *testing.T) {
// 	if net, imark, err := parser.PNreadFromFile("./data/spnp_example5.spn"); err == nil {
// 		mg := CreateMarkingGraphWithDFS(net, imark)
// 		if !(mg.GetSizeGroup(IMMGroup) == 72 &&
// 			mg.GetSizeGroup(GENGroup) == 44 &&
// 			mg.GetSizeGroup(ABSGroup) == 235) {
// 			t.Errorf("Error")
// 		}
// 		if !(mg.GetNNZSizeGroup(IMMGroup) == 144 &&
// 			mg.GetNNZSizeGroup(GENGroup) == 346 &&
// 			mg.GetNNZSizeGroup(ABSGroup) == 0) {
// 			t.Errorf("Error")
// 		}
// 	}
// }

// func TestGoSPNP6(t *testing.T) {
// 	if net, imark, err := parser.PNreadFromFile("./data/spnp_example6.spn"); err == nil {
// 		mg := CreateMarkingGraphWithDFS(net, imark)
// 		if !(mg.GetSizeGroup(IMMGroup) == 19868 &&
// 			mg.GetSizeGroup(GENGroup) == 26244 &&
// 			mg.GetSizeGroup(ABSGroup) == 0) {
// 			t.Errorf("Error")
// 		}
// 		if !(mg.GetNNZSizeGroup(IMMGroup) == 9844+10084 &&
// 			mg.GetNNZSizeGroup(GENGroup) == 11016+142560 &&
// 			mg.GetNNZSizeGroup(ABSGroup) == 0) {
// 			t.Errorf("Error")
// 		}
// 	}
// }

func TestGoSPNP7(t *testing.T) {
	net, imark := buildRaid6()
	mg := CreateMarkingGraphWithDFS(net, imark)
	mg.TransMatrix()
	// if !(mg.GetSizeGroup(IMMGroup) == 19868 &&
	// 	mg.GetSizeGroup(GENGroup) == 26244 &&
	// 	mg.GetSizeGroup(ABSGroup) == 0) {
	// 	t.Errorf("Error")
	// }
	// if !(mg.GetNNZSizeGroup(IMMGroup) == 9844+10084 &&
	// 	mg.GetNNZSizeGroup(GENGroup) == 11016+142560 &&
	// 	mg.GetNNZSizeGroup(ABSGroup) == 0) {
	// 	t.Errorf("Error")
	// }
	writer := bytes.NewBuffer(make([]byte, 0, 256))
	mg.ToGroupMarkDot(writer)
	fmt.Println(writer.String())
}

// func TestGoSPNP8(t *testing.T) {
// 	if net, imark, err := parser.PNreadFromFile("./data/raid10.spn"); err == nil {
// 		mg := CreateMarkingGraphWithDFS(net, imark)
// 		// if !(mg.GetSizeGroup(IMMGroup) == 19868 &&
// 		// 	mg.GetSizeGroup(GENGroup) == 26244 &&
// 		// 	mg.GetSizeGroup(ABSGroup) == 0) {
// 		// 	t.Errorf("Error")
// 		// }
// 		// if !(mg.GetNNZSizeGroup(IMMGroup) == 9844+10084 &&
// 		// 	mg.GetNNZSizeGroup(GENGroup) == 11016+142560 &&
// 		// 	mg.GetNNZSizeGroup(ABSGroup) == 0) {
// 		// 	t.Errorf("Error")
// 		// }
// 		e, i, g := mg.TransMatrix()
// 		label1 := mg.GroupLabels()
// 		label2 := mg.TransLabels()
// 		for tr, m := range e {
// 			fmt.Printf("%s%s%s ", label1[tr.GetSrc()], label1[tr.GetDest()], label2[tr])
// 			fmt.Println(m)
// 		}
// 		for tr, m := range i {
// 			fmt.Printf("%s%s%s ", label1[tr.GetSrc()], label1[tr.GetDest()], label2[tr])
// 			fmt.Println(m)
// 		}
// 		for tr, m := range g {
// 			fmt.Printf("%s%s%s ", label1[tr.GetSrc()], label1[tr.GetDest()], label2[tr])
// 			fmt.Println(m)
// 		}

// 		writer := bytes.NewBuffer(make([]byte, 0, 256))
// 		mg.ToGroupMarkDot(writer)
// 		fmt.Println(writer.String())

// 		iv := mg.InitVector()
// 		for g, v := range iv {
// 			fmt.Printf("init %s ", label1[g])
// 			fmt.Println(v)
// 		}

// 		rv := mg.RewardVector("rwd1")
// 		for g, v := range rv {
// 			fmt.Printf("reward %s ", label1[g])
// 			fmt.Println(v)
// 		}
// 	}
// }
