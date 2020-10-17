package petrinet

import (
	"fmt"
	"io"
	"sort"
)

type Place struct {
	label   string    // the label of node
	index   int       // the index of index
	max     MarkInt   // the maximum number of tokens
	inarcs  []*OutArc // the list of inarcs
	outarcs []*InArc  // the list of outarcs
}

func newPlace(label string, index int, max MarkInt) *Place {
	return &Place{
		label:   label,
		index:   index,
		max:     max,
		inarcs:  make([]*OutArc, 0),
		outarcs: make([]*InArc, 0),
	}
}

func (p *Place) getPlace() *Place {
	return p
}

func (p *Place) addInArc(arc *InArc) {
	p.outarcs = append(p.outarcs, arc)
}

func (p *Place) addOutArc(arc *OutArc) {
	p.inarcs = append(p.inarcs, arc)
}

type TransType int

const (
	TransIMM TransType = iota + 1
	TransEXP
	TransGEN
)

func (t TransType) String() string {
	switch t {
	case TransIMM:
		return "IMM"
	case TransEXP:
		return "EXP"
	case TransGEN:
		return "GEN"
	default:
		return "Unknown"
	}
}

type Trans struct {
	label      string    // the label of node
	index      int       // the index of index
	priority   int       // the priority of transition
	vanishable bool      // the transition can be vanished or not
	inarcs     []*InArc  // the list of inarcs
	outarcs    []*OutArc // the list of outarcs
}

type ImmTrans struct {
	*Trans
	weight float64 // the weight of transition
}

type ExpTrans struct {
	*Trans
	rate float64 // the rate of transition
}

type GenTransPolicy int

const (
	GenTransPolicyPRD GenTransPolicy = iota + 1 // PRD: Premenptive different
	GenTransPolicyPRS                           // PRS: Premptive resume
	GenTransPolicyPRI                           // PRI: Preemptive repeat
)

type GenTrans struct {
	*Trans
	dist   *Distribution  // the distribution
	policy GenTransPolicy // policy for preemption
}

func (p *Place) GetLabel() string {
	return p.label
}

func (p *Place) GetIndex() int {
	return p.index
}

func (tr *Trans) GetLabel() string {
	return tr.label
}

func (tr *Trans) GetIndex() int {
	return tr.index
}

func (tr *Trans) getTrans() *Trans {
	return tr
}

func (tr *ImmTrans) getTrans() *Trans {
	return tr.Trans
}

func (tr *ExpTrans) getTrans() *Trans {
	return tr.Trans
}

func (tr *GenTrans) getTrans() *Trans {
	return tr.Trans
}

func (tr *Trans) addInArc(arc *InArc) {
	tr.inarcs = append(tr.inarcs, arc)
}

func (tr *Trans) addOutArc(arc *OutArc) {
	tr.outarcs = append(tr.outarcs, arc)
}

func (tr *ImmTrans) addInArc(arc *InArc) {
	tr.Trans.addInArc(arc)
}

func (tr *ImmTrans) addOutArc(arc *OutArc) {
	tr.Trans.addOutArc(arc)
}

func (tr *ExpTrans) addInArc(arc *InArc) {
	tr.Trans.addInArc(arc)
}

func (tr *ExpTrans) addOutArc(arc *OutArc) {
	tr.Trans.addOutArc(arc)
}

func (tr *GenTrans) addInArc(arc *InArc) {
	tr.Trans.addInArc(arc)
}

func (tr *GenTrans) addOutArc(arc *OutArc) {
	tr.Trans.addOutArc(arc)
}

func newTrans(label string, index int, priority int, vanishable bool) *Trans {
	return &Trans{
		label:      label,
		index:      index,
		priority:   priority,
		vanishable: vanishable,
		inarcs:     make([]*InArc, 0),
		outarcs:    make([]*OutArc, 0),
	}
}

func newImmTrans(label string, index int, priority int, vanishable bool, weight float64) *ImmTrans {
	return &ImmTrans{
		Trans:  newTrans(label, index, priority, vanishable),
		weight: weight,
	}
}

func newExpTrans(label string, index int, priority int, vanishable bool, rate float64) *ExpTrans {
	return &ExpTrans{
		Trans: newTrans(label, index, priority, vanishable),
		rate:  rate,
	}
}

func newGenTrans(label string, index int, priority int, vanishable bool, dist *Distribution, policy GenTransPolicy) *GenTrans {
	return &GenTrans{
		Trans:  newTrans(label, index, priority, vanishable),
		dist:   dist,
		policy: policy,
	}
}

type InArc struct {
	src          *Place  // source node
	dest         *Trans  // destination node
	multiplicity MarkInt // multiplicity
	inhibit      bool    // the arc is inhibit or not
}

type OutArc struct {
	src          *Trans  // source node
	dest         *Place  // destination node
	multiplicity MarkInt // multiplicity
}

