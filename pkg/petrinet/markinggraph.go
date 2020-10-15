package petrinet

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
)

type linkTransInterface interface {
	getTrans() *Trans
	getValue(*Net, *Mark) float64
}

func (tr *ImmTrans) getValue(net *Net, mark *Mark) float64 {
	if f, ok := net.ratefunc[tr.getTrans()]; ok {
		return f(mark.toSlice())
	} else {
		return tr.weight
	}
}

func (tr *ExpTrans) getValue(net *Net, mark *Mark) float64 {
	if f, ok := net.ratefunc[tr.getTrans()]; ok {
		return f(mark.toSlice())
	} else {
		return tr.rate
	}
}

func (tr *GenTrans) getValue(net *Net, mark *Mark) float64 {
	return 1.0
}

type Link struct {
	src  *Mark              // source
	dest *Mark              // destination
	tr   linkTransInterface // transition
	tt   TransType          // type
}

// The enum to indicate group types.
type GroupType int

// GEN: There is no enabled IMM trans
// IMM: One or more IMM trans are enabled
// ABS: Absorbing marks (There are no enabled transitions)
const (
	GENGroup GroupType = iota + 1
	IMMGroup
	ABSGroup
)

// The structure of group for marks, which consists of GroupType and status vector of GEN transitions.
type Group struct {
	gtype GroupType
	gv    *GenVec
}

func (g *Group) String() string {
	return fmt.Sprintf("[%d %s]", g.gtype, g.gv)
}

// func (m *Group) String(net *Net) string {
// 	result := make([]string, m.gv.n)
// 	for i := 0; i < m.gv.n; i++ {
// 		result[i] = net.translist[i].label + "->" + m.gv.Get(i).String()
// 	}
// 	return "{" + strings.Join(result, ",") + "}"
// }

// A generator to make a unique instance of Group
type groupGenerator map[Group]*Group

func newGroupGenerator() *groupGenerator {
	m := make(groupGenerator)
	return &m
}

// The method to generate a unique instance of Group
func (s *groupGenerator) generate(g GroupType, gv *GenVec) *Group {
	m := Group{
		gtype: g,
		gv:    gv,
	}
	if mp, ok := (*s)[m]; ok {
		return mp
	} else {
		(*s)[m] = &m
		return &m
	}
}

type GroupTrans struct {
	src       *Group
	dest      *Group
	transtype TransType
	gentrans  *Trans
}

func (g GroupTrans) GetSrc() *Group {
	return g.src
}

func (g GroupTrans) GetDest() *Group {
	return g.dest
}

// The structure to store the result of analysis. The group represents the type of markings: IMM, GEN, ABS and
// the vector of status of GEN transitions (Enabled, Disabled and Premenpted).
type MarkingGraph struct {
	net              *Net                  // The object for PetriNet
	imark            *Mark                 // The initimal marking
	marks            []*Mark               // list of all marks (sorted)
	groups           []*Group              // list of all groups (sorted)
	links            []Link                // list of links between marks
	grouplinks       []GroupTrans          // list of grouptrans
	markToGroup      map[*Mark]*Group      // map from a mark to a group
	markToInt        map[*Mark]int         // map from a mark to an index
	groupToMark      map[*Group][]*Mark    // map from a group to a set of marks
	groupTransToLink map[GroupTrans][]Link // map from a group transition to a set of links
	groupGenerator   *groupGenerator       // a generator to make an instance of group
}

type makingGraphGenerator interface {
	create(net *Net, imark []MarkInt) (*Mark, []*Mark, map[*Mark]*GenVec, map[*Mark]GroupType, []Link)
}

func CreateMarkingGraph(net *Net, imark []MarkInt, method makingGraphGenerator) *MarkingGraph {
	m0, marks, markToGenvec, markToGroupType, links := method.create(net, imark)
	return newMarkingGraph(net, m0, marks, markToGenvec, markToGroupType, links)
}

func CreateMarkingGraphWithDFS(net *Net, imark []MarkInt) *MarkingGraph {
	return CreateMarkingGraph(net, imark, new(dfs))
}

func CreateMarkingGraphWithDFSTangible(net *Net, imark []MarkInt) *MarkingGraph {
	return CreateMarkingGraph(net, imark, new(dfstangible))
}

