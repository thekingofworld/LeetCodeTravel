package main

import "fmt"

func main() {
	fmt.Println(intToRoman(20))
}

var RomanMums = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var RomanStrs = map[int]string{1000:"M", 900:"CM", 500:"D", 400:"CD", 100:"C", 90:"XC", 50:"L", 40:"XL", 10:"X", 9:"IX", 5:"V", 4:"IV", 1:"I"}

func intToRoman(num int) string {
	Roman := make([]rune, 0)
	for num > 0 {
		for _, item := range RomanMums {
			if num - item >= 0 {
				Roman = append(Roman, []rune(RomanStrs[item])...)
				num -= item
				break
			}
		}
	}
	return string(Roman)
}