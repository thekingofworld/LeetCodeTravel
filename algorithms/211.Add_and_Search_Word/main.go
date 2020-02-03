package main

import "fmt"

func main() {
	w := Constructor()
	w.AddWord("at")
	w.AddWord("and")
	w.AddWord("an")
	w.AddWord("add")
	fmt.Println(w.Search("a"))
	fmt.Println(w.Search(".at"))
	w.AddWord("bat")
	fmt.Println(w.Search(".at"))
	fmt.Println(w.Search("an."))
	fmt.Println(w.Search("a.b."))
	fmt.Println(w.Search("b."))
	fmt.Println(w.Search("a.d"))
	fmt.Println(w.Search("."))
}

type WordDictionary struct {
	count  int
	prefix int
	nodes  []*WordDictionary
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	w := WordDictionary{}
	w.nodes = make([]*WordDictionary, 26)
	return w
}

/** Adds a word into the data structure. */
func (w *WordDictionary) AddWord(word string) {
	next := w
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if next.nodes[index] == nil {
			tmp := Constructor()
			next.nodes[index] = &tmp
		}
		next = next.nodes[index]
		next.prefix++
	}
	next.count++
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (w *WordDictionary) Search(word string) bool {
	next := w
	for i := 0; i < len(word); i++ {
		if word[i] == '.' {
			for j := 0; j < len(next.nodes); j++ {
				if next.nodes[j] != nil {
					if (i+1 > len(word)-1 && next.nodes[j].count > 0) ||
						next.nodes[j].Search(word[i+1:]) {
						return true
					}
				}
			}
			return false
		}
		index := word[i] - 'a'
		if next.nodes[index] == nil {
			return false
		}
		next = next.nodes[index]
	}
	if next.count > 0 {
		return true
	}
	return false
}
