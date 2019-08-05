package union_find

import (
	"fmt"
	"testing"
)

func TestWQU(t *testing.T) {
	w := WQU(10)
	fmt.Println(w)
}

func TestWeightedQuickUnion_Connected(t *testing.T) {
	w := WQU(10)
	fmt.Println(w.Connected(0, 1))
}

func TestWeightedQuickUnion_Find(t *testing.T) {
	w := WQU(10)
	fmt.Println(w.Find(1))
}

func TestWeightedQuickUnion_Union(t *testing.T) {
	w := WQU(10)
	w.Union(1, 5)
	w.Union(2, 5)
	fmt.Println(w)
}

func TestWeightedQuickUnion_Count(t *testing.T) {
	w := WQU(10)
	w.Union(1, 5)
	w.Union(2, 5)
	fmt.Println(w.Count())
}
