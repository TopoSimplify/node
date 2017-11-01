package node

import (
	"github.com/franela/goblin"
	"testing"
)

func TestNodes(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("test contig nodes", func() {
		g.It("should test nodes", func() {
			//checks if score is valid at threshold of constrained dp
			var coords = linearCoords("LINESTRING ( 960 840, 980 840, 980 880, 1020 900, 1080 880, 1120 860, 1160 800, 1160 760, 1140 700, 1080 700, 1040 720, 1060 760, 1120 800, 1080 840, 1020 820, 940 760 )")
			var hulls = createHulls([][]int{{8, 10}, {10, 12}, {0, 2}, {2, 6}, {6, 8}, {12, len(coords) - 1}}, coords)
			var ns = NewNodes(len(hulls))
			for _, n := range hulls {
				g.Assert(n.Id()).Equal(n.id)
				ida, idb := n.SubNodeIds()
				g.Assert(n.Id()).Equal(n.id)
				g.Assert(ida).Equal(n.id + "/a")
				g.Assert(idb).Equal(n.id + "/b")
				ns.Push(n)
			}
			ns.Sort()
			g.Assert(ns.Get(0).Range.AsArray()).Equal([2]int{0, 2})
			ns.Reverse()
			g.Assert(ns.Get(0).Range.AsArray()).Equal([2]int{12, len(coords) - 1})
			var n = ns.Get(len(hulls) - 1)
			g.Assert(ns).Equal(ns.Pop())
			g.Assert(n.Range.AsArray()).Equal([2]int{0, 2})
			g.Assert(ns.Len()).Equal(len(hulls) - 1)

			var que = ns.AsDeque()
			g.Assert(ns.Len()).Equal(que.Len())
			var set = ns.AsPointSet()
			g.Assert(set.Values()).Equal([]interface{}{2, 6, 8, 10, 12, len(coords) - 1})
			ns.Empty()
			g.Assert(ns.Len()).Equal(0)


		})
	})
}
