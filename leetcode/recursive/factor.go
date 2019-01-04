package recursive

func factor(n int) int {
	sum := 1
	for i := n; i > 0; i-- {
		sum *= i
	}
	return sum
}

func factorR(n int) int {
	if n == 1 {
		return n
	}
	return (n) * factorR(n-1)
}
