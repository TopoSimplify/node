package node

import (
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/sset"
)

type NodePtrs []*Node

func (self NodePtrs) Len() int {
	return len(self)
}

func (self NodePtrs) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self NodePtrs) Less(i, j int) bool {
	return self[i].Range.I < self[j].Range.I
}

type Nodes []Node

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

func Pop(self *[]Node) []Node {
	var nodes = *self
	if len(nodes) != 0 {
		n := len(nodes) - 1
		nodes[n] = Node{}
		*self = nodes[:n]
	}
	return nodes
}

func Clear(self *[]Node) []Node {
	var nodes = *self
	*self = make([]Node, 0, len(nodes))
	return nodes
}
