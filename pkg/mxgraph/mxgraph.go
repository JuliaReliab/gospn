package mxgraph

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"encoding/xml"
	_ "fmt"
	"io"
	"log"
	"net/url"
	_ "os"
	_ "strconv"
	"strings"
)

type mxFile struct {
	XMLName xml.Name `xml:"mxfile"`
	Diagram []byte   `xml:"diagram"`
}

type mxDiagram struct {
	XMLName    xml.Name `xml:"diagram"`
	GraphModel []byte   `xml:"mxGraphModel"`
}

type mxGraphModel struct {
	XMLName xml.Name `xml:"mxGraphModel"`
	Root    mxCells  `xml:"root"`
}

type mxCells struct {
	XMLName xml.Name `xml:"root"`
	Cells   []mxCell `xml:"mxCell"`
}

type mxCell struct {
	XMLName  xml.Name   `xml:"mxCell"`
	Id       string     `xml:"id,attr"`
	Parent   string     `xml:"parent,attr"`
	Value    string     `xml:"value,attr"`
	Style    string     `xml:"style,attr"`
	Vertex   bool       `xml:"vertex,attr"`
	Edge     bool       `xml:"edge,attr"`
	Source   string     `xml:"source,attr"`
	Target   string     `xml:"target,attr"`
	Geometry mxGeometry `xml:"mxGeometry"`
}

type mxGeometry struct {
	XMLName xml.Name `xml:"mxGeometry"`
	X       float64  `xml:"x,attr"`
	Y       float64  `xml:"y,attr"`
	Width   float64  `xml:"width,attr"`
	Height  float64  `xml:"height,attr"`
}

type Component map[string]interface{}

type Parser interface {
	GetNode(*mxCell, map[string]string) (Component, bool)
	GetArc(*mxCell, map[string]string) (Component, bool)
}

func ParseXML(data []byte, p Parser) []Component {
	comps := make([]Component, 0)
	cells := GetCells(data)
	for _, x := range cells {
		if comp, ok := x.GetObj(p); ok {
			comps = append(comps, comp)
		}
	}
	return comps
}

func decode(data string) (string, error) {
	var b []byte
	var s string
	var err error
	b, err = base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	r := flate.NewReader(bytes.NewReader(b))
	defer r.Close()
	b, err = io.ReadAll(r)
	if err != nil {
		return "", err
	}
	s, err = url.QueryUnescape(string(b))
	if err != nil {
		return "", err
	}
	return s, nil
}

func getGraphModel(data []byte) []byte {
	model := mxFile{}
	if err := xml.Unmarshal(data, &model); err == nil {
		diagram := mxDiagram{}
		if err := xml.Unmarshal(model.Diagram, &diagram); err == nil {
			return diagram.GraphModel
		} else {
			if str, err := decode(string(model.Diagram)); err == nil {
				return []byte(str)
			} else {
				log.Fatal(err)
			}
		}
	} else {
		log.Println("Error: XML cannot be read at mxfile. Try to read from mxGraphModel.")
		return data
	}
	return nil
}

func GetCells(data []byte) []mxCell {
	model := mxGraphModel{}
	if err := xml.Unmarshal(getGraphModel(data), &model); err == nil {
		return model.Root.Cells
	} else {
		log.Println("Error: XML cannot be read at mxGraphModel.")
		panic(err)
	}
}

func (c *mxCell) GetObj(p Parser) (Component, bool) {
	s := c.getStyle()
	switch {
	case c.Vertex:
		return p.GetNode(c, s)
	case c.Edge:
		return p.GetArc(c, s)
	default:
		log.Println("Skip mxCell id: ", c.Id)
		return nil, false
	}
}

func (c *mxCell) getStyle() map[string]string {
	styles := make(map[string]string)
	for _, x := range strings.Split(c.Style, ";") {
		e := strings.Split(x, "=")
		switch len(e) {
		case 1:
			styles[e[0]] = ""
		case 2:
			styles[e[0]] = e[1]
		default:
			panic("Error: Style in Cell")
		}
	}
	return styles
}
