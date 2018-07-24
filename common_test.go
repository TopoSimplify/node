package node

import (
	"github.com/intdxdt/geom"
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/rng"
	"github.com/intdxdt/rtree"
)

//hull geom
func hullGeom(coords []geom.Point) geom.Geometry {
	var g geom.Geometry

	if len(coords) > 2 {
		g = geom.NewPolygon(coords)
	} else if len(coords) == 2 {
		g = geom.NewLineString(coords)
	} else {
		g = coords[0]
	}
	return g
}

func linearCoords(wkt string) []geom.Point{
	return geom.NewLineStringFromWKT(wkt).Coordinates()
}

func createHulls(indxs [][]int, coords []geom.Point) []*Node {
	poly := pln.New(coords)
	hulls := make([]*Node, 0)
	for _, o := range indxs {
		r := rng.Range(o[0], o[1])
		hulls = append(hulls, New(poly.SubCoordinates(r), r, hullGeom))
	}
	return hulls
}

func createHullObjects(indxs [][]int, coords []geom.Point) []*rtree.Obj {
	poly := pln.New(coords)
	hulls := make([]*rtree.Obj, 0)
	for i := range indxs {
		o := indxs[i]
		r := rng.Range(o[0], o[1])
		h := New(poly.SubCoordinates(r), r, hullGeom)
		hulls = append(hulls, rtree.Object(i, h.Bounds(), h))
	}
	return hulls
}

//New Node
func newNodeFromPolyline(polyline *pln.Polyline, rng rng.Rng, geomFn geom.GeometryFn) *Node {
	return New(polyline.SubCoordinates(rng), rng, geomFn)
}

