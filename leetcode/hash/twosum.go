package hash

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		match := target - v
		if j, ok := m[match]; ok {
			return []int{i, j}
		}
		m[v] = i
	}
	return nil
}