// The method to create a marking graph. This is only called from CreateMarkingGraph
func newMarkingGraph(net *Net,
	m0 *Mark,
	marks []*Mark,
	markToGenvec map[*Mark]*GenVec,
	markToGroupType map[*Mark]GroupType,
	links []Link) *MarkingGraph {

	// make groupmarks
	generator := newGroupGenerator()
	groupToMark := make(map[*Group][]*Mark)
	markToGroup := make(map[*Mark]*Group)
	for _, m := range marks {
		g := generator.generate(markToGroupType[m], markToGenvec[m])
		markToGroup[m] = g
		if mset, ok := groupToMark[g]; ok {
			groupToMark[g] = append(mset, m)
		} else {
			groupToMark[g] = []*Mark{m}
		}
	}

	// numbering for marks
	markToInt := make(map[*Mark]int)
	for _, mset := range groupToMark {
		for i, m := range mset {
			markToInt[m] = i
		}
	}

	// make grouplists and sort
	groups := make([]*Group, 0, len(groupToMark))
	for g, _ := range groupToMark {
		groups = append(groups, g)
	}
	sort.Slice(groups, func(i, j int) bool {
		si := groups[i].gv.toSlice()
		sj := groups[j].gv.toSlice()
		for k := 0; k < len(si); k++ {
			if si[k] == sj[k] {
				continue
			} else {
				return si[k] < sj[k]
			}
		}
		return groups[i].gtype < groups[j].gtype
	})

	// make grouplinks
	groupTransToLink := make(map[GroupTrans][]Link)
	grouplinks := make([]GroupTrans, 0)
	for _, l := range links {
		var tr GroupTrans
		if l.tt == TransGEN {
			tr = GroupTrans{
				src:       generator.generate(markToGroupType[l.src], markToGenvec[l.src]),
				dest:      generator.generate(markToGroupType[l.dest], markToGenvec[l.dest]),
				transtype: TransGEN,
				gentrans:  l.tr.getTrans(),
			}
		} else {
			tr = GroupTrans{
				src:       generator.generate(markToGroupType[l.src], markToGenvec[l.src]),
				dest:      generator.generate(markToGroupType[l.dest], markToGenvec[l.dest]),
				transtype: l.tt,
				gentrans:  nil,
			}
		}
		if lset, ok := groupTransToLink[tr]; ok {
			groupTransToLink[tr] = append(lset, l)
		} else {
			groupTransToLink[tr] = []Link{l}
			grouplinks = append(grouplinks, tr)
		}
	}

	return &MarkingGraph{
		net:              net,
		imark:            m0,
		marks:            marks,
		groups:           groups,
		links:            links,
		grouplinks:       grouplinks,
		markToGroup:      markToGroup,
		markToInt:        markToInt,
		groupToMark:      groupToMark,
		groupTransToLink: groupTransToLink,
		groupGenerator:   generator,
	}
}

func (mg *MarkingGraph) ToMarkDot(writer io.Writer) {
	fmt.Fprintf(writer, "digraph { layout=dot; overlap=false; splines=true;\n")
	for _, mark := range mg.marks {
		switch markgroup := mg.markToGroup[mark]; markgroup.gtype {
		case IMMGroup:
			fmt.Fprintf(writer, "\"%p\" [shape=circle, label=\"%d\n%s\", style=filled];\n", mark, mg.markToInt[mark], markgroup.gv)
		case GENGroup:
			fmt.Fprintf(writer, "\"%p\" [shape=circle, label=\"%d\n%s\"];\n", mark, mg.markToInt[mark], markgroup.gv)
		case ABSGroup:
			fmt.Fprintf(writer, "\"%p\" [shape=circle, label=\"%d\n%s\"];\n", mark, mg.markToInt[mark], markgroup.gv)
		default:
		}
	}
	for _, link := range mg.links {
		fmt.Fprintf(writer, "\"%p\"->\"%p\" [label=\"%s\"];\n", link.src, link.dest, link.tr.getTrans().label)
	}
	fmt.Fprintf(writer, "}\n")
}

func (mg *MarkingGraph) ToMarkDotWithLabel(writer io.Writer) {
	fmt.Fprintf(writer, "digraph { layout=dot; overlap=false; splines=true;\n")
	for _, mark := range mg.marks {
		switch mg.markToGroup[mark].gtype {
		case IMMGroup:
			if mg.imark == mark {
				fmt.Fprintf(writer, "\"%p\" [label=\"%s\", style=filled, peripheries=2];\n", mark, mark)
			} else {
				fmt.Fprintf(writer, "\"%p\" [label=\"%s\", style=filled];\n", mark, mark)
			}
		case GENGroup:
			if mg.imark == mark {
				fmt.Fprintf(writer, "\"%p\" [label=\"%s\", peripheries=2];\n", mark, mark)
			} else {
				fmt.Fprintf(writer, "\"%p\" [label=\"%s\"];\n", mark, mark)
			}
		case ABSGroup:
			if mg.imark == mark {
				fmt.Fprintf(writer, "\"%p\" [label=\"%s\", peripheries=2];\n", mark, mark)
			} else {
				fmt.Fprintf(writer, "\"%p\" [label=\"%s\"];\n", mark, mark)
			}
		default:
		}
	}
	for _, link := range mg.links {
		fmt.Fprintf(writer, "\"%p\"->\"%p\" [label=\"%s\"];\n", link.src, link.dest, link.tr.getTrans().label)
	}
	fmt.Fprintf(writer, "}\n")
}

