package node

import (
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/geom"
	"github.com/TopoSimplify/rng"
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
func CreateNode(id *iter.Igen, coordinates geom.Coords, rng rng.Rng, geomFn func(geom.Coords)geom.Geometry) Node {
	var chull = geom.ConvexHull(coordinates)
	var g = geomFn(chull)
	return Node{
		Id:       id.Next(),
		Polyline: pln.New(coordinates),
		Range:    rng,
		MBR:      g.Bounds(),
		Geom:     g,
	}
}

//CreateNode From MBR
func CreateNodeFromMBR(id *iter.Igen, box mbr.MBR) Node {
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


//coordinates
func (self *Node) Coordinates() geom.Coords {
	return self.Polyline.Coordinates
}


//first point in coordinates
func (self *Node) First() *geom.Point {
	return self.Polyline.Coordinates.First()
}

//last point in coordinates
func (self *Node) Last() *geom.Point {
	return self.Polyline.Coordinates.Last()
}

//as segment
func (self *Node) Segment() *geom.Segment {
	return geom.NewSegment(self.Polyline.Coordinates, self.Range.I, self.Range.J)
}

//hull segment as polyline
func (self *Node) SegmentAsPolyline() *pln.Polyline {
	var n =  self.Polyline.Len() -1
	var coords = self.Polyline.Coordinates
	coords.Idxs = []int{coords.Idxs[0], coords.Idxs[n]}
	return pln.New(coords)
}

//segment points
func (self *Node) SegmentPoints() (*geom.Point, *geom.Point) {
	return self.First(), self.Last()
}
