package main

import "fmt"

func main() {
	s := "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"
	i := myAtoi(s)
	fmt.Println(i)
}

func myAtoi(str string) int {
	i := 0
	start := 0
	numRune := make([]rune, 0)
	data := []rune(str)
	isNegative := false
	for _, c := range data {
		if c == ' ' && start == 0 {
			continue
		}
		if c != ' ' && start == 0 && !(c == '-' || c == '+' || (c >= '0' && c <= '9')) {
			return 0
		}
		if start == 0 && (c == '-' || c == '+' || (c >= '0' && c <= '9')) {
			start = 1
			if c == '-' {
				isNegative = true
				continue
			}
			if c == '+' {
				continue
			}
			numRune = append(numRune, c)
			continue
		}
		if start == 1 && !(c >= '0' && c <= '9') {
			break
		}
		numRune = append(numRune, c)
	}
	count := len(numRune)
	if count == 0 {
		return 0
	}
	for _, c := range numRune {
		m := 1
		for j := 0; j < count-1; j++ {
			if uint64(m) * 10 > (1<<63 - 1) && c != '0' {
				i = (1<<63 - 1)
				break
			}
			m = m * 10
		}
		count--
		if (1<<63 - 1) - (int(c-48) * m) < i {
			i = (1<<63 - 1)
			break
		}
		i = i + (int(c-48) * m)
	}
	if isNegative {
		i = -i
		if i < (-1 << 31) {
			return -1 << 31
		}
	} else if i > (1<<31 - 1) {
		return (1<<31 - 1)
	}
	return i
}