package union_find

type QuickUnion struct {
	id    []int
	count int
}

func QU(n int) *QuickUnion {
	q := &QuickUnion{}
	q.id = make([]int, n)
	for i := range q.id {
		q.id[i] = i
	}
	q.count = n
	return q
}

func (q *QuickUnion) Connected(x, y int) bool {
	return q.Find(x) == q.Find(y)
}

func (q *QuickUnion) Find(x int) int {
	if x >= len(q.id) {
		return -1
	}
	t := x
	for q.id[t] != t {
		t = q.id[t]
	}
	return t
}

func (q *QuickUnion) Union(x, y int) {
	xRoot := q.Find(x)
	yRoot := q.Find(y)
	if xRoot == yRoot {
		return
	}
	q.id[xRoot] = yRoot
	q.count--
}

func (q *QuickUnion) Count() int {
	return q.count
}
