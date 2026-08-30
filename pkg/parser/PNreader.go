package parser

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/okamumu/gospn/pkg/petrinet"
)

// logger is initialised here rather than only in the entry points below: it was nil
// until PNreadFromText or PNreadFromFile ran, so calling makeNet or walking a parse tree
// directly -- which is what a test does -- dereferenced nil. Output is discarded, which
// is what the entry points set anyway.
var logger = log.New(io.Discard, "[PNparser] ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

// syntaxErrors collects what ANTLR reports instead of letting its default listener write
// to stderr, where nothing can act on it.
type syntaxErrors struct {
	*antlr.DefaultErrorListener
	msgs []string
}

func (c *syntaxErrors) SyntaxError(_ antlr.Recognizer, _ interface{}, line, col int, msg string, _ antlr.RecognitionException) {
	c.msgs = append(c.msgs, fmt.Sprintf("line %d:%d: %s", line, col, msg))
}

func (c *syntaxErrors) err() error {
	if len(c.msgs) == 0 {
		return nil
	}
	// Every error, not just the first: a definition with two mistakes in it should say
	// so once rather than over two runs.
	return fmt.Errorf("syntax error in the Petri net definition:\n  %s", strings.Join(c.msgs, "\n  "))
}

// read parses one definition. It is the single place the five steps live; PNreadFromText
// and PNreadFromFile used to be copies of them, differing only in the input stream.
//
// The AST builder and the compiler report problems by panicking (`logger.Panic`), which
// is convenient inside a recursive walk and useless to a caller: a mistyped definition
// reached the user as a Go stack trace. The panic is caught here and returned as an
// error, so the panics stay an internal convention and stop at this boundary.
func read(is antlr.CharStream) (net *petrinet.Net, imark []petrinet.MarkInt, err error) {
	errs := &syntaxErrors{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	lexer := NewJSPNLLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errs)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewJSPNLParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errs)

	tree := p.Prog()
	// Walking a tree that has error nodes in it builds a net from a definition nobody
	// wrote, so the syntax has to be sound first.
	if err := errs.err(); err != nil {
		return nil, nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			net, imark = nil, nil
			err = fmt.Errorf("%v", r)
		}
	}()
	listener := NewPNListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	net, imark = makeNet(listener.builder.labels, listener.builder.env)
	return net, imark, nil
}

// PNreadFromText reads a Petri net definition from a string.
func PNreadFromText(text string) (*petrinet.Net, []petrinet.MarkInt, error) {
	return read(antlr.NewInputStream(text))
}

// PNreadFromFile reads a Petri net definition from a file.
func PNreadFromFile(fileName string) (*petrinet.Net, []petrinet.MarkInt, error) {
	is, err := antlr.NewFileStream(fileName)
	if err != nil {
		return nil, nil, err
	}
	return read(is)
}
