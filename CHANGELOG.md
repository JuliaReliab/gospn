# gospn 0.23.0

- **The parser's tests assert something now.** `pkg/parser` had 51 `fmt.Println` calls
  and, in two of its five test files, not one `t.Error`: a walk that produced complete
  nonsense would still have passed. This is the gap that let the arithmetic bugs fixed
  in 0.11.1 survive, and it has been carried in the notes since 2026-08-27.

  Every test now checks a value. What they cover that nothing did before: operator
  precedence in a parsed expression, that declarations reach the environment tagged and
  in order, that an arc multiplicity and an update block reach the token game, that a
  distribution reaches the model with its parameters, that `PNreadFromFile` and
  `PNreadFromText` agree on all nine bundled nets, and that a missing file is an error.

  Checked by mutation: reintroducing the 0.11.1 `minus`-is-`plus` bug fails 4 tests,
  ignoring an error node fails 1, dropping an arc multiplicity fails 1. Statement
  coverage of the package is 59.4%, against 59.0% for the printing tests it replaces --
  which is the point: the old number came from walking nets and looking at nothing.

- **Fixed: a syntax error said only `Parser error. Stop to run`.** It now names the
  position: `syntax error at line 6:10, near "..."`. The definition still reaches the
  user as a panic rather than an error -- the reader entry points return none -- but it
  at least says where to look.

- **Fixed: `parser.logger` was nil until a reader entry point ran**, so calling
  `makeNet` or walking a parse tree directly dereferenced nil. It is initialised to a
  discarding logger, which is what the entry points set anyway.

# gospn 0.22.0

- **An MRSPN result file now says what its general blocks mean.** A `P<k>` block is a
  0/1 jump matrix, so the distribution never reached the file through it, and `P<k>` is
  a counter that names no transition -- `det(5)` and `det(99)` produced byte-identical
  files, and nothing said whether `P0` was `Trebuild` or `Trecon`. Both are needed to
  solve the regenerative process the file describes; this was found by the cross-check
  against `PetriAnalysis.jl`, which could compare every block and still not notice a
  changed distribution.

  Two tab-separated text elements are added, and only for a net that has general
  transitions:

  ```
  gentrans   P0  Trebuild  det(2)        which transition each P<k> block is
             P1  Trecon    det(24)

  groupgen   G1  Trecon    E  det(24)    per group, the general transitions that are
             G2  Trebuild  E  det(2)     aging (E) or preempted (P) in it -- what
             I2  Trebuild  E  det(2)     governs how long the group is occupied
  ```

  Distributions render in the definition language: `det(2)`, `unif(1,3)`,
  `expdist(0.5)`. `DistributionInterface` gained `String()` for this.

  New: `MarkingGraph.BlockGenTrans()` (keyed like `TransLabels`) and
  `MarkingGraph.GroupGens()`.

- `example/raid6.spn` joins the JSON goldens, so the GEN blocks and the two new elements
  are pinned in a file that diffs. It has two general transitions, which is where the
  `P<k>` numbering can go wrong.

# gospn 0.21.0

- **The marking-graph search uses about a third less memory and is about a third
  faster.** Two structures cost far more than what they held:

  - the links existed **twice** -- `newMarkingGraph` copied every link into a
    per-`GroupTrans` slice while the full list stayed on the graph for the DOT output.
    They are now partitioned in place, each group getting a subslice of the one array.
  - `markToGroup`, `markToInt` and the search's `markToGenvec` / `markToGroupType` were
    four `map[*Mark]...` side tables, about 200 bytes per state of map overhead where
    four fields on `Mark` cost 32 -- and every read was a hash lookup, which is where
    the speed came from.

  Measured on `example/iaas_cloud.spn` with `n = 5` (910,731 states, 3.1 M links):

  |  | before | after |
  |---|---|---|
  | retained | 635 B/state (578 MB) | 407 B/state (371 MB) |
  | peak | 1431 B/state | 921 B/state |
  | time | 9.0 s | 5.1 s |

  `n = 6` (3,811,992 states) completes either way in a 6 GB container: it needs about
  3.5 GB now instead of about 5.5 GB, and 12 s instead of 18 s. It still needs an
  explicit `-maxstates`. **This does not make `example/k8s.spn` enumerable** -- that net
  has 53 places and several with `max = 1000`, and no amount of memory reaches it;
  `gospn sim` is what it is for.

  No value changed: the golden tests are untouched, and the link partition is a *stable*
  sort for that reason -- `getTransMatrix` sums the links in this order and a float sum
  depends on it.

- **`BenchmarkMarkingGraphMemory`** (`test/`) reports the heap the graph retains and the
  heap in use before collecting, over `iaas_cloud.spn` at n=3,4,5. `-benchmem` reports
  neither, and the two differ by more than 2x: a net fails on the first number and is
  analysed out of the second.

