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
