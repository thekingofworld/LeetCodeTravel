package graph

import "container/list"

type BFP struct {
	marked []bool
	edgeTo []int
	s      int
}

func BFPPaths(g *Graph, s int) *BFP {
	b := &BFP{}
	b.marked = make([]bool, g.V())
	b.edgeTo = make([]int, g.V())
	b.s = s
	b.bfs(g, s)
	return b
}

func (b *BFP) HasPathTo(v int) bool {
	return b.marked[v]
}

func (b *BFP) PathTo(v int) []int {
	if !b.HasPathTo(v) {
		return nil
	}
	var path []int
	for x := v; x != b.s; x = b.edgeTo[v] {
		path = append(path, x)
	}
	path = append(path, b.s)
	length := len(path)
	for i := 0; i < length/2; i++ {
		path[i], path[length-i-1] = path[length-i-1], path[i]
	}
	return path
}

func (b *BFP) bfs(g *Graph, s int) {
	queue := NewQueue()
	b.marked[s] = true
	queue.Enqueue(s)
	for !queue.Empty() {
		v := queue.Dequeue()
		for _, w := range g.Adj(v) {
			if !b.marked[w] {
				b.marked[w] = true
				b.edgeTo[w] = v
				queue.Enqueue(w)
			}
		}
	}
}

type Queue struct {
	l *list.List
}

func NewQueue() *Queue {
	q := &Queue{}
	q.l = list.New()
	return q
}

func (q *Queue) Enqueue(e int) {
	q.l.InsertAfter(e, q.l.Back())
}

func (q *Queue) Dequeue() int {
	return q.l.Remove(q.l.Back()).(int)
}

func (q *Queue) Empty() bool {
	return q.l.Len() > 0
}
