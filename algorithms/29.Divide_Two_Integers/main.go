package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(-2147483648, 1))
}

func divide(dividend int, divisor int) int {
	negative := false
	if dividend < 0 && divisor < 0 {
		dividend = -dividend
		divisor = -divisor
	} else if dividend < 0 {
		negative = true
		dividend = -dividend
	} else if divisor < 0 {
		negative = true
		divisor = -divisor
	}
	res := 0
	for dividend >= divisor {
		dividend -= divisor
		res++
	}
	if negative {
		return -res
	}
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res
}