package symbol_table

import (
	"fmt"
	"testing"
)

func TestBST_Print(t *testing.T) {
	bst := &BST{}
	bst.Put(5, 5)
	bst.Put(3, 3)
	bst.Put(1, 1)
	bst.Put(2, 2)
	bst.Put(9, 9)
	bst.Put(7, 7)
	val, ok := bst.Get(1)
	if !ok {
		fmt.Println("not ok")
	}
	fmt.Println(val)
	bst.Delete(3)
	bst.Print()
}
