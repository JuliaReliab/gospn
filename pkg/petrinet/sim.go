package petrinet

import (
	// "log"
	"encoding/json"
	"fmt"
	"math"
	"strings"
)

func (tr *ImmTrans) getWeight(net *Net, m []MarkInt) float64 {
	if wfunc, ok := net.ratefunc[tr.Trans]; ok {
		return wfunc(m)
	}
	return tr.weight
}

type simTransInterface interface {
	firingInterface
	getTrans() *Trans
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

type PNSimConfig struct {
	EndingTime      float64  `json:"time"`
	NumOfFiring     int32    `json:"firings"`
	NumOfSimulation int      `json:"simulations"`
	Rewards         []string `json:"rewards"`
}

func ReadConfigFromJson(b []byte) (PNSimConfig, error) {
	var config PNSimConfig
	err := json.Unmarshal(b, &config)
	return config, err
}

type PNSimulation struct {
	PNSimConfig
	net *Net
}

func NewPNSimulation(net *Net, config PNSimConfig) *PNSimulation {
	return &PNSimulation{
		PNSimConfig: config,
		net:         net,
	}
}

type event struct {
	time float64
	mark []MarkInt
	tr   *Trans
}

func (e event) String(net *Net) string {
	str := make([]string, 0)
	for i, n := range e.mark {
		if n > 0 {
			str = append(str, fmt.Sprintf("%s:%d", net.placelist[i].label, n))
		}
	}
	if e.tr != nil {
		return fmt.Sprintf("%.4f {%s} %s", e.time, strings.Join(str, ","), e.tr.label)
	} else {
		return fmt.Sprintf("%.4f {%s} -", e.time, strings.Join(str, ","))
	}
}

func (sim *PNSimulation) RunSimulation(init []MarkInt, rng RandomNumberGenerator) ([]event, float64, int32) {
	net := sim.net
	elapsedtime := 0.0
	var count int32 = 0
	m := init
	weights := make([]float64, len(net.immlist))
	genstates := make([]TransStatus, len(net.genlist))
	genremain := make([]float64, len(net.genlist))
	geninit := make([]float64, len(net.genlist))
	events := make([]event, 0)

	// initialize for sim
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
	events = append(events, event{
		time: 0.0,
		mark: m,
	})
	for {
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
					m, _ = net.immlist[i].DoFiring(net, m)
					count++
					events = append(events, event{
						time: elapsedtime,
						mark: m,
						tr:   net.immlist[i].getTrans(),
					})
					break
				}
			}
		} else {
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
				if sim.EndingTime > 0.0 {
					events = append(events, event{
						time: sim.EndingTime,
						mark: m,
					})
				}
				break
			}

			for i, _ := range net.genlist {
				if genstates[i] == ENABLE {
					genremain[i] -= mintime
				}
			}
			elapsedtime += mintime

			if sim.EndingTime != 0.0 && elapsedtime > sim.EndingTime {
				elapsedtime = sim.EndingTime
				events = append(events, event{
					time: elapsedtime,
					mark: m,
				})
				break
			}

			m, _ = firingtr.DoFiring(net, m)
			count++
			events = append(events, event{
				time: elapsedtime,
				mark: m,
				tr:   firingtr.getTrans(),
			})
		}
		if sim.NumOfFiring != 0 && count >= sim.NumOfFiring {
			break
		}
	}
	return events, elapsedtime, count
}

func (sim *PNSimulation) calcReward(events []event, rfunc func([]MarkInt) float64) (float64, float64, float64) {
	irwd := 0.0
	crwd := 0.0
	lastrwd := 0.0
	prevtime := 0.0
	for _, e := range events {
		r := rfunc(e.mark)
		if e.tr != nil {
			irwd += r
		}
		crwd += r * (e.time - prevtime)
		if e.time == sim.EndingTime {
			lastrwd = r
		}
		prevtime = e.time
	}
	return irwd, crwd, lastrwd
}

func (sim *PNSimulation) RunAll(init []MarkInt, rng RandomNumberGenerator) (map[string][]float64, map[string][]float64, map[string][]float64, []float64, []int32) {
	irwd := make(map[string][]float64)
	crwd := make(map[string][]float64)
	lastrwd := make(map[string][]float64)
	for _, str := range sim.Rewards {
		irwd[str] = make([]float64, sim.NumOfSimulation)
		crwd[str] = make([]float64, sim.NumOfSimulation)
		lastrwd[str] = make([]float64, sim.NumOfSimulation)
	}
	nn := make([]int32, sim.NumOfSimulation)
	elapsedtime := make([]float64, sim.NumOfSimulation)

	for k := 0; k < sim.NumOfSimulation; k++ {
		events, time, count := sim.RunSimulation(init, rng)
		elapsedtime[k] = time
		nn[k] = count
		for _, str := range sim.Rewards {
			if rfunc, ok := sim.net.rewardfunc[str]; ok {
				irwd[str][k], crwd[str][k], lastrwd[str][k] = sim.calcReward(events, rfunc)
			}
		}
	}

	return irwd, crwd, lastrwd, elapsedtime, nn
}
