package mxgraph

import (
	"bytes"
	"fmt"
	"html"
	"io"
	"log"
	"regexp"
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
	comps := make([]Component, 0)
	if cells, err := GetGraphModel(data); err == nil {
		for _, x := range cells {
			if comp, ok := getObj(&x); ok {
				comps = append(comps, comp)
			}
		}
		var buf bytes.Buffer
		write(&buf, comps)
		return buf.Bytes(), nil
	} else {
		return nil, err
	}
}

func getObj(c *mxCell) (Component, bool) {
	switch {
	case c.Vertex:
		return getNode(c)
	case c.Edge:
		return getArc(c)
	default:
		log.Println("Skip mxCell id: ", c.Id)
		return nil, false
	}
}

func getNode(c *mxCell) (Component, bool) {
	s := c.StyleMap
	_, ellipse := s["ellipse"]
	color, fillcolor := s["fillColor"]
	_, dash := s["dashed"]
	_, group := s["group"]
	_, text := s["text"]
	_, edgeLabel := s["edgeLabel"]

	switch {
	case group == true:
		log.Printf("mxCell %s is detected as Group", c.Id)
		return Component{
			"type":     "group",
			"id":       c.Id,
			"parent":   c.Parent,
			"value":    toValue(c.Value),
			"geometry": []float64{c.Geometry.X, c.Geometry.Y, c.Geometry.Width, c.Geometry.Height},
		}, true
	case text == true:
		log.Printf("mxCell %s is detected as Label", c.Id)
		return Component{
			"type":     "label",
			"id":       c.Id,
			"parent":   c.Parent,
			"value":    toValue(c.Value),
			"geometry": []float64{c.Geometry.X, c.Geometry.Y, c.Geometry.Width, c.Geometry.Height},
		}, true
	case edgeLabel == true:
		log.Printf("mxCell %s is detected as EdgeLabel", c.Id)
		return Component{
			"type":   "edgeLabel",
			"id":     c.Id,
			"parent": c.Parent,
			"value":  toValue(c.Value),
		}, true
	case ellipse == true && fillcolor == false:
		log.Printf("mxCell %s is detected as Place", c.Id)
		return Component{
			"type":     "place",
			"id":       c.Id,
			"parent":   c.Parent,
			"value":    toValue(c.Value),
			"geometry": []float64{c.Geometry.X, c.Geometry.Y, c.Geometry.Width, c.Geometry.Height},
		}, true
	case ellipse == false && group == false && dash == true:
		log.Printf("mxCell %s is detected as IMM", c.Id)
		return Component{
			"type":     "imm",
			"id":       c.Id,
			"parent":   c.Parent,
			"value":    toValue(c.Value),
			"geometry": []float64{c.Geometry.X, c.Geometry.Y, c.Geometry.Width, c.Geometry.Height},
		}, true
	case ellipse == false && group == false && dash == false && fillcolor == true && color == "#000000":
		log.Printf("mxCell %s is detected as GEN", c.Id)
		return Component{
			"type":     "gen",
			"id":       c.Id,
			"parent":   c.Parent,
			"value":    toValue(c.Value),
			"geometry": []float64{c.Geometry.X, c.Geometry.Y, c.Geometry.Width, c.Geometry.Height},
		}, true
	case ellipse == false && group == false && dash == false && fillcolor == false:
		log.Printf("mxCell %s is detected as EXP", c.Id)
		return Component{
			"type":     "exp",
			"id":       c.Id,
			"parent":   c.Parent,
			"value":    toValue(c.Value),
			"geometry": []float64{c.Geometry.X, c.Geometry.Y, c.Geometry.Width, c.Geometry.Height},
		}, true
	default:
		log.Println("Cannot detect mxCell as a node: id=", c.Id)
		return nil, false
	}
}

func getArc(c *mxCell) (Component, bool) {
	s := c.StyleMap
	arrow, endarrow := s["endArrow"]
	src := c.Source
	dest := c.Target
	switch {
	case src != "" && dest != "" && endarrow == true && arrow == "oval":
		log.Printf("mxCell %s is detected as HARC", c.Id)
		return Component{
			"type":   "harc",
			"id":     c.Id,
			"parent": c.Parent,
			"value":  c.Value,
			"source": c.Source,
			"target": c.Target,
		}, true
	case src != "" && dest != "":
		log.Printf("mxCell %s is detected as ARC", c.Id)
		return Component{
			"type":   "arc",
			"id":     c.Id,
			"parent": c.Parent,
			"value":  c.Value,
			"source": c.Source,
			"target": c.Target,
		}, true
	default:
		log.Println("Cannot detect mxCell as an arc: id=", c.Id)
		return nil, false
	}
}

