package petrinet

// The structure for a stack of Marks
type markStack []*Mark

// The method to create a stack
func NewMarkStack() markStack {
	return make(markStack, 0)
}

// The function for push
func (stack *markStack) push(mark *Mark) {
	*stack = append(*stack, mark)
}

// The function for pop
func (stack *markStack) pop() *Mark {
	result := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return result
}

// The function for peek
func (stack *markStack) peek() *Mark {
	return (*stack)[len(*stack)-1]
}

// The function to check the stack is empty or not
func (stack markStack) isempty() bool {
	return len(stack) == 0
}

// The structure for a set of Marks
type markSet map[*Mark]struct{}

// The method to create a set
func NewMarkSet() markSet {
	return make(map[*Mark]struct{})
}

// The method to add a mark
func (set markSet) add(mark *Mark) {
	set[mark] = struct{}{}
}

// The method to check a given mark exists or not
func (set markSet) exist(mark *Mark) bool {
	_, ok := set[mark]
	return ok
}

// The structure of DFS (depth first search) to make a marking graph.
//
//	markGenerator: A generator to create a unique object of mark
//	genVecGenerator: A gnerator to create a unique object of GenVec
//	novisited: A stack for the mark that is visited next
//	visited: A set for the marks that are already visited
//
// The following data are required to create an instance of marking graph
//
//	marks: A slice for all the visited marks
//	links: A scile for links between all marks
type dfs struct {
	markGenerator   MarkGeneratorInterface
	genVecGenerator GenVecGeneratorInterface
	novisited       markStack
	visited         markSet
	marks           []*Mark
	links           []Link        // links
	clamps          clampRecorder // places clamped while firing (see clamp.go)
	firebuf         []MarkInt     // scratch destination reused by createNextMarking
	counter         *stateCounter // state limit and progress (see statelimit.go)
}

// The method to create a marking graph.
// This is an interface for markinggraphGenerator.
func (d *dfs) create(net *Net, imark []MarkInt, opts SearchOptions) (*Mark, []*Mark, []Link, error) {
	d.markGenerator = NewMarkGenerator(len(net.placelist))
	d.genVecGenerator = NewGenVecGenerator(len(net.genlist))
	d.visited = NewMarkSet()
	d.marks = make([]*Mark, 0)
	d.novisited = NewMarkStack()
	d.links = make([]Link, 0)

	d.firebuf = make([]MarkInt, len(net.placelist))
	d.counter = newStateCounter(opts)
	m0 := d.markGenerator.genMark(imark)
	d.novisited.push(m0)
	if err := d.createMarking(net); err != nil {
		return nil, nil, nil, err
	}

	return m0, d.marks, d.links, nil
}

// The method to report the places that were clamped while firing.
// This is an interface for markinggraphGenerator.
func (d *dfs) clampEvents() []ClampSummary {
	return d.clamps.events()
}

// The method to regist a mark as a member of IMMgroup (There is one or more enabled IMM trans)
func (d *dfs) addMarkAsImm(m *Mark, g *GenVec) {
	d.marks = append(d.marks, m)
	m.genvec = g
	m.gtype = IMMGroup
}

// The method to regist a mark as a member of GENgroup (There is no enabled IMM trans)
func (d *dfs) addMarkAsGen(m *Mark, g *GenVec) {
	d.marks = append(d.marks, m)
	m.genvec = g
	m.gtype = GENGroup
}

// The method to regist a mark as a member of ABSgroup (There is no enabled trans)
func (d *dfs) addMarkAsAbs(m *Mark, g *GenVec) {
	d.marks = append(d.marks, m)
	m.genvec = g
	m.gtype = ABSGroup
}

// The method to regist a link from a mark in IMMgroup
// Since one or more IMM trans are enabled, EXP/GEN trans never fires
func (d *dfs) addLinkAsImm(src *Mark, dest *Mark, tr *ImmTrans) {
	d.links = append(d.links,
		Link{
			src:  src,
			dest: dest,
			tr:   tr,
			tt:   TransIMM,
		})
}

