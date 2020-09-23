package petrinet

type Place struct {
	label   string    // the label of node
	index   int       // the index of index
	max     markInt   // the maximum number of tokens
	inarcs  []*OutArc // list of inarcs
	outarcs []*InArc  // list of outarcs
}

func newPlace(label string, index int, max markInt) *Place {
	return &Place{
		label:   label,
		index:   index,
		max:     max,
		inarcs:  make([]*OutArc, 0),
		outarcs: make([]*InArc, 0),
	}
}

type Trans struct {
	label      string    // the label of node
	index      int       // the index of index
	priority   int       // the priority of transition
	vanishable bool      // the transition can be vanished or not
	inarcs     []*InArc  // list of inarcs
	outarcs    []*OutArc // list of outarcs
	// guard
	// update
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

type InArc struct {
	src          *Place  // source node
	dest         *Trans  // destination node
	multiplicity markInt // multiplicity
	inhibit      bool    // the arc is inhibit or not
}

type OutArc struct {
	src          *Trans  // source node
	dest         *Place  // destination node
	multiplicity markInt // multiplicity
}

func newInArc(src *Place, dest *Trans, multiplicity markInt) *InArc {
	arc := &InArc{
		src:          src,
		dest:         dest,
		multiplicity: multiplicity,
		inhibit:      false,
	}
	src.outarcs = append(src.outarcs, arc)
	dest.inarcs = append(dest.inarcs, arc)
	return arc
}

func newInhibitArc(src *Place, dest *Trans, multiplicity markInt) *InArc {
	arc := &InArc{
		src:          src,
		dest:         dest,
		multiplicity: multiplicity,
		inhibit:      true,
	}
	src.outarcs = append(src.outarcs, arc)
	dest.inarcs = append(dest.inarcs, arc)
	return arc
}

func newOutArc(src *Trans, dest *Place, multiplicity markInt) *OutArc {
	arc := &OutArc{
		src:          src,
		dest:         dest,
		multiplicity: multiplicity,
	}
	src.outarcs = append(src.outarcs, arc)
	dest.inarcs = append(dest.inarcs, arc)
	return arc
}

type Net struct {
	places  []*Place                       // the list of places
	transes []*Trans                       // the list of transitions
	guards  map[*Trans]func(*Mark)bool     // guard functions
	updates map[*Trans]func(*Mark)*Mark    // update functions
	infunc  map[*InArc]func(*Mark)markInt  // update functions
	outfunc map[*OutArc]func(*Mark)markInt // update functions
}

func NewNet() *Net {
	return &Net{
		places:  make([]*Place, 0),
		transes: make([]*Trans, 0),
		guards:  make(map[*Trans]func(*Mark)bool),
		updates: make(map[*Trans]func(*Mark)*Mark),
		infunc:  make(map[*InArc]func(*Mark)markInt),
		outfunc: make(map[*OutArc]func(*Mark)markInt),
	}
}

func (net *Net) NewPlace(label string, max markInt) *Place {
	place := newPlace(label, 0, max)
	net.places = append(net.places, place)
	return place
}

func (net *Net) NewTrans(label string, priority int, vanishable bool) *Trans {
	tr := newTrans(label, 0, priority, vanishable)
	net.transes = append(net.transes, tr)
	return tr
}

func (net *Net) NewInArc(src *Place, dest *Trans, multiplicity markInt) *InArc {
	arc := newInArc(src, dest, multiplicity)
	return arc
}

func (net *Net) NewInhibitArc(src *Place, dest *Trans, multiplicity markInt) *InArc {
	arc := newInhibitArc(src, dest, multiplicity)
	return arc
}

func (net *Net) NewOutArc(src *Trans, dest *Place, multiplicity markInt) *OutArc {
	arc := newOutArc(src, dest, multiplicity)
	return arc
}

func (net *Net) Indexing() {
	for i,x := range net.places {
		x.index = i
	}
	for i,x := range net.transes {
		x.index = i
	}
}
