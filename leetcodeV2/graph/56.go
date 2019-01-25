package graph

import (
	"sort"

	"github.com/breakD/algorithms/leetcodeV2/util"
)

/*
	solution 1: Connected Components
	solution 2: Sort
*/

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}

	// 排序的时候用 等于 速度会变慢？
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	merged := []Interval{intervals[0]}
	j := 0

	for i := 1; i < len(intervals); i++ {
		if merged[j].End >= intervals[i].Start {
			merged[j].End = util.Max(merged[j].End, intervals[i].End)
		} else {
			merged = append(merged, intervals[i])
			j++
		}
	}
	return merged
}
