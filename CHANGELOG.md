# gospn 0.16.0

- run the simulation's replications in parallel. They are independent, so `RunAll` now
  spreads them over one worker per CPU by default. On `example/k8s.spn` (50 replications
  to time 10000) that is **2.9x** on this 10-core machine: 424 ms to 145 ms

  It does not scale linearly -- 82% efficiency at 2 workers, 56% at 4, and flat at about
  3x from 8 workers on -- and the reason took measuring rather than guessing. It is not
  the container (a pure-CPU control in the same image scales 9.4x on 10 workers), not the
  garbage collector (`GOGC=off` moves it from 3.25x to 3.33x), not uneven replications
  (their firing counts run 6,065 to 6,857, so perfect balance is available), and not
  memory bandwidth (the run allocates at 1.4 GB/s at its peak, orders of magnitude below
  what the machine can do).

  It is contention on the Go runtime's own allocator locks: at 10 workers the mutex
  profile is 100% `runtime._LostContendedRuntimeLock`. The allocation *rate* is what
  drives it -- roughly one 424-byte marking per firing, 318,000 of them per run, every
  one retained in the replication's event path so none can be reused. Reusing the
  marking buffers, or computing the rewards as events are produced instead of storing
  the path, would cut that rate; either is a separate change

- `sim` takes `-parallel`, and the configuration JSON takes `parallel`; the flag wins if
  both are given. Zero means one worker per CPU

- **the same `-s` now produces different numbers than it did up to 0.15.1.** Replications
  used to share one sequential random stream, which cannot be parallelised; each now has
  its own, derived from its index with `init_by_array` -- the form the Mersenne Twister
  authors recommend for deriving several streams, and already implemented in `pkg/mt`.
  Deriving streams by adding the index to the seed can leave nearby ones correlated,
  which `init_by_array`'s scrambling exists to prevent

  The results are statistically the same: over 20,000 replications of
  `spnp_example4.spn` the mean cumulative `reliab` moved from 1838.4 to 1847.0 with a
  standard error of 10.2 on the difference (z = 0.84)

- **API**: `RunAll` takes the master seed rather than a `RandomNumberGenerator`, since it
  now makes one stream per replication. `RunSimulation` is unchanged and still runs a
  single replication with a caller-supplied generator

- results do not depend on the worker count. Replication `k` takes the stream derived
  from `k` and writes to slot `k`, never sharing memory with another worker;
  `TestSimIsIndependentOfWorkerCount` checks the fingerprint is identical for 1, 2, 3,
  4, 8 and 16 workers, and `TestClampCountsSurviveTheMerge` checks the per-worker clamp
  recorders merge to the same counts. The test suite runs under `-race`

# gospn 0.15.0

- stop the reachability search at a state limit instead of running until the operating
  system kills the process. `gospn mark` on a net whose state space cannot be enumerated
  printed `Create marking...` and nothing else, then died with exit 137 and no output at
  all -- no state count, no indication of how far it had got

  `mark` takes `-maxstates`, default 1,000,000; `-maxstates 0` removes the limit. The
  limit is in states rather than in bytes so that the same input behaves the same way on
  every machine; the error reports the memory actually in use, which is what a limit has
  to be chosen against (about 1,218 bytes per state on `example/k8s.spn`, 483 on
  `spnp_example6.spn`)

- report progress while searching. A long run was completely silent between
  `Create marking...` and its result

- **API**: the `CreateMarkingGraph*` functions return `(*MarkingGraph, error)`. A search
  that stopped at its limit returns no graph on purpose: transition matrices taken from a
  truncated marking graph look valid and are not. `CreateMarkingGraphWithDFSOpts` and
  `...TangibleOpts` take a `SearchOptions` to set the limit and a progress callback

- refuse a stray positional argument instead of ignoring it. The subcommands take the net
  with `-i`, so `gospn mark -o out.mat net.spn` dropped the filename, read an empty
  standard input, and reported a one-state marking graph -- an answer that looks like a
  successful analysis of the wrong thing

