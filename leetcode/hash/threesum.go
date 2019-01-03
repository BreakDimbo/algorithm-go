package hash

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	res := [][]int{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// slice 的时候，右边是开区间
	for i, v := range nums[:len(nums)-1] {
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if v == 0 {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			s := v + nums[l] + nums[r]
			if s > 0 {
				r--
			} else if s < 0 {
				l++
			} else {
				res = append(res, []int{v, nums[l], nums[r]})
				// 去重
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}

	return res
}