func (mg *MarkingGraph) ToMarkDotWithLabelAndGroup(writer io.Writer) {
	fmt.Fprintf(writer, "digraph { layout=dot; overlap=false; splines=true;\n")
	for _, mark := range mg.marks {
		switch markgroup := mg.markToGroup[mark]; markgroup.gtype {
		case IMMGroup:
			fmt.Fprintf(writer, "\"%p\" [label=\"%s\n%s\", style=filled];\n", mark, mark, markgroup.gv)
		case GENGroup:
			fmt.Fprintf(writer, "\"%p\" [label=\"%s\n%s\"];\n", mark, mark, markgroup.gv)
		case ABSGroup:
			fmt.Fprintf(writer, "\"%p\" [label=\"%s\n%s\"];\n", mark, mark, markgroup.gv)
		default:
		}
	}
	for _, link := range mg.links {
		fmt.Fprintf(writer, "\"%p\"->\"%p\" [label=\"%s\"];\n", link.src, link.dest, link.tr.getTrans().label)
	}
	fmt.Fprintf(writer, "}\n")
}

func (mg *MarkingGraph) ToGroupMarkDot(writer io.Writer) {
	label1 := mg.GroupLabels()
	label2 := mg.TransLabels()
	fmt.Fprintf(writer, "digraph { layout=dot; overlap=false; splines=true;\n")
	for _, g := range mg.groups {
		switch g.gtype {
		case IMMGroup:
			fmt.Fprintf(writer, "\"%p\" [label=\"%s\", style=filled];\n", g, label1[g])
		case GENGroup:
			fmt.Fprintf(writer, "\"%p\" [label=\"%s\"];\n", g, label1[g])
		case ABSGroup:
			fmt.Fprintf(writer, "\"%p\" [label=\"%s\"];\n", g, label1[g])
		default:
		}
	}
	for _, link := range mg.grouplinks {
		fmt.Fprintf(writer, "\"%p\"->\"%p\" [label=\"%s\"];\n", link.src, link.dest, label2[link])
	}
	fmt.Fprintf(writer, "}\n")
}

type CSC struct {
	m      int
	n      int
	nnz    int
	colptr []int
	rowind []int
	value  []float64
}

func (mat *CSC) Get() ([]int32, int, []int, []int, []float64) {
	return []int32{int32(mat.m), int32(mat.n)}, mat.nnz, mat.rowind, mat.colptr, mat.value
}

// The function to generate CSC matrix for a transition matrix.
// src and dest are objects to indicate mark groups, and tt is a transition type to pick up
// as an element of transition.
//
// The behavior of this function is unknown if there are two or more transitions between the same marks.
//
func (mg *MarkingGraph) getTransMatrix(gtr GroupTrans) (*CSC, []float64) {

	// The structure to represent an element of COO matrix
	type matelem struct {
		i   int
		j   int
		val float64
	}

	m := len(mg.groupToMark[gtr.src])
	n := len(mg.groupToMark[gtr.dest])
	sum := make([]float64, m, m)
	elems := make([]matelem, 0)
	if gtr.src == gtr.dest && gtr.transtype == TransEXP {
		for i := 0; i < m; i++ {
			e := matelem{
				i:   i,
				j:   i,
				val: 0,
			}
			elems = append(elems, e)
		}
	}
	for _, lset := range mg.groupTransToLink[gtr] {
		e := matelem{
			i:   mg.markToInt[lset.src],
			j:   mg.markToInt[lset.dest],
			val: lset.tr.getValue(mg.net, lset.src),
		}
		// log.Print("add ", e)
		elems = append(elems, e)
		sum[e.i] += e.val
	}
	// log.Print("before ", elems)
	sort.Slice(elems, func(i, j int) bool {
		if elems[i].j == elems[j].j {
			return elems[i].i < elems[j].i
		} else {
			return elems[i].j < elems[j].j
		}
	})
	// log.Print("after ", elems)
	colptr := make([]int, n+1)
	rowind := make([]int, len(elems))
	value := make([]float64, len(elems))
	z := 0
	j := 0
	colptr[j] = z
	for _, e := range elems {
		if j != e.j {
			for u := j + 1; u <= e.j; u++ {
				colptr[u] = z
			}
			j = e.j
		}
		rowind[z] = e.i
		value[z] = e.val
		z += 1
	}
	for u := j + 1; u <= n; u++ {
		colptr[u] = z
	}
	// log.Print(m, " ", n, " ", gtr.src, gtr.dest, gtr.transtype, sum)
	// log.Print(m, " ", n, " ", rowind, colptr, value)
	return &CSC{
		m:      m,
		n:      n,
		nnz:    len(elems),
		colptr: colptr,
		rowind: rowind,
		value:  value,
	}, sum
}

