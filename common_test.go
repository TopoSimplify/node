package node

import (
	"simplex/pln"
	"simplex/rng"
	"github.com/intdxdt/geom"
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

func linear_coords(wkt string) []*geom.Point{
	return geom.NewLineStringFromWKT(wkt).Coordinates()
}

func create_hulls(indxs [][]int, coords []*geom.Point) []*Node {
	poly := pln.New(coords)
	hulls := make([]*Node, 0)
	for _, o := range indxs {
		r := rng.NewRange(o[0], o[1])
		hulls = append(hulls, New(poly.SubCoordinates(r), r, hullGeom))
	}
	return hulls
}
