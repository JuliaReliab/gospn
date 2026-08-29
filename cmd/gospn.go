package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/okamumu/gospn/pkg/matout"
	"github.com/okamumu/gospn/pkg/mt"
	"github.com/okamumu/gospn/pkg/mxgraph"
	"github.com/okamumu/gospn/pkg/parser"
	"github.com/okamumu/gospn/pkg/petrinet"
	"io"
	"os"
	"time"
)

func usage() {
	msg := `usage: gospn <command> [<args>]

commands: (command help: gospn command -h)
  view    Output a dot file to draw a Petrinet
  mark    Make a marking graph and output matrices
  sim     Monte Carlo simulation
  test    Simulate a path of markings
  gen     Generate Petrinet definition from XML file
  help    Display this message`

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
	case "gen":
		cmdgen(args)
	case "help":
		usage()
	default:
		usage()
	}
}

// reportClamped tells the user when a firing had to be clamped. The result is not exact
// in that case, and nothing else in the output would say so.
func reportClamped(events []petrinet.ClampSummary) {
	if report := petrinet.FormatClampEvents(events); report != "" {
		fmt.Fprint(os.Stderr, "warning: "+report)
	}
}

func cmdview(args []string) {
	infile := flag.String("i", "", "Petrinet definition file")
	outfile := flag.String("o", "", "Output file (dot file)")
	params0 := flag.String("pre", "", "Put a small Petrinet definition like parameters to the beginning of original PN definition")
	params := flag.String("post", "", "Put a small Petrinet definition like parameters to the end of original PN definition")
	flag.CommandLine.Parse(args)

	defs := readDefs(*infile, "view")
	if *params0 != "" {
		defs = *params0 + "\n" + defs + "\n"
	}
	if *params != "" {
		defs = defs + "\n" + *params + "\n"
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

// readDefs returns the net definition, from the file named by -i or from stdin.
//
// A leftover positional argument is an error rather than something to ignore. The
// subcommands take the file with -i, so `gospn mark -o out.mat net.spn` used to drop
// the filename, read an empty stdin, and report a one-state marking graph -- an answer
// that looks like a successful analysis of the wrong thing.
func readDefs(infile string, subcommand string) string {
	if rest := flag.CommandLine.Args(); len(rest) > 0 && infile == "" {
		fmt.Fprintf(os.Stderr,
			"%s: unexpected argument %q. The net is given with -i, as in `gospn %s -i %s`, "+
				"or on standard input.\n", subcommand, rest[0], subcommand, rest[0])
		os.Exit(2)
	}
	if infile != "" {
		b, err := os.ReadFile(infile)
		if err != nil {
			panic(err)
		}
		return string(b)
	}
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func cmdmark(args []string) {
	infile := flag.String("i", "", "Petrinet definition file")
	outfile := flag.String("o", "out.mat", "Nmae of a mat file")
	tangible := flag.Bool("t", false, "Create a (semi) tangible marking")
	maxstates := flag.Int("maxstates", petrinet.DefaultMaxStates, "Stop the reachability search after this many states (0 for no limit)")
	state := flag.String("s", "", "Output a state file")
	markgraph := flag.String("m", "", "Output a dot file to draw the marking graph")
	groupmarkgraph := flag.String("g", "", "Output a dot file to draw the group marking graph")
	params0 := flag.String("pre", "", "Put a small Petrinet definition like parameters to the beginning of original PN definition")
	params := flag.String("post", "", "Put a small Petrinet definition like parameters to the end of original PN definition")
	flag.CommandLine.Parse(args)

	defs := readDefs(*infile, "mark")
	if *params0 != "" {
		defs = *params0 + "\n" + defs + "\n"
	}
	if *params != "" {
		defs = defs + "\n" + *params + "\n"
	}
	net, imark := parser.PNreadFromText(defs)

	// A large search used to print "Create marking..." and then nothing at all until
	// the process was killed. Report the state count as it grows, on stderr so the
	// summary on stdout stays pipeable.
	opts := petrinet.SearchOptions{
		MaxStates: *maxstates,
		Progress: func(states int) {
			fmt.Fprintf(os.Stderr, "\r  %d states...", states)
		},
	}
	fmt.Print("Create marking...")
	var mg *petrinet.MarkingGraph
	var mgerr error
	start := time.Now()
	if *tangible {
		mg, mgerr = petrinet.CreateMarkingGraphWithDFSTangibleOpts(net, imark, opts)
	} else {
		mg, mgerr = petrinet.CreateMarkingGraphWithDFSOpts(net, imark, opts)
	}
	end := time.Now()
	if mgerr != nil {
		fmt.Println()
		fmt.Fprintf(os.Stderr, "\n%v\n", mgerr)
		os.Exit(1)
	}
	fmt.Println("done")
	fmt.Printf("computation time : %.4f (sec)\n", (end.Sub(start)).Seconds())
	reportClamped(mg.ClampEvents())
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
		// fmt.Printf("Write transition matrix %s\n", label)
	}
	for tr, m := range immmat {
		label := fmt.Sprintf("%s%s%s", grouplabel[tr.GetSrc()], grouplabel[tr.GetDest()], grouptranslabel[tr])
		dim, nnz, rowind, colptr, val := m.Get()
		data := matout.CreateMATLABSparseMatrix(dim, label, nnz, rowind, colptr, val)
		matfile.AddElement(data)
		// fmt.Printf("Write transition matrix %s\n", label)
	}
	for tr, m := range genmat {
		label := fmt.Sprintf("%s%s%s", grouplabel[tr.GetSrc()], grouplabel[tr.GetDest()], grouptranslabel[tr])
		dim, nnz, rowind, colptr, val := m.Get()
		data := matout.CreateMATLABSparseMatrix(dim, label, nnz, rowind, colptr, val)
		matfile.AddElement(data)
		// fmt.Printf("Write transition matrix %s\n", label)
	}
	iv := mg.InitVector()
	for g, v := range iv {
		label := fmt.Sprintf("init%s", grouplabel[g])
		data := matout.CreateMATLABMatrix(len(v), label, v)
		matfile.AddElement(data)
		// fmt.Printf("Write init vector %s\n", label)
	}
	rv := mg.RewardVector()
	for rewardlabel, rv := range rv {
		for g, v := range rv {
			label := fmt.Sprintf("%s%s", rewardlabel, grouplabel[g])
			data := matout.CreateMATLABMatrix(len(v), label, v)
			matfile.AddElement(data)
			// fmt.Printf("Write reward vector %s\n", label)
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
	params0 := flag.String("pre", "", "Put a small Petrinet definition like parameters to the beginning of original PN definition")
	params := flag.String("post", "", "Put a small Petrinet definition like parameters to the end of original PN definition")
	seed := flag.Int64("s", 1234, "A seed for random number generator")
	parallel := flag.Int("parallel", 0, "Number of replications to run at once (0 for one per CPU)")
	configfile := flag.String("f", "", "Configuration file for simulation")
	configure := flag.String("c", "", "JSON configuration (text)")
	flag.CommandLine.Parse(args)

	defs := readDefs(*infile, "sim")
	if *params0 != "" {
		defs = *params0 + "\n" + defs + "\n"
	}
	if *params != "" {
		defs = defs + "\n" + *params + "\n"
	}
	net, imark := parser.PNreadFromText(defs)

	var config petrinet.PNSimConfig
	var json []byte
	if *configfile != "" {
		if j, err := os.ReadFile(*configfile); err == nil {
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
	// The flag wins over the configuration file when both name a worker count.
	if *parallel != 0 {
		config.Parallel = *parallel
	}
	sim := petrinet.NewPNSimulation(net, config)

	fmt.Print("Run simulation...")
	start := time.Now()
	irwd, crwd, lastrwd, elapsedtime, count := sim.RunAll(imark, *seed)
	end := time.Now()
	fmt.Println("done")
	fmt.Printf("computation time : %.4f (sec)\n", (end.Sub(start)).Seconds())
	reportClamped(sim.ClampEvents())

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
	params0 := flag.String("pre", "", "Put a small Petrinet definition like parameters to the beginning of original PN definition")
	params := flag.String("post", "", "Put a small Petrinet definition like parameters to the end of original PN definition")
	seed := flag.Int64("s", 1234, "A seed for random number generator")
	elapsedtime := flag.Float64("t", 0.0, "Maximum elapsed time for simulation")
	maxcount := flag.Int("n", 100, "Maximum number of firings for simulation")
	flag.CommandLine.Parse(args)

	defs := readDefs(*infile, "test")
	if *params0 != "" {
		defs = *params0 + "\n" + defs + "\n"
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

func cmdgen(args []string) {
	infile := flag.String("i", "", "XML file for drawing Petrinet")
	outfile := flag.String("o", "", "Output file (spn file)")
	flag.CommandLine.Parse(args)

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
