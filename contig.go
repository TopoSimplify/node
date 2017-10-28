package node

//Checks if two nodes: nopde `a` and `b` are contiguous
// returns  bool(intersects), bool(is contig at vertex), int(number of intersections)
func IsContiguous(a, b *Node) (bool, bool, int) {
	//@formatter:off
	var ga          = a.Geom
	var gb          = b.Geom
	var contig      = false
	var interCount = 0

	var bln = ga.Intersects(gb)
	if bln {
		var interpts = ga.Intersection(gb)

		var aiPt, ajPt = a.SegmentPoints()
		var biPt, bjPt = b.SegmentPoints()

		interCount = len(interpts)

		for _, pt := range interpts {
			var blnAseg = pt.Equals2D(aiPt) || pt.Equals2D(ajPt)
			var blnBseg = pt.Equals2D(biPt) || pt.Equals2D(bjPt)

			if blnAseg || blnBseg {
				contig = ajPt.Equals2D(biPt) ||
					     ajPt.Equals2D(bjPt) ||
					     aiPt.Equals2D(bjPt) ||
					     aiPt.Equals2D(biPt)
			}

			if contig {
				break
			}
		}
	}

	return bln, contig, interCount
}


//Find neibours of node (prev , next)
func Neighbours(hull *Node, neighbs *Nodes) (*Node, *Node) {
	var prev, nxt *Node
	var i, j = hull.Range.I(), hull.Range.J()
	for _, h := range neighbs.DataView() {
		if h != hull {
			if i == h.Range.J() {
				prev = h
			} else if j == h.Range.I() {
				nxt = h
			}
		}
		if prev != nil && nxt != nil {
			break
		}
	}
	return prev, nxt
}
