package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/okamumu/gospn/pkg/mxgraph"
	"io"
	"os"
)

func main() {
	infile := flag.String("i", "", "XML file for drawing Petrinet")
	outfile := flag.String("o", "", "Output file (spn file)")
	flag.Parse()

	var defs string
	if *infile != "" {
		if xml, err := os.ReadFile(*infile); err == nil {
			p := &mxgraph.PetriParser{}
			b, err := p.ParseXML(xml)
			if err != nil {
				panic(err)
			}
			defs = string(b)
		} else {
			panic(err)
		}
	} else {
		if xml, err := io.ReadAll(os.Stdin); err == nil {
			p := &mxgraph.PetriParser{}
			b, err := p.ParseXML(xml)
			if err != nil {
				panic(err)
			}
			defs = string(b)
		} else {
			panic(err)
		}
	}

	if *outfile != "" {
		file, err := os.Create(*outfile)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		writer := bufio.NewWriter(file)
		fmt.Fprint(writer, defs)
		writer.Flush()
	} else {
		writer := bufio.NewWriter(os.Stdout)
		fmt.Fprint(writer, defs)
		writer.Flush()
	}
}
