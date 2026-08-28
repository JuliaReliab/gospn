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
