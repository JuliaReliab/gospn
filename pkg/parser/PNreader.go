package parser

import (
	"../petrinet"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"io/ioutil"
	"log"
	"os"
)

var logger *log.Logger

func PNreadFromText(text string) (*petrinet.Net, []petrinet.MarkInt) {
	logger = log.New(os.Stdout, "[Hello] ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	logger.SetOutput(ioutil.Discard)
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
	logger.SetOutput(ioutil.Discard)
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
