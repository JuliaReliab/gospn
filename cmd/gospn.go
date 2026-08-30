package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/okamumu/gospn/pkg/analysis"
	"github.com/okamumu/gospn/pkg/jsonout"
	"github.com/okamumu/gospn/pkg/matout"
	"github.com/okamumu/gospn/pkg/mt"
	"github.com/okamumu/gospn/pkg/mxgraph"
	"github.com/okamumu/gospn/pkg/npzout"
	"github.com/okamumu/gospn/pkg/parser"
	"github.com/okamumu/gospn/pkg/petrinet"
	"github.com/okamumu/gospn/pkg/result"
	"io"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
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
	net, _ := readNet(defs)

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

// version is set at build time with -ldflags "-X main.version=...". A binary built
// without it says so rather than claiming a version it cannot know.
var version = "dev"

// buildRevision returns the commit the binary was built from. The Go toolchain records
// it automatically when building from a checkout, so this needs no build plumbing.
func buildRevision() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}
	rev, dirty := "", false
	for _, s := range info.Settings {
		switch s.Key {
		case "vcs.revision":
			rev = s.Value
		case "vcs.modified":
			dirty = s.Value == "true"
		}
	}
	if rev != "" && dirty {
		return rev + "-dirty"
	}
	return rev
}

// readNet parses a definition, or reports what is wrong with it and stops. A syntax
// error and a mistyped option used to arrive as a Go panic with a stack trace; the
// reader returns them as errors now, and this is where they become a message.
func readNet(defs string) (*petrinet.Net, []petrinet.MarkInt) {
	net, imark, err := parser.PNreadFromText(defs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	return net, imark
}

// checkNet refuses a net whose expressions do not evaluate. Without this the first
// symptom of a typo like `rate = Tvreset.rate` is a Go stack trace after the whole
// reachability search has run.
func checkNet(net *petrinet.Net, imark []petrinet.MarkInt) {
	if err := net.CheckExpressions(imark); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

// formatUsage is the -format flag's help text, shared by the subcommands that write a
// result file.
const formatUsage = "Output format: mat (MATLAB v5), npz (NumPy) or json. Empty means guess from the -o extension, else mat"

// resolveFormat decides the output format. An explicit -format wins; otherwise the
// extension of the output file decides, so `-o out.npz` needs no second flag. The
// default stays mat, which is what every existing invocation gets.
func resolveFormat(format, outfile string) string {
	if format == "" {
		format = strings.TrimPrefix(strings.ToLower(filepath.Ext(outfile)), ".")
	}
	switch format {
	case "mat", "npz", "json":
		return format
	case "":
		return "mat"
	default:
		fmt.Fprintf(os.Stderr, "unknown output format %q (want mat, npz or json)\n", format)
		os.Exit(1)
		return ""
	}
}

// writeResult writes the collected result in the requested format. The subcommands
// collect numbers; deciding what the bytes look like happens only here.
func writeResult(outfile, format string, r *result.Result) {
	f := resolveFormat(format, outfile)
	file, err := os.Create(outfile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	switch f {
	case "mat":
		err = matout.WriteResult(writer, r)
	case "npz":
		err = npzout.Write(writer, r)
	case "json":
		err = jsonout.Write(writer, r)
	}
	if err != nil {
		panic(err)
	}
	if err := writer.Flush(); err != nil {
		panic(err)
	}
}

func cmdmark(args []string) {
	infile := flag.String("i", "", "Petrinet definition file")
	outfile := flag.String("o", "out.mat", "Name of the output file")
	format := flag.String("format", "", formatUsage)
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
	net, imark := readNet(defs)
	checkNet(net, imark)

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

	// Collect what was computed; writeResult decides what the file looks like.
	res := analysis.MarkResult(mg)

	// Until now only `gospn sim` recorded where its numbers came from. A marking graph
	// on disk said nothing about which binary or which net produced it.
	result.Provenance{
		Version:  version,
		Revision: buildRevision(),
		Net:      *infile,
		Command:  "mark",
	}.AddTo(res)

	writeResult(*outfile, *format, res)

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
	outfile := flag.String("o", "out.mat", "Name of the output file")
	format := flag.String("format", "", formatUsage)
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
	net, imark := readNet(defs)
	checkNet(net, imark)

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

	run := analysis.SimRun{
		Irwd:        irwd,
		Crwd:        crwd,
		Lastrwd:     lastrwd,
		ElapsedTime: elapsedtime,
		Count:       count,
		Config:      config,
		Seed:        *seed,
		Clamped:     len(sim.ClampEvents()),
	}
	res := analysis.SimResult(run)

	// What the numbers came from. Without this the file is just vectors: the seed and
	// the horizon are what a run needs to be repeated, and the version matters because
	// 0.16.0 changed which random stream a given seed produces.
	result.Provenance{
		Version:  version,
		Revision: buildRevision(),
		Net:      *infile,
		Command:  "sim",
	}.AddTo(res)
	analysis.SimRunInfo(res, run)

	writeResult(*outfile, *format, res)
}

func cmdtest(args []string) {
	infile := flag.String("i", "", "Petrinet definition file")
	outfile := flag.String("o", "", "Name of the output file (empty writes nothing but the path on stdout)")
	format := flag.String("format", "", formatUsage)
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
	net, imark := readNet(defs)
	checkNet(net, imark)

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

	// The path went to stdout and nowhere else, so this was the one subcommand whose
	// output nothing could read back. Writing a file stays opt-in: -o was not a flag
	// here before, and the printed path is what existing use depends on.
	if *outfile == "" {
		return
	}
	steps := make([]analysis.PathStep, len(path))
	for i, x := range path {
		steps[i] = analysis.PathStep{Time: x.Time(), Mark: x.Mark(), Trans: x.TransLabel()}
	}
	res := analysis.PathResult(net, steps)
	result.Provenance{
		Version:  version,
		Revision: buildRevision(),
		Net:      *infile,
		Command:  "test",
	}.AddTo(res)
	res.AddDense("seed", []int64{*seed})
	writeResult(*outfile, *format, res)
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
