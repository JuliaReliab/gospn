### Enhancement

- Sampling-base search

### Deferred

- Perfect sampling
  - the algorithm is known; fitting its coupling / monotonicity requirements onto the
    existing search is the hard part

### Decided against (2026-08-31)

- refactor tangible algorithm
  - the strategy seam is already there: `makingGraphGenerator` in markinggraph.go is a
    two-method interface and `CreateMarkingGraph` takes it, so another tangible
    algorithm needs no refactor first -- a type, a constructor and a CLI choice.
  - the two searches share no algorithm. `create`, `createMarking`, `visitGenMark`,
    `visitImmMark` and `addMarkAs*` all differ; the only identical code is ~40 lines of
    plumbing (`addLinkAs*`, `createGenVec`, `createNextMarking`, `clampEvents`).
    Keeping that duplicated is what keeps the strategies independent, which is what
    matters when the current tangible reduction is the cheap incomplete one and a
    complete one may sit beside it.
  - `-t` is a bool flag, so it becomes `-search plain|tangible|...` when there is a
    third. Worth doing with that change, not before it.
- refactor simulation
- compiler with go generate
