package parser

import (
	"../petrinet"
	"bytes"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
	"os"
	"testing"
)

func TestPNListener1(t *testing.T) {
	logger = log.New(os.Stdout, "[Hello] ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	fmt.Println("test1")
	is := antlr.NewInputStream("1 + 2 * 3")
	lexer := NewJSPNLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewJSPNLParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(NewPNListener(), p.Expression())
}

func TestPNListener2(t *testing.T) {
	logger = log.New(os.Stdout, "[Hello] ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	fmt.Println("test2")
	is := antlr.NewInputStream(`
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
gen Trebuild (guard = gfail) {
	#Pn = 1
}
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
gen Trecon
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

Trebuild.dist = exp(MTTR1)
Trecon.dist = log(MTTR2)

MTTF = 1.0e+6 // [hours]
lambda = 1/MTTF
MTTR1 = 2 // [hours]
MTTR2 = 24 // [hours]

reward r1 #Pc
`)
	lexer := NewJSPNLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewJSPNLParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(NewPNListener(), p.Prog())
}

func TestPNListener3(t *testing.T) {
	logger = log.New(os.Stdout, "[Hello] ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	fmt.Println("test4")
	is := antlr.NewInputStream(`
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
gen Trebuild (guard = gfail) {
	#Pn = 1
}
imm Tinit + (guard = ginit)
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
gen Trecon
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

Trebuild.dist = sqrt(MTTR1)
Trecon.dist = max(MTTR2, 1)

MTTF = 1.0e+6 // [hours]
lambda = 1/MTTF
MTTR1 = 2 // [hours]
MTTR2 = 24 // [hours]

reward r1 #Pc
`)
	lexer := NewJSPNLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	NewJSPNLParser(stream)
	// antlr.ParseTreeWalkerDefault.Walk(NewPNListener(), p.Prog())
}

func TestPNListener4(t *testing.T) {
	logger = log.New(os.Stdout, "[Hello] ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	fmt.Println("test4")
	is := antlr.NewInputStream(`
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
gen Trebuild (guard = gfail) {
	#Pn = 1
}
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
gen Trecon
imm Tend (guard = gend)
arc Po to Tstart
arc Tstart to Pr (multi = 10)
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

Trebuild.dist = exp(MTTR1)
Trecon.dist = log(MTTR2)

MTTF = 1.0e+6 // [hours]
lambda = 1/MTTF
MTTR1 = 2 // [hours]
MTTR2 = 24 // [hours]

reward r1 #Pc
`)
	lexer := NewJSPNLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewJSPNLParser(stream)
	listener := NewPNListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Prog())
	net, initmark := makeNet(listener.builder.labels, listener.builder.env)
	fmt.Println(net)
	fmt.Println(initmark)

	writer := bytes.NewBuffer(make([]byte, 0, 256))
	net.ToPNDot(writer)
	fmt.Println(writer.String())
	// for key, value := range listener.builder.env {
	// 	switch node := value.(type) {
	// 	case *PNNode:
	// 		fmt.Println(key, node)
	// 	}
	// }
}

func TestPNListener5(t *testing.T) {
	logger = log.New(os.Stdout, "[Hello] ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	fmt.Println("test5")
	is := antlr.NewInputStream(`
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
gen Trebuild (guard = gfail, dist = Trebuild.dist)
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
gen Trecon (dist = Trecon.dist)
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

Trebuild.dist = expdist(MTTR1)
Trecon.dist = unif(1.0, 2.0)

MTTF = 1.0e+6 // [hours]
lambda = 1/MTTF
MTTR1 = 2 // [hours]
MTTR2 = 24 // [hours]

reward r1 #Pc
`)
	lexer := NewJSPNLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewJSPNLParser(stream)
	listener := NewPNListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Prog())
	net, imark := makeNet(listener.builder.labels, listener.builder.env)

	writer := bytes.NewBuffer(make([]byte, 0, 256))
	net.ToPNDot(writer)
	fmt.Println(writer.String())

	mg := petrinet.CreateMarkingGraphWithDFS(net, imark)
	writer2 := bytes.NewBuffer(make([]byte, 0, 256))
	mg.ToMarkDotWithLabelAndGroup(writer2)
	fmt.Println(writer2.String())
}
