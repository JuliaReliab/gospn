package mxgraph

import (
	"bytes"
	"fmt"
	_ "html"
	"io"
	_ "log"
	_ "regexp"
	_ "strconv"
	"strings"
)

type Component map[string]interface{}
type Options map[string]string

func (o Options) toString() string {
	s := make([]string, 0, len(o))
	for k, v := range o {
		s = append(s, k+" = "+v)
	}
	return strings.Join(s, ", ")
}

type PetriParser struct{}

func (p *PetriParser) ParseXML(data []byte) ([]byte, error) {
	nodes := make(map[string]MxElement)
	arcs := make(map[string]MxElement)
	labels := make([]MxElement, 0)
	edgelabels := make([]MxElement, 0)
	if elems, err := GetGraphModel(data); err == nil {
		for _, x := range elems {
			switch x.getObj() {
			case "place", "imm", "exp", "gen":
				nodes[x.Id] = x
			case "arc", "harc":
				arcs[x.Id] = x
			case "label":
				labels = append(labels, x)
			case "edgeLabel":
				edgelabels = append(edgelabels, x)
			default:
			}
		}
		// label matching
		matchingNode(nodes, labels)
		for _, x := range edgelabels {
			a, ok := arcs[x.Parent]
			if ok {
				a.Properties["multi"] = x.Value
			}
		}

		var buf bytes.Buffer
		write(&buf, nodes, arcs)
		return buf.Bytes(), nil
	} else {
		return nil, err
	}
}

func matchingNode(nodes map[string]MxElement, labels []MxElement) {
	dists := make(map[string][]float64)
	for i, x := range nodes {
		dists[i] = make([]float64, len(labels))
		for j, y := range labels {
			dists[i][j] = dist(x, y)
		}
	}

	// matching
	for len(dists) > 0 {
		mink := ""
		minj := -1
		smin := 1.0e+200
		for k, x := range dists {
			for j, y := range x {
				if y < smin {
					smin = y
					minj = j
					mink = k
				}
			}
		}
		if smin == 1.0e+200 {
			for k, _ := range dists {
				nodes[k].Properties["label"] = strings.Replace(nodes[k].Id, "-", "_", -1)
			}
			break
		}
		nodes[mink].Properties["label"] = labels[minj].Value
		delete(dists, mink)
		for _, x := range dists {
			x[minj] = 1.0e+200
		}
	}
}

func (c *MxElement) getObj() string {
	switch c.Type {
	case "place", "imm", "exp", "gen":
		return c.Type
	case "vertex":
		return c.getNode()
	case "edge":
		return c.getArc()
	default:
		// log.Println("Skip object id: ", c.Id, " type: ", c.Type)
		return ""
	}
}

func (c *MxElement) getNode() string {
	s := c.Properties
	_, ellipse := s["ellipse"]
	color, fillcolor := s["fillColor"]
	_, dash := s["dashed"]
	_, group := s["group"]
	_, text := s["text"]
	_, edgeLabel := s["edgeLabel"]

	switch {
	case group == true:
		// log.Printf("mxCell %s is detected as Group", c.Id)
		c.Type = "group"
	case text == true:
		// log.Printf("mxCell %s is detected as Label", c.Id)
		c.Type = "label"
	case edgeLabel == true:
		// log.Printf("mxCell %s is detected as EdgeLabel", c.Id)
		c.Type = "edgeLabel"
	case ellipse == true && fillcolor == false:
		// log.Printf("mxCell %s is detected as Place", c.Id)
		c.Type = "place"
	case ellipse == false && group == false && dash == true:
		// log.Printf("mxCell %s is detected as IMM", c.Id)
		c.Type = "imm"
	case ellipse == false && group == false && dash == false && fillcolor == true && color == "#000000":
		// log.Printf("mxCell %s is detected as GEN", c.Id)
		c.Type = "gen"
	case ellipse == false && group == false && dash == false && fillcolor == false:
		// log.Printf("mxCell %s is detected as EXP", c.Id)
		c.Type = "exp"
	default:
		// log.Println("Cannot detect mxCell as a node: id=", c.Id)
		c.Type = ""
	}
	return c.Type
}

func (c *MxElement) getArc() string {
	s := c.Properties
	arrow, endarrow := s["endArrow"]
	src := s["source"]
	dest := s["target"]
	switch {
	case src != "" && dest != "" && endarrow == true && arrow == "oval":
		// log.Printf("mxCell %s is detected as HARC", c.Id)
		c.Type = "harc"
	case src != "" && dest != "":
		// log.Printf("mxCell %s is detected as ARC", c.Id)
		c.Type = "arc"
	default:
		// log.Println("Cannot detect mxCell as an arc: id=", c.Id)
		c.Type = ""
	}
	return c.Type
}

