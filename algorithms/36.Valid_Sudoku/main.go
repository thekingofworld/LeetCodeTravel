package main

import "fmt"

func main() {
	board := make([][]byte, 0)
	board = append(board, []byte{'5','3','.','.','7','.','.','.','.'})
	board = append(board, []byte{'6','.','.','1','9','5','.','.','.'})
	board = append(board, []byte{'.','9','8','.','.','.','.','6','.'})
	board = append(board, []byte{'8','.','.','.','6','.','.','.','3'})
	board = append(board, []byte{'4','.','.','8','.','3','.','.','1'})
	board = append(board, []byte{'7','.','.','.','2','.','.','.','6'})
	board = append(board, []byte{'.','6','.','.','.','.','2','8','.'})
	board = append(board, []byte{'.','.','.','4','1','9','.','.','5'})
	board = append(board, []byte{'.','.','.','.','8','.','.','7','9'})
	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	//check sub-boxes
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			checkMap := make(map[byte]bool)
			for k:=i*3; k<=i*3+2; k++ {
				for l:=j*3; l<=j*3+2; l++ {
					tmp := board[k][l]
					if tmp == '.' {
						continue
					}
					if _,ok := checkMap[tmp]; ok {
						return false
					}
					checkMap[tmp] = true
				}
			}
		}
	}

	//check row
	for i:=0; i<9; i++ {
		checkRowMap := make(map[byte]bool)
		for j:=0; j<9; j++ {
			tmpRow := board[i][j]
			if tmpRow == '.' {
				continue
			}
			if _,ok := checkRowMap[tmpRow]; ok {
				return false
			}
			checkRowMap[tmpRow] = true
		}
	}

	//check column
	for i:=0; i<9; i++ {
		checkColumnMap := make(map[byte]bool)
		for j:=0; j<9; j++ {
			tmpColumn := board[j][i]
			if tmpColumn == '.' {
				continue
			}
			if _,ok := checkColumnMap[tmpColumn]; ok {
				return false
			}
			checkColumnMap[tmpColumn] = true
		}
	}

	return true
}