package node

import (
    "sync"
    "github.com/intdxdt/deque"
)

//@formatter:off

type Queue struct {
    sync.RWMutex
    que *deque.Deque
}

func NewQueue() *Queue {
    return &Queue{que: deque.NewDeque()}
}

func (q *Queue) Append(o *Node) *Queue {
    q.Lock(); defer q.Unlock()
    q.que.Append(o)
    return q
}

func (q *Queue) AppendLeft(o *Node) *Queue {
    q.Lock(); defer q.Unlock()
    q.que.AppendLeft(o)
    return q
}

func (q *Queue) Pop() *Node {
    q.Lock(); defer q.Unlock()
    return q.que.Pop().(*Node)
}

func (q *Queue) PopLeft() *Node {
    q.Lock(); defer q.Unlock()
    return q.que.PopLeft().(*Node)
}

func (q *Queue) Clear() *Queue {
    q.Lock(); defer q.Unlock()
    q.que.Clear()
    return q
}

func (q *Queue) Size() int {
    q.RLock(); defer q.RUnlock()
    return q.que.Len()
}

func (q *Queue) First() *Node {
    q.RLock(); defer q.RUnlock()
	return q.que.First().(*Node)
}

func (q *Queue) Last() *Node {
    q.RLock(); defer q.RUnlock()
	return q.que.Last().(*Node)
}
