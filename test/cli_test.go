package test

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

// The command layer had no test at all: every path through cmd/gospn.go -- the flags,
// the format resolution, the exit statuses, the messages a bad definition produces --
// was verified by running it by hand, if at all. These build the binary once and run it.

var (
	buildOnce sync.Once
	binPath   string
	buildErr  error
)

// gospn returns the path to a freshly built binary. It is built once per test run, into
// a directory the go tool cleans up, so nothing is left in bin/.
func gospn(t *testing.T) string {
	t.Helper()
	buildOnce.Do(func() {
		dir, err := os.MkdirTemp("", "gospn-cli")
		if err != nil {
			buildErr = err
			return
		}
		binPath = filepath.Join(dir, "gospn")
		cmd := exec.Command("go", "build", "-o", binPath, "../cmd/gospn.go")
		if out, err := cmd.CombinedOutput(); err != nil {
			buildErr = err
			binPath = string(out)
		}
	})
	if buildErr != nil {
		t.Fatalf("building gospn: %v\n%s", buildErr, binPath)
	}
	return binPath
}

type cliResult struct {
	stdout, stderr string
	code           int
}

func (r cliResult) out() string { return r.stdout + r.stderr }

func run(t *testing.T, stdin string, args ...string) cliResult {
	t.Helper()
	cmd := exec.Command(gospn(t), args...)
	if stdin != "" {
		cmd.Stdin = strings.NewReader(stdin)
	}
	var out, errb bytes.Buffer
	cmd.Stdout, cmd.Stderr = &out, &errb
	err := cmd.Run()
	code := 0
	if ee, ok := err.(*exec.ExitError); ok {
		code = ee.ExitCode()
	} else if err != nil {
		t.Fatalf("running gospn %v: %v", args, err)
	}
	return cliResult{out.String(), errb.String(), code}
}

// A small net, written where the test can point at it.
func tinyNet(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "tiny.spn")
	def := "place P (init = 2)\nplace Q\nexp T (rate = 1.5)\narc P to T\narc T to Q\nreward n #Q\n"
	if err := os.WriteFile(path, []byte(def), 0644); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestCLIMarkWritesAResult(t *testing.T) {
	out := filepath.Join(t.TempDir(), "out.mat")
	r := run(t, "", "mark", "-i", tinyNet(t), "-o", out)
	if r.code != 0 {
		t.Fatalf("exit %d: %s", r.code, r.out())
	}
	if fi, err := os.Stat(out); err != nil || fi.Size() == 0 {
		t.Errorf("no result file: %v", err)
	}
	if !strings.Contains(r.stdout, "of total states") {
		t.Errorf("the summary is missing from stdout: %q", r.stdout)
	}
}

// -format defaults to the extension of -o, and to mat when that says nothing. This is
// what keeps every invocation from before 0.18.0 working.
func TestCLIFormatFollowsTheExtension(t *testing.T) {
	net := tinyNet(t)
	for _, tc := range []struct {
		name, file, format string
		check              func(t *testing.T, b []byte)
	}{
		{"mat by extension", "out.mat", "", wantMAT},
		{"npz by extension", "out.npz", "", wantNPZ},
		{"json by extension", "out.json", "", wantJSON},
		{"mat when the extension says nothing", "out.dat", "", wantMAT},
		{"the flag wins over the extension", "out.mat", "json", wantJSON},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			out := filepath.Join(t.TempDir(), tc.file)
			args := []string{"mark", "-i", net, "-o", out}
			if tc.format != "" {
				args = append(args, "-format", tc.format)
			}
			if r := run(t, "", args...); r.code != 0 {
				t.Fatalf("exit %d: %s", r.code, r.out())
			}
			b, err := os.ReadFile(out)
			if err != nil {
				t.Fatal(err)
			}
			tc.check(t, b)
		})
	}
}

