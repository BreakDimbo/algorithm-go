package array

/*
	solution 1: sort O(nlg)
	solution 2: hash
	solution 3: bit
	solution 4: 高斯方程
*/

/*
func missingNumber(nums []int) int {
	n := len(nums)
	exptedNum := (1 + n) * n / 2
	actualNum := 0
	for _, num := range nums {
		actualNum += num
	}
	return exptedNum - actualNum
}
*/

func missingNumber(nums []int) int {
	missing := len(nums)
	for i, num := range nums {
		missing ^= i ^ num
	}
	return missing
}
