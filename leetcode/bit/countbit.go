package bit

func countBits(num int) []int {
	count := make([]int, num+1)
	for i := 1; i <= num; i++ {
		count[i] = count[i&(i-1)] + 1
	}
	return count
}
