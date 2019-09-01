package graph

type DFP struct {
	marked []bool
	edgeTo []int
	s      int
}

func DFPPaths(g *Graph, s int) *DFP {
	d := &DFP{}
	d.marked = make([]bool, g.V())
	d.edgeTo = make([]int, g.V())
	d.s = s
	d.dfs(g, s)
	return d
}

func (d *DFP) HasPathTo(v int) bool {
	return d.marked[v]
}

func (d *DFP) PathTo(v int) []int {
	if !d.HasPathTo(v) {
		return nil
	}
	var path []int
	for x := v; x != d.s; x = d.edgeTo[v] {
		path = append(path, x)
	}
	path = append(path, d.s)
	length := len(path)
	for i := 0; i < length/2; i++ {
		path[i], path[length-i-1] = path[length-i-1], path[i]
	}
	return path
}

func (d *DFP) dfs(g *Graph, v int) {
	d.marked[v] = true
	for _, w := range g.Adj(v) {
		if !d.marked[w] {
			d.edgeTo[w] = v
			d.dfs(g, w)
		}
	}
}