func wantMAT(t *testing.T, b []byte) {
	t.Helper()
	// A MATLAB v5 file opens with a 116-byte text header.
	if !bytes.Contains(b[:116], []byte("MATLAB 5.0 MAT-file")) {
		t.Errorf("not a MATLAB v5 file: %q", b[:40])
	}
}

func wantNPZ(t *testing.T, b []byte) {
	t.Helper()
	if _, err := zip.NewReader(bytes.NewReader(b), int64(len(b))); err != nil {
		t.Errorf("not a zip archive: %v", err)
	}
}

func wantJSON(t *testing.T, b []byte) {
	t.Helper()
	var doc struct {
		Format string `json:"format"`
	}
	if err := json.Unmarshal(b, &doc); err != nil {
		t.Fatalf("not JSON: %v", err)
	}
	if doc.Format != "gospn-result" {
		t.Errorf("format is %q", doc.Format)
	}
}

func TestCLIUnknownFormatIsRefused(t *testing.T) {
	out := filepath.Join(t.TempDir(), "out.mat")
	r := run(t, "", "mark", "-i", tinyNet(t), "-format", "hdf5", "-o", out)
	if r.code == 0 {
		t.Fatal("an unknown format was accepted")
	}
	if !strings.Contains(r.out(), "unknown output format") {
		t.Errorf("the message is %q", r.out())
	}
}

// `gospn mark -o out.mat net.spn` used to read an empty stdin and report a one-state
// marking graph, which looks like a successful analysis.
func TestCLIPositionalNetIsRefused(t *testing.T) {
	out := filepath.Join(t.TempDir(), "out.mat")
	r := run(t, "", "mark", "-o", out, tinyNet(t))
	if r.code != 2 {
		t.Fatalf("exit %d, want 2: %s", r.code, r.out())
	}
	if !strings.Contains(r.out(), "-i") {
		t.Errorf("the message does not point at -i: %q", r.out())
	}
	if _, err := os.Stat(out); err == nil {
		t.Error("a result file was written for a refused invocation")
	}
}

// Without -i the definition comes from standard input.
func TestCLIReadsStandardInput(t *testing.T) {
	def, err := os.ReadFile(tinyNet(t))
	if err != nil {
		t.Fatal(err)
	}
	out := filepath.Join(t.TempDir(), "out.json")
	r := run(t, string(def), "mark", "-o", out)
	if r.code != 0 {
		t.Fatalf("exit %d: %s", r.code, r.out())
	}
	b, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	wantJSON(t, b)
}

// The messages a bad definition produces. Each of these used to be a Go panic with a
// stack trace; they are the reason this file exists.
func TestCLIRejectsABadDefinition(t *testing.T) {
	for _, tc := range []struct {
		name, def, want string
	}{
		{
			"syntax error",
			"place P (init = 1)\nimm T + (guard = g)\n",
			"syntax error",
		},
		{
			"undefined variable",
			"place P (init = 1)\nexp T (rate = nowhere)\narc P to T\n",
			"the net does not evaluate",
		},
		{
			"an arc to something that is not a transition",
			"place P (init = 1)\nexp T (rate = 1.0)\narc P to Q\n",
			"not a trans",
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			out := filepath.Join(t.TempDir(), "out.mat")
			r := run(t, tc.def, "mark", "-o", out)
			if r.code == 0 {
				t.Fatalf("accepted: %s", r.out())
			}
			if !strings.Contains(r.out(), tc.want) {
				t.Errorf("the message is %q, want it to mention %q", r.out(), tc.want)
			}
			if strings.Contains(r.out(), "goroutine ") {
				t.Errorf("the message carries a stack trace:\n%s", r.out())
			}
		})
	}
}

