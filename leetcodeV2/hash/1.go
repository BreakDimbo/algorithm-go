package hash

/*
	solution 1: 穷举所有两个数组合的可能
	solution 2: 使用 hash 表，nums[i] 和 target-nums[i]
		注意所有操作可以放到一个循环中完成，因为 
		1. 将元素加入 map 前进行判断，防止 targe-v 是当前元素，导致元素重复使用
		2. 因为只有一个解，v 和 target-v 是互补的，如果遍历到 v 时，target-v 还没有被添加进 map，遍历到 target-v 时 v 一定已经在 map 中
*/

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{i, j}
		}
		m[v] = i
	}
	return nil
}
