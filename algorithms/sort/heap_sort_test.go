package sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	data := []int{4, 5, 3, 6, 1, 2, 21, 32, 12, 15, 10, 9, 8, 7}
	HeapSort(data)
	fmt.Println(data)
}
