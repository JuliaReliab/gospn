package mt

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

// This file used to print a thousand numbers and check none of them, which for a random
// number generator means nothing was checked at all: the simulation goldens elsewhere
// pin what gospn produces, not whether the generator is the one it claims to be. A wrong
// InitByArray would make every simulation quietly wrong and move no golden, because the
// goldens would move with it.
//
// The reference output in testdata/ comes from Matsumoto and Nishimura's own C program;
// see testdata/README.md for how it is regenerated.

// numbers reads the reference output: a header line naming what follows, then lines of
// numbers, repeated. Splitting on the header text alone would fold the "1000" of the
// next header into the previous section, so this goes line by line.
func numbers(t *testing.T, file string, sections int) [][]string {
	t.Helper()
	b, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	var out [][]string
	for _, line := range strings.Split(string(b), "\n") {
		if strings.Contains(line, "outputs of ") {
			out = append(out, nil)
			continue
		}
		if len(out) == 0 {
			continue
		}
		out[len(out)-1] = append(out[len(out)-1], strings.Fields(line)...)
	}
	if len(out) != sections {
		t.Fatalf("%s: %d sections, want %d", file, len(out), sections)
	}
	return out
}

// The 1000 integers and the 1000 reals in the reference output come from one continuous
// stream: the C program does not reinitialise between them, so neither does this.
func TestInitByArrayMatchesTheReference(t *testing.T) {
	want := numbers(t, "testdata/mt19937-64.out.txt", 2)
	ints, reals := want[0], want[1]
	if len(ints) != 1000 || len(reals) != 1000 {
		t.Fatalf("reference has %d ints and %d reals, want 1000 each", len(ints), len(reals))
	}

	mt := NewMT64()
	mt.InitByArray([]uint64{0x12345, 0x23456, 0x34567, 0x45678})

	for i, s := range ints {
		w, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		if got := mt.UInt64(); got != w {
			t.Fatalf("UInt64 %d = %d, want %d", i, got, w)
		}
	}
	// Float64 is genrand64_real2: the top 53 bits over 2^53. The reference prints it
	// with %10.8f, so the comparison is against that rendering.
	for i, s := range reals {
		if got := fmt.Sprintf("%10.8f", mt.Float64()); strings.TrimSpace(got) != s {
			t.Fatalf("Float64 %d = %s, want %s", i, strings.TrimSpace(got), s)
		}
	}
}

// Seed is init_genrand64, which the reference main never calls -- `gospn test -s` uses
// it, so it is checked against a run of the same algorithm with only main changed.
func TestSeedMatchesTheReference(t *testing.T) {
	want := numbers(t, "testdata/mt19937-64-seed.out.txt", 1)[0]
	if len(want) != 1000 {
		t.Fatalf("reference has %d ints, want 1000", len(want))
	}

	mt := NewMT64()
	mt.Seed(1234)
	for i, s := range want {
		w, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		if got := mt.UInt64(); got != w {
			t.Fatalf("UInt64 %d = %d, want %d", i, got, w)
		}
	}
}

// Two generators seeded the same way produce the same stream, which is what makes
// `gospn sim -s` reproducible -- and what makes the per-replication streams in RunAll,
// derived with InitByArray([seed, k]), independent of the worker that runs them.
func TestSameSeedSameStream(t *testing.T) {
	a, b := NewMT64(), NewMT64()
	a.InitByArray([]uint64{1234, 7})
	b.InitByArray([]uint64{1234, 7})
	c := NewMT64()
	c.InitByArray([]uint64{1234, 8})

	same, differ := 0, 0
	for i := 0; i < 100; i++ {
		x, y, z := a.UInt64(), b.UInt64(), c.UInt64()
		if x == y {
			same++
		}
		if x != z {
			differ++
		}
	}
	if same != 100 {
		t.Errorf("the same seed gave a different stream at %d of 100 draws", 100-same)
	}
	if differ < 99 {
		t.Errorf("replication 7 and replication 8 agreed on %d of 100 draws", 100-differ)
	}
}
