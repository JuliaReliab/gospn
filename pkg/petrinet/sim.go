package petrinet

import (
	"encoding/json"
	"fmt"
	// "log"
	"math"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/okamumu/gospn/pkg/mt"
)

func (tr *ImmTrans) getWeight(net *Net, m []MarkInt) float64 {
	if wfunc := tr.ratefunc; wfunc != nil {
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
	if rfunc := tr.ratefunc; rfunc != nil {
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

	// Parallel is how many replications run at once. Zero or negative selects
	// runtime.NumCPU(). Results do not depend on it: replication k always uses the
	// random stream derived from k, whichever worker happens to run it.
	Parallel int `json:"parallel"`
}

func ReadConfigFromJson(b []byte) (PNSimConfig, error) {
	var config PNSimConfig
	err := json.Unmarshal(b, &config)
	return config, err
}

type PNSimulation struct {
	PNSimConfig
	net    *Net
	clamps clampRecorder // places clamped while firing (see clamp.go)
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

// Time, Mark and TransLabel are how a caller outside this package reads a path: event
// is unexported and its fields with it, but the values RunSimulation returns can still
// be asked these questions.
func (e event) Time() float64 { return e.time }

// Mark is the marking after the firing, indexed by place, in Net.PlaceLabels order.
func (e event) Mark() []MarkInt { return e.mark }

// TransLabel is the transition that fired, empty for the initial marking.
func (e event) TransLabel() string {
	if e.tr == nil {
		return ""
	}
	return e.tr.label
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

// A simulation produces one event per firing, and the reward it contributes depends
// only on that event and the one before it. Collecting the whole path first, as
// RunSimulation does, means one marking is allocated and kept per firing -- on
// example/k8s.spn that is 95% of the objects the program allocates, and the resulting
// pressure on the runtime's allocator locks is what stops RunAll scaling past about
// three cores. A sink that folds each event as it arrives keeps nothing, so the markings
// can be handed a buffer to reuse.
type eventSink interface {
	observe(t float64, m []MarkInt, tr *Trans)

	// retains reports whether the sink keeps a reference to m after observe returns.
	// A sink that does forces a fresh marking per firing.
	retains() bool
}

// pathSink collects the events, which is what RunSimulation returns and `gospn test`
// prints. It retains every marking, so no buffer is reused underneath it.
type pathSink struct {
	events []event
}

func (p *pathSink) observe(t float64, m []MarkInt, tr *Trans) {
	p.events = append(p.events, event{time: t, mark: m, tr: tr})
}

func (p *pathSink) retains() bool { return true }

// rewardSink folds the rewards as the events arrive, which is the same computation
// calcReward performs over a stored path -- for each reward: the cumulative reward is
// the previous rate integrated over the elapsed interval, and the instantaneous reward
// accumulates at every firing.
type rewardSink struct {
	rfuncs   []func([]MarkInt) float64
	irwd     []float64
	crwd     []float64
	r        []float64
	prevtime float64
}

func newRewardSink(rfuncs []func([]MarkInt) float64) *rewardSink {
	return &rewardSink{
		rfuncs: rfuncs,
		irwd:   make([]float64, len(rfuncs)),
		crwd:   make([]float64, len(rfuncs)),
		r:      make([]float64, len(rfuncs)),
	}
}

func (s *rewardSink) observe(t float64, m []MarkInt, tr *Trans) {
	dt := t - s.prevtime
	s.prevtime = t
	for i, f := range s.rfuncs {
		s.crwd[i] += s.r[i] * dt
		s.r[i] = f(m)
		if tr != nil {
			s.irwd[i] += s.r[i]
		}
	}
}

func (s *rewardSink) retains() bool { return false }

func (s *rewardSink) reset() {
	s.prevtime = 0
	for i := range s.rfuncs {
		s.irwd[i], s.crwd[i], s.r[i] = 0, 0, 0
	}
}

// RunSimulation runs one replication. It records any clamping into the simulation's
// own recorder, so it is not safe to call concurrently on the same PNSimulation; the
// parallel path in RunAll uses runSimulation with a recorder per worker.
func (sim *PNSimulation) RunSimulation(init []MarkInt, rng RandomNumberGenerator) ([]event, float64, int32) {
	sink := &pathSink{events: make([]event, 0)}
	t, n := sim.runSimulation(init, rng, &sim.clamps, sink)
	return sink.events, t, n
}

func (sim *PNSimulation) runSimulation(init []MarkInt, rng RandomNumberGenerator, clamps *clampRecorder, sink eventSink) (float64, int32) {
	net := sim.net
	elapsedtime := 0.0
	var count int32 = 0
	m := init
	weights := make([]float64, len(net.immlist))
	genstates := make([]TransStatus, len(net.genlist))
	genremain := make([]float64, len(net.genlist))
	geninit := make([]float64, len(net.genlist))
	// Two buffers, used alternately: a firing reads the current marking while writing
	// the next one, so they cannot be the same slice. Nil when the sink retains the
	// markings, which makes doFiringInto allocate as it did before.
	var buf [2][]MarkInt
	var bufi int
	if !sink.retains() {
		buf[0] = make([]MarkInt, len(init))
		buf[1] = make([]MarkInt, len(init))
	}
	nextbuf := func() []MarkInt {
		if buf[0] == nil {
			return nil
		}
		bufi = 1 - bufi
		return buf[bufi]
	}

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
	sink.observe(0.0, m, nil)
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
					var err error
					m, err = net.immlist[i].doFiringInto(net, m, nextbuf())
					clamps.record(err)
					count++
					sink.observe(elapsedtime, m, net.immlist[i].getTrans())
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
					sink.observe(sim.EndingTime, m, nil)
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
				sink.observe(elapsedtime, m, nil)
				break
			}

			var err error
			m, err = firingtr.doFiringInto(net, m, nextbuf())
			clamps.record(err)
			count++
			sink.observe(elapsedtime, m, firingtr.getTrans())
		}
		if sim.NumOfFiring != 0 && count >= sim.NumOfFiring {
			break
		}
	}
	return elapsedtime, count
}

// ClampEvents reports the places that had to be clamped while firing over every run
// made with this PNSimulation, collapsed per (transition, place, bound). A non-empty
// result means some sample paths did not follow the transitions' real destinations.
// See (*MarkingGraph).ClampEvents for the usual cause.
func (sim *PNSimulation) ClampEvents() []ClampSummary {
	return sim.clamps.events()
}

func (sim *PNSimulation) calcReward(events []event, rfunc func([]MarkInt) float64) (float64, float64, float64) {
	irwd := 0.0
	crwd := 0.0
	prevtime := 0.0
	r := 0.0
	for _, e := range events {
		crwd += r * (e.time - prevtime)
		prevtime = e.time
		r = rfunc(e.mark)
		if e.tr != nil {
			irwd += r
		}
	}
	return irwd, crwd, r
}

// RunAll runs NumOfSimulation independent replications and returns their per-run
// instantaneous, cumulative and final rewards, elapsed times and firing counts.
//
// Replications are independent, so they run on Parallel workers at once. seed is the
// master seed: replication k uses its own MT19937-64 initialised from (seed, k) with
// init_by_array, the form the Mersenne Twister authors recommend for deriving several
// streams. Deriving them by adding k to the seed instead can leave nearby streams
// correlated, which init_by_array's scrambling is there to prevent.
//
// The stream is fixed by k alone, never by which worker picks the replication up, so
// the result is identical for any Parallel and reproducible from seed.
func (sim *PNSimulation) RunAll(init []MarkInt, seed int64) (map[string][]float64, map[string][]float64, map[string][]float64, []float64, []int32) {
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

	// The reward closures are looked up once: the map is read-only from here on, and
	// so is everything else the workers touch except their own clamp recorder.
	rfuncs := make([]func([]MarkInt) float64, 0, len(sim.Rewards))
	rnames := make([]string, 0, len(sim.Rewards))
	for _, str := range sim.Rewards {
		if f, ok := sim.net.rewardfunc[str]; ok {
			rfuncs = append(rfuncs, f)
			rnames = append(rnames, str)
		}
	}

	workers := sim.Parallel
	if workers <= 0 {
		workers = runtime.NumCPU()
	}
	if workers > sim.NumOfSimulation {
		workers = sim.NumOfSimulation
	}
	if workers < 1 {
		workers = 1
	}

	// Every result is written at index k of a slice allocated up front, so the workers
	// never write to the same memory and need no lock. Only the clamp recorders are
	// merged afterwards.
	clamps := make([]clampRecorder, workers)
	next := int64(-1)
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			rng := mt.NewMT64()
			// One sink per worker, reset between replications: the rewards are folded
			// as the events arrive, so no path is kept and the markings are reused.
			sink := newRewardSink(rfuncs)
			for {
				k := int(atomic.AddInt64(&next, 1))
				if k >= sim.NumOfSimulation {
					return
				}
				rng.InitByArray([]uint64{uint64(seed), uint64(k)})
				sink.reset()
				time, count := sim.runSimulation(init, rng, &clamps[w], sink)
				elapsedtime[k] = time
				nn[k] = count
				for i := range rfuncs {
					irwd[rnames[i]][k] = sink.irwd[i]
					crwd[rnames[i]][k] = sink.crwd[i]
					lastrwd[rnames[i]][k] = sink.r[i]
				}
			}
		}(w)
	}
	wg.Wait()
	for w := range clamps {
		sim.clamps.merge(&clamps[w])
	}

	return irwd, crwd, lastrwd, elapsedtime, nn
}
