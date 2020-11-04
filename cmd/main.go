package main

import (
	"../pkg/matout"
	"../pkg/mt"
	"../pkg/parser"
	"../pkg/petrinet"
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	// "math/rand"
	"os"
	"time"
)

func usage() {
	msg := `usage: gospn <command> [<args>]

commands: (command help: gospn command -h)
  view  Output a dot file to draw a Petrinet
  mark  Make a marking graph and output matrices
  sim   Monte Carlo simulation
  test  Simulate a path of markings
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
	case "sim":
		cmdsim(args)
	case "test":
		cmdtest(args)
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

	var defs string
	if *infile != "" {
		if b, err := ioutil.ReadFile(*infile); err == nil {
			defs = string(b)
		} else {
			panic(err)
		}
	} else {
		if b, err := ioutil.ReadAll(os.Stdin); err == nil {
			defs = string(b)
		} else {
			panic(err)
		}
	}
	net, _ := parser.PNreadFromText(defs)

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
	state := flag.String("s", "", "Output a state file")
	markgraph := flag.String("m", "", "Output a dot file to draw the marking graph")
	groupmarkgraph := flag.String("g", "", "Output a dot file to draw the group marking graph")
	params := flag.String("p", "", "Put a small Petrinet definition like parameters to the end of original PN definition")
	flag.CommandLine.Parse(args)

	var defs string
	if *infile != "" {
		if b, err := ioutil.ReadFile(*infile); err == nil {
			defs = string(b)
		} else {
			panic(err)
		}
	} else {
		if b, err := ioutil.ReadAll(os.Stdin); err == nil {
			defs = string(b)
		} else {
			panic(err)
		}
	}
	if *params != "" {
		defs = defs + "\n" + *params + "\n"
	}
	net, imark := parser.PNreadFromText(defs)

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
	mg.Summary()

	// WriteMatrix
	expmat, immmat, genmat := mg.TransMatrix()
	grouplabel := mg.GroupLabels()
	grouptranslabel := mg.TransLabels()
	matfile := matout.CreateMATLABMatFile(true)
	for tr, m := range expmat {
		label := fmt.Sprintf("%s%s%s", grouplabel[tr.GetSrc()], grouplabel[tr.GetDest()], grouptranslabel[tr])
		dim, nnz, rowind, colptr, val := m.Get()
		data := matout.CreateMATLABSparseMatrix(dim, label, nnz, rowind, colptr, val)
		matfile.AddElement(data)
		fmt.Printf("Write transition matrix %s\n", label)
	}
	for tr, m := range immmat {
		label := fmt.Sprintf("%s%s%s", grouplabel[tr.GetSrc()], grouplabel[tr.GetDest()], grouptranslabel[tr])
		dim, nnz, rowind, colptr, val := m.Get()
		data := matout.CreateMATLABSparseMatrix(dim, label, nnz, rowind, colptr, val)
		matfile.AddElement(data)
		fmt.Printf("Write transition matrix %s\n", label)
	}
	for tr, m := range genmat {
		label := fmt.Sprintf("%s%s%s", grouplabel[tr.GetSrc()], grouplabel[tr.GetDest()], grouptranslabel[tr])
		dim, nnz, rowind, colptr, val := m.Get()
		data := matout.CreateMATLABSparseMatrix(dim, label, nnz, rowind, colptr, val)
		matfile.AddElement(data)
		fmt.Printf("Write transition matrix %s\n", label)
	}
	iv := mg.InitVector()
	for g, v := range iv {
		label := fmt.Sprintf("init%s", grouplabel[g])
		data := matout.CreateMATLABMatrix(len(v), label, v)
		matfile.AddElement(data)
		fmt.Printf("Write init vector %s\n", label)
	}
	rv := mg.RewardVector()
	for rewardlabel, rv := range rv {
		for g, v := range rv {
			label := fmt.Sprintf("%s%s", rewardlabel, grouplabel[g])
			data := matout.CreateMATLABMatrix(len(v), label, v)
			matfile.AddElement(data)
			fmt.Printf("Write reward vector %s\n", label)
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

	// Write groupmarking graph
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

	// WriteState
	if *state != "" {
		fmt.Print("Write state file...")
		file, err := os.Create(*state)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		writer := bufio.NewWriter(file)
		mg.WriteState(writer)
		writer.Flush()
		fmt.Println("done")
	}

	// Write marking graph
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

func cmdsim(args []string) {
	infile := flag.String("i", "", "Petrinet definition file")
	outfile := flag.String("o", "out.mat", "Nmae of a mat file")
	params := flag.String("p", "", "Put a small Petrinet definition like parameters to the end of original PN definition")
	seed := flag.Int64("s", 1234, "A seed for random number generator")
	configfile := flag.String("f", "", "Configuration file for simulation")
	configure := flag.String("c", "", "JSON configuration (text)")
	flag.CommandLine.Parse(args)

	var defs string
	if *infile != "" {
		if b, err := ioutil.ReadFile(*infile); err == nil {
			defs = string(b)
		} else {
			panic(err)
		}
	} else {
		if b, err := ioutil.ReadAll(os.Stdin); err == nil {
			defs = string(b)
		} else {
			panic(err)
		}
	}
	if *params != "" {
		defs = defs + "\n" + *params + "\n"
	}
	net, imark := parser.PNreadFromText(defs)

	var config petrinet.PNSimConfig
	var json []byte
	if *configfile != "" {
		if j, err := ioutil.ReadFile(*configfile); err == nil {
			json = j
		} else {
			panic(err)
		}
	} else if *configure != "" {
		json = []byte(*configure)
	} else {
		panic("Configuration JSON was not found")
	}
	if c, err := petrinet.ReadConfigFromJson([]byte(json)); err == nil {
		config = c
	} else {
		panic(err)
	}
	sim := petrinet.NewPNSimulation(net, config)
	rng := mt.NewMT64()
	// rng := rand.New(rand.NewSource(0))
	rng.Seed(*seed)

	fmt.Print("Run simulation...")
	start := time.Now()
	irwd, crwd, lastrwd, elapsedtime, count := sim.RunAll(imark, rng)
	end := time.Now()
	fmt.Println("done")
	fmt.Printf("computation time : %.4f (sec)\n", (end.Sub(start)).Seconds())

	// WriteMatrix
	matfile := matout.CreateMATLABMatFile(true)
	for rlabel, v := range irwd {
		label := fmt.Sprintf("%s_irwd", rlabel)
		data := matout.CreateMATLABMatrix(len(v), label, v)
		matfile.AddElement(data)
	}
	for rlabel, v := range crwd {
		label := fmt.Sprintf("%s_crwd", rlabel)
		data := matout.CreateMATLABMatrix(len(v), label, v)
		matfile.AddElement(data)
	}
	for rlabel, v := range lastrwd {
		label := fmt.Sprintf("%s_irwd", rlabel)
		data := matout.CreateMATLABMatrix(len(v), label, v)
		matfile.AddElement(data)
	}
	matfile.AddElement(matout.CreateMATLABMatrix(len(elapsedtime), "elapsedtime", elapsedtime))
	matfile.AddElement(matout.CreateMATLABMatrix(len(count), "count", count))

	mfile, err := os.Create(*outfile)
	if err != nil {
		panic(err)
	}
	defer mfile.Close()
	writer := bufio.NewWriter(mfile)
	matfile.ToBytes(matout.NewMATLABBuffer(writer, binary.LittleEndian))
	writer.Flush()
}

func cmdtest(args []string) {
	infile := flag.String("i", "", "Petrinet definition file")
	params := flag.String("p", "", "Put a small Petrinet definition like parameters to the end of original PN definition")
	seed := flag.Int64("s", 1234, "A seed for random number generator")
	elapsedtime := flag.Float64("t", 0.0, "Maximum elapsed time for simulation")
	maxcount := flag.Int("n", 100, "Maximum number of firings for simulation")
	flag.CommandLine.Parse(args)

	var defs string
	if *infile != "" {
		if b, err := ioutil.ReadFile(*infile); err == nil {
			defs = string(b)
		} else {
			panic(err)
		}
	} else {
		if b, err := ioutil.ReadAll(os.Stdin); err == nil {
			defs = string(b)
		} else {
			panic(err)
		}
	}
	if *params != "" {
		defs = defs + "\n" + *params + "\n"
	}
	net, imark := parser.PNreadFromText(defs)

	config := petrinet.PNSimConfig{
		EndingTime:  *elapsedtime,
		NumOfFiring: int32(*maxcount),
	}
	sim := petrinet.NewPNSimulation(net, config)
	rng := mt.NewMT64()
	// rng := rand.New(rand.NewSource(0))
	rng.Seed(*seed)
	path, _, _ := sim.RunSimulation(imark, rng)
	for i, x := range path {
		fmt.Println(i, x.String(net))
	}
}