// A search that hits the state limit reports it and writes nothing: matrices from a
// truncated marking graph look valid and are not.
func TestCLIStateLimitIsReported(t *testing.T) {
	out := filepath.Join(t.TempDir(), "out.mat")
	r := run(t, "", "mark", "-i", "../example/k8s.spn", "-maxstates", "500", "-o", out)
	if r.code == 0 {
		t.Fatal("a truncated search reported success")
	}
	if !strings.Contains(r.out(), "at the limit of 500") {
		t.Errorf("the message is %q", r.out())
	}
	if _, err := os.Stat(out); err == nil {
		t.Error("a result file was written for a truncated search")
	}
}

func TestCLISimAndTest(t *testing.T) {
	net := tinyNet(t)

	simOut := filepath.Join(t.TempDir(), "sim.json")
	r := run(t, "", "sim", "-i", net, "-c", `{"time":10,"simulations":5,"rewards":["n"]}`, "-o", simOut)
	if r.code != 0 {
		t.Fatalf("sim: exit %d: %s", r.code, r.out())
	}
	names := elementNames(t, simOut)
	for _, want := range []string{"n_irwd", "n_crwd", "n_lastrwd", "seed", "simulations"} {
		if !names[want] {
			t.Errorf("sim wrote no %q; it wrote %v", want, keysOf(names))
		}
	}

	testOut := filepath.Join(t.TempDir(), "path.json")
	r = run(t, "", "test", "-i", net, "-n", "3", "-o", testOut)
	if r.code != 0 {
		t.Fatalf("test: exit %d: %s", r.code, r.out())
	}
	names = elementNames(t, testOut)
	for _, want := range []string{"place", "path_time", "path_trans", "path_state"} {
		if !names[want] {
			t.Errorf("test wrote no %q; it wrote %v", want, keysOf(names))
		}
	}
	if !strings.Contains(r.stdout, "P:2") {
		t.Errorf("the path is missing from stdout: %q", r.stdout)
	}
}

func TestCLIViewAndGen(t *testing.T) {
	dot := filepath.Join(t.TempDir(), "net.dot")
	if r := run(t, "", "view", "-i", tinyNet(t), "-o", dot); r.code != 0 {
		t.Fatalf("view: exit %d: %s", r.code, r.out())
	}
	b, err := os.ReadFile(dot)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(b, []byte("digraph")) {
		t.Errorf("view did not write a dot file: %q", b[:min(40, len(b))])
	}

	spn := filepath.Join(t.TempDir(), "from-diagram.spn")
	if r := run(t, "", "gen", "-i", "data/test.drawio", "-o", spn); r.code != 0 {
		t.Fatalf("gen: exit %d: %s", r.code, r.out())
	}
	b, err = os.ReadFile(spn)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(b, []byte("place")) {
		t.Errorf("gen wrote no places:\n%s", b)
	}
}

// `gospn` on its own used to panic with an index-out-of-range.
func TestCLINoArgumentsAndUnknownCommand(t *testing.T) {
	r := run(t, "")
	if r.code == 0 || !strings.Contains(r.out(), "usage") {
		t.Errorf("bare `gospn`: exit %d, output %q", r.code, r.out())
	}
	if strings.Contains(r.out(), "goroutine ") {
		t.Errorf("bare `gospn` panicked:\n%s", r.out())
	}

	r = run(t, "", "marc")
	if r.code == 0 || !strings.Contains(r.out(), "unknown command") {
		t.Errorf("`gospn marc`: exit %d, output %q", r.code, r.out())
	}

	if r := run(t, "", "help"); r.code != 0 || !strings.Contains(r.stdout, "usage") {
		t.Errorf("`gospn help`: exit %d, output %q", r.code, r.stdout)
	}
}

func elementNames(t *testing.T, path string) map[string]bool {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var doc struct {
		Elements []struct {
			Name string `json:"name"`
		} `json:"elements"`
	}
	if err := json.Unmarshal(b, &doc); err != nil {
		t.Fatal(err)
	}
	names := make(map[string]bool, len(doc.Elements))
	for _, e := range doc.Elements {
		names[e.Name] = true
	}
	return names
}

func keysOf(m map[string]bool) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
