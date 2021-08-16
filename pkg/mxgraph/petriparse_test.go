package mxgraph

import (
	"fmt"
	"testing"
)

func TestPetri01(t *testing.T) {
	data := []byte(`
<mxGraphModel dx="1773" dy="646" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="1" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0">
  <root>
    <mxCell id="0" />
    <mxCell id="1" parent="0" />
    <mxCell id="tUs784MvHepMPJK3DbLS-9" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-1" target="tUs784MvHepMPJK3DbLS-3" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-1" value="1" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;container=0;" parent="1" vertex="1">
      <mxGeometry x="20" y="40" width="30" height="30" as="geometry" />
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
    <mxCell id="l5spLULfT4fBQ-QFcXco-3" value="#Pnormal" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="tUs784MvHepMPJK3DbLS-12" vertex="1" connectable="0">
      <mxGeometry x="-0.2596" relative="1" as="geometry">
        <mxPoint x="-1" y="13" as="offset" />
      </mxGeometry>
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-5" value="" style="group" parent="1" vertex="1" connectable="0">
      <mxGeometry x="145" y="150" width="70" height="85" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-6" value="" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;fillColor=#000000;" parent="tUs784MvHepMPJK3DbLS-5" vertex="1">
      <mxGeometry y="20" width="50" height="10" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-4" value="Treju&lt;br&gt;[#Pnormal &amp;gt;= 2]" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" parent="tUs784MvHepMPJK3DbLS-5" vertex="1">
      <mxGeometry x="-30" y="55" width="100" height="30" as="geometry" />
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
    <mxCell id="l5spLULfT4fBQ-QFcXco-1" value="Pnormal" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" parent="1" vertex="1">
      <mxGeometry x="-5" y="10" width="60" height="20" as="geometry" />
    </mxCell>
    <mxCell id="l5spLULfT4fBQ-QFcXco-2" value="Pfailure" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" parent="1" vertex="1">
      <mxGeometry x="270" y="120" width="60" height="20" as="geometry" />
    </mxCell>
    <mxCell id="l5spLULfT4fBQ-QFcXco-6" value="trepair" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" parent="1" vertex="1">
      <mxGeometry x="180" width="50" height="20" as="geometry" />
    </mxCell>
    <mxCell id="l5spLULfT4fBQ-QFcXco-7" value="Tfail" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" parent="1" vertex="1">
      <mxGeometry x="120" y="110" width="40" height="20" as="geometry" />
    </mxCell>
  </root>
</mxGraphModel>
`)
	p := &PetriParser{}
	b, err := p.ParseXML(data)
	if err != nil {
		t.Error("error")
	}
	fmt.Println(string(b))
}

func TestPetri02(t *testing.T) {
	data := []byte(`
<mxGraphModel dx="1865" dy="810" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="1" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0">
  <root>
    <mxCell id="0" />
    <mxCell id="1" parent="0" />
    <mxCell id="tUs784MvHepMPJK3DbLS-9" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-1" target="tUs784MvHepMPJK3DbLS-3" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-1" value="1" style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;container=0;" parent="1" vertex="1">
      <mxGeometry x="20" y="40" width="30" height="30" as="geometry" />
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
    <mxCell id="l5spLULfT4fBQ-QFcXco-3" value="#Pnormal" style="edgeLabel;html=1;align=center;verticalAlign=middle;resizable=0;points=[];" parent="tUs784MvHepMPJK3DbLS-12" vertex="1" connectable="0">
      <mxGeometry x="-0.2596" relative="1" as="geometry">
        <mxPoint x="-1" y="13" as="offset" />
      </mxGeometry>
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-5" value="" style="group" parent="1" vertex="1" connectable="0">
      <mxGeometry x="145" y="150" width="95" height="85" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-6" value="" style="rounded=0;whiteSpace=wrap;html=1;rotation=90;fillColor=#000000;" parent="tUs784MvHepMPJK3DbLS-5" vertex="1">
      <mxGeometry y="20" width="50" height="10" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-4" value="Treju&lt;br&gt;[#Pnormal &amp;gt;= 2, dist: unif]" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" parent="tUs784MvHepMPJK3DbLS-5" vertex="1">
      <mxGeometry x="-55" y="55" width="150" height="30" as="geometry" />
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
    <mxCell id="l5spLULfT4fBQ-QFcXco-1" value="Pnormal" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" parent="1" vertex="1">
      <mxGeometry x="-5" y="10" width="60" height="20" as="geometry" />
    </mxCell>
    <mxCell id="l5spLULfT4fBQ-QFcXco-2" value="Pfailure" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" parent="1" vertex="1">
      <mxGeometry x="270" y="120" width="60" height="20" as="geometry" />
    </mxCell>
    <mxCell id="l5spLULfT4fBQ-QFcXco-6" value="trepair" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" parent="1" vertex="1">
      <mxGeometry x="180" width="50" height="20" as="geometry" />
    </mxCell>
    <mxCell id="l5spLULfT4fBQ-QFcXco-7" value="Tfail&lt;br&gt;[rate: 10.0]" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" parent="1" vertex="1">
      <mxGeometry x="105" y="105" width="70" height="30" as="geometry" />
    </mxCell>
  </root>
</mxGraphModel>
`)
	p := &PetriParser{}
	b, err := p.ParseXML(data)
	if err != nil {
		t.Error("error")
	}
	fmt.Println(string(b))
}
