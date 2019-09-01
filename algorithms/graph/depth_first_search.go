package graph

type DFS struct {
	marked []bool
	count  int
}

func DFSSearch(g *Graph, s int) *DFS {
	d := &DFS{}
	d.marked = make([]bool, g.V())
	d.count = 0
	d.dfs(g, s)
	return d
}

func (d *DFS) Marked(v int) bool {
	return d.marked[v]
}

func (d *DFS) Count() int {
	return d.count
}

func (d *DFS) dfs(g *Graph, v int) {
	d.marked[v] = true
	d.count++
	for _, w := range g.Adj(v) {
		if !d.Marked(w) {
			d.dfs(g, w)
		}
	}
}