- add `example/k8s.spn`, a Kubernetes GSPN (53 places, 64 transitions, places holding up
  to 1000 tokens) that is far too large to enumerate and is analysed by simulation. It is
  the net the state limit was built against, and `BenchmarkSimK8s` measures the
  simulation path on it

# gospn 0.14.0

- compile guards, rates, multiplicities and rewards into typed Go closures instead of
  walking the AST at every evaluation. **The gain is uneven and depends on the net.**
  A net whose transitions share a guard built out of variables gains a lot;
  `spnp_example5.spn`, where every transition carries `enall`, is about **6x** faster.
  Nets with constant rates and simple guards gain little: 1.0x to 1.5x, and
  `spnp_example1.spn` and `spnp_example2.spn` have nothing to compile and are unchanged.
  Allocations drop by 50-97% wherever anything compiles
- **simulation gains more than the marking graph does**: about 3.7x on `raid6.spn`,
  9.2x on `spnp_example4.spn` and 18.2x on `spnp_example5.spn`, with allocations down
  84-99.6%. A reachability search visits each marking once; `RunAll` re-evaluates the
  same guards, rates and rewards at every event of every replication, and the rewards
  in those nets are all `ifelse(...)`, which was fully interpreted before
- what made the interpreter expensive was not the tree walk alone: every intermediate
  value was boxed into an `*ASTValue`, every variable was resolved through the
  environment by string, and every `#place` called `Net.GetPlace` by string -- none of
  which depend on the marking, so all of it now happens once
- an expression outside the compiled subset falls back to the interpreter, which is
  unchanged. `TestCompiledAgreesWithInterpreter` runs both against each other over
  every bundled net, so the two implementations cannot drift apart unnoticed
- `makeNet` creates every place before any transition. A transition whose guard names
  a place declared later in the file -- `raid6.spn` does -- could not have that guard
  compiled. Marking-graph output is unaffected
- report an undefined variable by name. It evaluates to its own name as a string, and
  the resulting failure said only `the value is neither int32 nor float64
  parser.ASTString`, naming neither the variable nor the expression

  That message is how a **pre-existing** defect was found: `example/spnp_example6.spn`
  had the definition of `ret_val` commented out while a reward still referred to it, so
  `gospn mark` panicked on that net when it wrote the reward vectors. No test called
  `RewardVector`, so nothing caught it. The definition is restored, and
  `TestCompiledAgreesWithInterpreter` now takes the reward vectors as well
- every expression in every bundled net compiles; nothing falls back
- add `BenchmarkSim` and `TestSimGolden`. `RunAll` had neither a benchmark nor a
  regression test, so nothing would have noticed it changing; the golden pins the
  reward vectors, elapsed times and firing counts of a seeded run, and matches what
  0.13.1 produced

  Update expressions (`#place = ...`) are still interpreted. They run only on a
  firing, which is not hot for a marking graph, but simulation is nothing but
  firings, so a net that uses them has more to gain here than this release delivers

# gospn 0.13.1

- commit `go.sum`. It was deleted and added to `.gitignore` in 0.11.0, when the antlr
  dependency was moved back to the 2018 revision, so since then nothing in the repository
  recorded which module contents the released binaries were built against and every build
  re-derived them. Committing it also lets `actions/setup-go` cache the module downloads,
  which it could not do before -- the release workflow logged "Dependencies file is not
  found ... Supported file pattern: go.sum" on every run
- drop `GOFLAGS=-mod=mod` from both workflows and the Dockerfile. It was there to let the
  build write the missing `go.sum`; leaving it would let a build silently rewrite the
  committed one, which is most of what committing it is for. The default `-mod=readonly`
  builds and tests every package without touching `go.mod` or `go.sum`
- the release workflow refuses a tag that is not on `master`. A tag is not tied to a
  branch and the workflow fires as soon as one is pushed, so tagging a feature branch
  published a release before the branch was merged -- and a squash or rebase merge then
  rewrote that commit, leaving the binaries corresponding to nothing in the history

