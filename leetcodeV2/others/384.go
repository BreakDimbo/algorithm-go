package others

import "math/rand"

/*
	solution 1: brute force
	solution 2: swap directly
*/

type Solution struct {
	data   []int
	origin []int
}

func ConstructorXXX(nums []int) Solution {
	origin := make([]int, len(nums))
	copy(origin, nums)
	return Solution{
		data:   nums,
		origin: origin,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	reset := make([]int, len(this.origin))
	copy(reset, this.origin)
	this.data = reset
	return reset
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	n := len(this.data)
	for i := n - 1; i >= 0; i-- {
		index := rand.Intn(i + 1)
		this.data[index], this.data[i] = this.data[i], this.data[index]
	}
	return this.data
}
