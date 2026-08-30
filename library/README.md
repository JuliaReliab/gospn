# Shape library for draw.io

`MRSPN.xml` is the draw.io shape library the graphical workflow needs: Place, EXP / IMM /
GEN transitions, Label, Postscripts and Comments, each carrying the `type` property and
the default option properties `gospn gen` reads.

Load it into [diagrams.net](https://www.diagrams.net/) with **File → Open Library from →
URL**:

```
https://raw.githubusercontent.com/JuliaReliab/gospn/master/library/MRSPN.xml
```

It lives here rather than only in a gist because it is the first step of the documented
workflow: the tool cannot read a diagram whose objects lack the `type` property, and
nothing in the repository recorded what the library was supposed to contain. The gist at
`gist.github.com/okamumu/d10aabf442905b51f627df803139bd87` is the same file and stays as
a mirror.

`gospn gen` does **not** require the library -- it recognises any object with the right
`type` property in its EditData. The library is what saves drawing them by hand.
