package node


//checks for node contiguity
// returns  bool (intersects), bool(is contig at vertex), int (number of intersections)
func IsContiguous(a, b *Node) (bool, bool, int) {
	//@formatter:off
	var pln         = a.Polyline
	var coords      = pln.Coordinates
	var ga          = a.Geom
	var gb          = b.Geom
	var contig      = false
	var inter_count = 0

	var bln = ga.Intersects(gb)
	if bln {
		var interpts = ga.Intersection(gb)

		var ai_pt = coords[a.Range.I()]
		var aj_pt = coords[a.Range.J()]

		var bi_pt = coords[b.Range.I()]
		var bj_pt = coords[b.Range.J()]

		inter_count = len(interpts)

		for _, pt := range interpts {
			var bln_aseg = pt.Equals2D(ai_pt) || pt.Equals2D(aj_pt)
			var bln_bseg = pt.Equals2D(bi_pt) || pt.Equals2D(bj_pt)

			if bln_aseg || bln_bseg {
				contig = aj_pt.Equals2D(bi_pt) ||
					     aj_pt.Equals2D(bj_pt) ||
					     ai_pt.Equals2D(bj_pt) ||
					     ai_pt.Equals2D(bi_pt)
			}

			if contig {
				break
			}
		}
	}

	return bln, contig, inter_count
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
