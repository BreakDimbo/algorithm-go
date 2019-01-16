package recursive

/*
	solution 1: brute force, n 次循环做乘法
	solution 2: 二分法，每次循环 n/2，乘 logn 次 x*x, 注意考虑负数的情况
*/

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	pow := 1.0
	for n != 0 {
		if n&1 == 1 {
			pow *= x
		}
		x *= x
		n >>= 1
	}
	return pow
}
