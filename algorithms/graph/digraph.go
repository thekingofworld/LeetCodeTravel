package graph

type Digraph struct {
	v   int
	e   int
	adj [][]int
}

func NewDigraph(v int) *Digraph {
	d := &Digraph{}
	d.v = v
	d.adj = make([][]int, v)
	for i := range d.adj {
		d.adj[i] = make([]int, 0)
	}
	return d
}

func (d *Digraph) V() int {
	return d.v
}

func (d *Digraph) E() int {
	return d.e
}

func (d *Digraph) AddEdge(v, w int) {
	d.adj[v] = append(d.adj[v], w)
	d.e++
}

func (d *Digraph) Adj(v int) []int {
	return d.adj[v]
}

func (d *Digraph) Reverse() *Digraph {
	rd := NewDigraph(d.V())
	for i := range d.adj {
		for _, w := range d.Adj(i) {
			rd.AddEdge(w, i)
		}
	}
	return rd
}
