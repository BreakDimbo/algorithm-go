package trie

type TrieNode struct {
	val      rune
	isWord   bool
	children []*TrieNode
}

type Trie struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			isWord:   false,
			children: make([]*TrieNode, 26),
		},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this.root
	for _, c := range word {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = &TrieNode{val: c, children: make([]*TrieNode, 26)}
		}
		node = node.children[c-'a']
	}
	node.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.root
	for _, c := range word {
		if node.children[c-'a'] == nil {
			return false
		}
		node = node.children[c-'a']
	}
	return node.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, c := range prefix {
		if node.children[c-'a'] == nil {
			return false
		}
		node = node.children[c-'a']
	}
	return true
}
