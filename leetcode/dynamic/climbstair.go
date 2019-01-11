package dynamic

// f(n) = f(n-1) + f(n-2)
func climbStairs(n int) int {
	a, b := 1, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

var memo map[int]int

func climbStairsV2(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	memo = make(map[int]int)
	return climb(n)
}

func climb(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	if _, ok := memo[n]; !ok {
		memo[n] = climb(n-1) + climb(n-2)
		return memo[n]
	}
	return memo[n]
}
