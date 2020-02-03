package main

import "fmt"

func main() {
	t := Constructor()
	t.Insert("apple")
	fmt.Println(t.Search("apple"))
	fmt.Println(t.Search("app"))
	fmt.Println(t.StartsWith("app"))
	t.Insert("app")
	fmt.Println(t.Search("app"))
}

type Trie struct {
	count  int
	prefix int
	nodes  []*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	t := Trie{}
	t.nodes = make([]*Trie, 26)
	return t
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	next := t
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

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	next := t
	for i := 0; i < len(word); i++ {
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

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	next := t
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		if next.nodes[index] == nil {
			return false
		}
		next = next.nodes[index]
	}
	return next.prefix > 0
}
