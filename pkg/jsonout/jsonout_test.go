package jsonout

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/okamumu/gospn/pkg/result"
)

func sample() *result.Result {
	r := &result.Result{}
	// 2x2 with entries (0,0)=1.5 and (1,1)=2.5, CSC 0-origin.
	r.AddSparse("A", []int32{2, 2}, 2, []int{0, 1}, []int{0, 1, 2}, []float64{1.5, 2.5})
	r.AddDense("initA", []float64{1, 0})
	r.AddDense("count", []int32{7})
	r.AddText("gospn_version", "1.2.3")
	return r
}

func TestWriteGolden(t *testing.T) {
	var buf bytes.Buffer
	if err := Write(&buf, sample()); err != nil {
		t.Fatal(err)
	}
	const want = `{
  "format": "gospn-result",
  "version": 1,
  "elements": [
    {
      "name": "A",
      "kind": "sparse",
      "dims": [
        2,
        2
      ],
      "dtype": "float64",
      "nnz": 2,
      "rowind": [
        0,
        1
      ],
      "colptr": [
        0,
        1,
        2
      ],
      "values": [
        1.5,
        2.5
      ]
    },
    {
      "name": "initA",
      "kind": "dense",
      "dims": [
        2
      ],
      "dtype": "float64",
      "values": [
        1,
        0
      ]
    },
    {
      "name": "count",
      "kind": "dense",
      "dims": [
        1
      ],
      "dtype": "int32",
      "values": [
        7
      ]
    },
    {
      "name": "gospn_version",
      "kind": "text",
      "text": "1.2.3"
    }
  ]
}
`
	if got := buf.String(); got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
}

func TestWriteIsValidJSON(t *testing.T) {
	var buf bytes.Buffer
	if err := Write(&buf, sample()); err != nil {
		t.Fatal(err)
	}
	var doc struct {
		Elements []struct{ Name string } `json:"elements"`
	}
	if err := json.Unmarshal(buf.Bytes(), &doc); err != nil {
		t.Fatal(err)
	}
	if len(doc.Elements) != 4 {
		t.Fatalf("got %d elements, want 4", len(doc.Elements))
	}
}

func TestWriteRejectsDuplicateNames(t *testing.T) {
	r := &result.Result{}
	r.AddDense("x", []float64{1})
	r.AddDense("x", []float64{2})
	err := Write(&bytes.Buffer{}, r)
	if err == nil || !strings.Contains(err.Error(), "duplicate") {
		t.Fatalf("got %v, want a duplicate-name error", err)
	}
}
