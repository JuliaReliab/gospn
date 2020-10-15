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

The tool has been developed with Go 1.10.4 on linux/amd64. A single binary `gospn` is built in `bin` directory.

## Usage

### Draw a perinet

```sh
gospn view -i <infile> -o <outfile>
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
gospn mark -i <infile> -o <outfile> [-t] [-m <dotfile>] [-g]
```

The tool analyzes the Petrinet and outputs MATLAB matrix for the transition matrix. The option `-t` creates a (semi) tangible marking.
The option `-m` outputs a dot file to draw the marking graph. The option `-g` displays the dot to draw a group marking.

## Definition of Petrinet

### Example of SPN

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

To be written

### Example of GSPN

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

exp Thr (guard = g2)
exp Thf (rate = r1, guard = g2)
exp Tbchf (guard = g1, rate = r1)
exp Tbwhf (rate = r1)
exp Tchm
exp Twhm
exp Twr (guard = g4)
exp Thcm
exp Thwm
exp Twf (guard = g3, rate = r2)
exp Tbcwf (rate = r2)
exp Tcf (rate = r3)
exp Tcwm
exp Tcr (guard = g5)
exp Twcm

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

Tchm.rate = #Pchm * gam.ch
Twhm.rate = #Pwhm * gam.wh
Thcm.rate = #Phcm * gam.hc
Thwm.rate = #Phwm * gam.hw
Tcwm.rate = #Pcwm * gam.cw
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

To be written

### Example of MRSPN

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

To be written
