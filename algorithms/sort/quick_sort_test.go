package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	data := []int{4, 5, 3, 6, 1, 2, 21, 32, 12, 15, 10, 9, 8, 7}
	QuickSort(data)
	fmt.Println(data)
}
