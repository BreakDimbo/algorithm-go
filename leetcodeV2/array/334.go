package array

import "math"

/*
	solution 1: brute force O(n^2)
	solution 2: small bit and bigger than both
*/

func increasingTriplet(nums []int) bool {
	small, big := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num <= small {
			small = num
		} else if num <= big {
			big = num
		} else {
			return true
		}
	}
	return false
}
