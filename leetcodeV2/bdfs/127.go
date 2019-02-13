package bdfs

/*
	solution 1: bfs
*/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	reached := make([]string, 0)
	unreached := make(map[string]bool)

	for _, word := range wordList {
		unreached[word] = true
	}
	if _, ok := unreached[endWord]; !ok {
		return 0
	}
	reached = append(reached, beginWord)
	distance := 1

	for unreached[endWord] {
		toVisit := make([]string, 0)
		for _, word := range reached {
			for i := 0; i < len(word); i++ {
				for j := 'a'; j <= 'z'; j++ {
					wordArr := []rune(word)
					wordArr[i] = j
					w := string(wordArr)
					if unreached[w] {
						toVisit = append(toVisit, w)
						unreached[w] = false
					}
				}
			}
		}
		distance++
		if len(toVisit) == 0 {
			return 0
		}
		reached = toVisit
	}
	return distance
}
