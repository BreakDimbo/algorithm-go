package bdfs

import "testing"

func Test_ladderLength(t *testing.T) {
	ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"})
}
