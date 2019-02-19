package bit

/*
explanation: https://leetcode.com/problems/sum-of-two-integers/discuss/132479/Simple-explanation-on-how-to-arrive-at-the-solution
*/

func getSum(a int, b int) int {
	for b != 0 {
		c := a & b // 记录需要进位的 bit
		a = a ^ b // 记录不需要进位的 bit
		b = c << 1 // 进位
	}
	return a
}