func (mg *MarkingGraph) TransMatrix() (map[GroupTrans]*CSC, map[GroupTrans]*CSC, map[GroupTrans]*CSC) {
	// type gengroup struct {
	// 	g  *Group
	// 	tr *Trans
	// }
	immsums := make(map[*Group][]float64)
	expsums := make(map[*Group][]float64)
	// gensums := make(map[gengroup][]float64)

	immmats := make(map[GroupTrans]*CSC)
	expmats := make(map[GroupTrans]*CSC)
	genmats := make(map[GroupTrans]*CSC)
	for gtr, _ := range mg.groupTransToLink {
		src := gtr.src
		switch gtr.transtype {
		case TransIMM:
			mat, sum := mg.getTransMatrix(gtr)
			immmats[gtr] = mat
			if _, ok := immsums[src]; ok {
				for i, s := range sum {
					immsums[src][i] += s
				}
			} else {
				immsums[src] = sum
			}
		case TransEXP:
			mat, sum := mg.getTransMatrix(gtr)
			expmats[gtr] = mat
			if _, ok := expsums[src]; ok {
				for i, s := range sum {
					expsums[src][i] += s
				}
			} else {
				expsums[src] = sum
			}
		case TransGEN:
			mat, _ := mg.getTransMatrix(gtr)
			genmats[gtr] = mat
			// geng := gengroup{
			// 	g:  src,
			// 	tr: gtr.gentrans,
			// }
			// if _, ok := gensums[geng]; ok {
			// 	for i, s := range sum {
			// 		gensums[geng][i] += s
			// 	}
			// } else {
			// 	gensums[geng] = sum
			// }
		default:
			log.Panic("Unknown transtype")
		}
	}
	for g, _ := range expsums {
		gtr := GroupTrans{
			src:       g,
			dest:      g,
			transtype: TransEXP,
			gentrans:  nil,
		}
		if _, ok := expmats[gtr]; ok == false {
			mat, _ := mg.getTransMatrix(gtr)
			expmats[gtr] = mat
			mg.grouplinks = append(mg.grouplinks, gtr)
		}
	}

	// diag
	for gtr, mat := range immmats {
		sum := immsums[gtr.src]
		for i := 0; i < mat.nnz; i++ {
			mat.value[i] /= sum[mat.rowind[i]]
		}
	}
	for gtr, mat := range expmats {
		if gtr.src == gtr.dest {
			sum := expsums[gtr.src]
			for j := 0; j < mat.n; j++ {
				for z := mat.colptr[j]; z < mat.colptr[j+1]; z++ {
					i := mat.rowind[z]
					if i == j {
						mat.value[z] = -sum[i]
						break
					}
				}
			}
		}
	}
	// for gtr, mat := range genmats {
	// 	geng := gengroup{
	// 		g:  gtr.src,
	// 		tr: gtr.gentrans,
	// 	}
	// 	sum := gensums[geng]
	// 	for i := 0; i < mat.nnz; i++ {
	// 		mat.value[i] /= sum[mat.rowind[i]]
	// 	}
	// }
	return expmats, immmats, genmats
}

func (mg *MarkingGraph) GroupLabels() map[*Group]string {
	labels := make(map[*Group]string)
	g2i := make(map[*GenVec]int)
	count := 0
	for _, g := range mg.groups {
		if v, ok := g2i[g.gv]; ok {
			switch g.gtype {
			case IMMGroup:
				labels[g] = fmt.Sprintf("I%d", v)
			case GENGroup:
				labels[g] = fmt.Sprintf("G%d", v)
			case ABSGroup:
				labels[g] = fmt.Sprintf("A%d", v)
			default:
				log.Panic("Unknown grouptype")
			}
		} else {
			g2i[g.gv] = count
			switch g.gtype {
			case IMMGroup:
				labels[g] = fmt.Sprintf("I%d", count)
			case GENGroup:
				labels[g] = fmt.Sprintf("G%d", count)
			case ABSGroup:
				labels[g] = fmt.Sprintf("A%d", count)
			default:
				log.Panic("Unknown grouptype")
			}
			count++
		}
	}
	return labels
}

