package petrinet

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"sort"
	"strconv"
	"strings"
)

type linkTransInterface interface {
	getTrans() *Trans
	getValue(*Net, *Mark) float64
}

func (tr *ImmTrans) getValue(net *Net, mark *Mark) float64 {
	if f := tr.getTrans().ratefunc; f != nil {
		return f(mark.toSlice())
	} else {
		return tr.weight
	}
}

func (tr *ExpTrans) getValue(net *Net, mark *Mark) float64 {
	if f := tr.getTrans().ratefunc; f != nil {
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
	label string
}

func newGroup(gtype GroupType, gv *GenVec, label string) *Group {
	return &Group{
		gtype: gtype,
		gv:    gv,
		label: label,
	}
}

func (g *Group) String() string {
	return fmt.Sprintf("[%d %s %s]", g.gtype, g.gv, g.label)
}

// A generator to make a unique instance of Group
type groupGenerator struct {
	key  []byte
	data map[string]*Group
}

func newGroupGenerator() *groupGenerator {
	return &groupGenerator{
		key:  make([]byte, 0, 20),
		data: make(map[string]*Group),
	}
}

// The method to generate a unique instance of Group
func (g *groupGenerator) generate(gtype GroupType, gv *GenVec, label string) *Group {
	g.key = g.key[:0]
	g.key = append(g.key, strconv.Itoa(int(gtype))...)
	g.key = append(g.key, gv.String()...)
	g.key = append(g.key, label...)
	key := string(g.key)
	if mg, ok := g.data[key]; ok {
		return mg
	} else {
		newmg := newGroup(gtype, gv, label)
		g.data[key] = newmg
		return newmg
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
	groupToMark      map[*Group][]*Mark    // map from a group to a set of marks
	groupTransToLink map[GroupTrans][]Link // map from a group transition to a set of links
	groupGenerator   *groupGenerator       // a generator to make an instance of group
	clamps           []ClampSummary        // places clamped while firing (see clamp.go)
}

type makingGraphGenerator interface {
	create(net *Net, imark []MarkInt, opts SearchOptions) (*Mark, []*Mark, []Link, error)
	clampEvents() []ClampSummary
}

// CreateMarkingGraph runs a reachability search under opts. It returns an error --
// a *StateLimitError -- rather than a partial graph when the search hits its state
// limit: transition matrices taken from a truncated graph look perfectly valid and
// are not, so there is nothing safe to hand back.
func CreateMarkingGraph(net *Net, imark []MarkInt, method makingGraphGenerator, opts SearchOptions) (*MarkingGraph, error) {
	m0, marks, links, err := method.create(net, imark, opts)
	if err != nil {
		return nil, err
	}
	mg := newMarkingGraph(net, m0, marks, links)
	mg.clamps = method.clampEvents()
	return mg, nil
}

// CreateMarkingGraphWithDFS searches with DefaultSearchOptions, which caps the state
// count at DefaultMaxStates. Use the Opts form to raise or remove that.
func CreateMarkingGraphWithDFS(net *Net, imark []MarkInt) (*MarkingGraph, error) {
	return CreateMarkingGraph(net, imark, new(dfs), DefaultSearchOptions())
}

func CreateMarkingGraphWithDFSOpts(net *Net, imark []MarkInt, opts SearchOptions) (*MarkingGraph, error) {
	return CreateMarkingGraph(net, imark, new(dfs), opts)
}

func CreateMarkingGraphWithDFSTangible(net *Net, imark []MarkInt) (*MarkingGraph, error) {
	return CreateMarkingGraph(net, imark, new(dfstangible), DefaultSearchOptions())
}

func CreateMarkingGraphWithDFSTangibleOpts(net *Net, imark []MarkInt, opts SearchOptions) (*MarkingGraph, error) {
	return CreateMarkingGraph(net, imark, new(dfstangible), opts)
}

func makeGroupString(net *Net, m []MarkInt) string {
	str := make([]string, 0, len(net.markgroupstring))
	for _, f := range net.markgroupstring {
		str = append(str, f(m))
	}
	return strings.Join(str, ",")
}

// The method to create a marking graph. This is only called from CreateMarkingGraph
func newMarkingGraph(net *Net,
	m0 *Mark,
	marks []*Mark,
	links []Link) *MarkingGraph {

	// make groupmarks
	generator := newGroupGenerator()
	groupToMark := make(map[*Group][]*Mark)
	for _, m := range marks {
		g := generator.generate(m.gtype, m.genvec, makeGroupString(net, m.toSlice()))
		m.group = g
		if mset, ok := groupToMark[g]; ok {
			groupToMark[g] = append(mset, m)
		} else {
			groupToMark[g] = []*Mark{m}
		}
	}

	// numbering for marks
	for _, mset := range groupToMark {
		for i, m := range mset {
			m.index = i
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
		if groups[i].label == groups[j].label {
			return groups[i].gtype < groups[j].gtype
		} else {
			return groups[i].label < groups[j].label
		}
	})

	// Group the links. The links are partitioned in place and each group is handed a
	// subslice of the one array: this used to copy every link into a per-group slice,
	// so the whole link list existed twice -- 26% of the graph's retained heap on
	// iaas_cloud with n=5, where there are 3.1 million links.
	//
	// The sort is stable, so the links within a group stay in the order they were
	// found. getTransMatrix sums them in that order, and a float sum depends on it:
	// an unstable sort here would move the last ULP of the generator diagonal.
	groupTransToLink := make(map[GroupTrans][]Link, len(links)/8+1)
	grouplinks := make([]GroupTrans, 0)
	groupOf := make(map[GroupTrans]int32)
	// linkGroup[i] is the index into grouplinks of links[i]'s group. int32 rather than
	// GroupTrans: 4 bytes per link instead of 40.
	linkGroup := make([]int32, len(links))
	for i, l := range links {
		tr := GroupTrans{
			src:       l.src.group,
			dest:      l.dest.group,
			transtype: l.tt,
			gentrans:  nil,
		}
		if l.tt == TransGEN {
			tr.gentrans = l.tr.getTrans()
		}
		id, ok := groupOf[tr]
		if !ok {
			// First appearance decides the order of grouplinks, and GroupLabels and
			// TransLabels number G0/I0/A0 and P0/P1 by walking it. Sorting the links
			// below must not disturb it: it would rename the output variables.
			id = int32(len(grouplinks))
			groupOf[tr] = id
			grouplinks = append(grouplinks, tr)
		}
		linkGroup[i] = id
	}
	order := make([]int, len(links))
	for i := range order {
		order[i] = i
	}
	sort.SliceStable(order, func(a, b int) bool { return linkGroup[order[a]] < linkGroup[order[b]] })
	permuteLinks(links, order)
	sort.SliceStable(linkGroup, func(a, b int) bool { return linkGroup[a] < linkGroup[b] })
	for start := 0; start < len(links); {
		end := start
		for end < len(links) && linkGroup[end] == linkGroup[start] {
			end++
		}
		groupTransToLink[grouplinks[linkGroup[start]]] = links[start:end:end]
		start = end
	}

	return &MarkingGraph{
		net:              net,
		imark:            m0,
		marks:            marks,
		groups:           groups,
		links:            links,
		grouplinks:       grouplinks,
		groupToMark:      groupToMark,
		groupTransToLink: groupTransToLink,
		groupGenerator:   generator,
	}
}

// permuteLinks rearranges links so that links[i] ends up holding what links[order[i]]
// held, in place. Sorting the links directly would need the group index to travel with
// each one, which is what the separate linkGroup array avoids.
func permuteLinks(links []Link, order []int) {
	done := make([]bool, len(order))
	for i := range order {
		if done[i] || order[i] == i {
			done[i] = true
			continue
		}
		// Walk the cycle this position belongs to, carrying one element around it.
		hold := links[i]
		j := i
		for {
			done[j] = true
			k := order[j]
			if k == i {
				links[j] = hold
				break
			}
			links[j] = links[k]
			j = k
		}
	}
}

// ClampEvents reports the places that had to be clamped while the marking graph was
// built, collapsed per (transition, place, bound). A non-empty result means at least one
// transition's destination marking was not the real one, so the graph -- and any
// generator matrix taken from it -- is not exact. The usual cause is a place that holds
// more tokens than its capacity allows; the default capacity is 255, so a place that is
// meant to hold more needs an explicit `max`.
//
// FormatClampEvents renders these for display.
func (mg *MarkingGraph) ClampEvents() []ClampSummary {
	return mg.clamps
}

func (mg *MarkingGraph) ToMarkDot(writer io.Writer) {
	fmt.Fprintf(writer, "digraph { layout=dot; overlap=false; splines=true;\n")
	for _, mark := range mg.marks {
		switch markgroup := mark.group; markgroup.gtype {
		case IMMGroup:
			fmt.Fprintf(writer, "\"%p\" [shape=circle, label=\"%d\n%s\", style=filled];\n", mark, mark.index, markgroup.gv)
		case GENGroup:
			fmt.Fprintf(writer, "\"%p\" [shape=circle, label=\"%d\n%s\"];\n", mark, mark.index, markgroup.gv)
		case ABSGroup:
			fmt.Fprintf(writer, "\"%p\" [shape=circle, label=\"%d\n%s\"];\n", mark, mark.index, markgroup.gv)
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
		switch mark.group.gtype {
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
		switch markgroup := mark.group; markgroup.gtype {
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
			i:   lset.src.index,
			j:   lset.dest.index,
			val: lset.tr.getValue(mg.net, lset.src),
		}
		// log.Print("add ", e)
		elems = append(elems, e)
		sum[e.i] += e.val
	}
	// log.Print("before ", elems)
	// Stable: entries with the same (i,j) are summed below in this order, and a
	// float sum depends on it. sort.Slice left the last ULP of the diagonal varying
	// from run to run on example/spnp_example5.spn.
	sort.SliceStable(elems, func(i, j int) bool {
		if elems[i].j == elems[j].j {
			return elems[i].i < elems[j].i
		} else {
			return elems[i].j < elems[j].j
		}
	})
	// log.Print("after ", elems)
	colptr := make([]int, n+1)
	rowind := make([]int, 0, len(elems))
	value := make([]float64, 0, len(elems))
	z := 0
	j := 0
	colptr[j] = z
	previ, prevj := -1, -1
	for _, e := range elems {
		if j != e.j {
			for u := j + 1; u <= e.j; u++ {
				colptr[u] = z
			}
			j = e.j
		}
		if e.i == previ && e.j == prevj {
			value[z-1] += e.val
		} else {
			rowind = append(rowind, e.i)
			value = append(value, e.val)
			previ, prevj = e.i, e.j
			z += 1
		}
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
	immsums := make(map[*Group][]float64)
	expsums := make(map[*Group][]float64)
	expgengroups := make(map[*Group]struct{})

	immmats := make(map[GroupTrans]*CSC)
	expmats := make(map[GroupTrans]*CSC)
	genmats := make(map[GroupTrans]*CSC)
	// grouplinks holds exactly the keys of groupTransToLink, in the order the links
	// were found. Ranging over the map instead made the order of the `expsums[src][i] +=`
	// accumulations vary per run, and with it the last ULP of the generator diagonal.
	for _, gtr := range mg.grouplinks {
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
			expgengroups[src] = struct{}{}
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
			expgengroups[src] = struct{}{}
			mat, _ := mg.getTransMatrix(gtr)
			genmats[gtr] = mat
		default:
			log.Panic("Unknown transtype")
		}
	}
	// group Gx should have GxGxE even if there is no EXP trans
	for _, g := range mg.groups {
		if _, ok := expgengroups[g]; !ok {
			continue
		}
		gtr := GroupTrans{
			src:       g,
			dest:      g,
			transtype: TransEXP,
			gentrans:  nil,
		}
		if _, ok := expmats[gtr]; ok == false {
			mat, sum := mg.getTransMatrix(gtr)
			expmats[gtr] = mat
			if _, ok := expsums[g]; ok == false {
				expsums[g] = sum
			}
			mg.grouplinks = append(mg.grouplinks, gtr)
		}
	}

	// normalize:
	//   immmat -> sum of row becomes 1
	//   expmat -> put diag elements
	//   genmat -> no need because genmat has only one element 1 for each row
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
	return expmats, immmats, genmats
}

func (mg *MarkingGraph) GroupLabels() map[*Group]string {
	type GG struct {
		gv    *GenVec
		label string
	}
	labels := make(map[*Group]string)
	g2i := make(map[GG]int)
	count := 0
	for _, g := range mg.groups {
		gg := GG{gv: g.gv, label: g.label}
		if v, ok := g2i[gg]; ok {
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
			g2i[gg] = count
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

// Net is the net the graph was built from. StateMarkings is indexed by its
// PlaceLabels order, and a caller outside the package needs that order to read it.
func (mg *MarkingGraph) Net() *Net { return mg.net }

// StateMarkings is which marking each row of a group's matrices is. The columns are
// places, in Net.PlaceLabels order.
//
// A result file used to carry the matrices and nothing that said what a row meant: the
// markings went to a separate text file, written only when -s was given. Anything
// reading the matrices back -- a cross-check against another implementation, most of
// all -- has to be able to key a row on its marking.
func (mg *MarkingGraph) StateMarkings() map[*Group][][]MarkInt {
	marks := make(map[*Group][][]MarkInt, len(mg.groups))
	for _, g := range mg.groups {
		tmp := make([][]MarkInt, len(mg.groupToMark[g]))
		for k, m := range mg.groupToMark[g] {
			// A copy: toSlice hands back the interned marking's own slice, and this
			// leaves the package.
			tmp[k] = append([]MarkInt(nil), m.toSlice()...)
		}
		marks[g] = tmp
	}
	return marks
}

func (mg *MarkingGraph) StateLabels() map[*Group][]string {
	labels := make(map[*Group][]string)
	for g, marks := range mg.StateMarkings() {
		tmp := make([]string, len(marks))
		for k, m := range marks {
			str := make([]string, 0)
			for i, n := range m {
				if n > 0 {
					str = append(str, fmt.Sprintf("%s:%d", mg.net.placelist[i].label, n))
				}
			}
			tmp[k] = strings.Join(str, ",")
		}
		labels[g] = tmp
	}
	return labels
}

func (mg *MarkingGraph) InitVector() map[*Group][]float64 {
	ivector := make(map[*Group][]float64)
	for g, mset := range mg.groupToMark {
		vec := make([]float64, len(mset), len(mset))
		if g == mg.imark.group {
			vec[mg.imark.index] = 1
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
				vec[m.index] = rewardfunc(m.toSlice())
			}
			rvector[g] = vec
		}
		result[label] = rvector
	}
	return result
}

/*
	Output interfaces
*/

func (mg *MarkingGraph) Summary() {
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
	prevgv := new(GenVec)
	for _, g := range mg.groups {
		if prevgv != g.gv {
			fmt.Fprintf(writer, "(%s)\n", g.gv.makeLabel(mg.net))
			prevgv = g.gv
		}
		switch g.gtype {
		case IMMGroup:
			immstates += len(mg.groupToMark[g])
			immnnz += nnz[g]
			fmt.Fprintf(writer, "  # of IMM states     (%3s) : %d (%d) %s\n", grouplabel[g], len(mg.groupToMark[g]), nnz[g], g.label)
		case GENGroup:
			genstates += len(mg.groupToMark[g])
			gennnz += nnz[g]
			fmt.Fprintf(writer, "  # of EXP/GEN states (%3s) : %d (%d) %s\n", grouplabel[g], len(mg.groupToMark[g]), nnz[g], g.label)
		case ABSGroup:
			absstates += len(mg.groupToMark[g])
			absnnz += nnz[g]
			fmt.Fprintf(writer, "  # of ABS states     (%3s) : %d (%d) %s\n", grouplabel[g], len(mg.groupToMark[g]), nnz[g], g.label)
		default:
			log.Panic("Unknown grouptype")
		}
	}
	fmt.Printf("# of total states         : %d (%d)\n", immstates+genstates+absstates, immnnz+gennnz+absnnz)
	fmt.Printf("# of total EXP/GEN states : %d (%d)\n", genstates, gennnz)
	fmt.Printf("# of total IMM states     : %d (%d)\n", immstates, immnnz)
	fmt.Printf("# of total ABS states     : %d (%d)\n", absstates, absnnz)
	fmt.Println(writer.String())
}

func (mg *MarkingGraph) WriteState(writer io.Writer) {
	statelabels := mg.StateLabels()
	grouplabel := mg.GroupLabels()
	for _, g := range mg.groups {
		fmt.Fprintf(writer, "# %s\n", grouplabel[g])
		for i, s := range statelabels[g] {
			fmt.Fprintf(writer, "%d : {%s}\n", i, s)
		}
	}
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

// GenInfo describes one general transition as a result file records it: the name it has
// in the definition, its firing-time distribution, and -- when it is being reported for a
// group -- whether it is aging, preempted or disabled there.
//
// A result file used to carry the general blocks and nothing else about the general
// transitions. A block is a 0/1 jump matrix, so `det(5)` and `det(99)` produce byte-
// identical files, and neither says which transition a `P<k>` block belongs to. Both are
// needed to solve the regenerative process the file describes.
type GenInfo struct {
	Label  string // the transition's name in the definition
	Dist   string // its distribution, in the definition language: det(2), unif(1,3), ...
	Status string // for a group: E aging, P preempted, D disabled. Empty otherwise.
}

// genTransOf finds the GenTrans whose base is tr. GroupTrans stores the base *Trans, so
// the distribution has to be looked up.
func (net *Net) genTransOf(tr *Trans) *GenTrans {
	for _, g := range net.genlist {
		if g.Trans == tr {
			return g
		}
	}
	return nil
}

// BlockGenTrans is which general transition each GEN block belongs to, keyed the same way
// TransLabels is -- so `TransLabels()[gtr]` names the block ("P0") and this says what it
// is.
func (mg *MarkingGraph) BlockGenTrans() map[GroupTrans]GenInfo {
	out := make(map[GroupTrans]GenInfo)
	for _, gtr := range mg.grouplinks {
		if gtr.gentrans == nil {
			continue
		}
		if g := mg.net.genTransOf(gtr.gentrans); g != nil {
			out[gtr] = GenInfo{Label: g.label, Dist: g.dist.String()}
		}
	}
	return out
}

// GroupGens is, per group, the general transitions that are aging or preempted in it --
// the ones whose distributions govern how long the group is occupied. A group where every
// general transition is disabled is absent rather than present and empty.
func (mg *MarkingGraph) GroupGens() map[*Group][]GenInfo {
	out := make(map[*Group][]GenInfo)
	for _, g := range mg.groups {
		if g.gv == nil {
			continue
		}
		var infos []GenInfo
		for i, st := range g.gv.toSlice() {
			if st == DISABLE || i >= len(mg.net.genlist) {
				continue
			}
			gen := mg.net.genlist[i]
			infos = append(infos, GenInfo{Label: gen.label, Dist: gen.dist.String(), Status: st.String()})
		}
		if len(infos) > 0 {
			out[g] = infos
		}
	}
	return out
}
