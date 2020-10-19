package petrinet

import (
	// "log"
	"math"
)

type PNSim struct {
	rng           RandomNumberGenerator
	remainingTime []float64
}

func (tr *ImmTrans) getWeight(net *Net, m []MarkInt) float64 {
	if wfunc, ok := net.ratefunc[tr.Trans]; ok {
		return wfunc(m)
	}
	return tr.weight
}

type simTransInterface interface {
	firingInterface
	nextTime(*Net, []MarkInt, RandomNumberGenerator) float64
}

func (tr *ExpTrans) nextTime(net *Net, m []MarkInt, rng RandomNumberGenerator) float64 {
	if rfunc, ok := net.ratefunc[tr.Trans]; ok {
		return -1 / rfunc(m) * math.Log(rng.Float64())
	}
	return -1 / tr.rate * math.Log(rng.Float64())
}

func (tr *GenTrans) nextTime(net *Net, m []MarkInt, rng RandomNumberGenerator) float64 {
	return tr.dist.Float64(rng)
}

type PNSimulation struct {
	net        *Net
	endingtime float64
}

type event struct {
	time float64
	mark []MarkInt
}

func (sim *PNSimulation) runSimulation(init []MarkInt, rng RandomNumberGenerator) []event {
	net := sim.net
	elapsedtime := 0.0
	m := init
	weights := make([]float64, len(net.immlist))
	genstates := make([]TransStatus, len(net.genlist))
	genremain := make([]float64, len(net.genlist))
	geninit := make([]float64, len(net.genlist))
	result := make([]event, 0)
	for i, tr := range net.genlist {
		genstates[i] = tr.IsEnabled(net, m)
		switch genstates[i] {
		case DISABLE:
			genremain[i] = 0.0
			geninit[i] = 0.0
		case ENABLE:
			genremain[i] = tr.nextTime(net, m, rng)
			geninit[i] = genremain[i]
		case PREEMPTION:
			switch tr.policy {
			case GenTransPolicyPRD:
				genremain[i] = 0.0
				geninit[i] = 0.0
			case GenTransPolicyPRS:
				genremain[i] = tr.nextTime(net, m, rng)
				geninit[i] = genremain[i]
			case GenTransPolicyPRI:
				genremain[i] = 0.0
				geninit[i] = tr.nextTime(net, m, rng)
			}
		}
	}
	for {
		result = append(result, event{
			time: elapsedtime,
			mark: m,
		})

		for i, tr := range net.genlist {
			switch tr.IsEnabled(net, m) {
			case DISABLE:
				genstates[i] = DISABLE
				genremain[i] = 0.0
				geninit[i] = 0.0
			case ENABLE:
				switch genstates[i] { // previous state
				case DISABLE:
					genremain[i] = tr.nextTime(net, m, rng)
					geninit[i] = genremain[i]
				case ENABLE:
					// pass
				case PREEMPTION:
					switch tr.policy {
					case GenTransPolicyPRD:
						genremain[i] = tr.nextTime(net, m, rng)
					case GenTransPolicyPRS:
						// pass
					case GenTransPolicyPRI:
						genremain[i] = geninit[i]
					}
				}
				genstates[i] = ENABLE
			case PREEMPTION:
				genstates[i] = PREEMPTION
			}
		}

		// IMM trans
		weightsum := 0.0
		for i, tr := range net.immlist {
			if tr.IsEnabled(net, m) == ENABLE {
				weights[i] = tr.getWeight(net, m)
				weightsum += weights[i]
			} else {
				weights[i] = 0
			}
		}
		if weightsum != 0 {
			u := weightsum * rng.Float64()
			s := 0.0
			for i, w := range weights {
				s += w
				if s > u {
					if next, err := net.immlist[i].DoFiring(net, m); err == nil {
						m = next
						break
					}
				}
			}
			continue // next simulation loop
		}

		mintime := math.MaxFloat64
		var firingtr simTransInterface
		// GEN trans
		for i, tr := range net.genlist {
			if genstates[i] == ENABLE && genremain[i] < mintime {
				mintime = genremain[i]
				firingtr = tr
			}
		}
		// EXP trans
		for _, tr := range net.explist {
			if tr.IsEnabled(net, m) == ENABLE {
				if t := tr.nextTime(net, m, rng); t < mintime {
					mintime = t
					firingtr = tr
				}
			}
		}

		if firingtr == nil { // absorbing state
			result = append(result, event{
				time: sim.endingtime,
				mark: m,
			})
			break
		}

		for i, _ := range net.genlist {
			if genstates[i] == ENABLE {
				genremain[i] -= mintime
			}
		}
		elapsedtime += mintime

		if elapsedtime > sim.endingtime {
			result = append(result, event{
				time: sim.endingtime,
				mark: m,
			})
			break
		}

		if next, err := firingtr.DoFiring(net, m); err == nil {
			m = next
		}
	}
	return result
}