func (mg *MarkingGraph) TransLabels() map[GroupTrans]string {
	labels := make(map[GroupTrans]string)
	tr2i := make(map[*Trans]int)
	count := 0
	for _, gtr := range mg.grouplinks {
		if gtr.gentrans == nil {
			if gtr.transtype == TransIMM {
				labels[gtr] = "I"
			} else {
				labels[gtr] = "E"
			}
		} else {
			tr := gtr.gentrans.getTrans()
			if v, ok := tr2i[tr]; ok {
				labels[gtr] = fmt.Sprintf("P%d", v)
			} else {
				tr2i[tr] = count
				labels[gtr] = fmt.Sprintf("P%d", count)
				count++
			}
		}
	}
	return labels
}

func (mg *MarkingGraph) InitVector() map[*Group][]float64 {
	ivector := make(map[*Group][]float64)
	for g, mset := range mg.groupToMark {
		vec := make([]float64, len(mset), len(mset))
		if g == mg.markToGroup[mg.imark] {
			vec[mg.markToInt[mg.imark]] = 1
		}
		ivector[g] = vec
	}
	return ivector
}

func (mg *MarkingGraph) RewardVector() map[string]map[*Group][]float64 {
	result := make(map[string]map[*Group][]float64)
	for label, _ := range mg.net.rewardfunc {
		rvector := make(map[*Group][]float64)
		rewardfunc := mg.net.rewardfunc[label]
		for g, mset := range mg.groupToMark {
			vec := make([]float64, len(mset), len(mset))
			for _, m := range mset {
				vec[mg.markToInt[m]] = rewardfunc(m.toSlice())
			}
			rvector[g] = vec
		}
		result[label] = rvector
	}
	return result
}

func (g *GenVec) makeLabel(net *Net) string {
	if g.IsAnyEnabled() == false {
		return "EXP"
	}
	s := g.toSlice()
	result := make([]string, 0, len(s))
	for i, x := range s {
		if x == ENABLE {
			result = append(result, fmt.Sprintf("%s->%s", net.genlist[i].label, x.String()))
		}
	}
	return strings.Join(result, ",")
}

func (mg *MarkingGraph) Print() {
	immstates := 0
	genstates := 0
	absstates := 0
	immnnz := 0
	gennnz := 0
	absnnz := 0
	grouplabel := mg.GroupLabels()
	nnz := make(map[*Group]int)
	for _, gtr := range mg.grouplinks {
		src := gtr.src
		nnz[src] = nnz[src] + len(mg.groupTransToLink[gtr])
	}
	writer := bytes.NewBuffer(make([]byte, 0, 1024))
	i := 0
	prevgv := new(GenVec)
	for i < len(mg.groups) {
		g := mg.groups[i]
		if prevgv != g.gv {
			fmt.Fprintf(writer, "# (%s)\n", g.gv.makeLabel(mg.net))
			prevgv = g.gv
		}
		switch g.gtype {
		case IMMGroup:
			immstates += len(mg.groupToMark[g])
			immnnz += nnz[g]
			fmt.Fprintf(writer, "#   # of IMM states     (%3s) : %d (%d)\n", grouplabel[g], len(mg.groupToMark[g]), nnz[g])
		case GENGroup:
			genstates += len(mg.groupToMark[g])
			gennnz += nnz[g]
			fmt.Fprintf(writer, "#   # of EXP/GEN states (%3s) : %d (%d)\n", grouplabel[g], len(mg.groupToMark[g]), nnz[g])
		case ABSGroup:
			absstates += len(mg.groupToMark[g])
			absnnz += nnz[g]
			fmt.Fprintf(writer, "#   # of ABS states     (%3s) : %d (%d)\n", grouplabel[g], len(mg.groupToMark[g]), nnz[g])
		default:
			log.Panic("Unknown grouptype")
		}
		i++
	}
	fmt.Printf("# # of total states         : %d (%d)\n", immstates+genstates+absstates, immnnz+gennnz+absnnz)
	fmt.Printf("# # of total EXP/GEN states : %d (%d)\n", genstates, gennnz)
	fmt.Printf("# # of total IMM states     : %d (%d)\n", immstates, immnnz)
	fmt.Printf("# # of total ABS states     : %d (%d)\n", absstates, absnnz)
	fmt.Println(writer.String())
}