type placeInterface interface {
	getPlace() *Place
	addInArc(arc *InArc)
	addOutArc(arc *OutArc)
}

type transInterface interface {
	getTrans() *Trans
	addInArc(arc *InArc)
	addOutArc(arc *OutArc)
}

func newInArc(src placeInterface, dest transInterface, multiplicity MarkInt) *InArc {
	arc := InArc{
		src:          src.getPlace(),
		dest:         dest.getTrans(),
		multiplicity: multiplicity,
		inhibit:      false,
	}
	src.addInArc(&arc)
	dest.addInArc(&arc)
	return &arc
}

func newInhibitArc(src placeInterface, dest transInterface, multiplicity MarkInt) *InArc {
	arc := InArc{
		src:          src.getPlace(),
		dest:         dest.getTrans(),
		multiplicity: multiplicity,
		inhibit:      true,
	}
	src.addInArc(&arc)
	dest.addInArc(&arc)
	return &arc
}

func newOutArc(src transInterface, dest placeInterface, multiplicity MarkInt) *OutArc {
	arc := OutArc{
		src:          src.getTrans(),
		dest:         dest.getPlace(),
		multiplicity: multiplicity,
	}
	src.addOutArc(&arc)
	dest.addOutArc(&arc)
	return &arc
}

// The structure for a petrinet.
type Net struct {
	placelist     []*Place                             // the list of places
	translist     []*Trans                             // the list of transitions
	immlist       []*ImmTrans                          // the list of imm transitions
	explist       []*ExpTrans                          // the list of exp transitons
	genlist       []*GenTrans                          // the list of gen transitions
	guards        map[*Trans]func([]MarkInt) bool      // guard functions
	updates       map[*Trans]func([]MarkInt) []MarkInt // update functions
	infunc        map[*InArc]func([]MarkInt) MarkInt   // multiplicity functions
	outfunc       map[*OutArc]func([]MarkInt) MarkInt  // multiplicity functions
	ratefunc      map[*Trans]func([]MarkInt) float64   // weight and rate functions
	placelabel    map[string]*Place                    // labels for places
	translabel    map[string]*Trans                    // labels for trans
	guardstring   map[*Trans]string                    // string for guard functions
	updatesstring map[*Trans]string                    // string for guard functions
	infuncstring  map[*InArc]string                    // multiplicity functions
	outfuncstring map[*OutArc]string                   // multiplicity functions
	rewardfunc    map[string]func([]MarkInt) float64   // reward functions
	// immlist   map[*Trans]*ImmTrans            // the list of imm transitions
	// explist   map[*Trans]*ExpTrans            // the list of exp transitons
	// genlist   map[*Trans]*GenTrans            // the list of gen transitions
}

func NewNet() *Net {
	return &Net{
		placelist:     make([]*Place, 0),
		translist:     make([]*Trans, 0),
		immlist:       make([]*ImmTrans, 0),
		explist:       make([]*ExpTrans, 0),
		genlist:       make([]*GenTrans, 0),
		guards:        make(map[*Trans]func([]MarkInt) bool),
		updates:       make(map[*Trans]func([]MarkInt) []MarkInt),
		infunc:        make(map[*InArc]func([]MarkInt) MarkInt),
		outfunc:       make(map[*OutArc]func([]MarkInt) MarkInt),
		ratefunc:      make(map[*Trans]func([]MarkInt) float64),
		placelabel:    make(map[string]*Place),
		translabel:    make(map[string]*Trans),
		guardstring:   make(map[*Trans]string),
		updatesstring: make(map[*Trans]string),
		infuncstring:  make(map[*InArc]string),
		outfuncstring: make(map[*OutArc]string),
		rewardfunc:    make(map[string]func([]MarkInt) float64),
	}
}

func (net *Net) NewPlace(label string, max MarkInt) *Place {
	place := newPlace(label, 0, max)
	net.placelist = append(net.placelist, place)
	net.placelabel[label] = place
	return place
}

func (net *Net) NewImmTrans(label string, priority int, vanishable bool, weight float64) *ImmTrans {
	tr := newImmTrans(label, 0, priority, vanishable, weight)
	net.immlist = append(net.immlist, tr)
	net.translabel[label] = tr.Trans
	return tr
}

func (net *Net) NewExpTrans(label string, priority int, vanishable bool, rate float64) *ExpTrans {
	tr := newExpTrans(label, 0, priority, vanishable, rate)
	net.explist = append(net.explist, tr)
	net.translabel[label] = tr.Trans
	return tr
}

func (net *Net) NewGenTrans(label string, priority int, vanishable bool, dist *Distribution, policy GenTransPolicy) *GenTrans {
	tr := newGenTrans(label, 0, priority, vanishable, dist, policy)
	net.genlist = append(net.genlist, tr)
	net.translabel[label] = tr.Trans
	return tr
}

