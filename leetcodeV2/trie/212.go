package trie

type TrieN struct {
	children []*TrieN
	isWord   bool
}

var (
	dx = []int{0, 0, 1, -1}
	dy = []int{1, -1, 0, 0}
)

func findWords(board [][]byte, words []string) []string {
	trie := &TrieN{
		children: make([]*TrieN, 26),
	}

	// pre process the words
	for _, word := range words {
		node := trie
		for _, r := range word {
			if node.children[r-'a'] == nil {
				node.children[r-'a'] = &TrieN{children: make([]*TrieN, 26)}
			}
			node = node.children[r-'a']
		}
		node.isWord = true
	}

	resMap := make(map[string]bool)
	// traverse the board
	for i := range board {
		for j := range board[i] {
			if trie.children[board[i][j]-'a'] != nil {
				dfs(board, i, j, string(board[i][j]), trie.children[board[i][j]-'a'], resMap)
			}
		}
	}

	res := []string{}
	for s := range resMap {
		res = append(res, s)
	}
	return res
}

func dfs(board [][]byte, i, j int, curStr string, trieNode *TrieN, resMap map[string]bool) {
	if trieNode.isWord {
		resMap[curStr] = true
	}

	tmp := board[i][j]
	board[i][j] = '@'
	for k := 0; k < 4; k++ {
		x := i + dx[k]
		y := j + dy[k]
		if x >= 0 && x < len(board) && y >= 0 && y < len(board[i]) &&
			board[x][y] != '@' && trieNode.children[board[x][y]-'a'] != nil {
			dfs(board, x, y, curStr+string(board[x][y]), trieNode.children[board[x][y]-'a'], resMap)
		}
	}
	board[i][j] = tmp
}
