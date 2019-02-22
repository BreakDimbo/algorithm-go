package bit

/*
	solution : x &= x-1
	solution 2 : dp
*/

/*
func countBits(num int) []int {
	counts := []int{}
	for i := 0; i <= num; i++ {
		count := 0
		k := i
		for k != 0 {
			k &= k - 1
			count++
		}
		counts = append(counts, count)
	}
	return counts
}
*/

func countBits(num int) []int {
	count := make([]int, num+1)
	count[0] = 0
	for i := 1; i <= num; i++ {
		count[i] = count[i&(i-1)] + 1
	}
	return count
}
