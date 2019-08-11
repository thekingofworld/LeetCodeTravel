package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	data := []int{4, 5, 3, 6, 1, 2}
	BubbleSort(data)
	fmt.Println(data)
}