- **Fixed: a net with an undefined variable ran the whole search and then panicked.**
  `gospn mark -i example/cold_vm_reju.spn` died with a Go stack trace out of
  `getTransMatrix`, on a net of 94 states. An undefined variable is not a parse error --
  the expression language resolves variables lazily and an unresolved one evaluates to
  its own name as a string -- and nothing asks a rate for a number until the matrices are
  built. `Net.CheckExpressions` now evaluates every guard, rate, weight and reward once
  at the initial marking, before anything else runs, and reports all of them together:

  ```
  the net does not evaluate:
    transition Thfail: rate: ... (Thfai.rate) -- an undefined variable evaluates to its own name as a string
    transition Tvrestart: rate: ... (Tvreset.rate) -- ...
  ```

  `mark`, `sim` and `test` run the check; `view` does not, since drawing a net should not
  require its rates to evaluate.

- **Fixed: two typos in `example/cold_vm_reju.spn`** -- `Thfai.rate` and `Tvreset.rate`,
  for `Thfail.rate` and `Tvrestart.rate` (`Tvreset` is an IMM transition and has no
  rate). The net analyses for the first time: 94 states, 49 tangible.

- **`TestEveryExampleEvaluates`** parses and evaluates every net in `example/`. This was
  the third bundled example to ship unevaluatable, and every one of them was found by a
  person running the tool rather than by a test, because nothing ran the bundled nets.

# gospn 0.20.0

- **`gospn mark` writes the markings.** A result file carried the matrices and nothing
  that said what a row meant: the markings went to a separate text file, and only when
  `-s` was given. The file now also holds `place` (the place order, one name per line)
  and a `mark<G>` matrix per group -- one row per state, one column per place -- where
  row *k* is the state of row *k* of every matrix of that group. Verified row for row
  against the `-s` output.

  This is what a reader needs to key a state on its marking, which is what comparing
  gospn's generator against another implementation's requires: the two enumerate the
  state space in different orders.

  New: `MarkingGraph.StateMarkings()` and `MarkingGraph.Net()`. `StateLabels()` is now
  written in terms of the former.

- **Fixed: a one-row matrix lost a dimension in `.npz`.** `dimsToShape` dropped a
  leading 1 so that a 1-by-n vector loads as a 1-D array, which also applied to a
  declared matrix -- a group with a single state would have loaded as 1-D while every
  other group loaded as 2-D. Vectors are unchanged.

# gospn 0.19.0

- **Fixed: the marking-graph matrices were not reproducible.** `example/spnp_example5.spn`
  produced two different `.mat` files over 20 runs of the same binary on the same net;
  the generator diagonal differed in the last ULP. Two map-iteration orders were behind
  it: `getTransMatrix` sorted the COO entries with `sort.Slice`, which is not stable, so
  entries with the same `(i,j)` were summed in a different order each run, and
  `TransMatrix` ranged over the `groupTransToLink` map, so the row-sum accumulations
  did too. All 11 enumerable nets in `example/` now write the same bytes every run.
  No value changed.

- **`gospn test` can write a result file**, with `-o` and `-format`, in any of the three
  formats. It printed its path to stdout and nowhere else, so it was the one subcommand
  whose output nothing could read back. The file holds `place` (the place order the
  marking columns are in), `path_time`, `path_trans` and `path_state` -- a
  firings-by-places matrix of token counts -- plus provenance and the seed. Without
  `-o` nothing changes.

  `path_state` is the first 2-D dense element gospn writes. It is column-major, and the
  `.npy` member says so (`fortran_order: True`): getting that wrong transposes the
  matrix with no error anywhere. Checked against `np.load`, `scipy.io.loadmat` and
  `NPZ.jl` on `example/raid6.spn`.

- **The output layer is tested end to end.** `test/output_golden_test.go` pins the JSON
  document of the marking graph of three small nets as files (`go test ./test -update`
  regenerates them), checks that the document is byte-identical over repeated runs, and
  writes one result in all three formats and reads each back, comparing exactly. The
  cross-format agreement used to be checked by hand in scipy and NPZ.jl and did not
  survive the session it was checked in. The JSON golden is the first output test where
  a variable-name collision -- the class of bug that hid `irwd` for several releases --
  shows up as a changed line in review.

- **`make test` now runs `go test ./...`.** It named four `pkg/` directories by hand, so
  the packages added in 0.18.0 and everything in `test/` were outside it: a passing
  local `make test` meant less than it looked. The per-package targets are still there.

- The code that turns an analysis into a result file moved out of `package main` into
  `pkg/analysis`, so a test can reach it.

