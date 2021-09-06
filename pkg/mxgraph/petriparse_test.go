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

func TestPetri03(t *testing.T) {
	data := []byte(`
<mxGraphModel dx="1865" dy="753" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="1" pageScale="1" pageWidth="827" pageHeight="1169" math="0" shadow="0">
  <root>
    <mxCell id="0" />
    <mxCell id="1" parent="0" />
    <mxCell id="tUs784MvHepMPJK3DbLS-9" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=1;exitY=0.5;exitDx=0;exitDy=0;entryX=0.5;entryY=1;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-1" target="tUs784MvHepMPJK3DbLS-3" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="1" type="place" id="tUs784MvHepMPJK3DbLS-1">
      <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;container=0;" parent="1" vertex="1">
        <mxGeometry x="20" y="40" width="30" height="30" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="tUs784MvHepMPJK3DbLS-11" style="edgeStyle=orthogonalEdgeStyle;rounded=0;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=0.5;entryY=0;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-2" target="tUs784MvHepMPJK3DbLS-6" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="place" id="tUs784MvHepMPJK3DbLS-2">
      <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;" parent="1" vertex="1">
        <mxGeometry x="230" y="100" width="40" height="40" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="tUs784MvHepMPJK3DbLS-10" style="edgeStyle=orthogonalEdgeStyle;rounded=1;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=0;exitDx=0;exitDy=0;entryX=0;entryY=0.5;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-3" target="tUs784MvHepMPJK3DbLS-2" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="exp" rate="10" id="tUs784MvHepMPJK3DbLS-3">
      <mxCell style="rounded=0;whiteSpace=wrap;html=1;rotation=90;" parent="1" vertex="1">
        <mxGeometry x="145" y="95" width="50" height="10" as="geometry" />
      </mxCell>
    </object>
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
    <object label="" type="gen" guard="#Pnormal &gt;= 2" dist="unif" id="tUs784MvHepMPJK3DbLS-6">
      <mxCell style="rounded=0;whiteSpace=wrap;html=1;rotation=90;fillColor=#000000;" parent="tUs784MvHepMPJK3DbLS-5" vertex="1">
        <mxGeometry y="20" width="50" height="10" as="geometry" />
      </mxCell>
    </object>
    <mxCell id="tUs784MvHepMPJK3DbLS-4" value="Treju" style="text;html=1;resizable=0;autosize=1;align=center;verticalAlign=middle;points=[];fillColor=none;strokeColor=none;rounded=0;" parent="tUs784MvHepMPJK3DbLS-5" vertex="1">
      <mxGeometry x="5" y="50" width="40" height="20" as="geometry" />
    </mxCell>
    <mxCell id="tUs784MvHepMPJK3DbLS-20" style="edgeStyle=orthogonalEdgeStyle;rounded=1;orthogonalLoop=1;jettySize=auto;html=1;exitX=0.5;exitY=1;exitDx=0;exitDy=0;entryX=1;entryY=0;entryDx=0;entryDy=0;" parent="1" source="tUs784MvHepMPJK3DbLS-18" target="tUs784MvHepMPJK3DbLS-1" edge="1">
      <mxGeometry relative="1" as="geometry" />
    </mxCell>
    <object label="" type="imm" id="tUs784MvHepMPJK3DbLS-18">
      <mxCell style="rounded=0;whiteSpace=wrap;html=1;rotation=90;strokeWidth=2;" parent="1" vertex="1">
        <mxGeometry x="145" y="30" width="50" height="2" as="geometry" />
      </mxCell>
    </object>
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

func TestPetri04(t *testing.T) {
	data := []byte(`
<mxGraphModel dx="712" dy="776" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="1" pageScale="1" pageWidth="1169" pageHeight="827" background="#ffffff">
  <root>
    <mxCell id="0"/>
    <mxCell id="1" parent="0"/>
    <mxCell id="6" style="edgeStyle=orthogonalEdgeStyle;rounded=0;html=1;exitX=0.5;exitY=0;entryX=0.5;entryY=1;jettySize=auto;orthogonalLoop=1;" parent="1" source="2" target="4" edge="1">
      <mxGeometry relative="1" as="geometry"/>
    </mxCell>
    <object label="" type="place" init="10" id="2">
      <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;" parent="1" vertex="1">
        <mxGeometry x="140" y="90" width="40" height="40" as="geometry"/>
      </mxCell>
    </object>
    <mxCell id="9" style="edgeStyle=orthogonalEdgeStyle;rounded=0;html=1;exitX=0.5;exitY=1;entryX=0.5;entryY=1;jettySize=auto;orthogonalLoop=1;" parent="1" source="3" target="2" edge="1">
      <mxGeometry relative="1" as="geometry"/>
    </mxCell>
    <object label="" type="exp" id="3">
      <mxCell style="rounded=0;whiteSpace=wrap;html=1;direction=south;" parent="1" vertex="1">
        <mxGeometry x="230" y="140" width="20" height="70" as="geometry"/>
      </mxCell>
    </object>
    <mxCell id="7" style="edgeStyle=orthogonalEdgeStyle;rounded=0;html=1;exitX=0.5;exitY=0;entryX=0.5;entryY=0;jettySize=auto;orthogonalLoop=1;" parent="1" source="4" target="5" edge="1">
      <mxGeometry relative="1" as="geometry"/>
    </mxCell>
    <object label="" type="gen" id="4">
      <mxCell style="rounded=0;whiteSpace=wrap;html=1;direction=south;fillColor=#000000;" parent="1" vertex="1">
        <mxGeometry x="230" y="20" width="20" height="70" as="geometry"/>
      </mxCell>
    </object>
    <mxCell id="8" style="edgeStyle=orthogonalEdgeStyle;rounded=0;html=1;exitX=0.5;exitY=1;entryX=0.5;entryY=0;jettySize=auto;orthogonalLoop=1;" parent="1" source="5" target="3" edge="1">
      <mxGeometry relative="1" as="geometry"/>
    </mxCell>
    <object label="" type="place" id="5">
      <mxCell style="ellipse;whiteSpace=wrap;html=1;aspect=fixed;" parent="1" vertex="1">
        <mxGeometry x="320" y="90" width="40" height="40" as="geometry"/>
      </mxCell>
    </object>
    <mxCell id="10" value="P1" style="text;html=1;strokeColor=none;fillColor=none;align=center;verticalAlign=middle;whiteSpace=wrap;rounded=0;" parent="1" vertex="1">
      <mxGeometry x="100" y="120" width="40" height="20" as="geometry"/>
    </mxCell>
    <mxCell id="11" value="P2" style="text;html=1;strokeColor=none;fillColor=none;align=center;verticalAlign=middle;whiteSpace=wrap;rounded=0;" parent="1" vertex="1">
      <mxGeometry x="370" y="110" width="40" height="20" as="geometry"/>
    </mxCell>
    <mxCell id="12" value="T2" style="text;html=1;strokeColor=none;fillColor=none;align=center;verticalAlign=middle;whiteSpace=wrap;rounded=0;" parent="1" vertex="1">
      <mxGeometry x="230" y="220" width="40" height="20" as="geometry"/>
    </mxCell>
    <mxCell id="13" value="T1" style="text;html=1;strokeColor=none;fillColor=none;align=center;verticalAlign=middle;whiteSpace=wrap;rounded=0;" parent="1" vertex="1">
      <mxGeometry x="250" y="10" width="40" height="20" as="geometry"/>
    </mxCell>
    <object label="// preprocess&#xa;&#xa;n=3&#xa;m=5" type="preprocess" id="16">
      <mxCell style="text;strokeColor=#000000;fillColor=none;align=left;verticalAlign=top;rounded=0;labelBorderColor=none;spacingTop=0;spacing=2;shadow=0;comic=0;dashed=1;" parent="1" vertex="1">
        <mxGeometry x="460" y="130" width="110" height="70" as="geometry"/>
      </mxCell>
    </object>
    <object label="// postprocess&#xa;&#xa;reward test ifelse(#P1 &gt;= 1, 1, 0)" type="postprocess" id="19">
      <mxCell style="text;strokeColor=#000000;fillColor=none;align=left;verticalAlign=top;rounded=0;labelBorderColor=none;spacingTop=0;spacing=2;shadow=0;comic=0;dashed=1;" parent="1" vertex="1">
        <mxGeometry x="430" y="255" width="260" height="70" as="geometry"/>
      </mxCell>
    </object>
    <object label="This is a comment" type="comment" id="21">
      <mxCell style="shape=callout;perimeter=calloutPerimeter;position2=0.68;direction=west;" vertex="1" parent="1">
        <mxGeometry x="210" y="245" width="120" height="80" as="geometry"/>
      </mxCell>
    </object>
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
