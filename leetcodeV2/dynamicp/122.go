package dynamicp

/*
	solution 1: greedy
	solution 2: dp
*/

func maxProfitX(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	profit := 0
	for i := 0; i < len(prices); i++ {
		if i > 0 && prices[i-1] < prices[i] {
			profit += prices[i]
		}
		if i < len(prices)-1 && prices[i+1] > prices[i] {
			profit -= prices[i]
		}
	}
	return profit
}
