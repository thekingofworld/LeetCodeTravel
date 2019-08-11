package sort

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	data := []int{4, 5, 3, 6, 1, 2}
	ShellSort(data)
	fmt.Println(data)
}
