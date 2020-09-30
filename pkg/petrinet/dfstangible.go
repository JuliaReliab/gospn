package petrinet

type exitMarkStatus int

const (
	initialized exitMarkStatus = iota + 1
	vanishable
	novanishable
)

type exitMark struct {
	self   *Mark
	gv     *GenVec
	next   *Mark
	status exitMarkStatus
}

func newExitMark(mark *Mark, gv *GenVec, status exitMarkStatus) *exitMark {
	return &exitMark{
		self:   mark,
		gv:     gv,
		next:   mark,
		status: status,
	}
}

func (m *exitMark) setNovanishable() {
	m.next = m.self
	m.status = novanishable
}

func (m *exitMark) union(child *exitMark) {
	switch m.status {
	case initialized:
		switch child.status {
		case vanishable, novanishable:
			if m.gv == child.gv {
				m.next = child.next
				m.status = vanishable
			} else {
				m.setNovanishable()
			}
		}
	case vanishable:
		switch child.status {
		case vanishable, novanishable:
			if m.next != child.next {
				m.setNovanishable()
			}
		}
	}
}

// The structure of DFS (depth first search) to make a marking graph.
//   markGenerator: A generator to create a unique object of mark
//   genVecGenerator: A gnerator to create a unique object of GenVec
//   novisited: A stack for the mark that is visited next
//   visited: A set for the marks that are already visited
// The following data are required to create an instance of marking graph
//   marks: A slice for all the visited marks
//   markToGenvec: A map to indicate a GenVec that a given mark belongs to
//   markToGroupType: A map to indicate a GroupType that a given mark belongs to
//   links: A scile for links between all marks
type dfstangible struct {
	markGenerator   MarkGeneratorInterface
	genVecGenerator GenVecGeneratorInterface
	novisited       markStack
	visited         markSet
	novisitedIMM    markStack
	visitedIMM      markSet
	exitMarks       map[*Mark]*exitMark
	markPath        markStack
	enabledIMMlist  []*ImmTrans
	marks           []*Mark
	markToGenvec    map[*Mark]*GenVec   // map from Mark to GenVec
	markToGroupType map[*Mark]GroupType // map from Mark to Gtype
	links           []Link              // links
}

// The method to create a marking graph.
// This is an interface for markinggraphGenerator.
func (d *dfstangible) create(net *Net, imark []MarkInt) (*Mark, []*Mark, map[*Mark]*GenVec, map[*Mark]GroupType, []Link) {
	d.markGenerator = NewMarkGenerator(len(net.placelist))
	d.genVecGenerator = NewGenVecGenerator(len(net.genlist))
	d.markToGenvec = make(map[*Mark]*GenVec)
	d.visited = NewMarkSet()
	d.novisited = NewMarkStack()
	d.novisitedIMM = NewMarkStack()
	d.visitedIMM = NewMarkSet()
	d.exitMarks = make(map[*Mark]*exitMark)
	d.markPath = NewMarkStack()
	d.enabledIMMlist = make([]*ImmTrans, 0, len(net.immlist))
	d.marks = make([]*Mark, 0)
	d.markToGroupType = make(map[*Mark]GroupType)
	d.links = make([]Link, 0)

	m0 := d.markGenerator.genMark(imark)
	d.novisited.push(m0)
	d.createMarking(net)

	// post processing
	newmarks := make([]*Mark, 0, len(d.marks))
	for _, m := range d.marks {
		if d.exitMarks[m].status != vanishable {
			newmarks = append(newmarks, m)
		}
	}
	newlinks := make([]Link, 0, len(d.links))
	for _, l := range d.links {
		src := l.src
		dest := l.dest
		if d.exitMarks[src].status != vanishable {
			em := d.exitMarks[dest]
			for em.status == vanishable {
				dest = em.next
				em = d.exitMarks[dest]
			}
			newlinks = append(newlinks, Link{
				src:  src,
				dest: dest,
				tr:   l.tr,
				tt:   l.tt,
			})
		}
	}
	em := d.exitMarks[m0]
	for em.status == vanishable {
		m0 = em.next
		em = d.exitMarks[m0]
	}
	return m0, newmarks, d.markToGenvec, d.markToGroupType, newlinks
}

// The method to regist a mark as a member of IMMgroup (There is one or more enabled IMM trans)
func (d *dfstangible) addMarkAsImm(m *Mark, g *GenVec) {
	d.marks = append(d.marks, m)
	d.markToGenvec[m] = g
	d.markToGroupType[m] = IMMGroup
	d.exitMarks[m] = newExitMark(m, g, initialized)
}

// The method to regist a mark as a member of GENgroup (There is no enabled IMM trans)
func (d *dfstangible) addMarkAsGen(m *Mark, g *GenVec) {
	d.marks = append(d.marks, m)
	d.markToGenvec[m] = g
	d.markToGroupType[m] = GENGroup
	d.exitMarks[m] = newExitMark(m, g, novanishable)
}

// The method to regist a mark as a member of ABSgroup (There is no enabled trans)
func (d *dfstangible) addMarkAsAbs(m *Mark, g *GenVec) {
	d.marks = append(d.marks, m)
	d.markToGenvec[m] = g
	d.markToGroupType[m] = ABSGroup
	d.exitMarks[m] = newExitMark(m, g, novanishable)
}

