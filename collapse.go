package node

import (
	"github.com/intdxdt/geom"
)

// Is node collapsible with respect to other
// self and other should be contiguous
func (self *Node) Collapsible(other *Node) bool {
	// segments are already collapsed
	if self.Range.Size() == 1 {
		return true
	}
	// or hull can be a linear for
	// colinear boundaries where self.range.size > 1
	if self.Geom.Type().IsLineString() {
		return true
	}

	var ai, aj = self.SegmentPoints()
	var bi, bj = other.SegmentPoints()

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
	if ply, ok := self.Geom.Geometry().(*geom.Polygon); ok {
		return !ply.Shell.PointCompletelyInRing(t)
	}
	panic("unimplemented : hull type is handled")
}
