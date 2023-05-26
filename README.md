# gospn

gospn is a tool for stochastic Petrinet. gospn can analyze
- SPN (Stochastic Petrinet): All transitions are EXP transitions
- GSPN (Generalized Stochastic Petrinet): The net includes both EXP and IMM transitions
- MRSPN (Markov ReGenerative Petrinet): The net includes EXP, IMM and GEN transitions
where
- IMM transition: A transition immediately fires when it enables
- EXP transition: A transition fires with time delay following the exponential distribution
- GEN transition: A transition fires with time delay following a general distribution

## Compile

```sh
make build
```

The tool has been developed with Go 1.16 on linux/amd64. A single binary `gospn` is built in `bin` directory.
If you want to do the cross compile, please uses

```sh
make build_all
```

Also we need ANTLR4 to compile the tool. Please install it before compiling.

```sh
sudo apt update
sudo apt install antlr4
```

## Usage

### Draw a perinet

```sh
gospn view [-i <infile>] [-o <outfile>] [-pre <string>] [-post <string>]
```

The inputfile is a text file that describes the definition of Perinet (see `Petrinet definition` in detail).
The output is a dot file for graphviz. To draw a picture, please run
```sh
dot -Tsvg -o petri.svg outfile
```
When the option `i` is omitted, the tool uses stdin to read the defition of Petrinet. Also, when the option `o` is omitted,
the tool displays the contens of dot file with stdout.

### Generate marking

```sh
gospn mark [-i <infile>] [-o <outfile>] [-t] [-m <filename>] [-g <filename>] [-pre <string>] [-post <string>]
```

The tool analyzes the Petrinet and outputs MATLAB matrix for the transition matrix. The option `-t` creates a (semi) tangible marking.
The option `-m` outputs a dot file to draw the marking graph. The option `-g` outputs a dot file to draw a group marking.
The option `-pre` is to put some additional defintion like parameters to the beginning of Petrinet definition.
The option `-post` is to put some additional defintion like parameters to the end of Petrinet definition.

### Monte Carlo simulation

```sh
gospn sim [-i <infile>] [-o <outfile>] [-s <int>] [-f <file>] [-c <string>] [-pre <string>] [-post <string>]
```

Simulate a given Petrinet and compute rewards based on Monte Carlo simulation. The option `-o` indicates the name of MATLAB
matrix file to store the vectors for rewards. The option `-s` is a seed of random number generator.
The option `-f` indicates a configuration file for the simulation that are written by
JSON. Also, the option `-c` provides the JSON configuration as strings. If both `-f` and `-c` are given, the option `-f` is used.
The option `-pre` is to put some additional defintion like parameters to the beginning of Petrinet definition.
The option `-post` is to put some additional defintion like parameters to the end of Petrinet definition.