func dist(x Component, y Component) float64 {
	v1, ok1 := x["geometry"].([]float64)
	v2, ok2 := y["geometry"].([]float64)
	if ok1 && ok2 {
		s := 0.0
		s += ((v1[0] + v1[2]/2) - (v2[0] + v2[2]/2)) * ((v1[0] + v1[2]/2) - (v2[0] + v2[2]/2))
		s += ((v1[1] + v1[3]/2) - (v2[1] + v2[3]/2)) * ((v1[1] + v1[3]/2) - (v2[1] + v2[3]/2))
		return s
	} else {
		return 1.0e+200
	}
}

func write(w io.Writer, cs []Component) {
	labels := make(map[string]string)

	fmt.Fprintln(w, "// Begin: This part has been generated automatically from XML file.")
	// place
	fmt.Fprintln(w)
	fmt.Fprintln(w, "// place")
	for _, x := range cs {
		id := x["id"].(string)
		if x["type"] == "place" {
			label, options := getLabel(x, cs, "init")
			labels[id] = label
			if i, ok := getIntValue(x); ok {
				options["init"] = i
			}
			if len(options) == 0 {
				fmt.Fprintf(w, "place %s\n", label)
			} else {
				fmt.Fprintf(w, "place %s (%s)\n", label, options.toString())
			}
		}
	}

	// transitions
	fmt.Fprintln(w)
	fmt.Fprintln(w, "// transition")
	for _, x := range cs {
		id := x["id"].(string)
		switch {
		case x["type"] == "imm":
			label, options := getLabel(x, cs, "guard")
			labels[id] = label
			if len(options) == 0 {
				fmt.Fprintf(w, "imm %s\n", label)
			} else {
				fmt.Fprintf(w, "imm %s (%s)\n", label, options.toString())
			}
		case x["type"] == "exp":
			label, options := getLabel(x, cs, "guard")
			labels[id] = label
			if len(options) == 0 {
				fmt.Fprintf(w, "exp %s\n", label)
			} else {
				fmt.Fprintf(w, "exp %s (%s)\n", label, options.toString())
			}
		case x["type"] == "gen":
			label, options := getLabel(x, cs, "guard")
			labels[id] = label
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
	for _, x := range cs {
		switch {
		case x["type"] == "arc":
			if m, one := getMulti(x, cs); one {
				fmt.Fprintf(w, "arc %s to %s\n", labels[x["source"].(string)], labels[x["target"].(string)])
			} else {
				fmt.Fprintf(w, "arc %s to %s (multi = %s)\n", labels[x["source"].(string)], labels[x["target"].(string)], m)
			}
		case x["type"] == "harc":
			if m, one := getMulti(x, cs); one {
				fmt.Fprintf(w, "harc %s to %s\n", labels[x["source"].(string)], labels[x["target"].(string)])
			} else {
				fmt.Fprintf(w, "harc %s to %s (multi = %s)\n", labels[x["source"].(string)], labels[x["target"].(string)], m)
			}
		}
	}

	// end
	fmt.Fprintln(w)
	fmt.Fprintln(w, "// End: This part has been generated automatically from XML file.")
}

func toValue(x string) string {
	a := html.UnescapeString(x)
	b := strings.Replace(a, `<br>`, "", -1)
	c := strings.Replace(b, `\n`, "", -1)
	return c
}

func getLabel(x Component, cs []Component, defaultOptionKey string) (string, Options) {
	value := ""
	mind := 1.0e+200
	for _, y := range cs {
		if x["parent"] == y["parent"] && y["type"] == "label" {
			d := dist(x, y)
			if d < mind {
				value = y["value"].(string)
				mind = d
			}
		}
	}
	tmp := regexp.MustCompile(`\s*([a-zA-Z0-9_\-\.@\+/\*\(\)&%$<>]+)\s*(\[\s*(.+)\s*\])?`).FindStringSubmatch(value)
	label := strings.Replace(tmp[1], " ", "", -1)
	opts := regexp.MustCompile(`[;,]`).Split(tmp[3], -1)
	options := Options{}
	kvregexp := regexp.MustCompile(`\s*(.+)+\s*:\s*(.+)\s*`)
	for _, x := range opts {
		if len(x) == 0 {
			continue
		}
		tmp := kvregexp.FindStringSubmatch(x)
		if len(tmp) == 3 && tmp[1] != "" {
			options[tmp[1]] = tmp[2]
		} else {
			options[defaultOptionKey] = x
		}
	}
	return label, options
}

func getIntValue(x Component) (string, bool) {
	v := x["value"].(string)
	if v == "" {
		return "0", false
	} else {
		return v, true
	}
}

func getMulti(x Component, cs []Component) (string, bool) {
	id := x["id"].(string)
	for _, y := range cs {
		if y["type"] == "edgeLabel" {
			if id == y["parent"].(string) {
				return y["value"].(string), false
			}
		}
	}
	return "", true
}
