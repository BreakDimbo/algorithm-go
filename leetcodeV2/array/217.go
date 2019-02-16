package array

/*
	solution 1: sort O(nlgn)/O(1)
	solution 2: hash O(n)/O(n)
*/

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		if m[v] {
			return true
		}
		m[v] = true
	}
	return false
}