An example of configuration is given by
```json
{ "time": 10, "firings": 1000, "simulations": 10, "rewards": ["avail", "unavail"] }
```
- `time` is the maximum time for simulation to stop the simulation (it's not real clock, but is the clock in simulation)
- `firings` is the maximum number of firings to stop the simulation
- `simulations` is the number of times of simulations
- `rewards` is a vector to indicate the rewards to be computed

The stop condition for one simulation is the AND condition for `time` and `firings`.
Also, if `time` and `firings` are set to zeros, these stop conditions are removed.
In `rewards`, we should set strings of rewards that are described in the Petrinet definition file
(It may be useful to use the option `-p` if there is no reward you want).

### Monte Carlo simulation for one path

```sh
gospn test [-i <infile>] [-s <int>] [-t <value>] [-n <int>] [-pre <string>] [-post <string>]
```

Simulate one path for a given Petrinet. For example, this is used for a testing of Petrinet definition.
The option `-s` is a seed of random number generator. The option `-t` indicates the stop condition with
elapsed time. The option `-n` indicates the stop condition based on the number of firrings.
The option `-pre` is to put some additional defintion like parameters to the beginning of Petrinet definition.
The option `-post` is to put some additional defintion like parameters to the end of Petrinet definition.
This mode just displays a path to stdout.

## Definition of Petrinet

#### Expressions

The expression in the definition file consists of values (bool, int, float), variables, operators and functions.

#### Values:

- bool: `true`, `false`
- int: `0`, `1`, ...
- float: `0.0`, `1.0`, `1.0e-4`, ...

#### Variables:

We use the characters; alphabets, digits, underline and period for the label of variable.
Variables can be defined without an explict declaration like a script language. 

```
a = 1   // a is assigned by 1
b = 1.0 // b is assigned by 2

c = a + b // c is assinged by a plus of 1 + 1.0. c becomes float

x = y
y = 10 // x also becomes 10 (lazy eval)
```

#### Operators

- `+`, `-`, `*`, `/`, `div`: Arithmetic operators. `div` is an integer division.
- `==`, `!=`, `<`, `<=`, `>`, `>=`: Comparison operators
- `!`, `&&`, `||`: Logical operators
- `#`: The operator to represent the number of tokens.  `#<place>` indicates the number of tokens in `<place>`.
- `?`: The operator to represent the enable condition. `?<transition>` indicates the enable condition of `<transition>`.

#### Functions

- `ifelse`: If-then-else function. `ifelse(<condition>, <expression1>, <expression2>)` means that if `<condition>` holds,
the function returns `<expression1>`. Otherwise, the function returns `<expression2>`.
- `exp`, `log`, `sqrt`, `pow`, `min`, `max`: Mathematical functions
- `det`, `unif`, `expdist`: Distribution functions.
  - `det`: Constant distribution. `det(10)` means the constant distribution whose sample is 10.
  - `unif`: Uniform distribution. `unif(0, 1)` means the uniform distribution on the domain (0,1).
  - `expdist`: Exponential distribution. `expdist(2)` means the exponential distribution whose rate is 2.

### Place

```
place <label> (init = <int>, max = <int>)
```

Define a place with name `<label>`.

Options:
- `init`: An integer means the number of initial tokes that put on the place. The default is 0.
- `max`: An integer means the maximum number of tokens. If the number of tokens exceeds the maximum number of
tokens by some firings, the number of tokens is forced to the maximum number of tokens.
The default is 255.

### Transitions

```
imm <label> (weight = <expression>, guard = <expression>, priority = <int>, vanishable = <bool>) {
  <update statements>
}
```

Define an IMM (immediate) transition with name `<label>`.
The parentheses indicates options, and the curly brackets indicates the procedure
after firing of this transition. The parts of parentheses and the curly brackets are
optional.

Options:
- `weight`: The expression to determine the probabilistic priority of
firing when there are some competitive IMM transitions.
The large weight has the high priority probabilistically to other IMM transitions.
The default is 1.0.
- `guard`: The expression to indicate the guard condition.
- `priority`: An integer indicating the deterministic priority of firing
when there are some competitive IMM transitions.
The large priority has the high priority to other competitive IMM transitions.
The default is 0.
- `vanishable`: A logical whether the IMM transition can be vanished in the (semi-)tangible marking.
If `vanishable` is false, the IMM transition is never vanished even if it satisfies the vanishing condition.
Otherwise, if `vanishable` is true, the IMM transition may be vanished when the option `-t' is selected in the marking mode.
But, note that the IMM transition is vanished by the option `-t` only when the vanishing condition holds.
The default is true.

Update:
The `<update statements>` describes an additional procedure to change the marking after firing.

```
{
  #P1 = #P1 + 1
  #P2 = #P2 - 1
}
```
For example, the above shows that the number of tokens in P1 is increases
by one and the number of tokens in P2 decreases by one for the marking just
after the firing of this transition.

```
exp <label> (rate = <expression>, guard = <expression>, priority = <int>) {
  <update statements>
}
```

Define an EXP transition with name `<label>`.
The parentheses indicates options, and the curly brackets indicates the procedure
after firing of this transition. The parts of parentheses and the curly brackets are
optional.

Options:
- `rate`: The expression of transition rate. The default is 1.0.
- `guard`: The expression to indicate the guard condition.
- `priority`: An integer indicating the deterministic priority of firing
when there are some competitive EXP/GEN transitions.
The large priority has the high priority to other competitive EXP/GEN transitions.
The default is 0.

```
gen <label> (dist = <function>, guard = <expression>, policy = prd | prs | pri, priority = <int>) {
  <update statements>
}
```

Define a GEN transition with name `<label>`.
The parentheses indicates options, and the curly brackets indicates the procedure
after firing of this transition. The parts of parentheses and the curly brackets are
optional.

Options:
- `dist`: The name of general distribution with parameters. The default is `det(1)`
- `guard`: The expression to indicate the guard condition.
- `policy`: The policy of age of GEN transition when the preemption occurs. The default is `prd`.
  - `prd`: Preemptive repeat different. The age of GEN transition is renewed when the preemption occurs.
  - `prs`: Preemptive resume. The age of GEN transition stops and resumes when the preemption occurs.
  - `pri`: Preemptive repeat identical. The firing time is fixed by a random variable.
  The age of GEN transition becomes zero when the preemption occurs. This policy is effective only in the simulation.
- `priority`: An integer indicating the deterministic priority of firing
when there are some competitive EXP/GEN transitions.
The large priority has the high priority to other competitive EXP/GEN transitions.
The default is 0.

### Arc

```
arc <source> to <target> (multi = <expression>)
iarc <place> to <transition> (multi = <expression>)
oarc <transition> to <place> (multi = <expression>)
```

Define an arc to connect between place and transition. `iarc` is the definition
of an input arc from a place to a transition. `oarc` is the definition of an
output arc from a transition to a place. `arc` is the definition of an input
or output arc that is detected by checking objects put on `<source>` and
`<target>` automatically.

Option:
- `multi`: The expression indicating the multiplicity of arc. The default is 1.

```
harc <place> to <transition> (multi = <expression>)
```

Define an inhibited arc from a place to a transition.

Option:
- `multi`: The expression indicating the multiplicity of arc. The default is 1.

### Reward

To be written.


### Block

To be written.

### Example of definition (SPN)

```
/*
  M/M/m/b queue
  (referred from an example of SPNP; https://trivedi.pratt.duke.edu/software_packages/spnp)
*/

place buf

exp trin (rate = lambda)
exp trserv (rate = rate.serv)

oarc trin to buf
harc buf to trin (multi = b)
iarc buf to trserv

rate.serv = ifelse(#buf < m, #buf*mu, m*mu)

qlength = #buf
util = ifelse(?trserv, 1, 0)
tput = ifelse(?trserv, rate.serv, 0)
probrej = ifelse(#buf == b, 1, 0)
probempty = ifelse(#buf == 0, 1, 0)
probhalffull = ifelse(#buf == b div 2, 1, 0)

// buffer size
b = 50

// # of servers
m = 10

lambda = 0.1
mu = 1.0

// reward
reward numOfCustomer #buf
```

### Example of definition (GSPN)

```
/*
  Example: IaaS Cloud (monolithic)
  R. Ghosh, F. Longo, F. Frattini, S. Russo and K.S. Trivedi,
  Scalable analytics for IaaS cloud availability,
  IEEE Transactions on Cloud Computing, 2:1, 57-70, 2014.
*/

// initial virtual machines. Note: the number of states exeeds 1 million if we set n >= 6
n = 3 

place Ph (init = n)
place Pfh
place Pchm
place Pwhm
place Pbc_d
place Pbw
place Phcm
place Phwm
place Pwr
place Pw (init = n)
place Pfw
place Pcwm
place Pbc_dd
place Pwcm
place Pcr
place Pc (init = n)
place Pfc

exp Thr (rate = Thr.rate)
exp Thf (guard = g2, rate = r1)
exp Tbchf (guard = g1, rate = r1)
exp Tbwhf (rate = r1)
exp Tchm (rate = Tchm.rate)
exp Twhm (rate = Twhm.rate)
exp Twr (guard = g4, rate = Twr.rate)
exp Thcm (rate = Thcm.rate)
exp Thwm (rate = Thwm.rate)
exp Twf (guard = g3, rate = r2)
exp Tbcwf (rate = r2)
exp Tcf (rate = r3)
exp Tcwm (rate = Tcwm.rate)
exp Tcr (guard = g5, rate = Tcr.rate)
exp Twcm (rate = Twcm.rate)

imm tcr1
imm twr1
imm twr2
imm tcr2
imm tcr3

iarc Ph to Thf
iarc Ph to Tbchf
iarc Ph to Tbwhf
iarc Pfh to Thr
iarc Pchm to Tchm
iarc Pwhm to Twhm
iarc Pbc_d to tcr1
iarc Pbw to twr1
iarc Phcm to Thcm
iarc Phwm to Thwm
iarc Pwr to twr1
iarc Pwr to twr2
iarc Pw to Twf
iarc Pw to Tbwhf
iarc Pw to Tbcwf
iarc Pfw to twr2
iarc Pcwm to Tcwm
iarc Pbc_dd to tcr2
iarc Pwcm to Twcm
iarc Pcr to tcr1
iarc Pcr to tcr2
iarc Pcr to tcr3
iarc Pfc to tcr3
iarc Pc to Tcf
iarc Pc to Tbcwf
iarc Pc to Tbchf

oarc Thr to Ph
oarc Thf to Pfh
oarc Tbchf to Pchm
oarc Tbwhf to Pwhm
oarc Tchm to Ph
oarc Tchm to Pbc_d
oarc Twhm to Ph
oarc Twhm to Pbw
oarc tcr1 to Phcm
oarc twr1 to Phwm
oarc twr2 to Pw
oarc Twr to Pwr
oarc Thcm to Pc
oarc Thwm to Pw
oarc Twf to Pfw
oarc Tbcwf to Pcwm
oarc Tcwm to Pbc_dd
oarc Tcwm to Pw
oarc tcr2 to Pwcm
oarc Tcr to Pcr
oarc tcr3 to Pc
oarc Tcf to Pfc
oarc Twcm to Pc

g1 = #Pw == 0
g2 = (#Pw == 0) && (#Pc == 0)
g3 = #Pc == 0
g4 = #Pfw + #Pbw > 0
g5 = #Pfc + #Pbc_d + #Pbc_dd > 0

r1 = #Ph * lambda_h
r2 = #Pw * lambda_w
r3 = #Pc * lambda_c

lambda_h = 1/(500*60) // MTTF 500hrs
lambda_w = 1/(1750*60) // MTTF 1750hrs
lambda_c = 1/(2500*60) // MTTF 2500hrs

nr = 1
mu = 1/(3*60) // MTTR 3hrs

gam.ch = 1/30 // MTTM 30min
gam.wh = 1/30 // MTTM 30min
gam.hc = 1/30 // MTTM 30min
gam.hw = 1/30 // MTTM 30min
gam.cw = 1/30 // MTTM 30min
gam.wc = 1/30 // MTTM 30min

Thr.rate = ifelse(#Pfh <= nr, #Pfh * mu, nr * mu)
Twr.rate = ifelse(#Pfw + #Pbw <= nr, mu*(#Pfw + #Pbw), nr * mu)
Tcr.rate = ifelse(#Pfc + #Pbc_d + #Pbc_dd <= nr, (#Pfc + #Pbc_d + #Pbc_dd) * mu, nr * mu)

Twhm.rate = #Pwhm * gam.wh
Tchm.rate = #Pchm * gam.ch
Tcwm.rate = #Pcwm * gam.cw
Thcm.rate = #Phcm * gam.hc
Thwm.rate = #Phwm * gam.hw
Twcm.rate = #Pwcm * gam.wc

reward rwd1 #Ph
reward rwd2 #Pw
reward rwd3 #Pc
reward avail1 ifelse(#Ph >= 1, 1.0, 0.0)
reward avail2 ifelse(#Ph >= 2, 1.0, 0.0)
reward avail3 ifelse(#Ph >= 3, 1.0, 0.0)
reward rwd5 ifelse(#Pw >= 1, 1.0, 0.0)
reward rwd6 ifelse(#Pc >= 1, 1.0, 0.0)
```

### Example of definition (MRSPN)

```
/*
  Example: RAID6
  F. Machida, R. Xia and K.S. Trivedi,
  Performability modeling for RAID storage systems by Markov regenerative process,
  IEEE Transactions on Dependable and Secure Computing
*/

// HDD model

place Pn (init = 6)
place Pdf
exp Tdfail (guard = gfail, rate = Tdfail_rate)
gen Trebuild (guard = gfail, dist = Trebuild_dist)
imm Tinit (guard = ginit)
arc Pn to Tdfail
arc Tdfail to Pdf
arc Pdf to Trebuild
arc Pdf to Tinit
arc Trebuild to Pn
arc Tinit to Pn

// Reconstruction model

place Po (init = 1)
place Pr
place Pc
imm Tstart (guard = gstart)
gen Trecon (dist = Trecon_dist)
imm Tend (guard = gend)
arc Po to Tstart
arc Tstart to Pr
arc Pr to Trecon
arc Trecon to Pc
arc Pc to Tend
arc Tend to Po

// rate and gurads

Tdfail_rate = #Pn * lambda
gfail = #Po == 1
gstart = #Pdf > 2
ginit = #Pc == 1
gend = #Pdf == 0

// params

Trebuild_dist = det(MTTR1)
Trecon_dist = det(MTTR2)

MTTF = 1.0e+6 // [hours]
lambda = 1/MTTF
MTTR1 = 2 // [hours]
MTTR2 = 24 // [hours]
```

## Graphical Definition

The gospn generates text-based Petrinet definition file from an XML diagram.

```sh
gospn gen [-i <infile>] [-o <outfile>]
```

The inputfile is an XML file (the extension of .drawio or .xml) that is generated by [draw.io/diagrams.net](https://www.diagrams.net/).
The output is a text-based Petrinet definition file this can be used for the inputfile of gospn directly.
When the option `i` is omitted, the tool uses stdin to read the XML.
Also, when the option `o` is omitted, the tool displays the Petrinet definition to stdout.

### How to draw a Petrinet in diagram.net

#### Use the shape library

Please add the shape library `MRSPN`

1. Access diagrams.net
2. File -> Open Library -> URL
3. Fill out the following URL

```
https://gist.githubusercontent.com/okamumu/d10aabf442905b51f627df803139bd87/raw/eee5ebf1d3730913afb58345804bf08527979bca/MRSPN.xml
```

The shape library includes the components; Place, IMM, EXP and GEN which have the required properties.

#### Place

To draw places, we use `EditData` for a component.

1. Select an object on diagrams.net
2. Select the `Arrange` tab in the righthand panel, or right click the object.
3. Select `EditData`
4. Add the property `type`
5. Put `place` into the property `type`

The tool recognizes a place object if it has the property `type=place` even if the object is not a circle.
The other properties `init` and `max` are also used by adding them as the properties of diagram object.
Furthermore, if we put a text into an object (the text becomes a label property of object), the text is used
for `init` of place in the text-based Petrinet definition.
The label of place uses the closest text object among the text objects having the same parent.
Therefore, if the place object and the text object are grouped, the text object is used as
the label of place.

#### IMM Transition

Similarly to place, we add the property `type` and put the value `imm` to the `type` property.
If we add the properties `weight`, `guard`, `priority` and `vanishable` to the property of object,
their values becomes the values of corresponding options in the text-based Petrinet definition.
The label of transition uses the closest text object among the text objects having the same parent.
Therefore, if the transition object and the text object are grouped, the text object is used as
the label of transition.

#### EXP Transition

Add the property `type` with the value `exp`.
If we add the properties `rate`, `guard` and `priority` to the property of object,
their values becomes the values of corresponding options in the text-based Petrinet definition.
The label of transition uses the closest text object among the text objects having the same parent.
Therefore, if the transition object and the text object are grouped, the text object is used as
the label of transition.

#### GEN Transition

Add the property `type` with the value `gen`.
If we add the properties `dist`, `guard`, `policy` and `priority` to the property of object,
their values becomes the values of corresponding options in the text-based Petrinet definition.
The label of transition uses the closest text object among the text objects having the same parent.
Therefore, if the transition object and the text object are grouped, the text object is used as
the label of transition.

#### Arc

If place and transition objects are connected by a directed edge,
the edge corresponds to `arc` in the text-based Petrinet definition.
When the edge has an edgelabel, the edgelabel gives the `multi` option.

#### Inhibit Arc

If place and transition objects are connected by a directed edge whose endarrow is oval,
the edge corresponds to `harc` in the text-based Petrinet definition.
When the edge has an edgelabel, the edgelabel gives the `multi` option.

#### Pre/Post

When we add the property `type` with the value `pre` or `post`, the text written in the
label of this object put to the beginning/end of Petrinet defition.

#### Comment

When we add the property `type` with the value `comment`, the object is ignored.

#### Template

Template URL:
[https://raw.githubusercontent.com/okamumu/gospn/master/diagrams/mrspn.xml](https://raw.githubusercontent.com/okamumu/gospn/master/diagrams/mrspn.xml)

