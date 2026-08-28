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
