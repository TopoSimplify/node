package node

import (
	"github.com/intdxdt/geom"
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/rng"
	"github.com/intdxdt/iter"
)

var idgen = iter.NewIntGen(0)

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

func createHulls(indxs [][]int, coords []geom.Point) []Node {
	poly := pln.New(coords)
	hulls := make([]Node, 0, len(indxs))
	for _, o := range indxs {
		r := rng.Range(o[0], o[1])
		hulls = append(hulls, CreateNode(idgen, poly.SubCoordinates(r), r, hullGeom))
	}
	return hulls
}

//CreateNode Node
func nodeFromPolyline(polyline *pln.Polyline, rng rng.Rng, geomFn geom.GeometryFn) Node {
	return CreateNode(idgen, polyline.SubCoordinates(rng), rng, geomFn)
}

