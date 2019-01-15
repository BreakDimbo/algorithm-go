package priorityQueue

import "testing"

func TestKthLargest_Add(t *testing.T) {
	obj := Constructor(3, []int{4, 5, 8, 2})
	obj.Add(5)
}
