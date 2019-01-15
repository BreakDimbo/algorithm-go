package priorityQueue

/*
	solution 1: 遍历，每次求最大值，O(n*k)
	solution 2: 使用大顶堆，每次移动 pop 旧元素，push 新元素，peek 堆顶元素
	solution 3: 双端队列
*/

/*
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	res := []int{}
	for i, j := 0, k; j <= len(nums); i, j = i+1, j+1 {
		res = append(res, util.Max(nums[i:j]...))
	}
	return res
}
*/

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	// 注意 window 存储的是下标
	window := make([]int, k)
	res := []int{}

	for i, v := range nums {
		// 判断 window 左边是否该移动, 注意边界条件
		if i >= k && window[0] <= i-k {
			window = window[1:]
		}

		// 新加入的元素从左向右依次与 window 中元素比较
		for len(window) > 0 && nums[window[len(window)-1]] <= v {
			window = window[:len(window)-1]
		}
		window = append(window, i)
		// 注意初始元素的添加
		if i >= k-1 {
			res = append(res, nums[window[0]])
		}
	}
	return res
}
