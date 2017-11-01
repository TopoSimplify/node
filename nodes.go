package node

import (
	"sort"
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/sset"
	"github.com/intdxdt/deque"
)

type Nodes struct {
	list []*Node
}

func NewNodes(size ...int) *Nodes {
	var n = 0
	if len(size) > 0 {
		n = size[0]
	}
	return &Nodes{list: make([]*Node, 0, n)}
}

func (self *Nodes) Len() int {
	return len(self.list)
}

func (self *Nodes) Swap(i, j int) {
	self.list[i], self.list[j] = self.list[j], self.list[i]
}

func (self *Nodes) Less(i, j int) bool {
	return self.list[i].Range.I() < self.list[j].Range.I()
}

//Get at index
func (self *Nodes) Get(index int) *Node {
	return self.list[index]
}

//Get underlying values
func (self *Nodes) DataView() []*Node {
	return self.list
}

func (self *Nodes) Sort() *Nodes {
	sort.Sort(self)
	return self
}

func (self *Nodes) Reverse() *Nodes {
	sort.Sort(sort.Reverse(self))
	return self
}

func (self *Nodes) Push(v *Node) *Nodes {
	self.list = append(self.list, v)
	return self
}

func (self *Nodes) Extend(values ...*Node) *Nodes {
	self.list = append(self.list, values...)
	return self
}

func (self *Nodes) Pop() *Nodes {
	if !self.IsEmpty() {
		n := len(self.list) - 1
		self.list[n] = nil
		self.list = self.list[:n]
	}
	return self
}

func (self *Nodes) IsEmpty() bool {
	return self.Len() == 0
}

func (self *Nodes) Empty() *Nodes {
	for i := range self.list {
		self.list[i] = nil
	}
	self.list = self.list[:0]
	return self
}

func (self *Nodes) AsDeque() *deque.Deque {
	queue := deque.NewDeque()
	for _, h := range self.list {
		queue.Append(h)
	}
	return queue
}

func (self *Nodes) AsPointSet() *sset.SSet {
	var set = sset.NewSSet(cmp.Int)
	for _, o := range self.list {
		set.Extend(o.Range.I(), o.Range.J())
	}
	return set
}