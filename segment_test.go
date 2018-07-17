package node

import (
	"testing"
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/rng"
	"github.com/intdxdt/geom"
	"github.com/franela/goblin"
)

func TestHullSeg(t *testing.T) {
	var g = goblin.Goblin(t)

	var createHulls = func(ranges [][]int, coords []geom.Point) []*Node {
		var polyline = pln.New(coords)
		var hulls    = make([]*Node, 0)
		for _, r := range ranges {
			var i, j = r[0], r[len(r)-1]
			h := newNodeFromPolyline(polyline, rng.Range(i, j), hullGeom)
			hulls = append(hulls, h)
		}
		return hulls
	}

	g.Describe("node decomposition", func() {
		g.It("should test decomposition of a line as nodes", func() {
			var wkt = "LINESTRING ( 670 550, 680 580, 750 590, 760 630, 830 640, 870 630, 890 610, 920 580, 910 540, 890 500, 900 460, 870 420, 860 390, 810 360, 770 400, 760 420, 800 440, 810 470, 850 500, 820 560, 780 570, 760 530, 720 530, 707.3112236920351 500.3928552814154, 650 450 )"
			var coords = geom.NewLineStringFromWKT(wkt).Coordinates()
			var ranges = [][]int{{0, 12}, {12, 18}, {18, len(coords) - 1}}
			var hulls = createHulls(ranges, coords)

			for i, r := range ranges {
				var s = hulls[i].Segment()
				var a, b = coords[r[0]][:2], coords[r[1]][:2]
				g.Assert(r).Equal(s.Range().AsSlice())
				g.Assert(s.A[:2]).Equal(a)
				g.Assert(s.B[:2]).Equal(b)

				g.Assert([]geom.Point{*s.A, *s.B}).Equal(hulls[i].SegmentAsPolyline().Coordinates)
				var cs = hulls[i].Coordinates()
				cs = append(cs, cs[0])
				g.Assert(hulls[i].Coordinates()).Equal(hulls[i].Polyline.Coordinates)
				g.Assert(hulls[i].String()).Equal(hulls[i].Geometry.WKT())
				g.Assert(hulls[i].Geometry).Equal(hulls[i].Geometry)
				g.Assert(hulls[i].Geometry.BBox()).Equal(hulls[i].BBox())
			}
		})
	})
}
