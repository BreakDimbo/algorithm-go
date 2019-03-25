package hash

import "sort"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		m[t[i]]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if v, ok := m[target-num]; ok {
			return []int{i, v}
		}
		m[num] = i
	}
	return nil
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			continue
		}
		l, r := i+1, len(nums)-1
		sum := nums[l] + nums[r] + nums[i]
		for l < r {
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
			}
		}
	}
	return res
}

func isValid(s string) bool {
	m := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}

	for _, r := range s {
		if _, ok := m[r]; ok {
			stack = append(stack, r)
		} else {
			if len(stack) == 0 {
				return false
			} else {
				top := stack[len(stack)-1]
				if m[top] != r {
					return false
				}
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}
