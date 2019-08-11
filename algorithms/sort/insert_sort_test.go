package sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	data := []int{4, 5, 3, 6, 1, 2}
	InsertSort(data)
	fmt.Println(data)
}
