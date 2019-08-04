package union_find

import (
	"fmt"
	"testing"
)

func TestQU(t *testing.T) {
	q := QU(10)
	fmt.Println(q)
}

func TestQuickUnion_Connected(t *testing.T) {
	q := QU(10)
	fmt.Println(q.Connected(0, 1))
}

func TestQuickUnion_Find(t *testing.T) {
	q := QU(10)
	fmt.Println(q.Find(1))
}

func TestQuickUnion_Union(t *testing.T) {
	q := QU(10)
	q.Union(1, 5)
	fmt.Println(q)
}

func TestQuickUnion_Count(t *testing.T) {
	q := QU(10)
	q.Union(1, 5)
	fmt.Println(q.Count())
}
