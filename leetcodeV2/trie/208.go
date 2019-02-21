package trie

type TrieNode struct {
	val      rune
	isWord   bool
	children []*TrieNode
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		&TrieNode{
			children: make([]*TrieNode, 26), // use hash if not just letter
		},
	}
}

func (t *Trie) Insert(s string) {
	node := t.root
	for _, r := range s {
		if node.children[r-'a'] == nil {
			node.children[r-'a'] = &TrieNode{
				val:      r,
				children: make([]*TrieNode, 26),
			}
		}
		node = node.children[r-'a']
	}
	node.isWord = true
}

func (t *Trie) Search(s string) bool {
	node := t.root
	for _, r := range s {
		if node.children[r-'a'] == nil {
			return false
		}
		node = node.children[r-'a']
	}
	return node.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, r := range prefix {
		if node.children[r-'a'] == nil {
			return false
		}
		node = node.children[r-'a']
	}
	return true
}
