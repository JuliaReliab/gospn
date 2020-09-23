package petrinet

type TransStatus int

const (
    DISABLE = iota
    ENABLE
    PREEMPTION
)

func getInMulti(net *Net, arc *InArc, mark *Mark) markInt {
    multifunc, ok := net.infunc[arc]
    if ok {
        return multifunc(mark)
    } else {
        return arc.multiplicity
    }
}

func getOutMulti(net *Net, arc *OutArc, mark *Mark) markInt {
    multifunc, ok := net.outfunc[arc]
    if ok {
        return multifunc(mark)
    } else {
        return arc.multiplicity
    }
}

func IsEnabled(net *Net, tr *Trans, mark *Mark) TransStatus {
    guard, ok := net.guards[tr]
    if ok && guard(mark) == false {
        return DISABLE
    }
    for _,arc := range tr.inarcs {
        multi := getInMulti(net, arc, mark)
        place := arc.src
        if arc.inhibit == false {
            if mark.Get(place.index) < multi {
                return DISABLE
            }
        } else {
            if mark.Get(place.index) >= multi {
                return DISABLE
            }
        }
    }
    return ENABLE
}

func DoFiring(net *Net, tr *Trans, mark *Mark) *Mark {
    m := toslice(mark)
    for _,arc := range tr.inarcs {
        if arc.inhibit == false {
            multi := getInMulti(net, arc, mark)
            place := arc.src
            m[place.index] -= multi
            if m[place.index] < 0 {
                return nil
            }
        }
    }    
    for _,arc := range tr.outarcs {
        multi := getOutMulti(net, arc, mark)
        place := arc.dest
        m[place.index] += multi
        if m[place.index] > place.max {
            return nil
        }
    }
    newmark := NewMark(m)
    update, ok := net.updates[tr]
    if ok {
        return update(newmark)
    } else {
        return newmark
    }
}