// The method to regist a link from a mark in IMMgroup
// Since one or more IMM trans are enabled, EXP/GEN trans never fires
func (d *dfstangible) addLinkAsImm(src *Mark, dest *Mark, tr *ImmTrans) {
	d.links = append(d.links,
		Link{
			src:  src,
			dest: dest,
			tr:   tr,
			tt:   TransIMM,
		})
}

// The method to regist a link from a mark in GENgroup by an EXP trans
func (d *dfstangible) addLinkAsExp(src *Mark, dest *Mark, tr *ExpTrans) {
	d.links = append(d.links,
		Link{
			src:  src,
			dest: dest,
			tr:   tr,
			tt:   TransEXP,
		})
}

// The method to regist a link from a mark in GENgroup by an GEN trans
func (d *dfstangible) addLinkAsGen(src *Mark, dest *Mark, tr *GenTrans) {
	d.links = append(d.links,
		Link{
			src:  src,
			dest: dest,
			tr:   tr,
			tt:   TransGEN,
		})
}

// The method to create a unique GenVec that a given mark belongs to
func (d *dfstangible) createGenVec(net *Net, mark *Mark) *GenVec {
	vec := make([]TransStatus, len(net.genlist), len(net.genlist))
	for i, tr := range net.genlist {
		vec[i] = tr.IsEnabled(net, mark.toSlice())
	}
	gv := d.genVecGenerator.genGenVec(vec)
	return gv
}

// The method to create a unique mark by firing of tr trans
// If the number of tokens is less than zero or greater than max, err is not nil
func (d *dfstangible) createNextMarking(net *Net, mark *Mark, tr firingInterface) (*Mark, error) {
	dest, err := tr.DoFiring(net, mark.toSlice())
	return d.markGenerator.genMark(dest), err
}

// The method to return a slice of IMM trans that are enabled
func (d *dfstangible) enabledIMM(net *Net, mark *Mark) {
	d.enabledIMMlist = d.enabledIMMlist[:0]
	highestPriority := 0
	for _, tr := range net.immlist {
		if highestPriority > tr.priority {
			break
		}
		if tr.IsEnabled(net, mark.toSlice()) == ENABLE {
			highestPriority = tr.priority
			d.enabledIMMlist = append(d.enabledIMMlist, tr)
		}
	}
}

// The method to regist all the next marks that are made by firing of all the enabled IMM trans
// to the stack (novisited)
func (d *dfstangible) visitImmMark(net *Net, mark *Mark) {
	d.markPath.push(mark)
	d.novisitedIMM.push(nil)
	for _, tr := range d.enabledIMMlist {
		dest, _ := d.createNextMarking(net, mark, tr)
		d.novisitedIMM.push(dest)
		d.addLinkAsImm(mark, dest, tr)
	}
	d.visitedIMM.add(mark)
}

// The method to regist all the next marks that are made by firing of all the enabled EXP/GEN trans
// to the stack (novisited)
func (d *dfstangible) visitGenMark(net *Net, mark *Mark) bool {
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
	d.visited.add(mark)
	return enabled
}

func (d *dfstangible) vanishing(net *Net) {
	for !d.novisitedIMM.isempty() {
		mark := d.novisitedIMM.pop()
		if mark == nil {
			child := d.markPath.pop()
			if !d.markPath.isempty() {
				parent := d.markPath.peek()
				d.exitMarks[parent].union(d.exitMarks[child])
			}
			d.visited.add(child)
			continue
		}
		if d.visited.exist(mark) {
			r := d.markPath.peek()
			d.exitMarks[r].union(d.exitMarks[mark])
			continue
		}
		if d.visitedIMM.exist(mark) {
			d.exitMarks[mark].setNovanishable()
			r := d.markPath.peek()
			d.exitMarks[r].union(d.exitMarks[mark])
			continue
		}

		gv := d.createGenVec(net, mark)
		d.enabledIMM(net, mark)
		if len(d.enabledIMMlist) > 0 {
			d.visitImmMark(net, mark)
			d.addMarkAsImm(mark, gv)
		} else {
			if d.visitGenMark(net, mark) {
				d.addMarkAsGen(mark, gv)
			} else {
				d.addMarkAsAbs(mark, gv)
			}
			r := d.markPath.peek()
			d.exitMarks[r].union(d.exitMarks[mark])
		}
	}
}

// The method to do the depth first search
func (d *dfstangible) createMarking(net *Net) {
	for !d.novisited.isempty() {
		mark := d.novisited.pop()
		if d.visited.exist(mark) {
			continue
		}

		gv := d.createGenVec(net, mark)
		d.enabledIMM(net, mark)
		if len(d.enabledIMMlist) > 0 {
			d.visitImmMark(net, mark)
			d.addMarkAsImm(mark, gv)
			d.vanishing(net)
		} else {
			if d.visitGenMark(net, mark) {
				d.addMarkAsGen(mark, gv)
			} else {
				d.addMarkAsAbs(mark, gv)
			}
		}
	}
}
