package bdfs

/*
	solution 1: dfs
	solution 2:
		如果 A 不能到达 B，那么 A-B 之间的任何节点都不能到达 B
		如果 sum(gas) - sum(cost) >= 0，则必有一个解
*/

/*
func canCompleteCircuit(gas []int, cost []int) int {
	for i := range gas {
		if travel(i, i, 0, gas, cost, true) {
			return i
		}
	}
	return -1
}

func travel(i, start, unit int, gas, cost []int, canMove bool) bool {
	if !canMove {
		return false
	}
	if i >= len(gas) {
		i -= len(gas)
	}

	unit += gas[i]
	if (i-start == -1 || i-start == len(gas)-1) && unit >= cost[i] {
		canMove = true
		return true
	}
	if unit < cost[i] {
		canMove = false
		return false
	}

	canMove = travel(i+1, start, unit-cost[i], gas, cost, canMove) && canMove
	return canMove
}
*/

func canCompleteCircuit(gas []int, cost []int) int {
	var sumGas, sumCost, tank, start int
	for i := 0; i < len(gas); i++ {
		sumCost += cost[i]
		sumGas += gas[i]
		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}
	if sumGas < sumCost {
		return -1
	}
	return start
}
