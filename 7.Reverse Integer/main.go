package main

import "fmt"

func main() {
	x := 120
	r := reverse(x)
	fmt.Println(r)
}

func reverse(x int) int {
	isNegative := false
	if x < 0 {
		isNegative = true
		x = -x
	}
	count := 0
	new := 0
	copyx := x
	for copyx != 0 {
		copyx = copyx / 10
		count++
	}
	for x != 0 {
		m := x % 10
		x = x / 10
		w := 1
		for j := 0; j < count-1; j++ {
			w = w*10
		}
		new = new + (m * w)
		count--
	}
	if new > (1<<31 - 1) {
		return 0
	}
	if isNegative {
		return -new
	}
	return new
}