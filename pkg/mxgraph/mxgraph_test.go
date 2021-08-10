package mxgraph

import (
	"fmt"
	"testing"
)

type TestParser struct{}

func (p *TestParser) GetNode(c *mxCell, s map[string]string) (Component, bool) {
	result := Component{
		"id":       c.Id,
		"parent":   c.Parent,
		"value":    c.Value,
		"geometry": []float64{c.Geometry.X, c.Geometry.Y, c.Geometry.Width, c.Geometry.Height},
	}
	for k, v := range s {
		result[k] = v
	}
	return result, true
}

func (p *TestParser) GetArc(c *mxCell, s map[string]string) (Component, bool) {
	result := Component{
		"id":     c.Id,
		"parent": c.Parent,
		"value":  c.Value,
		"source": c.Source,
		"target": c.Target,
	}
	for k, v := range s {
		result[k] = v
	}
	return result, true
}

func TestMxgraph01(t *testing.T) {
	data := []byte(`
<mxGraphModel dx="1257" dy="789" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="1" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0">
  <root>
    <mxCell id="0" />
    <mxCell id="1" parent="0" />
    <mxCell id="tUs784MvHepMPJK3DbLS-9" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-1" target="tUs784MvHepMPJK3DbLS-3" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-1" value="" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;container=1;" parent="1" vertex="1">
      <mxGeometry x="10" y="30" width="110" height="110" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-7" value="" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;fillColor=#000000;" vertex="1" parent="tUs784MvHepMPJK3DbLS-1">
      <mxGeometry x="30" y="30" width="20" height="20" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-11" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-2" target="tUs784MvHepMPJK3DbLS-6" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-19" style="edgeStyle=orthogonalEdgeStyle;rounded=1;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=0;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;" edge="1" parent="1" source="tUs784MvHepMPJK3DbLS-2" target="tUs784MvHepMPJK3DbLS-18">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-2" value="" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;" parent="1" vertex="1">
      <mxGeometry x="230" y="100" width="40" height="40" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-10" style="edgeStyle=orthogonalEdgeStyle;rounded=1;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=0;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-3" target="tUs784MvHepMPJK3DbLS-2" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-3" value="" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;" parent="1" vertex="1">
      <mxGeometry x="145" y="95" width="50" height="10" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-12" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-6" target="tUs784MvHepMPJK3DbLS-1" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-5" value="" style="group" vertex="1" connectable="0" parent="1">
      <mxGeometry x="145" y="150" width="40" height="80" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-6" value="tttttt" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;fillColor=#000000;" parent="tUs784MvHepMPJK3DbLS-5" vertex="1">
      <mxGeometry y="20" width="50" height="10" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-20" style="edgeStyle=orthogonalEdgeStyle;rounded=1;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=1;entryY=0;entryDx=0;entryDy=0;" edge="1" parent="1" source="tUs784MvHepMPJK3DbLS-18" target="tUs784MvHepMPJK3DbLS-1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-18" value="" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;dashed=1;" vertex="1" parent="1">
      <mxGeometry x="145" y="30" width="50" height="10" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-4" value="Treju" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" vertex="1" parent="1">
      <mxGeometry x="310" y="320" width="40" height="20" as="geometry" />
    </mxCell>
  </root>
</mxGraphModel>
`)
	result := GetCells(data)
	for _, x := range result {
		fmt.Println(x)
	}
}

func TestMxgraph02(t *testing.T) {
	data := []byte(`
<mxGraphModel dx="1286" dy="809" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="1" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0">
  <root>
    <mxCell id="0" />
    <mxCell id="1" parent="0" />
    <mxCell id="tUs784MvHepMPJK3DbLS-9" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-1" target="tUs784MvHepMPJK3DbLS-3" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-1" value="" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;container=1;" parent="1" vertex="1">
      <mxGeometry x="10" y="30" width="110" height="110" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-7" value="" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;fillColor=#000000;" parent="tUs784MvHepMPJK3DbLS-1" vertex="1">
      <mxGeometry x="30" y="30" width="20" height="20" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-11" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-2" target="tUs784MvHepMPJK3DbLS-6" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-2" value="" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;" parent="1" vertex="1">
      <mxGeometry x="230" y="100" width="40" height="40" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-10" style="edgeStyle=orthogonalEdgeStyle;rounded=1;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=0;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-3" target="tUs784MvHepMPJK3DbLS-2" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-3" value="" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;" parent="1" vertex="1">
      <mxGeometry x="145" y="95" width="50" height="10" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-12" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-6" target="tUs784MvHepMPJK3DbLS-1" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-5" value="" style="group" parent="1" vertex="1" connectable="0">
      <mxGeometry x="145" y="150" width="40" height="80" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-6" value="tttttt" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;fillColor=#000000;" parent="tUs784MvHepMPJK3DbLS-5" vertex="1">
      <mxGeometry y="20" width="50" height="10" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-4" value="Treju" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" parent="tUs784MvHepMPJK3DbLS-5" vertex="1">
      <mxGeometry y="60" width="40" height="20" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-20" style="edgeStyle=orthogonalEdgeStyle;rounded=1;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=1;entryY=0;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-18" target="tUs784MvHepMPJK3DbLS-1" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-18" value="" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;dashed=1;" parent="1" vertex="1">
      <mxGeometry x="145" y="30" width="50" height="10" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-19" style="edgeStyle=orthogonalEdgeStyle;rounded=1;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=0;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;endArrow=oval;endFill=0;" parent="1" source="tUs784MvHepMPJK3DbLS-2" target="tUs784MvHepMPJK3DbLS-18" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
  </root>
</mxGraphModel>
`)
	result := ParseXML(data, &TestParser{})
	for _, x := range result {
		fmt.Println(x)
	}
}
