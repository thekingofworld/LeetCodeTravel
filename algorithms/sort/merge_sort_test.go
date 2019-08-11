package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	data := []int{4, 5, 3, 6, 1, 2, 21, 32, 12, 15, 10, 9, 8, 7}
	MergeSort(data)
	fmt.Println(data)
}
