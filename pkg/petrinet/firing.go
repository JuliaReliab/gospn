package petrinet

type TransStatus int

const (
	DISABLE TransStatus = iota + 1
	ENABLE
	PREEMPTION
)

func (t TransStatus) String() string {
	switch t {
	case DISABLE:
		return "D"
	case ENABLE:
		return "E"
	case PREEMPTION:
		return "P"
	default:
		return "Unknown"
	}
}

func (arc *InArc) getMulti(net *Net, mark []MarkInt) MarkInt {
	if arc.multifunc != nil {
		return arc.multifunc(mark)
	}
	return arc.multiplicity
}

func (arc *OutArc) getMulti(net *Net, mark []MarkInt) MarkInt {
	if arc.multifunc != nil {
		return arc.multifunc(mark)
	}
	return arc.multiplicity
}

type firingInterface interface {
	IsEnabled(net *Net, mark []MarkInt) TransStatus
	DoFiring(net *Net, mark []MarkInt) ([]MarkInt, error)
	doFiringInto(net *Net, mark []MarkInt, dst []MarkInt) ([]MarkInt, error)
}

func (tr *Trans) IsEnabled(net *Net, mark []MarkInt) TransStatus {
	if tr.guard != nil && tr.guard(mark) == false {
		return DISABLE
	}
	for _, arc := range tr.inarcs {
		multi := arc.getMulti(net, mark)
		place := arc.src
		if arc.inhibit == false {
			if mark[place.index] < multi {
				return DISABLE
			}
		} else {
			if mark[place.index] >= multi {
				return DISABLE
			}
		}
	}
	return ENABLE
}

func (tr *ImmTrans) IsEnabled(net *Net, mark []MarkInt) TransStatus {
	return tr.Trans.IsEnabled(net, mark)
}

func (tr *ExpTrans) IsEnabled(net *Net, mark []MarkInt) TransStatus {
	return tr.Trans.IsEnabled(net, mark)
}

func (tr *GenTrans) IsEnabled(net *Net, mark []MarkInt) TransStatus {
	maybePreemption := false
	if tr.guard != nil && tr.guard(mark) == false {
		maybePreemption = true
	}
	for _, arc := range tr.inarcs {
		multi := arc.getMulti(net, mark)
		place := arc.src
		if arc.inhibit == false {
			if mark[place.index] < multi {
				return DISABLE
			}
		} else {
			if mark[place.index] >= multi {
				maybePreemption = true
			}
		}
	}
	if maybePreemption {
		if tr.policy == GenTransPolicyPRD {
			return DISABLE
		} else {
			return PREEMPTION
		}
	} else {
		return ENABLE
	}
}

// DoFiring returns the marking reached by firing tr. If the destination would put a
// place outside [0, max] the value is clamped and a *ClampError naming every place that
// was clamped is returned alongside it. Callers that build a marking graph or run a
// simulation should feed that error to a clampRecorder rather than discard it: the
// clamped marking is not the transition's real destination, so anything derived from it
// (a generator matrix, a sample path) is no longer exact.
func (tr *Trans) DoFiring(net *Net, m []MarkInt) ([]MarkInt, error) {
	return tr.doFiringInto(net, m, nil)
}

// doFiringInto is DoFiring with the destination written into dst when it is long enough,
// instead of into a freshly allocated slice. The search reuses one buffer across every
// firing: the result is handed straight to MarkGenerator.genMark, which copies it if the
// marking turns out to be new and otherwise discards it. Passing nil allocates, which is
// what the exported DoFiring does.
//
// The returned slice is not necessarily dst: an update expression produces its own.
func (tr *Trans) doFiringInto(net *Net, m []MarkInt, dst []MarkInt) ([]MarkInt, error) {
	var clamped *ClampError
	var mark []MarkInt
	if len(dst) == len(m) {
		mark = dst
	} else {
		mark = make([]MarkInt, len(m))
	}
	copy(mark, m)
	for _, arc := range tr.inarcs {
		if arc.inhibit == false {
			multi := arc.getMulti(net, m)
			place := arc.src
			mark[place.index] -= multi
			if mark[place.index] < 0 {
				mark[place.index] = 0
				if clamped == nil {
					clamped = &ClampError{}
				}
				clamped.add(ClampEvent{Trans: tr.label, Place: place.label, Bound: ClampMin, Limit: 0})
			}
		}
	}
	for _, arc := range tr.outarcs {
		multi := arc.getMulti(net, m)
		place := arc.dest
		mark[place.index] += multi
		if mark[place.index] > place.max {
			mark[place.index] = place.max
			if clamped == nil {
				clamped = &ClampError{}
			}
			clamped.add(ClampEvent{Trans: tr.label, Place: place.label, Bound: ClampMax, Limit: place.max})
		}
	}
	// Keep the returned error a true nil when nothing was clamped: a typed nil pointer
	// stored in an error interface would compare non-nil at every call site.
	var err error
	if clamped != nil {
		err = clamped
	}
	if tr.update != nil {
		return tr.update(mark), err
	}
	return mark, err
}

func (tr *ImmTrans) DoFiring(net *Net, mark []MarkInt) ([]MarkInt, error) {
	return tr.Trans.DoFiring(net, mark)
}

func (tr *ImmTrans) doFiringInto(net *Net, mark []MarkInt, dst []MarkInt) ([]MarkInt, error) {
	return tr.Trans.doFiringInto(net, mark, dst)
}

func (tr *ExpTrans) DoFiring(net *Net, mark []MarkInt) ([]MarkInt, error) {
	return tr.Trans.DoFiring(net, mark)
}

func (tr *ExpTrans) doFiringInto(net *Net, mark []MarkInt, dst []MarkInt) ([]MarkInt, error) {
	return tr.Trans.doFiringInto(net, mark, dst)
}

func (tr *GenTrans) DoFiring(net *Net, mark []MarkInt) ([]MarkInt, error) {
	return tr.Trans.DoFiring(net, mark)
}

func (tr *GenTrans) doFiringInto(net *Net, mark []MarkInt, dst []MarkInt) ([]MarkInt, error) {
	return tr.Trans.doFiringInto(net, mark, dst)
}
