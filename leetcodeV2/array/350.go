package array

import "sort"

/*
	solution 1: sort and find
	solution 2: brute force
	solution 3: hash
*/

/*
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m := make(map[int]int)
	for _, num := range nums2 {
		m[num]++
	}

	res := []int{}
	for _, num := range nums1 {
		if count, ok := m[num]; ok && count > 0 {
			res = append(res, num)
			m[num]--
		}
	}
	return res
}
*/

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	res := []int{}
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return res
}
