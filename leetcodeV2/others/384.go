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
	tmp := make([]int, len(this.data))
	copy(tmp, this.data)

	for i := 0; i < len(this.data); i++ {
		index := rand.Intn(len(tmp))
		this.data[i] = tmp[index]
		tmp[index], tmp[len(tmp)-1] = tmp[len(tmp)-1], tmp[index]
		tmp = tmp[:len(tmp)-1]
	}
	return this.data
}
