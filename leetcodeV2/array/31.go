package array

/*
	solution 1: 从最后一个元素 i 开始，依次比较其和前 i-1 个元素的大小
		当发现元素 i 大于其前面的某个元素 j 时，说明发现候选元素，记录最高位 j，以及交换到 j 的最小元素
	solution 2:
*/

/*
func nextPermutation(nums []int) {
	high, low := 0, 0
	minToHigh := make([]int, len(nums))
	for j := range minToHigh {
		minToHigh[j] = math.MaxInt32
	}

	for i := len(nums) - 1; i >= 1; i-- {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] && j >= high && nums[i] < minToHigh[j] {
				high = j
				low = i
				minToHigh[j] = nums[i]
			}
		}
	}

	nums[low], nums[high] = nums[high], nums[low]
	if low == 0 {
		// not found next Permutation，sort the num
		high = -1
	}
	sort.Ints(nums[high+1:])
}
*/

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	// reverse i+1..n
	m, n := i+1, len(nums)-1
	for m < n {
		nums[m], nums[n] = nums[n], nums[m]
		m++
		n--
	}
}
