package bdfs

import "testing"

func Test_canCompleteCircuit(t *testing.T) {
	canCompleteCircuit([]int{3, 1, 1}, []int{1, 2, 2})
}
