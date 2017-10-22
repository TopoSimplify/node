package node

import (
	"simplex/rng"
	"simplex/seg"
	"simplex/pln"
	"simplex/lnr"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/mbr"
)


//hull node
type Node struct {
	Polyline *pln.Polyline
	Range    *rng.Range
	Geom     geom.Geometry
	Instance lnr.SimpleAlgorithm
}

//New Hull Node
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

//implements igeom interface
func (h *Node) Geometry() geom.Geometry {
	return h.Geom
}

//implements bbox interface
func (h *Node) BBox() *mbr.MBR {
	return h.Geom.BBox()
}

//stringer interface
func (h *Node) String() string {
	return h.Geom.WKT()
}

//stringer interface
func (h *Node) Coordinates() []*geom.Point {
	return h.Polyline.Coordinates
}

//as segment
func (h *Node) Segment() *seg.Seg {
	return h.Polyline.Segment(h.Range)
}

//as segment
func (h *Node) SubPolyline() *pln.Polyline {
	return h.Polyline.SubPolyline(h.Range)
}
