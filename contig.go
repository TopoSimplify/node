package node

import "github.com/intdxdt/rtree"

//Checks if two nodes: nopde `a` and `b` are contiguous
// returns  bool(intersects), bool(is contig at vertex), int(number of intersections)
func IsContiguous(a, b *Node) (bool, bool, int) {
	//@formatter:off
	var ga = a.Geometry
	var gb = b.Geometry
	var contig = false
	var interCount = 0

	var bln = ga.Intersects(gb)
	if bln {
		var interpts = ga.Intersection(gb)

		var aiPt, ajPt = a.SegmentPoints()
		var biPt, bjPt = b.SegmentPoints()

		interCount = len(interpts)

		for _, pt := range interpts {
			var blnAseg = pt.Equals2D(&aiPt) || pt.Equals2D(&ajPt)
			var blnBseg = pt.Equals2D(&biPt) || pt.Equals2D(&bjPt)

			if blnAseg || blnBseg {
				contig = ajPt.Equals2D(&biPt) ||
					ajPt.Equals2D(&bjPt) ||
					aiPt.Equals2D(&bjPt) ||
					aiPt.Equals2D(&biPt)
			}

			if contig {
				break
			}
		}
	}

	return bln, contig, interCount
}

//Find neibours of node (prev , next)
func Neighbours(obj *rtree.Obj, neighbs []*rtree.Obj) (*rtree.Obj, *rtree.Obj) {
	var hull = obj.Object.(*Node)
	var prev, nxt *rtree.Obj
	var  h *Node
	var i, j = hull.Range.I, hull.Range.J
	for k := range neighbs {
		h  = neighbs[k].Object.(*Node)
		if h != hull {
			if i == h.Range.J {
				prev = neighbs[k]
			} else if j == h.Range.I {
				nxt = neighbs[k]
			}
		}
		if prev != nil && nxt != nil {
			break
		}
	}
	return prev, nxt
}
