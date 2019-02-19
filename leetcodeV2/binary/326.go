package binary

/*
	solution 1: recursive
	solution 2: loop
	solution 3: log(n)/log(3) 是整数
*/

func isPowerOfThree(n int) bool {
	i := 1
	for i <= n {
		if i == n {
			return true
		}
		i *= 3
	}
	return false
}