func (net *Net) NewInArc(src placeInterface, dest transInterface, multiplicity MarkInt) *InArc {
	arc := newInArc(src, dest, multiplicity)
	return arc
}

func (net *Net) NewInhibitArc(src placeInterface, dest transInterface, multiplicity MarkInt) *InArc {
	arc := newInhibitArc(src, dest, multiplicity)
	return arc
}

func (net *Net) NewOutArc(src transInterface, dest placeInterface, multiplicity MarkInt) *OutArc {
	arc := newOutArc(src, dest, multiplicity)
	return arc
}

func (net *Net) GetPlace(label string) (*Place, bool) {
	result, ok := net.placelabel[label]
	return result, ok
}

func (net *Net) GetTrans(label string) (*Trans, bool) {
	result, ok := net.translabel[label]
	return result, ok
}

func (net *Net) LenPlaceList() int {
	return len(net.placelist)
}

func (net *Net) SetGuard(tr transInterface, str string, guard func([]MarkInt) bool) {
	net.guardstring[tr.getTrans()] = str
	net.guards[tr.getTrans()] = guard
}

func (net *Net) SetUpdate(tr transInterface, str string, update func([]MarkInt) []MarkInt) {
	net.updatesstring[tr.getTrans()] = str
	net.updates[tr.getTrans()] = update
}

func (net *Net) SetWeightRate(tr transInterface, rate func([]MarkInt) float64) {
	net.ratefunc[tr.getTrans()] = rate
}

func (net *Net) SetInArcMulti(arc *InArc, str string, multi func([]MarkInt) MarkInt) {
	net.infuncstring[arc] = str
	net.infunc[arc] = multi
}

func (net *Net) SetOutArcMulti(arc *OutArc, str string, multi func([]MarkInt) MarkInt) {
	net.outfuncstring[arc] = str
	net.outfunc[arc] = multi
}

func (net *Net) SetReward(str string, rwd func([]MarkInt) float64) {
	net.rewardfunc[str] = rwd
}

func (net *Net) Finalize() {
	// net.sortPlaceList()
	net.sortTransList()
	net.makeTransList()
	net.indexing()
}

func (net *Net) makeTransList() {
	for _, tr := range net.immlist {
		net.translist = append(net.translist, tr.Trans)
	}
	for _, tr := range net.explist {
		net.translist = append(net.translist, tr.Trans)
	}
	for _, tr := range net.genlist {
		net.translist = append(net.translist, tr.Trans)
	}
}

func (net *Net) indexing() {
	for i, x := range net.placelist {
		x.index = i
	}
	for i, x := range net.translist {
		x.index = i
	}
}

func (net *Net) sortPlaceList() {
	sort.SliceStable(net.placelist,
		func(i, j int) bool {
			return net.placelist[i].label < net.placelist[j].label
		})
}

func (net *Net) sortTransList() {
	sort.Slice(net.immlist, func(i, j int) bool {
		if net.immlist[i].priority == net.immlist[j].priority {
			return net.immlist[i].label < net.immlist[j].label
		} else {
			return net.immlist[i].priority < net.immlist[j].priority
		}
	})
	sort.Slice(net.explist, func(i, j int) bool {
		if net.explist[i].priority == net.explist[j].priority {
			return net.explist[i].label < net.explist[j].label
		} else {
			return net.explist[i].priority < net.explist[j].priority
		}
	})
	sort.Slice(net.genlist, func(i, j int) bool {
		if net.genlist[i].priority == net.genlist[j].priority {
			return net.genlist[i].label < net.genlist[j].label
		} else {
			return net.genlist[i].priority < net.genlist[j].priority
		}
	})
}

func (net *Net) ToPNDot(writer io.Writer) {
	pnbuf := newpndot(writer)
	transtype := makeTransType(net)
	fmt.Fprintf(pnbuf.buf, "digraph { layout=dot; overlap=false; splines=true; node [fontsize=10];\n")
	number := 0
	for _, place := range net.placelist {
		if _, ok := pnbuf.visitedPlace[place]; ok == false {
			fmt.Fprintf(pnbuf.buf, "subgraph cluster%d {\n", number)
			pnbuf.drawPlace(net, transtype, place)
			fmt.Fprintf(pnbuf.buf, "}\n")
			number++
		}
	}
	for _, tr := range net.translist {
		if _, ok := pnbuf.visitedTrans[tr]; ok == false {
			fmt.Fprintf(pnbuf.buf, "subgraph cluster%d {\n", number)
			pnbuf.drawTrans(net, transtype, tr)
			fmt.Fprintf(pnbuf.buf, "}\n")
			number++
		}
	}
	fmt.Fprintf(pnbuf.buf, "}\n")
}

