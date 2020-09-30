package petrinet

import (
	"fmt"
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
