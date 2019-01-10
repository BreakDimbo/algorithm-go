package trie

type TrieNode struct {
	val      rune
	isWord   bool
	children map[rune]*TrieNode
}

type Trie struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			isWord:   false,
			children: make(map[rune]*TrieNode),
		},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this.root
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			node.children[c] = &TrieNode{val: c, children: make(map[rune]*TrieNode)}
		}
		node = node.children[c]
	}
	node.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.root
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			return false
		}
		node = node.children[c]
	}
	return node.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, c := range prefix {
		if _, ok := node.children[c]; !ok {
			return false
		}
		node = node.children[c]
	}
	return true
}
