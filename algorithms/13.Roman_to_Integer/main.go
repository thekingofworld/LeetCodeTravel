package main

import "fmt"

func main() {
	fmt.Println(romanToInt("LVIII"))
}

var RomanStrs = map[string]int{"M":1000, "CM":900, "D":500, "CD":400, "C":100, "XC":90, "L":50, "XL":40, "X":10, "IX":9, "V":5, "IV":4, "I":1}

func romanToInt(s string) int {
	var i int
	var length = len(s)
	var sum int
	for i < length {
		var tmp string
		tmp += string(s[i])
		if i+1 < length {
			if x,ok := RomanStrs[tmp+string(s[i+1])]; ok {
				sum += x
				i += 2
				continue
			}
		}
		sum += RomanStrs[tmp]
		i++
	}
	return sum
}