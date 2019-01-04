package recursive

func fab(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	a := 1
	b := 1

	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

func fabR(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fabR(n-1) + fabR(n-2)
}
