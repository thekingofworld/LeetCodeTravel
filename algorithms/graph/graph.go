package graph

type Graph struct {
	v   int
	e   int
	adj [][]int
}

func NewGraph(v int) *Graph {
	g := &Graph{
		v: v,
	}
	g.adj = make([][]int, v)
	for i := range g.adj {
		g.adj[i] = make([]int, 0)
	}
	return g
}

func (g *Graph) V() int {
	return g.v
}

func (g *Graph) E() int {
	return g.e
}

func (g *Graph) AddEdge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.e++
}

func (g *Graph) Adj(v int) []int {
	return g.adj[v]
}
