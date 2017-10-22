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
	Instance lnr.SimpleAlgorithm
}

//New Node
func New(polyline *pln.Polyline, rng *rng.Range, gfn geom.GeometryFn) *Node {
	var pt *geom.Point
	var chull []*geom.Point
	var coords = make([]*geom.Point, 0, rng.Size()+1)

	for _, i := range rng.Stride() {
		pt = polyline.Coordinates[i].Clone()
		pt[geom.Z] = float64(i)
		coords = append(coords, pt)
	}

	chull = geom.ConvexHull(coords, false)

	return &Node{
		Polyline: polyline,
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

//as segment
func (self *Node) Segment() *seg.Seg {
	return self.Polyline.Segment(self.Range)
}

//as segment
func (self *Node) SubPolyline() *pln.Polyline {
	return self.Polyline.SubPolyline(self.Range)
}
