package symbol_table

import (
	"fmt"
	"testing"
)

func TestRedBlackBST_Print(t *testing.T) {
	bst := &RedBlackBST{}
	bst.Put(5, 5)
	bst.Put(3, 3)
	bst.Put(1, 1)
	bst.Put(2, 2)
	bst.Put(9, 9)
	bst.Put(7, 7)
	_, ok := bst.Get(1)
	if !ok {
		fmt.Println("not ok")
	}
	bst.Print()
}
