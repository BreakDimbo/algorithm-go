package bdfs

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
	solution 1: sort
*/

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	strNums := make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = fmt.Sprintf("%d", num)
	}

	sort.Slice(strNums, func(i, j int) bool {
		s1 := strNums[i] + strNums[j]
		s2 := strNums[j] + strNums[i]
		s1Int, _ := strconv.Atoi(s1)
		s2Int, _ := strconv.Atoi(s2)
		return s1Int > s2Int
	})

	if strNums[0][0] == '0' {
		return "0"
	}

	return strings.Join(strNums, "")
}
