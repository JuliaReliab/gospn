package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/okamumu/gospn/pkg/petrinet"
	"io"
	"log"
	"os"
)

// logger is initialised here rather than only in the entry points below: it was nil
// until PNreadFromText or PNreadFromFile ran, so calling makeNet or walking a parse tree
// directly -- which is what a test does -- dereferenced nil. Output is discarded, which
// is what the entry points set anyway.
var logger = log.New(io.Discard, "[PNparser] ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

func PNreadFromText(text string) (*petrinet.Net, []petrinet.MarkInt) {
	logger = log.New(os.Stdout, "[PNparser] ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	logger.SetOutput(io.Discard)
	is := antlr.NewInputStream(text)
	lexer := NewJSPNLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewJSPNLParser(stream)
	listener := NewPNListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Prog())
	net, imark := makeNet(listener.builder.labels, listener.builder.env)
	return net, imark
}

func PNreadFromFile(fileName string) (*petrinet.Net, []petrinet.MarkInt, error) {
	logger = log.New(os.Stdout, "[Hello] ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	logger.SetOutput(io.Discard)
	is, err := antlr.NewFileStream(fileName)
	if err != nil {
		return nil, make([]petrinet.MarkInt, 0), err
	}
	lexer := NewJSPNLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewJSPNLParser(stream)
	listener := NewPNListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Prog())
	net, imark := makeNet(listener.builder.labels, listener.builder.env)
	return net, imark, nil
}
