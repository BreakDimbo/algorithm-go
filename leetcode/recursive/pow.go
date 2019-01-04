package recursive

func pow(x, n int) int {
	for i := 0; i < n; i++ {
		x *= x
	}
	return x
}

func powR(x, n int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 1 / powR(x, -n)
	}

	// x 是偶数和奇数的情况下，都是取整操作
	result := powR(x, n/2)
	if n%2 == 0 {
		return result * result
	} else {
		return result * result * x
	}
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	pow := 1.0
	for n > 0 {
		if n&1 == 1 {
			pow *= x
		}
		x *= x
		n >>= 1
	}
	return pow
}
