package petrinet

// TODO: this should be implemented to ensure the marking result.
// When parser is imported, a cycle is happen.

// import (
// 	"bufio"
// 	"bytes"
// 	"fmt"
// 	"os"
// 	"testing"
// )

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

// func TestGoSPNP7(t *testing.T) {
// 	if net, imark, err := parser.PNreadFromFile("./data/raid6.spn"); err == nil {
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
// 		writer := bytes.NewBuffer(make([]byte, 0, 256))
// 		mg.ToGroupMarkDot(writer)
// 		fmt.Println(writer.String())
// 	}
// }

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
