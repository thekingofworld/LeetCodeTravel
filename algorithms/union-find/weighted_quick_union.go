package union_find

type WeightedQuickUnion struct {
	id     []int
	weight []int
	count  int
}

func WQU(n int) *WeightedQuickUnion {
	w := &WeightedQuickUnion{}
	w.id = make([]int, n)
	w.weight = make([]int, n)
	for i := range w.id {
		w.id[i] = i
		w.weight[i] = 1
	}
	w.count = n
	return w
}

func (w *WeightedQuickUnion) Connected(x, y int) bool {
	return w.Find(x) == w.Find(y)
}

func (w *WeightedQuickUnion) Find(x int) int {
	if x >= len(w.id) {
		return -1
	}
	t := x
	for w.id[t] != t {
		t = w.id[t]
	}
	//路径压缩
	//p := x
	//for w.id[p] != p {
	//	tmp := w.id[p]
	//	w.id[p] = t
	//	p = w.id[tmp]
	//}
	return t
}

func (w *WeightedQuickUnion) Union(x, y int) {
	xRoot := w.Find(x)
	yRoot := w.Find(y)
	if xRoot == yRoot {
		return
	}
	if w.weight[xRoot] < w.weight[yRoot] {
		w.id[xRoot] = yRoot
		w.weight[yRoot] += w.weight[xRoot]
	} else {
		w.id[yRoot] = xRoot
		w.weight[xRoot] += w.weight[yRoot]
	}
	w.count--
}

func (w *WeightedQuickUnion) Count() int {
	return w.count
}