# gospn 0.13.0

- performance: the marking graph is built with fewer allocations. On the largest bundled
  example (`spnp_example6.spn`) allocations drop about 25% and construction about 8%
- `clampRecorder.record` no longer allocated on every firing. It is called once per firing
  and reached `errors.As`, which boxes its target into an `interface{}`; the common case
  where nothing was clamped now returns before that
- the compiled guard/rate/multiplicity closures no longer build an error on their normal
  path. `createRateFunc` asked `GetInt` first and fell through to `GetFloat` for every
  float-valued rate, and the discarded mismatch was reported with `fmt.Errorf`, so a rate
  expression allocated a formatted string per edge
- markings are interned on their raw bytes instead of a decimal rendering, and the map
  lookup is written in the form the compiler can keep allocation-free
- the destination of a firing is written into a buffer the search reuses, rather than a
  fresh slice per firing; `MarkGenerator` copies it only when the marking is new
- guard/rate/update/multiplicity closures live on `Trans`, `InArc` and `OutArc` instead of
  in maps keyed by their pointer, removing a map lookup per arc per enabling check
- the benchmarks in `test/` never actually ran: the `b.N` loop was commented out in all
  eight, and `BenchmarkGoSPNP5` read `".../example/..."` so it silently measured nothing.
  They are now one table-driven `BenchmarkGoSPN` that fails on a read error
- add `TestMarkingGraphGolden`, which pins the state labels and transition matrices of
  every bundled example, so changes of this kind can be shown not to alter the output

# gospn 0.12.0

- report the places that a firing had to clamp, instead of discarding the error. A clamped
  marking is not the transition's real destination, so the marking graph -- and any
  generator matrix taken from it -- is not exact, and nothing used to say so. Easy to hit
  because the default place capacity is 255.
- `DoFiring` returns a typed `*ClampError` naming every place that firing clamped, instead
  of a string error that kept only the last one
- add `(*MarkingGraph).ClampEvents`, `(*PNSimulation).ClampEvents` and `FormatClampEvents`;
  `gospn mark` and `gospn sim` print the report on stderr
- publish releases from a GitHub Actions workflow when a version tag is pushed
- add a linux/arm64 binary and a SHA256SUMS file to the release assets

# gospn 0.11.1

- bugfix: integer subtraction (e.g. `reward r #P1 - #P2`) was evaluated as addition
- bugfix: `!=` between an integer and a float was evaluated as `==`
- bugfix: `pow` with an integer base and a negative integer exponent looped forever
- bugfix: missing closing parenthesis in the symbolic form of `pow`/`max`/`min`
- bugfix: TestPNreader4 read a misspelled path and always failed
- rename MATLABBuffer.WriteByte to PutByte (it conflicted with io.ByteWriter)
- add a Dockerfile / docker-compose.yml for the build and test environment
- apply gofmt

# gospn 0.11.0

- integrate the conversion from XML to Petrinet definition file into gospn binary

# gospn 0.10.2

- add a binary to convert MxGraph data to Petrinet definition file

# gospn 0.10.1

- build with go 1.16
- use go.mod
- provide Apple m1 binary

# gospn 0.10.0

- enhancement: generate block matrix

# gospn 0.9.5

- bugfix: fix rewards in sim (Do not use sim before this version)
- enhancement: Use MT64

# gospn 0.9.4

- bugfix: add rates for transitions with the same source and destination

# gospn 0.9.3

- enhancement: Implement updata function

# gospn 0.9.2

- enhancement: add option `-p` to put additional definition
- enhancement: remove the prefix `rwd`
- enhancement: implement sim

# gospn 0.9.1

- bugfix: the styple of IMM and EXP trans in view mode
- bugfix: output GxGxE for all EXP/GEN groups even if there is no transitions

# gospn 0.9.0

- First release
    - Draw a Petrinet graph with graphviz from Petrinet definition file
    - Generate a marking graph and the transition matrix (continuous-time Markov chain)
    - Write MATLAB matrix for the transition matrix
