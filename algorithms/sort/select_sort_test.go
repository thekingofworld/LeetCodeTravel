package sort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	data := []int{4, 5, 3, 6, 1, 2}
	SelectSort(data)
	fmt.Println(data)
}
