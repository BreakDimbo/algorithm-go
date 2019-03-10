package array

import (
	"sort"
)

// solution 1: two hash
// solution 2: 3 pointer
func threeSum(nums []int) [][]int {
	// sort
	sort.Ints(nums)
	r := make([][]int, 0)

	for i := range nums {
		// compare nums[i] and nums[i-1], not nums[i] and nums[i+1]
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		m, n := i+1, len(nums)-1
		for m < n {
			// sum is here not out of for
			sum := nums[i] + nums[m] + nums[n]
			if sum < 0 {
				m++
			} else if sum > 0 {
				n--
			} else {
				r = append(r, []int{nums[i], nums[m], nums[n]})
				for m < n && nums[m] == nums[m+1] {
					m++
				}
				for n > m && nums[n] == nums[n-1] {
					n--
				}
				// remember update m and n
				m++
				n--
			}
		}
	}
	return r
}

// solution 1: use hash O(n)
// solution 2: sort and return nums[n/2] O(nlgn)
// solution 3: vote with candidate O(n)
func majorityElement(nums []int) int {
	count, maj := 0, 0
	for _, num := range nums {
		if count == 0 {
			maj = num
		}
		if num == maj {
			count++
		} else {
			count--
		}
	}
	return maj
}

// solution 1: sort and find the first positive
// solution 2: put nums[i] > 0 -- nums[i] < len(nums) in right place
func firstMissingPositive(nums []int) int {
	for i := range nums {
		for nums[i] > 0 && nums[i] < len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	i := 0
	for ; i < len(nums) && nums[i] == i+1; i++ {
	}
	return i + 1
}
