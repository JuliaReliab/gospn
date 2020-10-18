package main

import (
	"../pkg/matout"
	"../pkg/parser"
	"../pkg/petrinet"
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func usage() {
	msg := `usage: gospn <command> [<args>]

commands: (command help: gospn command -h)
  view  Output a dot file to draw a Petrinet
  mark  Make a marking graph and output matrices
  sim   Monte Carlo simulation (not yet)
  help  Display this message`

	fmt.Println(msg)
}

func main() {
	mode := os.Args[1]
	args := os.Args[2:]
	switch mode {
	case "view":
		cmdview(args)
	case "mark":
		cmdmark(args)
	case "help":
		usage()
	default:
		usage()
	}
}

func cmdview(args []string) {
	infile := flag.String("i", "", "Petrinet definition file")
	outfile := flag.String("o", "", "Output file (dot file)")
	flag.CommandLine.Parse(args)
	var net *petrinet.Net
	if *infile != "" {
		if n, _, err := parser.PNreadFromFile(*infile); err == nil {
			net = n
		} else {
			panic(err)
		}
	} else {
		if b, err := ioutil.ReadAll(os.Stdin); err == nil {
			net, _ = parser.PNreadFromText(string(b))
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
		net.ToPNDot(writer)
		writer.Flush()
	} else {
		writer := bufio.NewWriter(os.Stdout)
		net.ToPNDot(writer)
		writer.Flush()
	}
}

func cmdmark(args []string) {
	infile := flag.String("i", "", "Petrinet definition file")
	outfile := flag.String("o", "out.mat", "Nmae of a mat file")
	tangible := flag.Bool("t", false, "Create a (semi) tangible marking")
	markgraph := flag.String("m", "", "Output a dot file to draw the marking graph")
	groupmarkgraph := flag.String("g", "", "OUtput a dot file to draw the group marking graph")
	flag.CommandLine.Parse(args)

	var net *petrinet.Net
	var imark []petrinet.MarkInt
	if *infile != "" {
		if n, i, err := parser.PNreadFromFile(*infile); err == nil {
			net, imark = n, i
		} else {
			panic(err)
		}
	} else {
		if b, err := ioutil.ReadAll(os.Stdin); err == nil {
			net, imark = parser.PNreadFromText(string(b))
		} else {
			panic(err)
		}
	}

	fmt.Print("Create marking...")
	var mg *petrinet.MarkingGraph
	start := time.Now()
	if *tangible {
		mg = petrinet.CreateMarkingGraphWithDFSTangible(net, imark)
	} else {
		mg = petrinet.CreateMarkingGraphWithDFS(net, imark)
	}
	end := time.Now()
	fmt.Println("done")
	fmt.Printf("computation time : %.4f (sec)\n", (end.Sub(start)).Seconds())
	mg.Print()

	expmat, immmat, genmat := mg.TransMatrix()
	grouplabel := mg.GroupLabels()
	grouptranslabel := mg.TransLabels()
	matfile := matout.CreateMATLABMatFile(true)
	for tr, m := range expmat {
		label := fmt.Sprintf("%s%s%s", grouplabel[tr.GetSrc()], grouplabel[tr.GetDest()], grouptranslabel[tr])
		dim, nnz, rowind, colptr, val := m.Get()
		data := matout.CreateMATLABSparseMatrix(dim, label, nnz, rowind, colptr, val)
		matfile.AddElement(data)
	}
	for tr, m := range immmat {
		label := fmt.Sprintf("%s%s%s", grouplabel[tr.GetSrc()], grouplabel[tr.GetDest()], grouptranslabel[tr])
		dim, nnz, rowind, colptr, val := m.Get()
		data := matout.CreateMATLABSparseMatrix(dim, label, nnz, rowind, colptr, val)
		matfile.AddElement(data)
	}
	for tr, m := range genmat {
		label := fmt.Sprintf("%s%s%s", grouplabel[tr.GetSrc()], grouplabel[tr.GetDest()], grouptranslabel[tr])
		dim, nnz, rowind, colptr, val := m.Get()
		data := matout.CreateMATLABSparseMatrix(dim, label, nnz, rowind, colptr, val)
		matfile.AddElement(data)
	}
	iv := mg.InitVector()
	for g, v := range iv {
		label := fmt.Sprintf("init%s", grouplabel[g])
		data := matout.CreateMATLABMatrix(len(v), label, v)
		matfile.AddElement(data)
	}
	rv := mg.RewardVector()
	for rewardlabel, rv := range rv {
		for g, v := range rv {
			label := fmt.Sprintf("%s%s", rewardlabel, grouplabel[g])
			data := matout.CreateMATLABMatrix(len(v), label, v)
			matfile.AddElement(data)
		}
	}

	mfile, err := os.Create(*outfile)
	if err != nil {
		panic(err)
	}
	defer mfile.Close()
	writer := bufio.NewWriter(mfile)
	matfile.ToBytes(matout.NewMATLABBuffer(writer, binary.LittleEndian))
	writer.Flush()

	if *groupmarkgraph != "" {
		fmt.Print("Write group marking graph...")
		file, err := os.Create(*groupmarkgraph)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		writer := bufio.NewWriter(file)
		mg.ToGroupMarkDot(writer)
		writer.Flush()
		fmt.Println("done")
	}

	if *markgraph != "" {
		fmt.Print("Write marking graph...")
		file, err := os.Create(*markgraph)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		writer := bufio.NewWriter(file)
		mg.ToMarkDotWithLabel(writer)
		writer.Flush()
		fmt.Println("done")
	}
}