- README: `.npz` strings are documented as the UTF-8 `uint8` arrays they have been since
  0.18.0 (the text still described NumPy's `<U`), and a short section says that solving
  the matrices is out of scope -- that is what NMarkov.jl and scipy are for.

# gospn 0.18.0

- `gospn mark` and `gospn sim` can write **NumPy `.npz`** and **JSON** as well as MATLAB
  v5, selected with `-format mat|npz|json`. Without the flag the extension of `-o`
  decides and a name with no useful extension is written as `mat`, so every existing
  invocation is unchanged -- verified variable by variable against 0.17.0 over every net
  in `example/`

  `.npz` is a zip of `.npy` members, so it needs nothing outside the standard library.
  That matters: MATLAB v7.3 is HDF5, which has no usable pure-Go implementation, and
  cgo would break the cross-compiled release builds. A sparse matrix `M` becomes
  `M.data` / `M.indices` / `M.indptr` / `M.shape`, holding CSC with a 0-origin, which
  `scipy.sparse.csc_matrix` and `SparseMatrixCSC` each take directly. Strings are stored
  as UTF-8 byte arrays rather than NumPy's own `<U` dtype, because NPZ.jl rejects the
  *whole* archive when it meets one

  JSON is the format that can be read without a library and diffed in a review. A `.mat`
  says nothing to `git diff`, which is how the reward-vector name collision fixed in
  0.17.0 survived several releases

- the subcommands no longer assemble MATLAB elements themselves. They fill in a
  format-neutral `result.Result` and hand it to a writer, so the file format is decided
  in one place instead of being interleaved with collecting the numbers

- **elements are sorted by name**, so the same net produces the same file twice. Go map
  iteration order is random and the order varied per run before this. MATLAB does not
  care; a text golden does

- **duplicate element names are now an error** rather than a silently dropped variable.
  MATLAB's `load` keeps the last variable of a given name, `np.load` keeps the last
  member, and a JSON object key collides -- this is the check that would have caught the
  0.17.0 collision

- `gospn mark` records `gospn_version`, `gospn_revision`, `gospn_command` and `net`. It
  previously recorded nothing at all about the run that produced it; only `gospn sim`
  did, and its provenance was built inline where `mark` could not reach it

# gospn 0.17.0

- fold the rewards as the simulation produces events, instead of storing the whole
  sample path and walking it afterwards. `RunAll` allocates **186 objects per run
  instead of 318,243**, and 0.1 MB instead of 184 MB

  0.16.0 identified the ceiling on parallel scaling as contention on the Go runtime's
  allocator locks, driven by one marking allocated per firing and retained in the event
  path. Removing the retention lets those markings use a reused buffer, and the
  diagnosis holds up: the mutex profile at 10 workers falls from 1.10 s of contention to
  11.95 ms, and what is left in the CPU profile is the compiled guard and rate closures
  doing the actual work

  | workers | 0.16.0 | 0.17.0 | |
  |---|---|---|---|
  | 1 | 224 ms | 175 ms | 1.28x |
  | 2 | 124 ms | 90 ms | 1.38x |
  | 4 | 81 ms | 53 ms | 1.54x |
  | 8 | 70 ms | 39 ms | 1.79x |
  | 10 | 69 ms | 36 ms | 1.90x |

  Parallel scaling improves with it, from 3.27x to **4.83x** on ten cores, because the
  contention it was hitting is gone

- **the numbers are unchanged.** `TestSimGolden` matches 0.16.0 exactly: folding the
  rewards event by event is the same arithmetic `calcReward` performed over a stored
  path, in the same order

- **bugfix: the final reward was written over the instantaneous one.** `gospn sim`
  named the `lastrwd` vectors `<reward>_irwd`, the same name it had already used for
  `irwd`, so the file contained that variable twice and `load` returned whichever was
  written last. The instantaneous reward could not be recovered from a result file at
  all, and what came back under its name was something else. The vectors are now
  `_irwd`, `_crwd` and `_lastrwd`

  Scripts reading `<reward>_irwd` have been reading the final reward; they need
  changing to keep doing so, or leaving alone to start reading what the name says

- record what a result file came from. It held only the vectors, so a `.mat` on disk
  said nothing about the run that produced it. It now also carries `gospn_version`,
  `gospn_revision` (when the toolchain stamped one), `net`, `seed`, `simulations`,
  `endingtime`, `firings`, `parallel` and `clamped`

  The version matters more than it looks: 0.16.0 changed which random stream a given
  seed produces, so a seed alone no longer identifies a run. `clamped` is there because
  clamping is reported on stderr, and whoever reads the file later did not see it

- `matout.CreateMATLABCharMatrix` writes a string as a MATLAB char array, which is what
  the provenance above needs. The version reaches the binary through
  `-ldflags -X main.version`, set by the Makefile and the release workflow; a build
  without it reports `dev` rather than claiming a version it cannot know

- `RunSimulation` still returns the whole path, which is what `gospn test` prints, and
  still allocates a marking per firing -- a sink that keeps the markings suppresses the
  buffer reuse underneath it. Only `RunAll` streams

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
