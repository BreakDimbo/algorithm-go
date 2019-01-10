package trie

// 使用 map[rune]*TrieNode

type TrieN struct {
	children map[byte]*TrieN
	isWord   bool
}

var result map[string]bool

// 方向数组
var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func findWords(board [][]byte, words []string) []string {
	if board == nil || words == nil {
		return nil
	}

	// 注意每次调用重新创建 result
	result = make(map[string]bool)

	root := &TrieN{children: make(map[byte]*TrieN)}
	// 将字符放入 Trie 表
	for _, w := range words {
		node := root
		for _, c := range w {
			if _, ok := node.children[byte(c)]; !ok {
				node.children[byte(c)] = &TrieN{children: make(map[byte]*TrieN)}
			}
			node = node.children[byte(c)]
		}
		node.isWord = true
	}

	for i, v := range board {
		for j, c := range v {
			// 注意只关心每个字符串的起始字母
			if _, ok := root.children[c]; ok {
				dfs(board, i, j, "", root) // 开始递归深度搜索
			}
		}
	}
	str := []string{}
	for w := range result {
		str = append(str, w)
	}
	return str
}

func dfs(board [][]byte, i, j int, c string, curDic *TrieN) {
	c += string(board[i][j])
	curDic = curDic.children[board[i][j]]

	if curDic.isWord {
		result[c] = true
	}

	// @ 表示当前格子被占用
	var tmp byte
	tmp, board[i][j] = board[i][j], byte('@')
	for k := 0; k < 4; k++ {
		x, y := i+dx[k], j+dy[k]
		if 0 <= x && x < len(board) && 0 <= y && y < len(board[0]) &&
			board[x][y] != byte('@') {
			if _, ok := curDic.children[board[x][y]]; ok {
				dfs(board, x, y, c, curDic)
			}
		}
	}
	// 恢复
	board[i][j] = tmp
}
