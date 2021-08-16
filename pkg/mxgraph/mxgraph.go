package mxgraph

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"encoding/xml"
	"io"
	_ "log"
	"net/url"
	_ "os"
	_ "strconv"
	"strings"
)

type mxFile struct {
	XMLName     xml.Name `xml:"mxfile"`
	Diagram     []byte   `xml:"diagram"`
	Description []byte   `xml:",innerxml"`
}

type mxDiagram struct {
	XMLName     xml.Name `xml:"diagram"`
	GraphModel  []byte   `xml:"mxGraphModel"`
	Description []byte   `xml:",innerxml"`
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
	XMLName  xml.Name `xml:"mxCell"`
	Id       string   `xml:"id,attr"`
	Parent   string   `xml:"parent,attr"`
	Value    string   `xml:"value,attr"`
	Style    string   `xml:"style,attr"`
	StyleMap map[string]string
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

func decode(data []byte) ([]byte, error) {
	var b []byte
	var s string
	var err error
	b, err = base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return nil, err
	}
	r := flate.NewReader(bytes.NewReader(b))
	defer r.Close()
	b, err = io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	s, err = url.QueryUnescape(string(b))
	if err != nil {
		return nil, err
	}
	return []byte(s), nil
}

func getCells(data []byte) ([]mxCell, error) {
	model := mxGraphModel{}
	if err := xml.Unmarshal(data, &model); err == nil {
		for i, _ := range model.Root.Cells {
			retval := getStyle(&model.Root.Cells[i])
			model.Root.Cells[i].StyleMap = retval
		}
		return model.Root.Cells, nil
	} else {
		return nil, err
	}
}

func getStyle(c *mxCell) map[string]string {
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

func GetGraphModel(data []byte) ([]mxCell, error) {
	model := mxFile{}
	if err := xml.Unmarshal(data, &model); err == nil {
		diagram := mxDiagram{}
		if s, err := decode(model.Diagram); err == nil {
			return getCells(s)
		} else if err := xml.Unmarshal(model.Description, &diagram); err == nil {
			return getCells(diagram.Description)
		} else {
			return nil, err
		}
	} else {
		return getCells(data)
	}
}
