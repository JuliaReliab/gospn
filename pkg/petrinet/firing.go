package petrinet

import (
	"fmt"
	"os"
	"sync"
)

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
	multifunc, ok := net.infunc[arc]
	if ok {
		return multifunc(mark)
	} else {
		return arc.multiplicity
	}
}

func (arc *OutArc) getMulti(net *Net, mark []MarkInt) MarkInt {
	multifunc, ok := net.outfunc[arc]
	if ok {
		return multifunc(mark)
	} else {
		return arc.multiplicity
	}
}

type firingInterface interface {
	IsEnabled(net *Net, mark []MarkInt) TransStatus
	DoFiring(net *Net, mark []MarkInt) ([]MarkInt, error)
}

func (tr *Trans) IsEnabled(net *Net, mark []MarkInt) TransStatus {
	guard, ok := net.guards[tr]
	if ok && guard(mark) == false {
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
	guard, ok := net.guards[tr.Trans]
	if ok && guard(mark) == false {
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

// Firing errors (a token count below zero, or above the place capacity) are returned
// by DoFiring, but every caller in the marking-graph search discards them. The marking
// is then silently clamped, which drops the transition's real destination and leaves a
// generator matrix whose rows do not sum to zero -- with no indication that anything
// went wrong. Warn on stderr instead, once per (transition, place, kind).
//
// This bites easily because the default place capacity is 255 (see PNcompile.go), so a
// place that legitimately holds more tokens but has no explicit `max` is truncated.
var (
	firingWarnMu   sync.Mutex
	firingWarnSeen = map[string]bool{}
)

func warnFiring(kind, trans, place string, limit MarkInt) {
	if os.Getenv("GOSPN_NO_WARN") != "" {
		return
	}
	key := kind + "\x00" + trans + "\x00" + place
	firingWarnMu.Lock()
	defer firingWarnMu.Unlock()
	if firingWarnSeen[key] {
		return
	}
	firingWarnSeen[key] = true
	fmt.Fprintf(os.Stderr,
		"gospn: warning: transition %s clamped place %s at %s %d; "+
			"the marking graph is no longer exact (set an explicit `max` on the place)\n",
		trans, place, kind, limit)
}

func (tr *Trans) DoFiring(net *Net, m []MarkInt) ([]MarkInt, error) {
	var err error
	mark := make([]MarkInt, len(m))
	for i, x := range m {
		mark[i] = x
	}
	for _, arc := range tr.inarcs {
		if arc.inhibit == false {
			multi := arc.getMulti(net, m)
			place := arc.src
			mark[place.index] -= multi
			if mark[place.index] < 0 {
				mark[place.index] = 0
				err = fmt.Errorf("The number of tokens is less than zero: tr %s, place %s", tr.label, place.label)
				warnFiring("min", tr.label, place.label, 0)
			}
		}
	}
	for _, arc := range tr.outarcs {
		multi := arc.getMulti(net, m)
		place := arc.dest
		mark[place.index] += multi
		if mark[place.index] > place.max {
			mark[place.index] = place.max
			err = fmt.Errorf("The number of tokens is greater than max: tr %s, place %s", tr.label, place.label)
			warnFiring("max", tr.label, place.label, place.max)
		}
	}
	update, ok := net.updates[tr]
	if ok {
		return update(mark), err
	} else {
		return mark, err
	}
}

func (tr *ImmTrans) DoFiring(net *Net, mark []MarkInt) ([]MarkInt, error) {
	return tr.Trans.DoFiring(net, mark)
}

func (tr *ExpTrans) DoFiring(net *Net, mark []MarkInt) ([]MarkInt, error) {
	return tr.Trans.DoFiring(net, mark)
}

func (tr *GenTrans) DoFiring(net *Net, mark []MarkInt) ([]MarkInt, error) {
	return tr.Trans.DoFiring(net, mark)
}
