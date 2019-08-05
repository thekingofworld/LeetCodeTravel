package union_find

import (
	"fmt"
	"testing"
)

func TestQF(t *testing.T) {
	q := QF(10)
	fmt.Println(q)
}

func TestQuickFind_Connected(t *testing.T) {
	q := QF(10)
	fmt.Println(q.Connected(0, 1))
}

func TestQuickFind_Find(t *testing.T) {
	q := QF(10)
	fmt.Println(q.Find(1))
}

func TestQuickFind_Union(t *testing.T) {
	q := QF(10)
	q.Union(1, 5)
	fmt.Println(q)
}

func TestQuickFind_Count(t *testing.T) {
	q := QF(10)
	q.Union(1, 5)
	fmt.Println(q.Count())
}