func dist(x MxElement, y MxElement) float64 {
	if x.Parent != y.Parent {
		return 1.0e+200
	} else {
		v1 := []float64{x.Geometry.X, x.Geometry.Y, x.Geometry.Width, x.Geometry.Height}
		v2 := []float64{y.Geometry.X, y.Geometry.Y, y.Geometry.Width, y.Geometry.Height}
		s := 0.0
		s += ((v1[0] + v1[2]/2) - (v2[0] + v2[2]/2)) * ((v1[0] + v1[2]/2) - (v2[0] + v2[2]/2))
		s += ((v1[1] + v1[3]/2) - (v2[1] + v2[3]/2)) * ((v1[1] + v1[3]/2) - (v2[1] + v2[3]/2))
		return s
	}
}

func write(w io.Writer, nodes map[string]MxElement, arcs map[string]MxElement) {
	fmt.Fprintln(w, "// Begin: This part has been generated automatically from XML file.")
	// place
	fmt.Fprintln(w)
	fmt.Fprintln(w, "// place")
	for _, x := range nodes {
		if x.Type == "place" {
			// make label
			label := x.Properties["label"]
			// make options
			options := Options{}
			if x.Properties["init"] != "" {
				options["init"] = x.Properties["init"]
			}
			if x.Value != "" {
				options["init"] = x.Value
			}
			// write
			if len(options) == 0 {
				fmt.Fprintf(w, "place %s\n", label)
			} else {
				fmt.Fprintf(w, "place %s (%s)\n", label, options.toString())
			}
		}
	}

	// trans
	fmt.Fprintln(w)
	fmt.Fprintln(w, "// trans")
	for _, x := range nodes {
		if x.Type == "imm" {
			// make label
			label := x.Properties["label"]
			// make options
			options := Options{}
			if x.Properties["guard"] != "" {
				options["guard"] = x.Properties["guard"]
			}
			// write
			if len(options) == 0 {
				fmt.Fprintf(w, "imm %s\n", label)
			} else {
				fmt.Fprintf(w, "imm %s (%s)\n", label, options.toString())
			}
		}
	}
	for _, x := range nodes {
		if x.Type == "exp" {
			// make label
			label := x.Properties["label"]
			// make options
			options := Options{}
			if x.Properties["guard"] != "" {
				options["guard"] = x.Properties["guard"]
			}
			if x.Properties["rate"] != "" {
				options["rate"] = x.Properties["rate"]
			}
			// write
			if len(options) == 0 {
				fmt.Fprintf(w, "exp %s\n", label)
			} else {
				fmt.Fprintf(w, "exp %s (%s)\n", label, options.toString())
			}
		}
	}
	for _, x := range nodes {
		if x.Type == "gen" {
			// make label
			label := x.Properties["label"]
			// make options
			options := Options{}
			if x.Properties["guard"] != "" {
				options["guard"] = x.Properties["guard"]
			}
			if x.Properties["dist"] != "" {
				options["dist"] = x.Properties["dist"]
			}
			// write
			if len(options) == 0 {
				fmt.Fprintf(w, "gen %s\n", label)
			} else {
				fmt.Fprintf(w, "gen %s (%s)\n", label, options.toString())
			}
		}
	}

	// arc
	fmt.Fprintln(w)
	fmt.Fprintln(w, "// arc")
	for _, x := range arcs {
		switch {
		case x.Type == "arc":
			// make labels
			src := nodes[x.Properties["source"]].Properties["label"]
			dest := nodes[x.Properties["target"]].Properties["label"]
			// make options
			options := Options{}
			if x.Properties["multi"] != "" {
				options["multi"] = x.Properties["multi"]
			}
			// write
			if len(options) == 0 {
				fmt.Fprintf(w, "arc %s to %s\n", src, dest)
			} else {
				fmt.Fprintf(w, "arc %s to %s (%s)\n", src, dest, options.toString())
			}
		case x.Type == "harc":
			// make labels
			src := nodes[x.Properties["source"]].Properties["label"]
			dest := nodes[x.Properties["target"]].Properties["label"]
			// make options
			options := Options{}
			if x.Properties["multi"] != "" {
				options["multi"] = x.Properties["multi"]
			}
			// write
			if len(options) == 0 {
				fmt.Fprintf(w, "harc %s to %s\n", src, dest)
			} else {
				fmt.Fprintf(w, "harc %s to %s (%s)\n", src, dest, options.toString())
			}
		}
	}
	// end
	fmt.Fprintln(w)
	fmt.Fprintln(w, "// End: This part has been generated automatically from XML file.")
}
