package main

import "fmt"

func main() {
	s := convert("PAYPALISHIRING", 4)
	fmt.Println(s)
}

func convert(s string, numRows int) string {
	length := len(s)
	if length == 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}
	r := make([][]rune, numRows)
	i := 0
	for i <= length {
		var size int
		if i + numRows > length {
			size = length - i
		} else {
			size = numRows
		}
		data := []rune(s[i:i+size])
		for j := 0; j < size; j++ {
			if len(r[j]) == 0 {
				r[j] = make([]rune, 0)
			}
			r[j] = append(r[j], data[j])
		}
		i = i + size

		if i + numRows - 2 > length {
			size = length - i
		} else {
			size = numRows - 2
		}
		other := []rune(s[i:i+size])
		for j := 0; j < size; j++ {
			r[numRows-2-j] = append(r[numRows-2-j], other[j])
		}
		i = i + size

		if i == length {
			break
		}
	}

	res := make([]rune, 0)
	for j := 0; j < numRows; j++ {
		for _, c := range r[j] {
			res = append(res, c)
		}
	}
	return string(res)
}