package greedy

func maxProfit(prices []int) int {
	profile := 0

	for i := 0; i < len(prices); i++ {
		if i != 0 && prices[i-1] < prices[i] {
			profile += prices[i]
		}

		if i != len(prices)-1 && prices[i+1] > prices[i] {
			profile -= prices[i]
		}
	}
	return profile
}
