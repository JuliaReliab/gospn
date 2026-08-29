package result

import "testing"

func TestSortFromOrdersOnlyTheTail(t *testing.T) {
	r := &Result{}
	r.AddDense("z", []float64{1})
	first := r.Len()
	r.AddDense("c", []float64{1})
	r.AddDense("a", []float64{1})
	r.AddDense("b", []float64{1})
	r.SortFrom(first)
	want := []string{"z", "a", "b", "c"}
	for i, e := range r.Elements {
		if e.Name != want[i] {
			t.Fatalf("element %d: got %q, want %q", i, e.Name, want[i])
		}
	}
}

// A duplicate name is silently lossy in every format gospn writes: MATLAB's load keeps
// the last, np.load keeps the last, a JSON object key collides. This is the check that
// would have caught the reward-vector collision fixed in 0.17.0.
func TestValidateRejectsDuplicateNames(t *testing.T) {
	r := &Result{}
	r.AddDense("rwd", []float64{1})
	if err := r.Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	r.AddText("rwd", "x")
	if err := r.Validate(); err == nil {
		t.Fatal("duplicate name accepted")
	}
}

func TestAddDenseRejectsUnsupportedType(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected a panic for []string")
		}
	}()
	(&Result{}).AddDense("bad", []string{"x"})
}
