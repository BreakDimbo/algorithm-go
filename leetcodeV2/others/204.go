package others

import "math"

/*
	solution 1: 试除法：1..n 的平方根
	solution 2: 埃拉托斯特尼筛法
*/

/*
func countPrimes(n int) int {
	count := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
*/

func countPrimes(n int) int {
	count := 0
	notPrime := make([]bool, n)
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if !notPrime[i] {
			for j := 2; i*j < n; j++ {
				notPrime[i*j] = true
			}
		}
	}
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			count++
		}
	}
	return count
}