type pndot struct {
	buf           io.Writer
	visitedPlace  map[*Place]struct{}
	visitedTrans  map[*Trans]struct{}
	visitedInArc  map[*InArc]struct{}
	visitedOutArc map[*OutArc]struct{}
}

func newpndot(writer io.Writer) *pndot {
	return &pndot{
		buf:           writer,
		visitedPlace:  make(map[*Place]struct{}),
		visitedTrans:  make(map[*Trans]struct{}),
		visitedInArc:  make(map[*InArc]struct{}),
		visitedOutArc: make(map[*OutArc]struct{}),
	}
}

func makeTransType(net *Net) map[*Trans]int {
	transtype := make(map[*Trans]int)
	for _, tr := range net.immlist {
		transtype[tr.Trans] = 1 // IMM
	}
	for _, tr := range net.explist {
		transtype[tr.Trans] = 2 // EXP
	}
	for _, tr := range net.genlist {
		transtype[tr.Trans] = 3 // GEN
	}
	return transtype
}

func (b *pndot) drawPlace(net *Net, transtype map[*Trans]int, p *Place) {
	if _, ok := b.visitedPlace[p]; ok == true {
		return
	}
	fmt.Fprintf(b.buf, "\"%p\" [shape=circle,label=\"%s\"];\n", p, p.makeLabel(net))
	b.visitedPlace[p] = struct{}{}
	for _, arc := range p.inarcs {
		b.drawOutArc(net, transtype, arc)
	}
	for _, arc := range p.outarcs {
		b.drawInArc(net, transtype, arc)
	}
}

func (b *pndot) drawTrans(net *Net, transtype map[*Trans]int, tr *Trans) {
	if _, ok := b.visitedTrans[tr]; ok == true {
		return
	}
	switch transtype[tr] {
	case 1: // IMM
		fmt.Fprintf(b.buf, "\"%p\" [shape=box,label=\"%s\", width=0.8, height=0.02, style=\"filled,dashed\"];\n", tr, tr.makeLabel(net))
	case 2: // EXP
		fmt.Fprintf(b.buf, "\"%p\" [shape=box,label=\"%s\", width=0.8, height=0.2];\n", tr, tr.makeLabel(net))
	case 3: // GEN
		fmt.Fprintf(b.buf, "\"%p\" [shape=box,label=\"%s\", width=0.8, height=0.2, style=\"filled\"];\n", tr, tr.makeLabel(net))
	default:
		panic("error")
	}
	b.visitedTrans[tr] = struct{}{}
	for _, arc := range tr.inarcs {
		b.drawInArc(net, transtype, arc)
	}
	for _, arc := range tr.outarcs {
		b.drawOutArc(net, transtype, arc)
	}
}

func (b *pndot) drawInArc(net *Net, transtype map[*Trans]int, arc *InArc) {
	if _, ok := b.visitedInArc[arc]; ok == true {
		return
	}
	if arc.inhibit == false {
		fmt.Fprintf(b.buf, "\"%p\"->\"%p\" [label=\"%s\"];\n", arc.src, arc.dest, arc.makeLabel(net))
	} else {
		fmt.Fprintf(b.buf, "\"%p\"->\"%p\" [label=\"%s\", arrowhead=odot];\n", arc.src, arc.dest, arc.makeLabel(net))
	}
	b.visitedInArc[arc] = struct{}{}
	b.drawPlace(net, transtype, arc.src)
	b.drawTrans(net, transtype, arc.dest)
}

func (b *pndot) drawOutArc(net *Net, transtype map[*Trans]int, arc *OutArc) {
	if _, ok := b.visitedOutArc[arc]; ok == true {
		return
	}
	fmt.Fprintf(b.buf, "\"%p\"->\"%p\" [label=\"%s\"];\n", arc.src, arc.dest, arc.makeLabel(net))
	b.visitedOutArc[arc] = struct{}{}
	b.drawTrans(net, transtype, arc.src)
	b.drawPlace(net, transtype, arc.dest)
}

func (place *Place) makeLabel(net *Net) string {
	// TODO
	return place.label
}

func (trans *Trans) makeLabel(net *Net) string {
	if s, ok := net.guardstring[trans]; ok {
		return fmt.Sprintf("%s\n[%s]", trans.label, s)
	} else {
		return trans.label
	}
}

func (arc *InArc) makeLabel(net *Net) string {
	if s, ok := net.infuncstring[arc]; ok {
		return s
	} else {
		if arc.multiplicity == 1 {
			return ""
		} else {
			return fmt.Sprintf("%d", arc.multiplicity)
		}
	}
}

func (arc *OutArc) makeLabel(net *Net) string {
	if s, ok := net.outfuncstring[arc]; ok {
		return s
	} else {
		if arc.multiplicity == 1 {
			return ""
		} else {
			return fmt.Sprintf("%d", arc.multiplicity)
		}
	}
}
