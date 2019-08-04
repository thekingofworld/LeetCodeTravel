package union_find

type QuickFind struct {
	id    []int
	count int
}

func QF(n int) *QuickFind {
	q := &QuickFind{}
	q.id = make([]int, n)
	for i := range q.id {
		q.id[i] = i
	}
	q.count = n
	return q
}

func (q *QuickFind) Connected(x, y int) bool {
	return q.Find(x) == q.Find(y)
}

func (q *QuickFind) Find(x int) int {
	if x >= len(q.id) {
		return -1
	}
	return q.id[x]
}

func (q *QuickFind) Union(x, y int) {
	if q.Connected(x, y) {
		return
	}
	xID := q.id[x]
	yID := q.id[y]
	for i := range q.id {
		if q.id[i] == xID {
			q.id[i] = yID
		}
	}
	q.count--
}

func (q *QuickFind) Count() int {
	return q.count
}
