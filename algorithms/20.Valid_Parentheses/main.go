package main

import "fmt"

func main() {
	if isValid("{[]}") {
		fmt.Println("valid str")
	} else {
		fmt.Println("invalid")
	}
}

type stack struct {
	items []uint8
	length int
}

func (s *stack) push(item uint8) {
	if s.length == 0 {
		s.items = make([]uint8, 0)
	}
	s.items = append(s.items, item)
	s.length++
}

func (s *stack) pop() uint8 {
	p := s.items[s.length-1]
	s.items = s.items[0:s.length-1]
	s.length--
	return p
}

func (s *stack) len() int {
	return s.length
}

func isValid(s string) bool {
	stack := new(stack)
	for i:=0; i<len(s); i++ {
		var needCheck bool
		var needCheckItem uint8
		if s[i] == '}' {
			needCheck = true
			needCheckItem = '{'
		}
		if s[i] == ']' {
			needCheck = true
			needCheckItem = '['
		}
		if s[i] == ')' {
			needCheck = true
			needCheckItem = '('
		}
		if needCheck {
			if stack.len() == 0 {
				return false
			}
			item := stack.pop()
			if item != needCheckItem {
				return false
			}
			continue
		}
		stack.push(s[i])
	}
	if stack.len() > 0 {
		return false
	}
	return true
}