// The method to regist a link from a mark in GENgroup by an EXP trans
func (d *dfs) addLinkAsExp(src *Mark, dest *Mark, tr *ExpTrans) {
	d.links = append(d.links,
		Link{
			src:  src,
			dest: dest,
			tr:   tr,
			tt:   TransEXP,
		})
}

// The method to regist a link from a mark in GENgroup by an GEN trans
func (d *dfs) addLinkAsGen(src *Mark, dest *Mark, tr *GenTrans) {
	d.links = append(d.links,
		Link{
			src:  src,
			dest: dest,
			tr:   tr,
			tt:   TransGEN,
		})
}

// The method to create a unique GenVec that a given mark belongs to
func (d *dfs) createGenVec(net *Net, mark *Mark) *GenVec {
	vec := make([]TransStatus, len(net.genlist), len(net.genlist))
	for i, tr := range net.genlist {
		vec[i] = tr.IsEnabled(net, mark.toSlice())
	}
	gv := d.genVecGenerator.genGenVec(vec)
	return gv
}

// The method to create a unique mark by firing of tr trans
// If the number of tokens is less than zero or greater than max, err is not nil
func (d *dfs) createNextMarking(net *Net, mark *Mark, tr firingInterface) (*Mark, error) {
	// One buffer is reused for every firing; genMark copies it when the marking is new.
	dest, err := tr.doFiringInto(net, mark.toSlice(), d.firebuf)
	// A clamped destination is not the transition's real one. Record it so the caller
	// can be told the marking graph is not exact; see (*MarkingGraph).ClampEvents.
	d.clamps.record(err)
	return d.markGenerator.genMark(dest), err
}

// The method to regist all the next marks that are made by firing of all the enabled IMM trans
// to the stack (novisited)
func (d *dfs) visitImmMark(net *Net, mark *Mark) bool {
	enabled := false
	highestPriority := 0
	for _, tr := range net.immlist {
		if highestPriority > tr.priority {
			break
		}
		if tr.IsEnabled(net, mark.toSlice()) == ENABLE {
			enabled = true
			highestPriority = tr.priority
			dest, _ := d.createNextMarking(net, mark, tr)
			d.novisited.push(dest)
			d.addLinkAsImm(mark, dest, tr)
		}
	}
	return enabled
}

// The method to regist all the next marks that are made by firing of all the enabled EXP/GEN trans
// to the stack (novisited)
func (d *dfs) visitGenMark(net *Net, mark *Mark) bool {
	enabled := false
	highestPriority := 0
	for _, tr := range net.genlist {
		if highestPriority > tr.priority {
			break
		}
		if tr.IsEnabled(net, mark.toSlice()) == ENABLE {
			enabled = true
			highestPriority = tr.priority
			dest, _ := d.createNextMarking(net, mark, tr)
			d.novisited.push(dest)
			d.addLinkAsGen(mark, dest, tr)
		}
	}
	for _, tr := range net.explist {
		if highestPriority > tr.priority {
			break
		}
		if tr.IsEnabled(net, mark.toSlice()) == ENABLE {
			enabled = true
			highestPriority = tr.priority
			dest, _ := d.createNextMarking(net, mark, tr)
			d.novisited.push(dest)
			d.addLinkAsExp(mark, dest, tr)
		}
	}
	return enabled
}

// The method to do the depth first search
func (d *dfs) createMarking(net *Net) error {
	// A search assembled by hand rather than through create() has no counter; give it
	// the default limit rather than none, so no path silently runs unbounded.
	if d.counter == nil {
		d.counter = newStateCounter(DefaultSearchOptions())
	}
	for !d.novisited.isempty() {
		if err := d.counter.check(d.markGenerator.size()); err != nil {
			return err
		}
		mark := d.novisited.pop()
		if d.visited.exist(mark) {
			continue
		}

		gv := d.createGenVec(net, mark)
		if d.visitImmMark(net, mark) {
			d.addMarkAsImm(mark, gv)
		} else {
			if d.visitGenMark(net, mark) {
				d.addMarkAsGen(mark, gv)
			} else {
				d.addMarkAsAbs(mark, gv)
			}
		}
		d.visited.add(mark)
	}
	return nil
}
