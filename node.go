package node

import (
	"fmt"
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/geom"
	"github.com/TopoSimplify/rng"
	"github.com/TopoSimplify/seg"
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/lnr"
	"github.com/intdxdt/random"
)

type NodeQueue interface {
	NodeQueue() []*Node
}

//Node Type
type Node struct {
	id       string
	Polyline *pln.Polyline
	Range    *rng.Range
	Geom     geom.Geometry
	Instance lnr.Linegen
}

//New Node
func New(coordinates []*geom.Point, rng *rng.Range, gfn geom.GeometryFn, nodeId ...string) *Node {
	var chull []*geom.Point
	var n = len(coordinates)
	var coords = make([]*geom.Point, n, n)

	copy(coords, coordinates)
	chull = geom.ConvexHull(coords, false)

	var nd = &Node{
		Polyline: pln.New(coordinates),
		Range:    rng,
		Geom:     gfn(chull),
	}

	var id string
	if len(nodeId) > 0 {
		id = nodeId[0]
	} else {
		id = random.String(8)
	}
	return nd.SetId(id)
}


//Id
func (self *Node) Id() string {
	return self.id
}

//Id
func (self *Node) SubNodeIds() (string, string) {
	return fmt.Sprintf("%v/a", self.id), fmt.Sprintf("%v/b", self.id)
}

//Set id
func (self *Node) SetId(key string) *Node {
	self.id = key
	return self
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
	var n = self.Polyline.Len()
	return self.Polyline.Coordinates[:n:n]
}

//first point in coordinates
func (self *Node) First() *geom.Point {
	return self.Polyline.Coordinates[0]
}

//last point in coordinates
func (self *Node) Last() *geom.Point {
	return self.Polyline.Coordinates[self.Polyline.Len()-1]
}

//as segment
func (self *Node) Segment() *seg.Seg {
	var a, b = self.SegmentPoints()
	return seg.NewSeg(a, b, self.Range.I, self.Range.J)
}

//hull segment as polyline
func (self *Node) SegmentAsPolyline() *pln.Polyline {
	var a, b = self.SegmentPoints()
	return pln.New([]*geom.Point{a, b})
}

//segment points
func (self *Node) SegmentPoints() (*geom.Point, *geom.Point) {
	return self.First(), self.Last()
}
