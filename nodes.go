package node

import (
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/sset"
)

type Nodes []*Node


func (self Nodes) Len() int {
	return len(self)
}

func (self Nodes) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self Nodes) Less(i, j int) bool {
	return self[i].Range.I < self[j].Range.I
}

func (self Nodes) AsPointSet() *sset.SSet {
	var set = sset.NewSSet(cmp.Int)
	for _, o := range self {
		set.Extend(o.Range.I, o.Range.J)
	}
	return set
}

func Pop(self *[]*Node) []*Node {
	var nodes = *self
	if len(nodes) != 0 {
		n := len(nodes) - 1
		nodes[n] = nil
		*self = nodes[:n]
	}
	return nodes
}

func Clear(self *[]*Node) []*Node{
	var nodes = *self
	for i := range nodes {
		nodes[i] = nil
	}
	*self = nodes[:0]
	return nodes
}

