package others

import "math"

/*
	solution 1: n % 10
*/

func reverse(x int) int {
	ans := 0
	for x != 0 {
		pop := x % 10
		ans = ans*10 + pop
		x /= x
		if ans > math.MaxInt32 || ans == math.MaxInt32 && pop > 7 {
			return 0
		}
		if ans < math.MinInt32 || ans == math.MinInt32 && pop < -8 {
			return 0
		}
	}
	return ans
}
