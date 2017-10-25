package node

import (
	"simplex/rng"
	"simplex/seg"
	"simplex/pln"
	"simplex/lnr"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/mbr"
)

//Node Type
type Node struct {
	Polyline *pln.Polyline
	Range    *rng.Range
	Geom     geom.Geometry
	Instance lnr.Linegen
}

//New Node
func New(coordinates []*geom.Point, rng *rng.Range, gfn geom.GeometryFn) *Node {
	var chull []*geom.Point
	var n = len(coordinates)
	var coords = make([]*geom.Point, n, n)
	copy(coords, coordinates)
	chull = geom.ConvexHull(coords, false)

	return &Node{
		Polyline: pln.New(coordinates),
		Range:    rng,
		Geom:     gfn(chull),
	}
}

//Implements igeom interface
func (self *Node) Geometry() geom.Geometry {
	return self.Geom
}

//Implements bbox interface
func (self *Node) BBox() *mbr.MBR {
	return self.Geom.BBox()
}

//stringer interface
func (self *Node) String() string {
	return self.Geom.WKT()
}

//stringer interface
func (self *Node) Coordinates() []*geom.Point {
	return self.Polyline.Coordinates
}

//first point in coordinates
func (self *Node) First() *geom.Point {
	return self.Polyline.Coordinates[0]
}

//last point in coordinates
func (self *Node) Last() *geom.Point {
	var n = len(self.Polyline.Coordinates)
	return self.Polyline.Coordinates[n-1]
}

//as segment
func (self *Node) Segment() *seg.Seg {
	var a, b = self.SegmentPoints()
	return seg.NewSeg(a, b, self.Range.I(), self.Range.J())
}

//segment points
func (self *Node) SegmentPoints() (*geom.Point, *geom.Point) {
	return self.First(), self.Last()
}
