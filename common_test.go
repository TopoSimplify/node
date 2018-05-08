package node

import (
	"github.com/intdxdt/geom"
	"github.com/intdxdt/random"
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/rng"
)

//hull geom
func hullGeom(coords []*geom.Point) geom.Geometry {
	var g geom.Geometry

	if len(coords) > 2 {
		g = geom.NewPolygon(coords)
	} else if len(coords) == 2 {
		g = geom.NewLineString(coords)
	} else {
		g = coords[0].Clone()
	}
	return g
}

func linearCoords(wkt string) []*geom.Point{
	return geom.NewLineStringFromWKT(wkt).Coordinates()
}

func createHulls(indxs [][]int, coords []*geom.Point) []*Node {
	poly := pln.New(coords)
	hulls := make([]*Node, 0)
	for _, o := range indxs {
		r := rng.NewRange(o[0], o[1])
		hulls = append(hulls, New(poly.SubCoordinates(r), r, hullGeom, random.String(4)))
	}
	return hulls
}

//New Node
func newNodeFromPolyline(polyline *pln.Polyline, rng *rng.Range, gfn geom.GeometryFn) *Node {
	return New(polyline.SubCoordinates(rng), rng, gfn)
}

