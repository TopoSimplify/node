package node

import "github.com/intdxdt/geom"

//Is node collapsible with respect to other
//self and other should be contiguous
func (self *Node) Collapsible( other *Node) bool {
	//segments are already collapsed
	if self.Range.Size() == 1 {
		return true
	}

	var pln = self.Coordinates()
	var pt_at = func(i int) *geom.Point {
		return geom.NewPoint(pln[i][:2])
	}

	var ra = self.Range
	var rb = other.Range
	var ai, aj = pt_at(ra.I()), pt_at(ra.J())
	var bi, bj = pt_at(rb.I()), pt_at(rb.J())

	var c *geom.Point
	if ai.Equals2D(bi) || aj.Equals2D(bi) {
		c = bi
	} else if ai.Equals2D(bj) || aj.Equals2D(bj) {
		c = bj
	} else {
		return true
	}

	var t = bj
	if c.Equals2D(t) {
		t = bi
	}
	var ply = self.Geom.(*geom.Polygon)
	return !ply.Shell.PointCompletelyInRing(t)
}
