package others

func myPow(x float64, n int) float64 {
	// n is negative
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

func majorityElement(nums []int) int {
	// solution 1: hash
	// solution 2: sort and return n/2
	// solution 3: divide and conq
	// solution 4: vote

	candidate, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func maxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices); i++ {
		if i > 0 && prices[i] > prices[i-1] {
			profit += prices[i]
		}
		if i < len(prices)-1 && prices[i] < prices[i+1] {
			profit -= prices[i]
		}
	}
	return profit
}
