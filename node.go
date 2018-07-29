package node

import (
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/geom"
	"github.com/TopoSimplify/rng"
	"github.com/TopoSimplify/seg"
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/lnr"
	"github.com/intdxdt/iter"
)

//Node Type
type Node struct {
	Id       int
	Polyline *pln.Polyline
	Range    rng.Rng
	MBR      mbr.MBR
	Geom     geom.Geometry
	Instance lnr.Linegen
}

//CreateNode Node
func CreateNode(idgen *iter.IntGen, coordinates []geom.Point, rng rng.Rng, geomFn geom.GeometryFn) Node {
	var chull []geom.Point
	var n = len(coordinates)
	var coords = make([]geom.Point, 0, n)
	for i := range coordinates {
		coords = append(coords, coordinates[i])
	}
	chull = geom.ConvexHull(coords, false)
	var g = geomFn(chull)

	return Node{
		Id:       idgen.Next(),
		Polyline: pln.New(coordinates),
		Range:    rng,
		MBR:      g.Bounds(),
		Geom:     g,
	}
}

//CreateNode From MBR
func CreateNodeFromMBR(id *iter.IntGen, box mbr.MBR) Node {
	return Node{Id: id.Next(), MBR: box}
}

//Implements bbox interface
func (self *Node) BBox() *mbr.MBR {
	return self.Geom.BBox()
}

//Implements bbox interface
func (self *Node) Bounds() mbr.MBR {
	return self.Geom.Bounds()
}

//Geometry bbox interface
func (self *Node) Geometry() geom.Geometry {
	return self.Geom
}

//stringer interface
func (self *Node) String() string {
	return self.Geom.WKT()
}

//stringer interface
func (self *Node) Coordinates() []geom.Point {
	var n = self.Polyline.Len()
	return self.Polyline.Coordinates[:n:n]
}

//first point in coordinates
func (self *Node) First() geom.Point {
	return self.Polyline.Coordinates[0]
}

//last point in coordinates
func (self *Node) Last() geom.Point {
	return self.Polyline.Coordinates[self.Polyline.Len()-1]
}

//as segment
func (self *Node) Segment() *seg.Seg {
	var a, b = self.SegmentPoints()
	return seg.NewSeg(&a, &b, self.Range.I, self.Range.J)
}

//hull segment as polyline
func (self *Node) SegmentAsPolyline() *pln.Polyline {
	var a, b = self.SegmentPoints()
	return pln.New([]geom.Point{a, b})
}

//segment points
func (self *Node) SegmentPoints() (geom.Point, geom.Point) {
	return self.First(), self.Last()
}
