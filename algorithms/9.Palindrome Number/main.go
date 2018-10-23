package main

import (
	"fmt"
)

func main() {
	if isPalindrome(1) {
		fmt.Println("is a palindrome")
	} else {
		fmt.Println("is not a palindrome")
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	data := make([]uint8, 0)
	for x != 0 {
		data = append(data, uint8(x % 10))
		x = x / 10
	}
	length := len(data)
	border := len(data)/2
	for i := 0; i < border; i++ {
		if data[i] != data[length-i-1] {
			return false
		}
	}
	return true
}